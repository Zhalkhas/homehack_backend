package parsers

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/Zhalkhas/homehack_backend/models"
	"github.com/Zhalkhas/homehack_backend/utils"
	"net/http"
	"regexp"
	"strconv"
)

var whitewindRegex = regexp.MustCompile("https://shop\\.kz/bitrix/tools/track_qr\\.php\\?art=([0-9]+)")

const urlFmt = "https://shop.kz/search/?q=%s&s="

type WhitewindParser struct {
}

func (w *WhitewindParser) Detect(qrCode string) bool {
	return whitewindRegex.MatchString(qrCode)
}

func (w *WhitewindParser) Parse(qrCode string) (*models.ProductInfo, error) {
	matches := whitewindRegex.FindStringSubmatch(qrCode)
	if len(matches) < 2 {
		fmt.Println(matches)
		return nil, fmt.Errorf("could not detect SKU")
	}
	sku := matches[1]
	url := fmt.Sprintf(urlFmt, sku)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	reader, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	priceStr := reader.Find(".item_current_price").First().Text()
	price, err := strconv.ParseFloat(utils.ExtractPrice(priceStr), 64)
	if err != nil {
		return nil, err
	}
	title := reader.Find("#pagetitle").First().Text()
	urlSelector := reader.Find(".afkl-lazy-wrapper").First().Children().First()
	for _, node := range urlSelector.Nodes {
		for _, attr := range node.Attr {
			if attr.Key == "src" {
				return &models.ProductInfo{Price: price, Name: title, ImgUrl: "https://shop.kz/" + attr.Val}, nil
			}

		}
	}
	return &models.ProductInfo{Price: price, Name: title}, nil
}
