package lib

import (
	"net/http"
	"strings"

	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/gin-gonic/gin"
)

func ProtectedMiddlware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionToken := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		claims, err := jwt.Verify(c.Request.Context(), &jwt.VerifyParams{
			Token: sessionToken,
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"access": "unauthorized"})
			c.Abort()
			return
		}
		usr, err := user.Get(c.Request.Context(), claims.Subject)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"access": "some internal error with decoding the token"})
			c.Abort()
			return
		}
		c.Set("user", usr)
		c.Next()
	}
}
