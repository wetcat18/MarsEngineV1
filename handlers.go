package main

import "github.com/hajimehoshi/ebiten/v2"

// Хандлер клавиатуры

func KeyBoardHandler() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		MoveCameraUp()
		MiniMapMoveUp()
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		MoveCameraDown()
		MiniMapMoveDown()
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		MoveCameraRight()
		MiniMapMoveRight()
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		MoveCameraLeft()
		MiniMapMoveLeft()
	}
}

// Тестовый хандлер управления юнитом

func MouseHandler() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		CatchAndSendUnit(x, y)
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		CanalCatch()
	}
}
