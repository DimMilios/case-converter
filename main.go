package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Parse()

	if help {
		fmt.Println("Case Converter")
		flag.PrintDefaults()
		os.Exit(0)
	}

	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("no text to convert provided")
		os.Exit(1)
	}

	converter := NewConverter(caseType)

	if len(args) > 1 {
		conv := converter.convert(&args)
		fmt.Println(conv)
		os.Exit(0)
	}

	// Handle input string contained in quotation marks ("")
	var argFields [][]string
	for _, s := range args {
		fields := strings.Fields(s)
		argFields = append(argFields, fields)
	}

	for _, f := range argFields {
		conv := converter.convert(&f)
		fmt.Println(conv)
	}
	os.Exit(0)
}

var caseType string
var help bool

func init() {
	flag.StringVar(&caseType, "case", CamelCase, "case to convert to")
	flag.StringVar(&caseType, "c", CamelCase, "case to convert to")
	flag.BoolVar(&help, "help", false, "print help")
	flag.BoolVar(&help, "h", false, "print help")
}
