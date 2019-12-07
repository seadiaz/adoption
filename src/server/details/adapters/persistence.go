package adapters

import "encoding"

// Persistence ...
type Persistence interface {
	Create(kind string, id string, obj PersistedData) error
	Update(kind string, id string, obj PersistedData) error
	Delete(kind string, id string) error
	GetAll(kind string, proto PersistedData) ([]interface{}, error)
	Find(kind string, id string, proto PersistedData) error
}

// PersistedData ...
type PersistedData interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}
