package main

func MoveCameraUp() {
	if GameMap[0].PosY != 0 {
		CameraPosY -= CameraSpeed
		for i := 0; i < len(GameMap); i++ {
			GameMap[i].PosY += CameraSpeed
		}
		for i := 0; i < len(AllWarFog); i++ {
			AllWarFog[i].posY += CameraSpeed
		}
		for i := 0; i < len(AllUnits); i++ {
			AllUnits[i].posY += CameraSpeed
			AllUnits[i].wayY += CameraSpeed
		}
		for i := 0; i < len(AllAnimations); i++ {
			AllAnimations[i].posY += CameraSpeed
		}
	}
}

func MoveCameraDown() {
	if GameMap[len(GameMap)-1].PosY+float64(ChunkSize) >= float64(1080) {
		CameraPosY += CameraSpeed
		for i := 0; i < len(GameMap); i++ {
			GameMap[i].PosY -= CameraSpeed
		}
		for i := 0; i < len(AllWarFog); i++ {
			AllWarFog[i].posY -= CameraSpeed
		}
		for i := 0; i < len(AllUnits); i++ {
			AllUnits[i].posY -= CameraSpeed
			AllUnits[i].wayY -= CameraSpeed
		}
		for i := 0; i < len(AllAnimations); i++ {
			AllAnimations[i].posY -= CameraSpeed
		}
	}
}

func MoveCameraLeft() {
	if GameMap[0].posX != 0 {
		CameraPosX -= CameraSpeed
		for i := 0; i < len(GameMap); i++ {
			GameMap[i].posX += CameraSpeed
		}
		for i := 0; i < len(AllWarFog); i++ {
			AllWarFog[i].posX += CameraSpeed
		}
		for i := 0; i < len(AllUnits); i++ {
			AllUnits[i].posX += CameraSpeed
			AllUnits[i].wayX += CameraSpeed
		}
		for i := 0; i < len(AllAnimations); i++ {
			AllAnimations[i].posX += CameraSpeed
		}
	}
}

func MoveCameraRight() {
	if GameMap[len(GameMap)-1].posX+float64(ChunkSize) >= float64(1920) {
		CameraPosX += CameraSpeed
		for i := 0; i < len(GameMap); i++ {
			GameMap[i].posX -= CameraSpeed
		}
		for i := 0; i < len(AllWarFog); i++ {
			AllWarFog[i].posX -= CameraSpeed
		}
		for i := 0; i < len(AllUnits); i++ {
			AllUnits[i].posX -= CameraSpeed
			AllUnits[i].wayX -= CameraSpeed
		}
		for i := 0; i < len(AllAnimations); i++ {
			AllAnimations[i].posX -= CameraSpeed
		}
	}
}
