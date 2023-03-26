package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, "welcome to go-gql study")
}
