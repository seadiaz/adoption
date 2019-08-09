package adapters

// Persistence ...
type Persistence interface {
	Create(id string, obj interface{}) error
	Delete(id string) error
	GetAll() []interface{}
}
