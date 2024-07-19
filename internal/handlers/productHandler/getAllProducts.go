package producthandler

import (
	"net/http"

	productrepository "github.com/aurindo10/sol_store/internal/repositories/productRepository"
	"github.com/gin-gonic/gin"
)

func GetAllProducts(db productrepository.ProductMethods) gin.HandlerFunc {
	return func(c *gin.Context) {
		allproducts, error := db.GetAllProducts()
		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		}
		c.JSON(http.StatusOK, allproducts)
		return
	}
}
