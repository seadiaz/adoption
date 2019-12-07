package adapters

import "encoding"

// Persistence ...
type Persistence interface {
	Create(kind string, id string, obj encoding.BinaryMarshaler) error
	Update(kind string, id string, obj encoding.BinaryMarshaler) error
	Delete(kind string, id string) error
	GetAll(kind string, proto encoding.BinaryUnmarshaler) ([]interface{}, error)
	Find(kind string, id string, proto encoding.BinaryUnmarshaler) error
}
