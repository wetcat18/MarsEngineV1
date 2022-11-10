package MarsEngine

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"log"
)

// Логика игрового движка

type Game struct{}

func (g *Game) Update() error {
	if len(activeQueue) != 0 {
		index := findObject(activeQueue[0].name)
		allObjects[index] = activeQueue[0]
		activeQueue = removeObject(0, activeQueue)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if len(allObjects) != 0 {
		for i := 0; i < len(allObjects); i++ {
			if allObjects[i].demonstration {
				src := &ebiten.DrawImageOptions{}
				sizeX, sizeY := ObjectSize(allObjects[i].name)
				rotateTexture(src, allObjects[i].angle, -float64(sizeX)/2, -float64(sizeY)/2)
				src.GeoM.Translate(allObjects[i].posX+float64(sizeX)/2, allObjects[i].posY+float64(sizeY)/2)
				screen.DrawImage(allObjects[i].texture, src)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 600, 600
}

// Запуск движка

func Start() {
	ebiten.SetWindowSize(winWidth, winHeight)
	ebiten.SetWindowTitle(winName)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
