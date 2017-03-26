package main

import (
	"fmt"
	"os"
)

func main() {
	fileName, err := parseArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		usage()
		os.Exit(1)
	}
	fmt.Println(fileName)
	os.Exit(0)
}

func parseArgs() (string, error) {
	if len(os.Args) == 1 {
		return "", fmt.Errorf("Invalid number of arguments (%d)", len(os.Args)-1)
	}
	return "", fmt.Errorf("TODO")
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage:\n\tcat file")
}
