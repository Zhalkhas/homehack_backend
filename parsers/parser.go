package parsers

import "github.com/Zhalkhas/homehack_backend/models"

type Parser interface {
	Detect(qrCode string) bool
	Parse(qrCode string) (*models.ProductInfo, error)
}
