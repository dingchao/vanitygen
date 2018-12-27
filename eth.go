package main

import (
	"encoding/hex"
//	"strings"
	"time"
        "runtime"
        "sync"

	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
)

type EthCmd struct {
	Pattern string `long:"pattern" default:"444" description:"match pattern"`
}

func (ec *EthCmd) Execute(args []string) ([]string, error) {
        if len(args) >= 1 {
                ec.Pattern = args[0]
		//fmt.Println("eth Execute:",ec)
        }
	beginTime := time.Now()
	//prefix := strings.ToLower(ec.Pattern)
	prefix := ec.Pattern

//	fmt.Println("--------------------------")
//	debug.PrintStack()
//	fmt.Println("--------------------------")

        ch := make(chan string, 2)
        lock := &sync.Mutex{}
        
        loop := runtime.NumCPU() 

        for i := 0; i < loop; i++ {
                go func(){

	var numAttempts int64 = 0
	addrStr := ""
	keyStr := ""

	for {
		numAttempts++

		key, _ := crypto.GenerateKey()
		addr := crypto.PubkeyToAddress(key.PublicKey)
		addrStr = hex.EncodeToString(addr[:])
		if matchPrefix(addrStr, prefix) {
			keyStr = hex.EncodeToString(crypto.FromECDSA(key))
			// fmt.Println("pub:", hex.EncodeToString(crypto.FromECDSAPub(&key.PublicKey)))
                        lock.Lock()
                        ch <-addrStr
                        ch <-keyStr
                        lock.Unlock()

			break
		}
	}
        }()

        }

        Addr :=  <-ch
        Wif := <-ch
        //defer close(ch)

        var result []string
        //result = append(result, time.Since(beginTime))
        result = append(result, Addr)
        result = append(result, Wif)

	log.Infof("\nElapsed: %s\naddr: 0x%s\npvt: 0x%s\n",
		time.Since(beginTime), Addr, Wif)

	return result,nil
}

var ethCmd EthCmd

func init() {
	parser.AddCommand("eth", "get a ETH vanity address", "The ticker command get a ETH vanity address", &ethCmd)
}
