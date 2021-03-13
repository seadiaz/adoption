package global

// ActionType ...
type ActionType string

// KindType ...
type KindType string

const (
	// Load ...
	Load ActionType = "load"

	// Display ...
	Display ActionType = "display"

	// People ...
	People KindType = "people"

	// Teams ...
	Teams KindType = "teams"

	// Memberships ...
	Memberships KindType = "memberships"

	// Adoptables ...
	Adoptables KindType = "adoptables"
)

// CommandHandler ...
type CommandHandler struct {
	BaseURL   string
	APIKey    string
	Executors []func(*CommandHandler, *CommandHandlerParams) error
}

// CommandHandlerParams ...
type CommandHandlerParams struct {
	Action   ActionType
	Kind     KindType
	Parent   string
	Filename string
}

// CreateCommandHandler ...
func CreateCommandHandler(baseURL, apiKey string) *CommandHandler {
	return &CommandHandler{
		baseURL,
		apiKey,
		make([]func(*CommandHandler, *CommandHandlerParams) error, 0),
	}
}

// AddExecutor ...
func (c *CommandHandler) AddExecutor(f func(*CommandHandler, *CommandHandlerParams) error) {
	c.Executors = append(c.Executors, f)
}

// Execute ...
func (c *CommandHandler) Execute(params *CommandHandlerParams) {
	for _, v := range c.Executors {
		v(c, params)
	}
}
