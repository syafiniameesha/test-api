package service

import "student-api/src/dto"

type TestServiceImpl struct{}

func (u TestServiceImpl) TestList(testid string) dto.ResponseDto {
	tests := []dto.TestDto{
		{
			Id:   "TEST-001",
			Name: "Dummy Test One",
		},
		{
			Id:   "TEST-002",
			Name: "Dummy Test Two",
		},
	}

	if testid != "" {
		for _, test := range tests {
			if test.Id == testid {
				return dto.Success(test)
			}
		}

		return dto.Success([]dto.TestDto{})
	}

	return dto.Success(tests)
}
