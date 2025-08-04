package routes

import (
	"github.com/gin-gonic/gin"
	"microservice/internal/adapter/locale"
	"microservice/internal/driver/delivery"
)

func TodoRoutes(r *gin.RouterGroup, h delivery.ITodoHandler, l locale.ILocale) {
	todo := r.Group("/todo") //.Use(middlewares.CheckAuth(l))
	todo.POST("/create", h.Create)
	todo.GET("/:uuid", h.GetDetails)
	todo.GET("/list", h.GetList)
}
