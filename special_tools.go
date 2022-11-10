package MarsEngine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"time"
)

// Текущий FPS

func FPS() float64 {
	return ebiten.CurrentFPS()
}

// Синхронизация c кадрами

func SyncFrame() {
	if time.Duration(ebiten.CurrentFPS()) != 0 {
		time.Sleep(time.Second / time.Duration(ebiten.CurrentFPS()))
	}
}

// Поворот текстуры

func rotateTexture(g *ebiten.DrawImageOptions, angle int, x float64, y float64) {
	g.GeoM.Translate(x, y)
	g.GeoM.Rotate(float64(angle%360) * 2 * math.Pi / 360)
}
