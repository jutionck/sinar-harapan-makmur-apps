package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/delivery/api/response"
)

type BaseApi struct{}

func (b *BaseApi) ParseRequestBody(c *gin.Context, payload interface{}) error {
	if err := c.ShouldBindJSON(payload); err != nil {
		return err
	}
	return nil
}

func (b *BaseApi) NewSuccessSingleResponse(c *gin.Context, data interface{}, description string) {
	response.SendSingleResponse(c, data, description)
}

func (b *BaseApi) NewFailedResponse(c *gin.Context, code int, description string) {
	response.SendErrorResponse(c, code, description)
}
