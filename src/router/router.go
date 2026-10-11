package router

import (
	"test-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	testbe := r.Group("/test")

	testbe.POST("/create", controllers.TestCtrl.TestCreate)
	testbe.POST("/list", controllers.TestCtrl.TestList)
	testbe.POST("/update", controllers.TestCtrl.TestUpdate)
	testbe.POST("/delete", controllers.TestCtrl.TestDelete)
}
