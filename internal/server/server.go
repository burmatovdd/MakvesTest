package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"makves_app/internal/server/methods"
)

type Service struct {
	service *Handler
}

type Handler interface {
	HandleFunc()
}

func (service *Service) HandleFunc() {
	method := methods.Service{}
	router := gin.Default()
	router.GET("get-items", method.GetItems)
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("err: ", err)
	}
}
