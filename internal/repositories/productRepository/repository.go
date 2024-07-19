package productrepository

import (
	"github.com/aurindo10/sol_store/internal/db/entities"
	"gorm.io/gorm"
)

type ProductMethods interface {
	CreateProduct(p *[]entities.Products) error
	GetAllProducts() (*[]entities.Products, error)
}
type ProductRepository struct {
	db *gorm.DB
}

func (c *ProductRepository) CreateProduct(p *[]entities.Products) error {
	result := c.db.Create(p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *ProductRepository) GetAllProducts() (*[]entities.Products, error) {
	var allProducts []entities.Products
	error := c.db.Find(&allProducts).Error
	if error != nil {
		return nil, error
	}
	return &allProducts, nil
}
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}
