package main

import (
//	"strings"
	"time"
	"fmt"
	"reflect"

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
			break
		}
	}

        var result []string
        //result = append(result, time.Since(beginTime))
        result = append(result, foundAddr)
        result = append(result, foundWif)

	log.Infof("\nElapsed: %s\naddr: %s\nwif: %s\nattempts: %d.\n",
		time.Since(beginTime), foundAddr, foundWif, numAttempts)

	return result,nil
}


var ulordCmd UlordCmd

func init() {
	parser.AddCommand("ulord", "get a ulord vanity address", "The ticker command get a ulord vanity address", &ulordCmd)
//	fmt.Println("ulord.go init:",ulordCmd)
}
