package entities

// StrategyType ...
type StrategyType string

const (
	// StrategyTypeSingleMember ...
	StrategyTypeSingleMember StrategyType = "single-member"
	// StrategyTypeAllMember ...
	StrategyTypeAllMember StrategyType = "all-member"
)

// Adoptable ...
type Adoptable struct {
	ID       *ID
	Name     string
	Labels   []*Label
	Strategy StrategyType
}

// CreateAdoptableWithNameAndStrategy ...
func CreateAdoptableWithNameAndStrategy(name string, strategy StrategyType) *Adoptable {
	return &Adoptable{
		Name:     name,
		ID:       generateID(),
		Strategy: strategy,
	}
}

// AddLabel ...
func (a *Adoptable) AddLabel(label *Label) {
	for _, item := range a.Labels {
		if item.Kind == label.Kind {
			item.Value = label.Value
			return
		}
	}

	a.Labels = append(a.Labels, label)
}

// RemoveLabel ...
func (a *Adoptable) RemoveLabel(label *Label) {
	for i, item := range a.Labels {
		if item.Kind == label.Kind {
			a.Labels = append(a.Labels[:i], a.Labels[i+1:]...)
			return
		}
	}
}

// HasLabelKindEqualToValue ...
func (a *Adoptable) HasLabelKindEqualToValue(kind, value string) bool {
	for _, item := range a.Labels {
		if item.Kind == kind && item.Value == value {
			return true
		}
	}

	return false
}
