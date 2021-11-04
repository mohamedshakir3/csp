package main

import (
	"flag"
	"io/ioutil"
	"fmt"
	"encoding/json"
	config "csp"
)

var filename = flag.String("config", "config.json", "Location of the config file.")

func main(){
	flag.Parse()
	data, err := ioutil.ReadFile(*filename)
	if err != nil {
		fmt.Print(err)
	}

	var config config.Config
	err = json.Unmarshal(data, &config)
	theta.NewCspFinder(config).Execute()

}
