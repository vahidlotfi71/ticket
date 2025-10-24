package main

import (
	"fmt"
	"os"
	"vahid/Commands"
	"vahid/Config"
)

func main() {
	if err := Config.GetEnv(); err != nil {
		fmt.Println("Error while loading the .env file")
		fmt.Printf("\t%s\n", err.Error())
		os.Exit(2)
	} else {
		fmt.Printf("Loaded .env file successfully ...\n")
	}

	if err := Config.Connect(); err != nil {
		fmt.Printf("Error while connecting to the database :\n")
		fmt.Printf("\t%s\n", err.Error())
		os.Exit(2)
	} else {
		fmt.Printf("Connected to the database successfully ...\n")
	}
	Commands.AutoMigrate()
}
