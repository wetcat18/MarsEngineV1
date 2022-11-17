package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

func RotateTexture(g *ebiten.DrawImageOptions, angle int, x float64, y float64) {
	g.GeoM.Translate(x, y)
	g.GeoM.Rotate(float64(angle%360) * 2 * math.Pi / 360)
}

// Расчет длины хп бара

func CountBar(max int, now int, lineWeight int) float64 {
	return (float64(now) * float64(lineWeight)) / float64(max)
}

func RemoveAnimation(i int, mass []Animation) []Animation {
	return append(mass[:i], mass[i+1:]...)
}

func RemoveUnit(i int, mass []Unit) []Unit {
	return append(mass[:i], mass[i+1:]...)
}
