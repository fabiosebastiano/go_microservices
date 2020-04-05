package main

import (
	"log"
	"os"
)

func main() {

	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	ph := handlers


	l.Println("Server startato su porta 9090")
}
