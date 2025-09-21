package product

//go:generate mockery --name=Repository --structname=ProductRepository --output=../../../test/mocks --filename=product_repository.go
type Repository interface {
	Create(args *Product) error
	List() ([]Product, error)
	GetByID(id int64) (*Product, error)
	Update(args *Product) error
	Delete(id int64) error

	// TODO: need to implement
	// GetProductsByCategoryID(categoryID int64) ([]*Product, error)
}
