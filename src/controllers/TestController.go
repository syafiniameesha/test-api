package controllers

import (
	"net/http"

	"test-api/src/dto"
	"test-api/src/service"

	"github.com/gin-gonic/gin"
)

type TestController struct{}

// POST /test/create
func (u *TestController) TestCreate(c *gin.Context) {
	var input dto.TestCreateRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail("Invalid JSON request body"))
		return
	}

	result := service.ITestService.TestCreate(input)

	if result.Code != 0 {
		c.JSON(http.StatusBadRequest, result)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// POST /test/list
func (u *TestController) TestList(c *gin.Context) {
	var input dto.TestListRequest

	if c.Request.ContentLength != 0 {
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, dto.Fail("Invalid JSON request body"))
			return
		}
	}

	result := service.ITestService.TestList(input.TestId)
	c.JSON(http.StatusOK, result)
}

// POST /test/update
func (u *TestController) TestUpdate(c *gin.Context) {
	var input dto.TestUpdateRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail("Invalid JSON request body"))
		return
	}

	result := service.ITestService.TestUpdate(input)

	if result.Code != 0 {
		c.JSON(http.StatusBadRequest, result)
		return
	}

	c.JSON(http.StatusOK, result)
}

// POST /test/delete
func (u *TestController) TestDelete(c *gin.Context) {
	var input dto.TestListRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail("Invalid JSON request body"))
		return
	}

	result := service.ITestService.TestDelete(input.TestId)

	if result.Code != 0 {
		c.JSON(http.StatusBadRequest, result)
		return
	}

	c.JSON(http.StatusOK, result)
}
