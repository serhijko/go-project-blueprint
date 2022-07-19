package apis

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/serhijko/go-project-blueprint/cmd/blueprint/daos"
	"github.com/serhijko/go-project-blueprint/cmd/blueprint/services"
)

// GetUser godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	s := services.NewUserService(daos.NewUserDAO())   // Create service
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32) // Parse ID from URL
	if user, err := s.Get(uint(id)); err != nil {     // Try to get user from database
		c.AbortWithStatus(http.StatusNotFound) // Abort if not found
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user) // Send back data
	}
}
