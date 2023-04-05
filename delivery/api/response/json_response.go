package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendSingleResponse(c *gin.Context, data interface{}, description string) {
	c.JSON(http.StatusOK, &SingleResponse{
		Status: Status{
			Code:        http.StatusOK,
			Description: description,
		},
		Data: data,
	})
}

func SendErrorResponse(c *gin.Context, code int, description string) {
	c.AbortWithStatusJSON(code, &Status{
		Code:        code,
		Description: description,
	})
}
