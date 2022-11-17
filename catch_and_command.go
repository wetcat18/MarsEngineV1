package main

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Инициализация функций захвата и команд юнита

func InisCatchAndSend() {
	CatchUnitTexture, _, _ = ebitenutil.NewImageFromFile("textures/interface/catch_ram.png")
	SendTexture, _, _ = ebitenutil.NewImageFromFile("textures/interface/send_ram.png")
}

// Захват юнита

func CatchAndSendUnit(posX int, posY int) {
	for i := 0; i < len(GameMap); i++ {
		UpLeftX, UpLeftY := int(GameMap[i].posX), int(GameMap[i].PosY)
		DownRightX, DownRightY := UpLeftX+ChunkSize, UpLeftY+ChunkSize
		if posX >= UpLeftX && posX <= DownRightX {
			if posY >= UpLeftY && posY <= DownRightY {
				chunkX := int(GameMap[i].posX) / ChunkSize
				chunkY := int(GameMap[i].PosY) / ChunkSize
				for k := 0; k < len(AllUnits); k++ {
					if !AllUnits[k].Enemy {
						if AllUnits[k].posChunkX == chunkX && AllUnits[k].posChunkY == chunkY {
							CatchObjectType = "unit"
							CatchObjectIndex = k
							ArtNow = true
							InterfaceUnitArtTexture, _, _ = ebitenutil.NewImageFromFile(AllUnits[k].art)
							InterfaceUnitWindowTexture2, _, _ = ebitenutil.NewImageFromFile(AllUnits[k].window)
							break
						}
					}
				}
				if CatchObjectIndex != -1 {
					SendUnit(CatchObjectIndex, chunkX, chunkY)
				}
			}
		}
	}
}

// Отмена захвата

func CanalCatch() {
	CatchObjectIndex = -1
	CatchObjectType = ""
	ArtNow = false
}

// Логика вывода текстур захвата и точки назначения

func CatchAndSendLogic() {
	if CatchObjectIndex != -1 {
		if CatchObjectType == "unit" {
			CatchTexturePosX = AllUnits[CatchObjectIndex].posX
			CatchTexturePosY = AllUnits[CatchObjectIndex].posY
			unitX, unitY := AllUnits[CatchObjectIndex].posX, AllUnits[CatchObjectIndex].posY
			wayX, wayY := AllUnits[CatchObjectIndex].wayX, AllUnits[CatchObjectIndex].wayY
			if unitX != wayX || unitY != wayY {
				SendTexturePosX = wayX
				SendTexturePosY = wayY
			} else {
				SendTexturePosX = -1000000
				SendTexturePosY = -1000000
			}
		}
	}
}
