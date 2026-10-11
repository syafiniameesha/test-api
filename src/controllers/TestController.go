package controllers

import (
	"net/http"
	"student-api/src/service"

	"github.com/gin-gonic/gin"
)

type TestController struct{}

func (u *TestController) TestList(context *gin.Context) {
	testid := context.PostForm("testid")

	r := service.ITestService.TestList(testid)
	context.JSON(http.StatusOK, r)
}
