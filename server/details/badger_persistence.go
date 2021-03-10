package details

import (
	badger "github.com/dgraph-io/badger/v3"
	"github.com/golang/glog"
	"github.com/seadiaz/adoption/server/details/adapters"
)

// BadgerPersistence is a badger implementantion of persistence
type BadgerPersistence struct {
	db *badger.DB
}

// BuildBadgerPersistence ...
func BuildBadgerPersistence(path string) adapters.Persistence {
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		glog.Fatal(err)
	}

	// defer db.Close()
	return &BadgerPersistence{
		db: db,
	}
}

// Create ...
func (p *BadgerPersistence) Create(kind string, id string, obj adapters.PersistedData) error {
	err := p.db.Update(func(txn *badger.Txn) error {
		data, _ := obj.MarshalBinary()
		return txn.Set([]byte(kind+"-"+id), data)
	})

	return err
}

// Update ...
func (p *BadgerPersistence) Update(kind string, id string, obj adapters.PersistedData) error {
	err := p.db.Update(func(txn *badger.Txn) error {
		data, _ := obj.MarshalBinary()
		return txn.Set([]byte(kind+"-"+id), data)
	})

	return err
}

// Delete ...
func (p *BadgerPersistence) Delete(kind string, id string) error {
	err := p.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(kind + "-" + id))
	})

	return err
}

// GetAll ...
func (p *BadgerPersistence) GetAll(kind string, proto adapters.PersistedData) ([]interface{}, error) {
	list := make([]interface{}, 0, 0)
	p.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte(kind)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			entity := proto.Clone()
			list = append(list, entity)
			err := item.Value(func(v []byte) error {
				entity.UnmarshalBinary(v)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	return list, nil
}

// Find ...
func (p *BadgerPersistence) Find(kind string, id string, proto adapters.PersistedData) (interface{}, error) {
	entity := proto.Clone()
	err := p.db.View(func(txn *badger.Txn) error {
		item, _ := txn.Get([]byte(kind + "-" + id))
		err := item.Value(func(val []byte) error {
			entity.UnmarshalBinary(val)
			return nil
		})

		return err
	})

	return entity, err
}

// BadgerTransaction ...
type BadgerTransaction struct{}

// BeginTransaction ...
func (p *BadgerPersistence) BeginTransaction() adapters.Transaction {
	return &BadgerTransaction{}
}

// Commit ...
func (t *BadgerTransaction) Commit() error {
	return nil
}

// Rollback ...
func (t *BadgerTransaction) Rollback() error {
	return nil
}
