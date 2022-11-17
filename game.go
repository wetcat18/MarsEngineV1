package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

// Конструкция игры

type Game struct{}

func (g *Game) Update() error {
	KeyBoardHandler()   // Хандлер клавиатуры
	MouseHandler()      // Хандлер мыши
	WarFogLogic()       // Логика тумана войны
	CatchAndSendLogic() // Логика вывода текстуры захвата
	ContestUnitLogic()  // Постоянная логика юнита
	UnitMoveLogic()     // Логика передвижения юнитов
	UnitStrikeLogic()   // Логика стрельбы юнитов
	AnimationLogic()    // Логика анимаций
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawMap(screen)          // Отрисовка карты
	DrawWarFog(screen)       // Отрисовка тумана войны
	DrawCatchAndSend(screen) // Отрисовка текстур захвата и точки назначения
	DrawUnits(screen)        // Отрисовка юнитов
	DrawAnimations(screen)   // Отрисовка анимаций
	DrawMiniMap(screen)      // Отрисовка миникарты
	DrawInterface(screen)    // Отрисовка интерфейса
	// Вывод ФПС
	fps := ebiten.CurrentFPS()
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%f", fps), 10, 10)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}

func main() {
	ebiten.SetFullscreen(true)               // Полный экран
	ebiten.SetWindowTitle("New Home")        // Название окна
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOn) // Установка режима отображения
	InisInterface()                          // Инициализация интерфейса
	GenerateMap()                            // Генерация карты
	InisMiniMap()                            // Инициализация миникарты
	InisCatchAndSend()                       // Инициализация функций захвата и команд юнита
	InisFrames()                             // Инициализация кадров анимации
	CreateUnit(1, 10, 8)
	CreateUnit(1, 11, 8)
	CreateUnit(-1, 20, 8)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
