package router

import (
	"student-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	studentbe := r.Group("student/be")

	// user
	studentbe.POST("/test/list", controllers.TestCtrl.TestList)

}
