package main

import (
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

type Converter struct {
	caseType string
}

func NewConverter(caseType string) *Converter {
	return &Converter{
		caseType: caseType,
	}
}

func (c *Converter) convert(words *[]string) string {
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

func convertCamel(words *[]string) string {
	out := (*words)[0]
	for i := 1; i < len((*words)); i++ {
		runes := []rune((*words)[i])
		runes[0] = unicode.ToUpper(runes[0])
		out += string(runes)
	}
	return out
}

func convertLower(words *[]string, sep string) string {
	for i, word := range *words {
		(*words)[i] = strings.ToLower(word)
	}
	return strings.Join(*words, sep)
}

func convertUpper(words *[]string, sep string) string {
	for i, word := range *words {
		(*words)[i] = strings.ToUpper(word)
	}
	return strings.Join(*words, sep)
}

func convertPascal(words *[]string, sep string) string {
	for i, word := range *words {
		runes := []rune(word)
		runes[0] = unicode.ToUpper(runes[0])
		(*words)[i] = string(runes)
	}
	return strings.Join(*words, sep)
}
