package cli

// Client ...
type Client struct {
	URL      string
	Filename string
	APIKey   string
}

// CreateClient ...
func CreateClient(url string, filename string, apiKey string) *Client {
	return &Client{
		URL:      url,
		Filename: filename,
		APIKey:   apiKey,
	}
}
