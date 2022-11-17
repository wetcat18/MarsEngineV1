package main

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Инициализация интерфейса

func InisInterface() {
	InterfaceUnitWindowTexture, _, _ = ebitenutil.NewImageFromFile("textures/interface/inter_unit_window.png")
	InterfaceUnitWindowSrc.GeoM.Translate(InterfaceUnitWindowTexturePosX, InterfaceUnitWindowTexturePosY)
	InterfaceUnitArtSrc.GeoM.Translate(InterfaceUnitWindowTexturePosX+4, InterfaceUnitWindowTexturePosY+4)
}
