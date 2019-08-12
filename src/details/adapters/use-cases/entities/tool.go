package entities

// Tool ...
type Tool struct {
	ID   string
	Name string
}

// CreateToolWithName ...
func CreateToolWithName(name string) *Tool {
	return &Tool{Name: name}
}
