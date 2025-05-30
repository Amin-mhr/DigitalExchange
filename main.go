/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"DigitalExchange/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
