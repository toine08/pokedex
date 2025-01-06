# Pokedex

This is a project in Golang from the course on [boot.dev](boot.dev).

As I did for the HTTP-server course, I will do the same and document everything I do, what I have to do, my code, and some notes about how I feel and things like that.
You can check the [HTTP-server](https://github.com/toine08/http-server) to see an example.

***

## Assignment 1.3

Create a new `cleanInput(text string) []string` function. For now, it should just return an empty slice of strings. And create some tests.

```go
func CleanInput(text string) []string {
	var sliceText []string
	textStrings := strings.Fields(strings.ToLower(text))
	sliceText = append(sliceText, textStrings...)
	return sliceText
}

//cleanInput_test.go:
unc TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    " I am toto",
			expected: []string{"i", "am", "toto"}, // note case sensitivity
		},
		{
			input:    " I am toto and toine",
			expected: []string{"i", "am", "toto", "and", "toine"}, // note case sensitivity
		},
	}
	for _, c := range cases {
		actual := utils.CleanInput(c.input) // note the capital letter for exported function
		if len(actual) != len(c.expected) {
			t.Errorf("Length of slices are not matching")
			continue
		}
		for i, word := range actual {
			if word != c.expected[i] {
				t.Errorf("Mismatch at index %d: got %s, want %s", i, word, c.expected[i])
			}
		}
	}
}

```

### Notes:
Easy, didn't understand everything about the test. When I started learning Go a few years ago, it was with [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests), and the tests were done differently, so I was confused. I also had some issues with the function, but yeah, it's just some logic I missed.


## Assignment 1.4

Create support for a simple REPL

```go
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			slicedInput := utils.CleanInput(input)
			fmt.Printf("Your command was: %s\n", slicedInput[0])
		}
	}
}
```

### Notes:

This one was a bit challenging because I wasn't familiar with the `bufio` package, and the documentation wasn't very clear to me. However, with the help of the chatbox, I was able to understand it better. Once I grasped how to use it, the logic was straightforward to implement.


