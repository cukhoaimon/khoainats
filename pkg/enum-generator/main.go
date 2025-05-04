package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var file string
	var pkgName string
	var showHelp bool

	flag.StringVar(&file, "file", "", "Path to the Go file containing enum definitions (required)")
	flag.StringVar(&pkgName, "pkgName", "", "Generated package. Default is 'enum'")
	flag.BoolVar(&showHelp, "help", false, "Show help message")
	flag.Parse()

	if showHelp || file == "" {
		printUsage()
		return
	}

	if pkgName == "" {
		pkgName = "enum"
	}

	if err := processFile(file, pkgName); err != nil {
		fmt.Fprintln(os.Stderr, "❌", err)
		os.Exit(1)
	}
}
