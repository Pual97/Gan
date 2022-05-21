package main

import (
	"gan/Context"
	"gan/HttpBase"
	"log"
	"os"
	"time"
)

func main() {
	name := os.Args[1]
	log.Printf("========================================================\n")
	log.Printf("|  server Info:\n")
	log.Printf("|  Server Name: %s\n", name)
	log.Printf("|  Time: %s\n", time.Now().Format("2003-01-20 15:04:05"))
	log.Printf("|  VERSION: %s\n", name+"-1.0.0")
	log.Printf("========================================================\n")
	switch name {
	case "http-base":
		HttpBase.Run()
		return
	case "context":
		Context.Run()
	default:
		log.Println("invalid server args")
		os.Exit(1)
	}
}
