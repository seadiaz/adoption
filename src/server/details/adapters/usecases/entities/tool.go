package entities

// Tool ...
type Tool struct {
	ID     *ID
	Name   string
	Labels []*Label
}

// CreateToolWithName ...
func CreateToolWithName(name string) *Tool {
	return &Tool{
		Name: name,
		ID:   newID(),
	}
}

// AddLabel ...
func (t *Tool) AddLabel(label *Label) {
	for _, item := range t.Labels {
		if item.Kind == label.Kind {
			item.Value = label.Value
			return
		}
	}

	t.Labels = append(t.Labels, label)
}

// RemoveLabel ...
func (t *Tool) RemoveLabel(label *Label) {
	for i, item := range t.Labels {
		if item.Kind == label.Kind {
			t.Labels = append(t.Labels[:i], t.Labels[i+1:]...)
			return
		}
	}
}
