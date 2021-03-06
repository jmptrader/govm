package main

import (
	"fmt"
	"os"
)

func exit(msg string) {
	fmt.Printf("\033[1m\033[31mgovm\033[39m\033[0m: %s\n", msg)
	os.Exit(1)
}

const (
	regCount      = 16
	stackSize     = 256
	callStackSize = 64
)

func main() {
	if len(os.Args) < 2 {
		exit("incorrect number of arguments")
	}

	compile := false

	if os.Args[1] == "version" {
		fmt.Println("GoVM 0.2")
		os.Exit(0)
	} else if os.Args[1] == "compile" {
		compile = true
	}

	if compile {
		if len(os.Args) < 4 {
			exit("compilation needs at least 2 extra arguments")
		}

		// Make sure all files exist
		for _, arg := range os.Args[3:] {
			if _, err := os.Stat(arg); os.IsNotExist(err) {
				exit(fmt.Sprintf("file \"%s\" does not exist", arg))
			}
		}

		// Main file
		file, err := os.Open(os.Args[3])
		if err != nil {
			exit(fmt.Sprintf("file \"%s\" could not be opened", os.Args[3]))
		}

		code, count := parser(file, os.Args[3], 0)
		file.Close()

		// Extra files
		for _, arg := range os.Args[4:] {
			file, err := os.Open(arg)
			if err != nil {
				exit(fmt.Sprintf("file \"%s\" could not be opened", arg))
			}

			var newCode []float64
			newCode, count = parser(file, arg, count)
			code = append(code, newCode...)
			file.Close()
		}

		// Debug
		// fmt.Println(decode(code))

		evaluateLabels(code)

		writeFile(code, os.Args[2])
	} else {
		code := loadFile(os.Args[1])

		// Debug
		// fmt.Println(decode(code))

		run(code)
	}
}
