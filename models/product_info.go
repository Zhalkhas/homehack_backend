package models

type ProductInfo struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	ImgUrl string  `json:"img_url"`
}
