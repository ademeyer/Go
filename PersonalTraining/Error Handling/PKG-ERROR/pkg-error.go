package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

type Config struct {
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, errors.Wrap(err, "Unable to open configuration file")
	}
	defer file.Close()

	cfg := &Config{}
	//Parse data into config here

	return cfg, nil
}

func setupLogging() {
	out, err := os.OpenFile("./app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return
	}
	log.SetOutput(out)

}

func main() {
	setupLogging()
	cfg, err := readConfig("/path/config.toml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		log.Printf("error: %+v", err)
		os.Exit(1)
	}

	fmt.Println(cfg)

}
