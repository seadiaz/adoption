package adapters

import "encoding"

// Persistence ...
type Persistence interface {
	BeginTransaction() Transaction
	Create(kind string, id string, obj PersistedData) error
	Update(kind string, id string, obj PersistedData) error
	Delete(kind string, id string) error
	GetAll(kind string, proto PersistedData) ([]interface{}, error)
	Find(kind string, id string, proto PersistedData) (interface{}, error)
}

// Transaction ...
type Transaction interface {
	Commit() error
	Rollback() error
}

// PersistedData ...
type PersistedData interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
	Clone() PersistedData
}
