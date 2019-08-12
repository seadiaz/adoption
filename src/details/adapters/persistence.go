package adapters

// Persistence ...
type Persistence interface {
	Create(id string, obj interface{}) error
	Update(id string, obj interface{}) error
	Delete(id string) error
	GetAll() []interface{}
	Find(id string) interface{}
}
