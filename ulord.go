package main

import (
	"reflect"
	"time"
	"fmt"
        "runtime"
	"sync"

	"github.com/ulordsuite/ulord/ulordec"
	"github.com/ulordsuite/ulord/chaincfg"
	"github.com/ulordsuite/ulordutil"
	log "github.com/sirupsen/logrus"
)

type UlordCmd struct {
	Network string `long:"network" default:"mainnet" description:"network"`
	Pattern string `long:"pattern" default:"Uss" description:"match pattern"`
}

func isEmpty(a interface{}) bool {
    v := reflect.ValueOf(a)
    if v.Kind() == reflect.Ptr {
        v=v.Elem()
    }
    return v.Interface() == reflect.Zero(v.Type()).Interface()
}

func (ulord *UlordCmd) Execute(args []string) ([]string, error) {
	if len(args) >= 1 {
		ulord.Pattern = args[0]
	}
	if len(ulord.Pattern) > 6 {
		//log.Fatal("Quitting, this pattern would take too much time.")
		log.Info("Quitting, this pattern would take too much time.")
	} else if len(ulord.Pattern) > 4 {
		log.Info("This pattern could take awhile, please wait.")
	}

	beginTime := time.Now()
	//prefix := strings.ToLower(ulord.Pattern)
	prefix := ulord.Pattern
	chainParams := &chaincfg.MainNetParams
	
	fmt.Println("*****", ulord, ulord.Pattern)

	switch ulord.Network {
	case "mainnet":
		chainParams = &chaincfg.MainNetParams
	case "testnet":
		chainParams = &chaincfg.TestNet3Params
	default:
		chainParams = &chaincfg.MainNetParams
	}
	
	ch := make(chan string, 2)
	lock := &sync.Mutex{}
	
	loop := runtime.NumCPU() 
	fmt.Println("loop:",loop)

	for i := 0; i < loop; i++ {
		go func(){
		var numAttempts int64 = 0
		foundAddr := ""
		foundWif := ""
		for {
		numAttempts++

		privKey, err := ulordec.NewPrivateKey(ulordec.S256())
		if err != nil {
			log.Fatalf("Failed to create private key, err: %v", err)
		}

		addrPubKey, err := ulordutil.NewAddressPubKey(
			privKey.PubKey().SerializeUncompressed(), chainParams)
		if err != nil {
			log.Fatalf("Failed to calculate public key, err: %v", err)
		}

		rcvAddr := addrPubKey.AddressPubKeyHash().EncodeAddress()
		 //log.Infof("rcvAddr:%s  prefix:%s \n", rcvAddr, prefix)
		if matchPrefix(rcvAddr, prefix) {
			foundAddr = rcvAddr
			wif, err := ulordutil.NewWIF(privKey, chainParams, true )//false)
			if err != nil {
				log.Fatalf("failed to get wif: %s", err)
			}
			foundWif = wif.String()
			log.Infof("privkey:%s\n",privKey)
			lock.Lock()
			ch <-foundAddr
			ch <-foundWif
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

	log.Infof("\nElapsed: %s\naddr: %s\nwif: %s\n",
		time.Since(beginTime), Addr, Wif)

	return result, nil
}


var ulordCmd UlordCmd

func init() {
	parser.AddCommand("ulord", "get a ulord vanity address", "The ticker command get a ulord vanity address", &ulordCmd)
//	fmt.Println("ulord.go init:",ulordCmd)
}
