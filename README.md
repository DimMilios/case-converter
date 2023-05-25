# Case Converter

Simple CLI tool for text conversion to various multiple-word identifier formats (cases).

## Usage

```sh
# Convert "Random word" to snake case
./case-converter -c snake Random word

# Prints
random_word

```

### Supported formats

Pass `-c`, `--case` argument to convert to one of the supported formats

| Format        | Value           |
| --------------|-----------------|
| camelCase     | camel (default) |
| lowercase     | lower           |
| UPPERCASE     | upper           |
| PascalCase    | pascal          |
| snake_case    | snake	          |
| kebab-case    | kebab	          |
| CONSTANT_CASE | const	          |
| Train-Case    | train	          |

**Note**: `camel` is the default case if one is not provided

### Pass a text file and convert case line by line

```sh
while read -r line; do ./case-converter --case snake $line; done < input.txt
```

Prints:

```
random_words
a_b_c_d_e_f
υποστηρίζει_ελληνικά;
many_words_test_weather_keyboard_mousepad_cup
words_containing_numerics123
numerics_123
lower_case_words
```

### Build the project

```sh
go build -o build/case-converter .
```
