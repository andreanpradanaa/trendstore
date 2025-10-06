package category

//go:generate mockery --name=Repository --structname=ProductRepository --output=../../../test/mocks --filename=product_repository.go
type Repository interface {
	Create(category *Category) error

	// TODO: add this feature
	// GetByID(id int64) (*Category, error)
	// List() ([]Category, error)
	// Update(category *Category) error
	// Delete(id int64) error
}
