package main

import (
	"errors"
	"strings" 
	"fmt"
	"os"
	"time"

        "github.com/Akagi201/utilgo/conflag"
        log "github.com/sirupsen/logrus"
)

type Vanitygen interface {
	Getvanityaddress(string, string) (string,string, error)
}

type  vanitygen struct{}

func help(){
        fmt.Println("\nfunction: get a vanity address")
        fmt.Println("Usage example:\"./bin CoinName  AddressPrefix\"")
        fmt.Println("----------------------------------------------")
        fmt.Println("btc:      Address Prefix: 1;  example:./bin btc 111")
        fmt.Println("eth:      Address Prefix: 4;  example:./bin eth 433")
        fmt.Println("ulord:    Address Prefix: U;  example:./bin ulord Uss")
        fmt.Println("----------------------------------------------\n")
}

var Errparam = errors.New("parameter error!!! please check it")

func checkparam(coin string, frefix string) bool{
         if coin == "" || frefix == "" {
                return  false
        }
        
        switch (coin){
                case "btc":
                        if !strings.HasPrefix(frefix, "1"){
                                help()  
                                return false    
                        }
                        break

                case "eth":
                        if !strings.HasPrefix(frefix, "4"){
                                help()
                                return false
                        }
                        break
                case "ulord":
                         if !strings.HasPrefix(frefix, "U"){
                                help()
                                return false
                        }
                        break
		default : 
			return false
        }       

        return true
}



func (vanitygen) Getvanityaddress(coin string, prefix string) (string, string, error) {
	if !checkparam(coin, prefix) {
		return "","", Errparam
	}
	if len(os.Args) < 3 {
		os.Args = append(os.Args, coin)
        	os.Args = append(os.Args, prefix)
		os.Args = append(os.Args, "-p")
		os.Args = append(os.Args, "20")
	}else{
		os.Args[1] = coin
		os.Args[2] = prefix
	}

	val, erro :=parser.Parse()

	fmt.Println(val, erro)

        if opts.Conf != "" {
                conflag.LongHyphen = true
                conflag.BoolValue = false
                args, err := conflag.ArgsFrom(opts.Conf)
                if err != nil {
                        panic(err)
                }

                parser.ParseArgs(args)
                fmt.Println("\nparser.ParseArgs\n")
        }

        log.Debugf("opts: %+v", opts)
        log.Info("opts: %+v", opts)

        if opts.LogLevel == "" {
                return "","", nil
        }

        level, err := log.ParseLevel(strings.ToLower(opts.LogLevel))
        if err != nil {
                log.Fatalf("log level error: %v", err)
        }

        log.SetLevel(level)

        log.SetFormatter(&log.TextFormatter{
                FullTimestamp:   true,
                TimestampFormat: time.RFC3339,
        })


	return val[0], val[1],  nil
}

type ServiceMiddleware func(Vanitygen) Vanitygen

