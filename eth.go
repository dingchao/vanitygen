package main

import (
	"encoding/hex"
//	"strings"
	"time"
	"runtime/debug"
	"fmt"

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

	fmt.Println("--------------------------")
	debug.PrintStack()
	fmt.Println("--------------------------")

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
			break
		}
	}

        var result []string
        //result = append(result, time.Since(beginTime))
        result = append(result, addrStr)
        result = append(result, keyStr)

	log.Infof("\nElapsed: %s\naddr: 0x%s\npvt: 0x%s\nattempts: %d.\n",
		time.Since(beginTime), addrStr, keyStr, numAttempts)

	return result,nil
}

var ethCmd EthCmd

func init() {
	parser.AddCommand("eth", "get a ETH vanity address", "The ticker command get a ETH vanity address", &ethCmd)
}
