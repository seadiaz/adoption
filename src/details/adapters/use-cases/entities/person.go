package entities

// Person ...
type Person struct {
	Name  string
	Tools []Tool
}

// BuildPerson ...
func BuildPerson(name string) *Person {
	return &Person{
		Name:  name,
		Tools: make([]Tool, 0),
	}
}

// AdoptTool ...
func (p *Person) AdoptTool(tool Tool) error {
	p.Tools = append(p.Tools, tool)
	return nil
}
