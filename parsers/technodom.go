package parsers

import (
	"encoding/json"
	"fmt"
	"github.com/Zhalkhas/homehack_backend/models"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const apiUrlFmt = "https://api.technodom.kz/katalog/api/v1/products/%s/show?city_id=5f5f1e3b4c8a49e692fefd70"

var technodomRegex = regexp.MustCompile("[0-9]+,\\b[0-9a-f]{8}\\b-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-\\b[0-9a-f]{12}\\b")

type TechnodomParser struct {
}

func (t *TechnodomParser) Detect(qrCode string) bool {
	return technodomRegex.MatchString(qrCode)
}

type TechnodomApiResp struct {
	Price string `json:"price"`
	Title string `json:"title"`
	Uri   string `json:"uri"`
}

func (t *TechnodomParser) Parse(qrCode string) (*models.ProductInfo, error) {
	qrSplitted := strings.Split(qrCode, ",")
	if len(qrSplitted) != 2 {
		return nil, fmt.Errorf("invalid qr format")
	}
	fmt.Println("techno parse started")
	sku := qrSplitted[0]
	storeID := qrSplitted[1]
	//storeID = strings.ReplaceAll(storeID, "-", "")
	fmt.Println("sku", sku)
	fmt.Println("storeID", storeID)

	resp, err := http.Get(fmt.Sprintf(apiUrlFmt, sku))
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(resp.Body)
	apiResp := &TechnodomApiResp{}
	err = decoder.Decode(apiResp)
	if err != nil {
		return nil, err
	}
	price, err := strconv.ParseFloat(apiResp.Price, 64)
	if err != nil {
		return nil, err
	}
	return &models.ProductInfo{Price: price, Name: apiResp.Title}, nil
}
