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
		printHelp()
		os.Exit(0)
	}

	args := flag.Args()
	converter, err := NewConverter(caseType)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := runCmd(converter, args); err != nil {
		fmt.Println(err)

		if err.Error() == ErrNoText {
			printHelp()
		}
		os.Exit(1)
	}

	os.Exit(0)
}

func runCmd(c *Converter, args []string) error {
	var err error
	if len(args) == 0 && !isFlagPassed("f", "file") {
		return fmt.Errorf("%v", ErrNoText)
	}

	if len(fileInput) > 0 {
		if err = c.convertFileLines(fileInput); err != nil {
			return err
		}
	}

	if len(args) > 1 {
		conv := c.convert(args)
		fmt.Println(conv)
	} else {
		// Handle input string contained in quotation marks ("")
		for _, s := range args {
			fields := strings.Fields(s)
			conv := c.convert(fields)
			fmt.Println(conv)
		}
	}

	return nil
}

var (
	caseType  string
	help      bool
	fileInput string
)

const ErrNoText = "no text to convert provided"

func init() {
	flag.StringVar(&caseType, "case", CamelCase, "case to convert to")
	flag.StringVar(&caseType, "c", CamelCase, "case to convert to")
	flag.BoolVar(&help, "help", false, "print help")
	flag.BoolVar(&help, "h", false, "print help")
	flag.StringVar(&fileInput, "file", "", "input file to convert")
	flag.StringVar(&fileInput, "f", "", "input file to convert")
}

func isFlagPassed(names ...string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		for _, name := range names {
			if f.Name == name {
				found = true
			}
		}
	})
	return found
}

func printHelp() {
	os.Stderr.WriteString(`Usage: case-converter [options...] <text>
-c, case string
    case to convert to (default "camel")
-f, file string
    input file to convert
-h, help
print help`)
}
