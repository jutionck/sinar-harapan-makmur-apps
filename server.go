package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type appServer struct {
	// semua usecase
	userUC UserUseCase
	engine *gin.Engine
	host   string
}

func (a *appServer) initController() {
	// semua controller
	NewUserController(a.engine, a.userUC)
}

func (a *appServer) Run() {
	a.initController()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}

func Server() *appServer {
	r := gin.Default()
	r.Use(Logger())

	// semua kebutuhan service (repo, usecase)
	repo := NewUserRepository()
	userUC := NewUserUseCase(repo)
	host := fmt.Sprintf("%s:%s", "localhost", "8888")
	return &appServer{
		userUC: userUC,
		engine: r,
		host:   host,
	}
}
