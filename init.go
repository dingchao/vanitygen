package main
import "fmt"
import (
	"runtime"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	Conf     string `long:"conf" description:"Config file"`
	LogLevel string `long:"log_level" default:"info" description:"log level"`
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("set user cpu numb:", runtime.NumCPU())
}

var parser = flags.NewParser(&opts, flags.Default|flags.IgnoreUnknown)
