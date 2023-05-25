# Case Converter

Pass a text file and convert case line by line

```sh
while read -r line; do case-converter --case snake $line; done < input.txt
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
