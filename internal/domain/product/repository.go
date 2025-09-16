package product

//go:generate mockery --name=ProductRepository --output=./mocks
type Repository interface {
	Create(args *Product) error
	List() ([]Product, error)

	// TODO: need to implement
	// GetProductByID(id int64) (*model.Product, error)
	// GetProductsByCategoryID(categoryID int64) ([]*Product, error)
	// GetProductsByIDs(ids []int64) ([]*Product, error)
	// UpdateProduct(args *Product) error
	// DeleteProduct(id int64) error
}
