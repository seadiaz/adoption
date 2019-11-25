package entities

// Label ...
type Label struct {
	Kind  string
	Value string
}

// CreateLabelWithKindAndValue ...
func CreateLabelWithKindAndValue(kind string, value string) *Label {
	return &Label{
		Kind:  kind,
		Value: value,
	}
}
