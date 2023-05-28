package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const (
	LowerCase    = "lower"  // lowercase
	UpperCase    = "upper"  // UPPERCASE
	CamelCase    = "camel"  // camelCase
	PascalCase   = "pascal" // PascalCase
	SnakeCase    = "snake"  // snake_case
	KebabCase    = "kebab"  // kebab-case
	ConstantCase = "const"  // CONSTANT_CASE
	TrainCase    = "train"  // Train-Case
)

var cases = [8]string{LowerCase, UpperCase, CamelCase, PascalCase, SnakeCase, KebabCase, ConstantCase, TrainCase}

type Converter struct {
	caseType string
}

func NewConverter(caseType string) (*Converter, error) {
	if !isCaseSupported(caseType) {
		return nil, fmt.Errorf("case \"%s\" is not supported", caseType)
	}

	return &Converter{
		caseType: caseType,
	}, nil
}

func (c *Converter) convert(words []string) string {
	var out string
	switch c.caseType {
	case CamelCase:
		out = convertCamel(words)
		break
	case SnakeCase:
		out = convertLower(words, "_")
		break
	case KebabCase:
		out = convertLower(words, "-")
		break
	case LowerCase:
		out = convertLower(words, " ")
		break
	case UpperCase:
		out = convertUpper(words, " ")
		break
	case PascalCase:
		out = convertPascal(words, "")
		break
	case ConstantCase:
		out = convertUpper(words, "_")
		break
	case TrainCase:
		out = convertPascal(words, "-")
		break
	}
	return out
}

func (c *Converter) convertLine(line string) string {
	words := strings.Fields(line)
	return c.convert(words)
}

func (c *Converter) convertFileLines(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("unable to read file: %v", err)
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		conv := c.convertLine(s)
		fmt.Println(conv)
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func isCaseSupported(caseType string) bool {
	for _, c := range cases {
		if c == caseType {
			return true
		}
	}
	return false
}

func convertCamel(words []string) string {
	out := strings.ToLower(words[0])
	for i := 1; i < len(words); i++ {
		runes := []rune(strings.ToLower(words[i]))
		runes[0] = unicode.ToUpper(runes[0])
		out += string(runes)
	}
	return out
}

func convertLower(words []string, sep string) string {
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return strings.Join(words, sep)
}

func convertUpper(words []string, sep string) string {
	for i, word := range words {
		words[i] = strings.ToUpper(word)
	}
	return strings.Join(words, sep)
}

func convertPascal(words []string, sep string) string {
	for i, word := range words {
		runes := []rune(strings.ToLower(word))
		runes[0] = unicode.ToUpper(runes[0])
		words[i] = string(runes)
	}
	return strings.Join(words, sep)
}
