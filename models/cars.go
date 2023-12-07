package models


type Cars struct {
	Brand string `json:"brand"`
	CarNo  string `json:"carno"`
	Model string `json:"model"`
	Year int `json:"year"`
	Price int `json:"price"`
	IsNew bool `json:"is_new"`
}