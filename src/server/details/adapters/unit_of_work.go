package adapters

// UnitOfWorkFactory ...
type UnitOfWorkFactory struct {
	adoptableRepository AdoptableRepository
}

// CreateUnitOfWorkFactory ...
func CreateUnitOfWorkFactory(aRepository AdoptableRepository) *UnitOfWorkFactory {
	return &UnitOfWorkFactory{aRepository}
}

// Create ...
func (f *UnitOfWorkFactory) Create() *UnitOfWork {
	return &UnitOfWork{f.adoptableRepository}
}

// UnitOfWork ...
type UnitOfWork struct {
	adoptableRepository AdoptableRepository
}

// Commit ...
func (u *UnitOfWork) Commit() error {
	return nil
}
