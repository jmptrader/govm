package main

import (
	"fmt"
	"os"
)

func exit(msg string) {
	fmt.Printf("govm: %s\n", msg)
	os.Exit(1)
}

const (
	maxCodeSize   = 256
	regCount      = 16
	dataSize      = 64
	stackSize     = 256
	callStackSize = 64
)

func main() {
	if len(os.Args) != 2 {
		exit("incorrect number of arguments")
	}

	if os.Args[1] == "--version" {
		fmt.Println("GoVM d0.0")
		os.Exit(0)
	}

	if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		exit("input file does not exist")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		exit("could not open input file for reading")
	}

	code := parser(file)
	file.Close()

	// Debug
	// fmt.Println(decode(code))

	run(code)
}
