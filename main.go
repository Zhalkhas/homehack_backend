package main

import "log"

func main() {
	app := NewApp()
	log.Fatalln(app.StartApp(":8080"))
	//testTechno := "138006,711e3234-d82f-11e8-812a-901b0e0fa5bd"
	//parser := parsers.WhitewindParser{}
	//parsedData := parser.Detect("https://shop.kz/bitrix/tools/track_qr.php?art=153003")
	//fmt.Println(parsedData)
}
