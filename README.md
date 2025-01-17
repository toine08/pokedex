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

## Assignment 1.5

Implementation of two commands:
- `help`: prints a help message describing how to use the REPL
- `exit`: exits the program

```go
func CommandHelp(commands map[string]CliCommand) error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Printf("Usage:\n\n")
    commandOrder := []string{"help", "exit"}

    for _, cmdName := range commandOrder {
        cmd := commands[cmdName]
        fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
    }
    return nil
}

func CommandExit() error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func main() {
    var commands map[string]utils.CliCommand
    callBackHelp := func() error {
        return utils.CommandHelp(commands)
    }

    commands = map[string]utils.CliCommand{
        "exit": {
            Name:        "exit",
            Description: "Exit the program",
            Callback:    utils.CommandExit,
        },
        "help": {
            Name:        "help",
            Description: "Displays a help message",
            Callback:    callBackHelp,
        },
    }
//..rest of the code...    
}
```

### Notes: 

Not everything was clear, but it's nice to practice Go a bit every day. I feel like I am making some progress, but I still have some issues with logic and resolving problems in the code. If you read this and have an idea about how to resolve these issues, do not hesitate to contact me on [bluesky](https://bsky.app/profile/togido.xyz)!


## Assignment 2.1

Add `map` and `mapb` commands to show the next 20 locations and the previous 20 locations, respectively.

```go
var index int32

func CommandMap() error {
    if index == 0 {
        index = 20
    } else {
        index += 20
    }
    url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?limit=%v", index)
    res, err := http.Get(url)
    if err != nil {
        return fmt.Errorf("sorry, couldn't resolve the URL: %v", err)
    }
    defer res.Body.Close()

    data, err := io.ReadAll(res.Body)
    if err != nil {
        return fmt.Errorf("couldn't get the data: %v", err)
    }
    var locationResponse LocationResponse

    if err := json.Unmarshal(data, &locationResponse); err != nil {
        return fmt.Errorf("sorry, couldn't resolve the data: %v", err)
    }
    for _, value := range locationResponse.Results {
        fmt.Printf("%s\n", value.Name)
    }

    return nil
}

func CommandMapB() error {
    if index < 20 {
        fmt.Println("You're on the first page")
    } else {
        index -= 40
        CommandMap()
    }
    return nil
}

// types.go

type Location struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}

type LocationResponse struct {
    Results []Location `json:"results"`
}

// I have also added the new commands in help.go and main.go. You can check this on GitHub!
```

### Notes:
It was nice to redo some HTTP client tasks in Golang. Thankfully, I have some cheatsheets from [boot.dev](boot.dev) about how to deal with data from requests. However, I also feel less confident because I had to use AI for some questions, but I mostly did it by myself. Also, because I didn't read the assignment before starting (that's my main issue, I would say haha), I didn't see that they gave some tips like using [JSON-to-GO](https://mholt.github.io/json-to-go/), which generates a type automatically from the JSON.


## Assignement 2.2

Add cache functionnalities

```go
type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	Mu    *sync.Mutex
	Cache map[string]CacheEntry
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{ // Assuming you have a map
		Mu:    &sync.Mutex{},               // Proper initialization of mutex as necessary
		Cache: make(map[string]CacheEntry), // Initialize the map
	}
	go cache.reapLoop(interval) // Start the reapLoop in a new goroutine
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	value := CacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
	c.Cache[key] = value
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	if cacheEntry, exists := c.Cache[key]; exists {
		return cacheEntry.Val, true
	} else {
		return nil, false
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.Mu.Lock()
		for key, entry := range c.Cache {
			if entry.CreatedAt.Add(interval).Before(time.Now()) {
				delete(c.Cache, key)
			}
		}
		c.Mu.Unlock()
	}
}


//func commandMap()

func CommandMap() error {
	if index == 0 {
		index = 20
	} else {
		index += 20
	}

	baseUrl := "https://pokeapi.co/api/v2/location-area"
	url := fmt.Sprintf("%s?limit=%v", baseUrl, index)

	if data, ok := cache.Get(url); ok {
		return processData(data)
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("sorry couldn't resolve the url: %v", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("couldn't get the data: %v", err)
	}
	cache.Add(url, data)

	return processData(data)
}
```


### Notes:

I encountered several challenges while working on this project. The source files were structured differently than I expected, which made it difficult to follow along. I had to refer to multiple sources and use AI to understand and implement certain parts of the code. This experience highlighted I need to improve and coding isn't easy. 

## Assignment 2.3

Add an explore command. It takes the name of a location area as an argument.

```go

func CommandExplore(zone string) error {
	baseUrl := "https://pokeapi.co/api/v2/location-area"
	url := fmt.Sprintf("%s/%s", baseUrl, zone)

	fmt.Printf("Exploring %s...\n", zone)

	if data, ok := cache.Get(url); ok {
		var locationArea LocationArea
		if err := json.Unmarshal(data, &locationArea); err != nil {
			return fmt.Errorf("sorry couldn't resolve the data: %v", err)
		}

		for _, value := range locationArea.PokemonEncounters {
			fmt.Printf("- %s\n", value.Pokemon.Name)
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("sorry couldn't resolve the url: %v", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("couldn't get the data: %v", err)
	}
	cache.Add(url, data)

	var locationArea LocationArea
	if err := json.Unmarshal(data, &locationArea); err != nil {
		return fmt.Errorf("sorry couldn't resolve the data: %v", err)
	}
	for _, value := range locationArea.PokemonEncounters {
		fmt.Printf("- %s\n", value.Pokemon.Name)
	}
	return nil
}

//I also had to rewrite main.go
func main() {

	var commands map[string]utils.CliCommand
	callBackHelp := func() error {
		return utils.CommandHelp(commands)
	}

	commands = map[string]utils.CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the program",
			Callback: func(args ...string) error {
				if len(args) > 0 {
					return fmt.Errorf("this command doesn't accept any arguments")
				}
				return utils.CommandExit()
			},
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback: func(args ...string) error {
				if len(args) > 0 {
					return fmt.Errorf("this command doesn't accept any arguments")
				}
				return callBackHelp()
			},
		},
		"map": {
			Name:        "map",
			Description: "Displays 20 names of locations",
			Callback: func(args ...string) error {
				if len(args) > 0 {
					return fmt.Errorf("this command doesn't accept any arguments")
				}
				return utils.CommandMap()
			},
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the 20 previous names of locations",
			Callback: func(args ...string) error {
				if len(args) > 0 {
					return fmt.Errorf("this command doesn't accept any arguments")
				}
				return utils.CommandMapB()
			},
		},
		"explore": {
			Name:        "explore",
			Description: "Explore an area",
			Callback: func(args ...string) error {
				if len(args) < 1 {
					return fmt.Errorf("please provide a zone to explore")
				}
				return utils.CommandExplore(args[0])
			},
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		// Split input into command and arguments
		inputParts := strings.Fields(input)
		if len(inputParts) == 0 {
			continue // Skip empty input
		}

		command := inputParts[0] // First part is the command
		args := inputParts[1:]   // Remaining parts are arguments (if any)

		if cmd, exists := commands[command]; exists {
			if err := cmd.Callback(args...); err != nil { // Pass args to the callback
				fmt.Println("Error: ", err)
			}
		} else {
			fmt.Println("Unknown command: ", command)
		}
	}
}

```

### Notes:

I was not feeling well, which made this exercise more challenging. I felt lost at times and realized it might not have been the best idea to tackle this while being sick. I hope to improve my skills over time. After completing this project, I plan to work on building projects without strict guidelines to enhance my learning experience.
