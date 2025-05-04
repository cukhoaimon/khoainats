package main

import "fmt"

func printUsage() {
	fmt.Println(`Usage:
  enumgen --file <filename>

Options:
  --file     Path to the Go source file with enum constants
  --pkgName  Package name of generated enum
  --help     Show this help message

Example:
  enumgen --file principal_type.go`)
}
