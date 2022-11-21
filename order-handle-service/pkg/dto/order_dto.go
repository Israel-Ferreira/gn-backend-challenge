package dto

type QuestionDTO map[string]string

type AddressDTO struct {
	City         string `json:"city"`
	Street       string `json:"street"`
	ZipCode      string `json:"zip_code"`
	Uf           string `json:"uf"`
	Neighborhood string `json:"neighborhood"`
}

type UserInfo struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type OrderDTO struct {
	UserInfoAttr UserInfo    `json:"user_info"`
	AddressAttr  AddressDTO  `json:"address_attributes"`
	RequestInfo  QuestionDTO `json:"request_info"`
}
