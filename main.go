package main

import (
	"github.com/flopp/go-findfont"
	"github.com/goki/freetype/truetype"
	"newTarot/fyneGui"
	"newTarot/tarot"
	"os"
)

func main() {

	//序列化塔罗牌json到map中
	tarot.JsonToMap()

	fyneGui.GuiStart()
}

func init() {
	fontPath, err := findfont.Find("/FontLibrary/simhei.ttf")
	if err != nil {
		panic(err)
	}

	fontData, err := os.ReadFile(fontPath)
	if err != nil {
		panic(err)
	}
	_, err = truetype.Parse(fontData)
	if err != nil {
		panic(err)
	}
	os.Setenv("FYNE_FONT", fontPath)
}
