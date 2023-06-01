package main

import (
	"errors"
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

	if listAll {
		printList()
		os.Exit(0)
	}

	args := flag.Args()
	converter, err := NewConverter(caseType)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error() + "\n")
        printList()
		os.Exit(1)
	}

	if err := runCmd(converter, args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())

		if strings.Contains(err.Error(), ErrNoText) {
			printHelp()
		}
		os.Exit(1)
	}

	os.Exit(0)
}

func runCmd(c *Converter, args []string) error {
	if len(args) == 0 && !isFlagPassed("f", "file") {
		return fmt.Errorf("%v\n", ErrNoText)
	}

	var errs []error
	if len(fileInput) > 0 {
		go c.convertFileLines(fileInput)
		for result := range c.outch {
			if result.error != nil {
				errs = append(errs, result.error)
				continue
			}
			fmt.Println(result.text)
		}

		if len(errs) > 1 {
			return errors.Join(errs...)
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
	listAll   bool
	fileInput string
)

const ErrNoText = "no text to convert provided"

func init() {
	flag.StringVar(&caseType, "case", CamelCase, "case to convert to")
	flag.StringVar(&caseType, "c", CamelCase, "case to convert to")
	flag.BoolVar(&help, "help", false, "print help")
	flag.BoolVar(&help, "h", false, "print help")
	flag.BoolVar(&listAll, "list", false, "print all supported cases")
	flag.BoolVar(&listAll, "l", false, "print all supported cases")
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
	usage := `Usage: case-converter [options...] <text>
-c, case string
    case to convert to
-l, list
    print all supported cases
-f, file string
    file to convert line by line
-h, help
    print help`

	fmt.Fprintln(os.Stderr, usage)
}

func printList() {
	fmt.Fprintf(os.Stderr, "supported cases:\n%s\n", strings.Join(cases[:], "\n"))
}
