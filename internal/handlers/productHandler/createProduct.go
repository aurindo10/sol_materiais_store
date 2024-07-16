package producthandler

import (
	"database/sql"
	"net/http"

	"github.com/aurindo10/sol_store/pkg/lib"
	"github.com/gin-gonic/gin"
)

func CreateProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, error := lib.GetUserFromContext(c)
		if error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Olá": "Mundo"})
	}
}
