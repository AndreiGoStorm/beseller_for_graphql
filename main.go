package main

import (
	"flag"
	"fmt"
	"os"

	"beseller/internal/app"
	"beseller/internal/config"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "config/config.yml", "Path to configuration file")
}

func main() {
	conf := config.New(configFile)
	application := app.NewApp(conf)

	err := application.DoRequest()
	if err != nil {
		fmt.Println("Error executing request:", err)
		os.Exit(1)
	}
	err = application.Write()
	if err != nil {
		fmt.Println("Error writing file:", err)
		os.Exit(1)
	}
	err = application.Close()
	if err != nil {
		fmt.Println("Error closing file:", err)
		os.Exit(1)
	}

	fmt.Printf("File written successfully - %s\n", application.File)
}
