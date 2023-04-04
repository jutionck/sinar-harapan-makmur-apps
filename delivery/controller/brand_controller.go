package controller

import (
	"github.com/gin-gonic/gin"
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
		// AbortWithStatusJSON -> untuk menghentikan proses request dan otomatis akan keluar dari function
		// jadi tidak memerlukan return
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := b.useCase.SaveData(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, payload)
}

func (b *BrandController) listHandler(c *gin.Context) {
	brands, err := b.useCase.FindAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":        200,
		"description": "Ok",
		"data":        brands,
	})
}

func (b *BrandController) getHandler(c *gin.Context) {
	id := c.Param("id")
	brand, err := b.useCase.FindById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":        200,
		"description": "Ok",
		"data":        brand,
	})
}

func (b *BrandController) updateHandler(c *gin.Context) {
	var payload model.Brand
	if err := c.ShouldBind(&payload); err != nil {
		// AbortWithStatusJSON -> untuk menghentikan proses request dan otomatis akan keluar dari function
		// jadi tidak memerlukan return
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := b.useCase.SaveData(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, payload)
}

func (b *BrandController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := b.useCase.DeleteData(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	r.POST("/brands", controller.createHandler)
	r.PUT("/brands", controller.updateHandler)
	r.DELETE("/brands/:id", controller.deleteHandler)
	return controller
}
