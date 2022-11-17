package main

import "github.com/hajimehoshi/ebiten/v2"

// Переменные игровой камеры

var (
	CameraSpeed            = 50.0  // Скорость движения камеры
	CameraPosX, CameraPosY float64 // Позиция камеры
)

// Переменные карты и чанков

var (
	MapSizeX, MapSizeY             = 100, 100    // Размер карты в чанках
	ChunkSize                      = 50          // Размер чанка в пикселях
	MiniMapRatX, MiniMapRatY       = 16.2, 20.0  // Коэффицент между размерами реальной карты и миникарты, вычисляется делением одного на другое
	GameMap                        []Chunk       // Массив карты из чанков
	MiniMapImg                     *ebiten.Image // Текстура миникарты
	MiniMapRamImg                  *ebiten.Image // Текстура рамки миникарты
	MiniMapRamPosX, MiniMapRamPosY float64       // Координаты рамки миникарты
)

// Переменные тумана войны

var (
	WarFogTexture *ebiten.Image
	AllWarFog     []WarFog
)

// Переменные юнитов

var (
	AllUnits []Unit // Массив со всеми существующими на карте юнитами
)

// Переменные захвата

var (
	CatchObjectType                    string        // Индекс в массиве захваченного объекта
	CatchObjectIndex                   = -1          // Тип захваченного объекта
	CatchUnitTexture                   *ebiten.Image // Тектура захвата юнита
	SendTexture                        *ebiten.Image // Текстура точки назначения юнита
	CatchTexturePosX, CatchTexturePosY float64       // Позиции текстуры захвата
	SendTexturePosX, SendTexturePosY   float64       // Позиция текстуры точки назначения юнита
)

// Переменные интерфейса

var (
	InterfaceUnitWindowTexture     *ebiten.Image                                             // Текстура окошка юнита пустое
	InterfaceUnitWindowSrc                                      = &ebiten.DrawImageOptions{} // Серфейс окошка юнита пустое
	InterfaceUnitWindowTexturePosX float64                      = 1720                       // Позиция окошка юнита X
	InterfaceUnitWindowTexturePosY float64                      = 880                        // Позиция окошка юнита Y
	ArtNow                         bool                         = false                      // Выведен ли какой-то арт
	InterfaceUnitArtTexture        *ebiten.Image                                             // Текстура арта юнита
	InterfaceUnitArtSrc            = &ebiten.DrawImageOptions{}                              // Серфейс арта юнита
	InterfaceUnitWindowTexture2    *ebiten.Image                                             // Текстура окна юнита
)
