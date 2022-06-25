// pkg/errors example
package main

import (
	"fmt"
	"os"
	"log"
	"github.com/pkg/errors"
)
// Config holds configuration
type Config struct {
	// configuration fields go here 
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, " can't open the configuration file")
	}

	defer file.Close()

	cfg := &Config{}
	return cfg, nil

}

func setupLogging() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.SetOutput(out)
}


func main(){
	setupLogging()
	cfg, err := readConfig("/path/to/config.toml")
	if err != nil {
		fmt.Printf("error: %s\n", err)
		log.Printf("error: %+v\n",err)
		os.Exit(1)
	}

	// Normal operation 
	fmt.Println(cfg)
}
