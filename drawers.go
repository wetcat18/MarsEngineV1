package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

// Отрисовщик карты

func DrawMap(screen *ebiten.Image) {
	for i := 0; i < len(GameMap); i++ {
		src := &ebiten.DrawImageOptions{}
		src.GeoM.Translate(GameMap[i].posX, GameMap[i].PosY)
		screen.DrawImage(GameMap[i].texture, src)
	}
}

// Отрисовщик тумана войны

func DrawWarFog(screen *ebiten.Image) {
	WarFogTexture, _, _ = ebitenutil.NewImageFromFile("textures/map/war_fog.png")
	for i := 0; i < len(AllWarFog); i++ {
		if AllWarFog[i].demo == true {
			src := &ebiten.DrawImageOptions{}
			src.GeoM.Translate(AllWarFog[i].posX, AllWarFog[i].posY)
			screen.DrawImage(WarFogTexture, src)
		}
	}
}

// Отрисовка миникарты

func DrawMiniMap(screen *ebiten.Image) {
	src := &ebiten.DrawImageOptions{}
	src.GeoM.Translate(0, 880)
	screen.DrawImage(MiniMapImg, src)
	src.GeoM.Translate(MiniMapRamPosX, MiniMapRamPosY)
	screen.DrawImage(MiniMapRamImg, src)
}

// Отрисовка юнитов

func DrawUnits(screen *ebiten.Image) {
	for i := 0; i < len(AllUnits); i++ {
		src := &ebiten.DrawImageOptions{}
		if AllUnits[i].angle != 0 {
			RotateTexture(src, AllUnits[i].angle, -float64(ChunkSize)/2, -float64(ChunkSize)/2)
			src.GeoM.Translate(AllUnits[i].posX+float64(ChunkSize)/2, AllUnits[i].posY+float64(ChunkSize)/2)
		} else {
			src.GeoM.Translate(AllUnits[i].posX, AllUnits[i].posY)
		}
		screen.DrawImage(AllUnits[i].texture, src)
	}
}

// Отрисовка текстур завхата и точки назначения юнита

func DrawCatchAndSend(screen *ebiten.Image) {
	if CatchObjectIndex != -1 {
		src := &ebiten.DrawImageOptions{}
		src2 := &ebiten.DrawImageOptions{}
		src.GeoM.Translate(CatchTexturePosX, CatchTexturePosY)
		screen.DrawImage(CatchUnitTexture, src)
		if SendTexturePosX != -1000000 && SendTexturePosY != -1000000 {
			src2.GeoM.Translate(SendTexturePosX, SendTexturePosY)
			screen.DrawImage(SendTexture, src2)
		}
	}
}

// Отрисовка анимаций

func DrawAnimations(screen *ebiten.Image) {
	for i := 0; i < len(AllAnimations); i++ {
		src := &ebiten.DrawImageOptions{}
		if AllAnimations[i].angle != 0 {
			RotateTexture(src, AllAnimations[i].angle, -float64(ChunkSize)/2, -float64(ChunkSize)/2)
			src.GeoM.Translate(AllAnimations[i].posX+float64(ChunkSize)/2, AllAnimations[i].posY+float64(ChunkSize)/2)
		} else {
			src.GeoM.Translate(AllAnimations[i].posX, AllAnimations[i].posY)
		}
		screen.DrawImage(AllAnimations[i].nowFrame, src)
	}
}

// Отрисовка интерфейса

func DrawInterface(screen *ebiten.Image) {
	if ArtNow {
		screen.DrawImage(InterfaceUnitArtTexture, InterfaceUnitArtSrc)        // Арт объекта для окошка
		screen.DrawImage(InterfaceUnitWindowTexture2, InterfaceUnitWindowSrc) // Окошко для объектов снизу справа
		// ХП бар
		if AllUnits[CatchObjectIndex].hp > (AllUnits[CatchObjectIndex].maxHP / 2) {
			ebitenutil.DrawRect(screen, InterfaceUnitWindowTexturePosX+16, InterfaceUnitWindowTexturePosY+185,
				CountBar(AllUnits[CatchObjectIndex].maxHP, AllUnits[CatchObjectIndex].hp, 168),
				10, color.RGBA{34, 139, 34, 200})
		} else if AllUnits[CatchObjectIndex].hp < (AllUnits[CatchObjectIndex].maxHP/2) && AllUnits[CatchObjectIndex].hp > (AllUnits[CatchObjectIndex].maxHP/4) {
			ebitenutil.DrawRect(screen, InterfaceUnitWindowTexturePosX+16, InterfaceUnitWindowTexturePosY+185,
				CountBar(AllUnits[CatchObjectIndex].maxHP, AllUnits[CatchObjectIndex].hp, 168),
				10, color.RGBA{184, 134, 11, 200})
		} else {
			ebitenutil.DrawRect(screen, InterfaceUnitWindowTexturePosX+16, InterfaceUnitWindowTexturePosY+185,
				CountBar(AllUnits[CatchObjectIndex].maxHP, AllUnits[CatchObjectIndex].hp, 168),
				10, color.RGBA{165, 42, 42, 200})
		}
		// Прогресс бар
		ebitenutil.DrawRect(screen, InterfaceUnitWindowTexturePosX+16, InterfaceUnitWindowTexturePosY+166,
			CountBar(int(AllUnits[CatchObjectIndex].rechargeRateSeconds), int(AllUnits[CatchObjectIndex].rechargeStatus), 168),
			10, color.RGBA{36, 100, 100, 200})

	} else {
		screen.DrawImage(InterfaceUnitWindowTexture, InterfaceUnitWindowSrc) // Окошко для объектов снизу справа
	}
}
