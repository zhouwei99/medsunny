package main

import (
	_ "./src/server"
	"log"
)

func init() {
	log.SetPrefix("INFO")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	log.Println("med sunny server is running ...")
}
