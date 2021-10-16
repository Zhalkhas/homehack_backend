package main

import (
	"fmt"
	"github.com/Zhalkhas/homehack_backend/parsers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type App struct {
	Parsers []parsers.Parser
	Server  *gin.Engine
}

func initParsers() []parsers.Parser {
	var res []parsers.Parser
	res = append(res, &parsers.WhitewindParser{})
	res = append(res, &parsers.TechnodomParser{})
	return res
}

func (app *App) StartApp(port string) error {
	return app.Server.Run(port)
}

func NewApp() *App {
	app := &App{}
	app.Parsers = initParsers()
	server := gin.New()
	server.POST("/qr", app.handleQR)
	app.Server = server
	return app
}

type QrDetectReq struct {
	QR string `json:"qr"`
}

func (app *App) handleQR(ctx *gin.Context) {
	req := &QrDetectReq{}
	err := ctx.Bind(req)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	for _, parser := range app.Parsers {
		if parser.Detect(req.QR) {
			productInfo, err := parser.Parse(req.QR)
			if err != nil {
				fmt.Printf("parse err %+v\n", err)
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			ctx.JSON(http.StatusOK, productInfo)
		}
	}
	ctx.AbortWithStatus(http.StatusNotFound)
	return
}
