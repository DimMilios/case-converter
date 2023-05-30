package main

import (
	"strings"
	"testing"
)

func TestConvertCamel(t *testing.T) {
	c, _ := NewConverter(CamelCase)

	var testCases = []struct {
		input string
		want  string
	}{
		{input: "Random words", want: "randomWords"},
		{input: "Υποστηρίζει ελληνικά;", want: "υποστηρίζειΕλληνικά;"},
		{input: "MANY WORDS TEST WEATHER KEYBOARD MOUSEPAD CUP", want: "manyWordsTestWeatherKeyboardMousepadCup"},
		{input: "Words containing numerics123", want: "wordsContainingNumerics123"},
		{input: "Numerics 123", want: "numerics123"},
		{input: "lower case words", want: "lowerCaseWords"},
	}

	for _, v := range testCases {
		converted := c.convert(strings.Split(v.input, " "))
		if v.want != converted {
			t.Errorf("got \"%s\", want \"%s\" for input \"%s\"", converted, v.want, v.input)
		}
	}
}

func TestConvertPascal(t *testing.T) {
	c, _ := NewConverter(PascalCase)

	var testCases = []struct {
		input string
		want  string
	}{
		{input: "Random words", want: "RandomWords"},
		{input: "Υποστηρίζει ελληνικά;", want: "ΥποστηρίζειΕλληνικά;"},
		{input: "MANY WORDS TEST WEATHER KEYBOARD MOUSEPAD CUP", want: "ManyWordsTestWeatherKeyboardMousepadCup"},
		{input: "Words containing numerics123", want: "WordsContainingNumerics123"},
		{input: "Numerics 123", want: "Numerics123"},
		{input: "lower case words", want: "LowerCaseWords"},
	}

	for _, v := range testCases {
		converted := c.convert(strings.Split(v.input, " "))
		if v.want != converted {
			t.Errorf("got \"%s\", want \"%s\" for input \"%s\"", converted, v.want, v.input)
		}
	}
}

func TestConvertSnake(t *testing.T) {
	c, _ := NewConverter(SnakeCase)

	var testCases = []struct {
		input string
		want  string
	}{
		{input: "Random words", want: "random_words"},
		{input: "Υποστηρίζει ελληνικά;", want: "υποστηρίζει_ελληνικά;"},
		{input: "MANY WORDS TEST WEATHER KEYBOARD MOUSEPAD CUP", want: "many_words_test_weather_keyboard_mousepad_cup"},
		{input: "Words containing numerics123", want: "words_containing_numerics123"},
		{input: "Numerics 123", want: "numerics_123"},
		{input: "lower case words", want: "lower_case_words"},
	}

	for _, v := range testCases {
		converted := c.convert(strings.Split(v.input, " "))
		if v.want != converted {
			t.Errorf("got \"%s\", want \"%s\" for input \"%s\"", converted, v.want, v.input)
		}
	}
}

func TestConvertKebab(t *testing.T) {
	c, _ := NewConverter(KebabCase)

	var testCases = []struct {
		input string
		want  string
	}{
		{input: "Random words", want: "random-words"},
		{input: "Υποστηρίζει ελληνικά;", want: "υποστηρίζει-ελληνικά;"},
		{input: "MANY WORDS TEST WEATHER KEYBOARD MOUSEPAD CUP", want: "many-words-test-weather-keyboard-mousepad-cup"},
		{input: "Words containing numerics123", want: "words-containing-numerics123"},
		{input: "Numerics 123", want: "numerics-123"},
		{input: "lower case words", want: "lower-case-words"},
	}

	for _, v := range testCases {
		converted := c.convert(strings.Split(v.input, " "))
		if v.want != converted {
			t.Errorf("got \"%s\", want \"%s\" for input \"%s\"", converted, v.want, v.input)
		}
	}
}

func TestConvertConst(t *testing.T) {
	c, _ := NewConverter(ConstantCase)

	var testCases = []struct {
		input string
		want  string
	}{
		{input: "Random words", want: "RANDOM_WORDS"},
		{input: "A b c d e f", want: "A_B_C_D_E_F"},
		{input: "Υποστηρίζει ελληνικά;", want: "ΥΠΟΣΤΗΡΊΖΕΙ_ΕΛΛΗΝΙΚΆ;"},
		{input: "MANY WORDS TEST WEATHER KEYBOARD MOUSEPAD CUP", want: "MANY_WORDS_TEST_WEATHER_KEYBOARD_MOUSEPAD_CUP"},
		{input: "Words containing numerics123", want: "WORDS_CONTAINING_NUMERICS123"},
		{input: "Numerics 123", want: "NUMERICS_123"},
		{input: "lower case words", want: "LOWER_CASE_WORDS"},
	}

	for _, v := range testCases {
		converted := c.convert(strings.Split(v.input, " "))
		if v.want != converted {
			t.Errorf("got \"%s\", want \"%s\" for input \"%s\"", converted, v.want, v.input)
		}
	}
}

func TestConvertTrain(t *testing.T) {
	c, _ := NewConverter(TrainCase)

	var testCases = []struct {
		input string
		want  string
	}{
		{input: "Random words", want: "Random-Words"},
		{input: "Υποστηρίζει ελληνικά;", want: "Υποστηρίζει-Ελληνικά;"},
		{input: "MANY WORDS TEST WEATHER KEYBOARD MOUSEPAD CUP", want: "Many-Words-Test-Weather-Keyboard-Mousepad-Cup"},
		{input: "Words containing numerics123", want: "Words-Containing-Numerics123"},
		{input: "Numerics 123", want: "Numerics-123"},
		{input: "lower case words", want: "Lower-Case-Words"},
	}

	for _, v := range testCases {
		converted := c.convert(strings.Split(v.input, " "))
		if v.want != converted {
			t.Errorf("got \"%s\", want \"%s\" for input \"%s\"", converted, v.want, v.input)
		}
	}
}
