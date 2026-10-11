package dto

type TestDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TestCreateRequest struct {
	Name string `json:"name"`
}

type TestListRequest struct {
	TestId string `json:"testid"`
}

type TestUpdateRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
