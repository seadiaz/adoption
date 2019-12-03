package adapters

// Persistence ...
type Persistence interface {
	Create(kind string, id string, obj interface{}) error
	Update(kind string, id string, obj interface{}) error
	Delete(kind string, id string) error
	GetAll(kind string) ([]interface{}, error)
	Find(kind string, id string) (interface{}, error)
}
