package controllers

import (
	"net/http"

	"test-api/src/dto"
	"test-api/src/service"

	"github.com/gin-gonic/gin"
)

type TestController struct{}

// POST /test/create
func (u *TestController) TestCreate(context *gin.Context) {
	var input dto.TestCreateRequest

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, dto.Fail("Invalid JSON request body"))
		return
	}

	result := service.ITestService.TestCreate(input)

	if result.Code != 0 {
		context.JSON(http.StatusBadRequest, result)
		return
	}

	context.JSON(http.StatusCreated, result)
}

// POST /test/list
func (u *TestController) TestList(context *gin.Context) {
	var input dto.TestListRequest

	// Allow an empty body to list all tests.
	// If a body is provided, parse its JSON.
	if context.Request.ContentLength != 0 {
		if err := context.ShouldBindJSON(&input); err != nil {
			context.JSON(http.StatusBadRequest, dto.Fail("Invalid JSON request body"))
			return
		}
	}

	result := service.ITestService.TestList(input.TestId)

	context.JSON(http.StatusOK, result)
}
