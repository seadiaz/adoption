package entities

// Tool ...
type Tool struct {
	ID   string
	Name string
}

// BuildToolWithName ...
func BuildToolWithName(name string) *Tool {
	return &Tool{Name: name}
}
