package producthandler

import (
	"net/http"

	"github.com/aurindo10/sol_store/internal/db/entities"
	productrepository "github.com/aurindo10/sol_store/internal/repositories/productRepository"
	"github.com/aurindo10/sol_store/pkg/lib"
	"github.com/gin-gonic/gin"
)

func CreateProduct(db productrepository.ProductMethods) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, error := lib.GetUserFromContext(c)
		if error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
			return
		}
		var products []entities.Products
		if error := c.ShouldBindJSON(&products); error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		}
		if error := db.CreateProduct(&products); error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		}
		c.JSON(http.StatusCreated, gin.H{"Sucess": "Produto criado com sucesso"})
		return
	}
}
