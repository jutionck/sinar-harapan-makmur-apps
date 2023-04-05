package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/delivery/api/response"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/usecase"
	"net/http"
)

type BrandController struct {
	router  *gin.Engine
	useCase usecase.BrandUseCase
}

func (b *BrandController) createHandler(c *gin.Context) {
	var payload model.Brand
	if err := c.ShouldBind(&payload); err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := b.useCase.SaveData(&payload); err != nil {
		response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SendSingleResponse(c, payload, "OK")
}

func (b *BrandController) listHandler(c *gin.Context) {
	brands, err := b.useCase.FindAll()
	if err != nil {
		response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SendSingleResponse(c, brands, "OK")
}

func (b *BrandController) getHandler(c *gin.Context) {
	id := c.Param("id")
	brand, err := b.useCase.FindById(id)
	if err != nil {
		response.SendErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	response.SendSingleResponse(c, brand, "OK")
}

func (b *BrandController) searchHandler(c *gin.Context) {
	//name := c.Query("name")
	name := c.DefaultQuery("name", "Honda") // memberikan default query -> Honda (case sensitive)
	filter := map[string]interface{}{"name": name}
	brands, err := b.useCase.SearchBy(filter)
	if err != nil {
		response.SendErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	response.SendSingleResponse(c, brands, "OK")
}

func (b *BrandController) updateHandler(c *gin.Context) {
	var payload model.Brand
	if err := c.ShouldBind(&payload); err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := b.useCase.SaveData(&payload); err != nil {
		response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SendSingleResponse(c, payload, "OK")
}

func (b *BrandController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := b.useCase.DeleteData(id)
	if err != nil {
		response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusNoContent, "")
}

func NewBrandController(r *gin.Engine, useCase usecase.BrandUseCase) *BrandController {
	controller := &BrandController{
		router:  r,
		useCase: useCase,
	}
	r.GET("/brands", controller.listHandler)
	r.GET("/brands/:id", controller.getHandler)
	r.GET("/brands/search", controller.searchHandler)
	r.POST("/brands", controller.createHandler)
	r.PUT("/brands", controller.updateHandler)
	r.DELETE("/brands/:id", controller.deleteHandler)
	return controller
}
