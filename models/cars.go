package models


type Cars struct {
	Id string `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year int `json:"year"`
	Price int `json:"price"`
	IsNew bool `json:"is_new"`
}