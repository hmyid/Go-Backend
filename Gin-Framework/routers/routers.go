package routers

import (
	"github.com/gin-gonic/gin"
	"Gin-Framework/controllers"
)

func RegistRoutesServer() *gin.Engine {
	router := gin.Default()

	router.GET("/getproduct", controllers.GetProduct)
	router.POST("/postproduct", controllers.CreateProduct)
	router.PUT("/updateproduct/:prodID", controllers.UpdateProduct)
	router.DELETE("/deleteproduct/:prodID", controllers.DeleteProduct)


	return router
}