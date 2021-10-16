package parsers

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/Zhalkhas/homehack_backend/models"
	"github.com/Zhalkhas/homehack_backend/utils"
	"net/http"
	"regexp"
	"strconv"
)

var sulpakRegex = regexp.MustCompile("https://www\\.sulpak\\.kz/g/[0-9]+")

type SulpakParser struct {
}

func (s *SulpakParser) Detect(qrCode string) bool {
	return sulpakRegex.MatchString(qrCode)
}

func (s *SulpakParser) Parse(qrCode string) (*models.ProductInfo, error) {
	resp, err := http.Get(qrCode)
	if err != nil {
		return nil, err
	}
	reader, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	title := reader.Find(".product-container-title").First().Text()
	priceStr := utils.ExtractPrice(reader.Find(".sum").First().Text())
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return nil, err
	}
	urlSelector := reader.Find(".product-photo")
	for _, node := range urlSelector.Nodes {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				return &models.ProductInfo{Name: title, Price: price, ImgUrl: attr.Val}, nil
			}
		}
	}
	return &models.ProductInfo{Name: title, Price: price}, nil
}
