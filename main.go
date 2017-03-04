package main

import (
	"flag"
	"fmt"

	"github.com/markhuge/status/config"
)

func main() {
	// Setup config
	configFile := flag.String("f", "config.yml", "path to config file")
	flag.Parse()

	conf := config.Read(*configFile)
	fmt.Printf("conf = %+v\n", conf)

}
