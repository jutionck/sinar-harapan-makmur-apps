package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	router      *gin.Engine // method + handler
	userUseCase UserUseCase // kebutuhan service
}

func (u *UserController) createHandler(c *gin.Context) {
	// define model (struct)
	var user UserCredential
	// cek error kalo terjadi kesalahan saat bind (400) -> Client
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// cek error ketika menyimpan terjadi kegagalan -> 500 (server)
	if err := u.userUseCase.RegisterNewUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Balikan ini ketika tidak terjadi error
	c.JSON(http.StatusCreated, gin.H{
		"code":        201,
		"description": "OK",
		"data":        user.Username,
	})
}

func (u *UserController) listHandler(c *gin.Context) {
	users, err := u.userUseCase.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Balikan ini ketika tidak terjadi error
	c.JSON(http.StatusOK, gin.H{
		"code":        200,
		"description": "OK",
		"data":        users,
	})
}

func NewUserController(router *gin.Engine, userUseCase UserUseCase) *UserController {
	uc := &UserController{
		router:      router,
		userUseCase: userUseCase,
	}
	// ENDPOINT -> /api/v1/users
	router.POST("/users", uc.createHandler)
	router.GET("/users", uc.listHandler)
	return uc
}
