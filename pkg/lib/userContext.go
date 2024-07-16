package lib

import (
	"fmt"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/gin-gonic/gin"
)

func GetUserFromContext(c *gin.Context) (*clerk.User, error) {
	userInterface, exists := c.Get("user")
	if !exists {
		return nil, fmt.Errorf("User not found in context")
	}
	usr, ok := userInterface.(*clerk.User)
	if !ok {
		return nil, fmt.Errorf("Failed to cast user")
	}
	fmt.Fprintf(c.Writer, `{"user_id": "%s", "user_banned": "%t"}`, usr.ID, usr.Banned)
	return usr, nil
}
