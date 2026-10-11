package service

import (
	"fmt"
	"strings"
	"sync"

	"test-api/src/dto"
)

type ITestServiceInterface interface {
	TestCreate(input dto.TestCreateRequest) dto.ResponseDto
	TestList(testid string) dto.ResponseDto
	TestUpdate(input dto.TestUpdateRequest) dto.ResponseDto
	TestDelete(testid string) dto.ResponseDto
}

type TestServiceImpl struct{}

var (
	testMu sync.RWMutex

	tests = []dto.TestDto{
		{Id: "TEST-001", Name: "Dummy Test One"},
		{Id: "TEST-002", Name: "Dummy Test Two"},
	}

	nextTestID = 3
)

// Create a test
func (u TestServiceImpl) TestCreate(
	input dto.TestCreateRequest,
) dto.ResponseDto {
	name := strings.TrimSpace(input.Name)

	if name == "" {
		return dto.Fail("Test name is required")
	}

	testMu.Lock()
	defer testMu.Unlock()

	test := dto.TestDto{
		Id:   fmt.Sprintf("TEST-%03d", nextTestID),
		Name: name,
	}

	nextTestID++
	tests = append(tests, test)

	return dto.SuccessMessage("Test created successfully", test)
}

// List all tests or find a test by ID
func (u TestServiceImpl) TestList(testid string) dto.ResponseDto {
	testMu.RLock()
	defer testMu.RUnlock()

	if testid != "" {
		for _, test := range tests {
			if test.Id == testid {
				return dto.SuccessCount(test, 1)
			}
		}

		return dto.SuccessCount([]dto.TestDto{}, 0)
	}

	result := append([]dto.TestDto{}, tests...)
	return dto.SuccessCount(result, int64(len(result)))
}

// Update a test by ID
func (u TestServiceImpl) TestUpdate(
	input dto.TestUpdateRequest,
) dto.ResponseDto {
	id := strings.TrimSpace(input.Id)
	name := strings.TrimSpace(input.Name)

	if id == "" {
		return dto.Fail("Test ID is required")
	}

	if name == "" {
		return dto.Fail("Test name is required")
	}

	testMu.Lock()
	defer testMu.Unlock()

	for i := range tests {
		if tests[i].Id == id {
			tests[i].Name = name

			return dto.SuccessMessage(
				"Test updated successfully",
				tests[i],
			)
		}
	}

	return dto.Fail("Test not found")
}

// Delete a test by ID
func (u TestServiceImpl) TestDelete(testid string) dto.ResponseDto {
	id := strings.TrimSpace(testid)

	if id == "" {
		return dto.Fail("Test ID is required")
	}

	testMu.Lock()
	defer testMu.Unlock()

	for i := range tests {
		if tests[i].Id == id {
			// Remove the test at index i
			tests = append(tests[:i], tests[i+1:]...)

			return dto.SuccessMessage(
				"Test deleted successfully",
				dto.NullDto{},
			)
		}
	}

	return dto.Fail("Test not found")
}
