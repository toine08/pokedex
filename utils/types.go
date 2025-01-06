package utils

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationResponse struct {
	Results []Location `json:"results"`
}
