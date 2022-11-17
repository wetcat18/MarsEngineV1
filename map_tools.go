package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"math/rand"
	"time"
)

// Структура чанка карты

type Chunk struct {
	sort       int
	texture    *ebiten.Image
	posX, PosY float64
}

// Структура чанка тумана войны

type WarFog struct {
	posX, posY float64
	demo       bool
}

// Инициализация миникарты

func InisMiniMap() {
	MiniMapImg, _, _ = ebitenutil.NewImageFromFile("textures/map/minimap.png")
	MiniMapRamImg, _, _ = ebitenutil.NewImageFromFile("textures/map/minimap_ram.png")
}

// Сдвиг рамки миникарты

func MiniMapMoveUp() {
	if GameMap[0].PosY != 0 {
		MiniMapRamPosY -= CameraSpeed / MiniMapRatY
	}
}

func MiniMapMoveDown() {
	if GameMap[len(GameMap)-1].PosY+float64(ChunkSize) >= float64(1080) {
		MiniMapRamPosY += CameraSpeed / MiniMapRatY
	}
}

func MiniMapMoveLeft() {
	if GameMap[0].posX != 0 {
		MiniMapRamPosX -= CameraSpeed / MiniMapRatX
	}
}

func MiniMapMoveRight() {
	if GameMap[len(GameMap)-1].posX+float64(ChunkSize) >= float64(1920) {
		MiniMapRamPosX += CameraSpeed / MiniMapRatX
	}
}

// Создание карты из чанков

func GenerateMap() {
	rand.Seed(time.Now().UnixNano())
	chunkSort1, _, _ := ebitenutil.NewImageFromFile("textures/map/sand1.png")
	chunkSort2, _, _ := ebitenutil.NewImageFromFile("textures/map/sand2.png")
	chunkSort3, _, _ := ebitenutil.NewImageFromFile("textures/map/sand3.png")
	// Генерация тумана войны
	for j := 0; j < MapSizeY; j++ {
		for i := 0; i < MapSizeX; i++ {
			AllWarFog = append(AllWarFog, WarFog{float64(i * ChunkSize), float64(j * ChunkSize), true})
		}
	}
	// Генерация чанков карты
	for j := 0; j < MapSizeY; j++ {
		for i := 0; i < MapSizeX; i++ {
			switch Random(1, 3) {
			case 1:
				GameMap = append(GameMap, Chunk{1, chunkSort1, float64(i * ChunkSize), float64(j * ChunkSize)})
			case 2:
				GameMap = append(GameMap, Chunk{2, chunkSort2, float64(i * ChunkSize), float64(j * ChunkSize)})
			case 3:
				GameMap = append(GameMap, Chunk{3, chunkSort3, float64(i * ChunkSize), float64(j * ChunkSize)})
			}
		}
	}
}

// Логика тумана войны

func WarFogLogic() {
	for i := 0; i < len(AllUnits); i++ {
		unitPosX, unitPosY := AllUnits[i].posX, AllUnits[i].posY
		vision := AllUnits[i].visionRange
		disX1 := unitPosX + float64(vision*ChunkSize)
		disX2 := unitPosX - float64(vision*ChunkSize)
		disY1 := unitPosY + float64(vision*ChunkSize)
		disY2 := unitPosY - float64(vision*ChunkSize)
		for j := 0; j < len(AllWarFog); j++ {
			if (AllWarFog[j].posX <= disX1 && AllWarFog[j].posX >= disX2) && (AllWarFog[j].posY <= disY1 && AllWarFog[j].posY >= disY2) {
				AllWarFog[j].demo = false
			}

		}
	}
}
