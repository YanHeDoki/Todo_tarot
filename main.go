package main

import (
	"github.com/flopp/go-findfont"
	"newTarot/fyneGui"
	"newTarot/tarot"
	"os"
	"strings"
)

func main() {

	//序列化塔罗牌json到map中
	tarot.JsonToMap()

	fyneGui.GuiStart()

	os.Unsetenv("FYNE_FONT")
}

func init() {

	fontPaths := findfont.List()

	for _, path := range fontPaths {
		if strings.Contains(path, "simkai.ttf") || strings.Contains(path, "simkai.ttc") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}

}
