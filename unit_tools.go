package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Конструкция юнита

type Unit struct {
	id                  int     // id Юнита
	name                string  // Название юнита
	description         string  // Описание юнита
	art                 string  // Арт юнита
	window              string  // Окно юнита
	hp                  int     // Количество жизней юнита
	maxHP               int     // Максимальное количество жизней
	armor               int     // Количество брони юнита
	damage              int     // Урон юнита
	penetration         int     // Бронепробиваемость юнита
	accuracy            int     // Точность юнита от 1 до 5
	velocity            int     // Скорость юнита
	rechargeRateSeconds float64 // Скороксть перезарядки в секундах
	rechargeStatus      float64 // Текущий статус перезарядки
	visionRange         int     // Дистанция зрения юнита
	fireRange           int     // Дистанция максимальной стрельбы юнита
	posChunkX           int     // Позиция юнита чанком по X
	posChunkY           int     // Позиция юнита чанком по Y
	angle               int     // Угол поворота юнита
	intellect           int     // Интеллект юнита, условная единица для сюжета, от 1 до 10
	//voice       string     // Голос юнита, обозначается путь папки с озвучками
	texture        *ebiten.Image // Текстура юнита
	wayX, wayY     float64       // Точка куда ехать юниту по X и Y, если 0, то юнит никуда не едет
	posX, posY     float64       // Реальная позиция по X и Y
	Enemy          bool          // Является ли юнит врагом
	hitAnimIndex   int16         // Индекс анимации попадения
	shootAnimIndex int16         // Индекс анимации стрельбы
}

// Создать юнита, указывается id юнита

func CreateUnit(id int, chunkPosX int, chunkPosY int) {
	switch id {
	case 1:
		texture, _, _ := ebitenutil.NewImageFromFile("textures/units/unit_T141.png")
		unit := Unit{id,
			"Танк T141",
			"", "textures/unit_arts/art_unit_T141.png",
			"textures/unit_windows/T141_window.png", 30, 30, 30, 5,
			20, 3, 1, 4, 4, 4, 4,
			chunkPosX, chunkPosY,
			0, 1, texture,
			float64(chunkPosX * ChunkSize), float64(chunkPosY * ChunkSize),
			float64(chunkPosX * ChunkSize), float64(chunkPosY * ChunkSize), false, 0, 2}
		AllUnits = append(AllUnits, unit)
	case -1:
		texture, _, _ := ebitenutil.NewImageFromFile("textures/units/unit_desert_fox.png")
		unit := Unit{id,
			"Tank Desert Fox", "", "",
			"", 50, 50, 5,
			8, 10, 3, 2, 3, 3, 5, 5,
			chunkPosX, chunkPosY,
			0, 0, texture,
			float64(chunkPosX * ChunkSize), float64(chunkPosY * ChunkSize),
			float64(chunkPosX * ChunkSize), float64(chunkPosY * ChunkSize), true, 0, 1}
		AllUnits = append(AllUnits, unit)
	}
}

// Отправить юнита на указанную точку

func SendUnit(i, chunkPosX int, chunkPosY int) {
	AllUnits[i].wayX = float64(chunkPosX * ChunkSize)
	AllUnits[i].wayY = float64(chunkPosY * ChunkSize)
}

// Постоянная логика юнита

func ContestUnitLogic() {
	for i := 0; i < len(AllUnits); i++ {
		if AllUnits[i].rechargeStatus < AllUnits[i].rechargeRateSeconds {
			AllUnits[i].rechargeStatus += 0.010
		}
		if AllUnits[i].hp <= 0 {
			AllUnits = RemoveUnit(i, AllUnits)
			break
		}
	}
}

// Логика движения юнита

func UnitMoveLogic() {
	for i := 0; i < len(AllUnits); i++ {
		AllUnits[i].posChunkX = int(AllUnits[i].posX / float64(ChunkSize))
		AllUnits[i].posChunkY = int(AllUnits[i].posY / float64(ChunkSize))
		unitX := AllUnits[i].posX
		unitY := AllUnits[i].posY
		wayX := AllUnits[i].wayX
		wayY := AllUnits[i].wayY
		velocity := AllUnits[i].velocity
		if wayX != unitX || wayY != unitY {
			if unitX > wayX && unitY > wayY {
				AllUnits[i].angle = -45
			} else if unitX < wayX && unitY < wayY {
				AllUnits[i].angle = 135
			} else if unitX > wayX && unitY < wayY {
				AllUnits[i].angle = -135
			} else if unitX < wayX && unitY > wayY {
				AllUnits[i].angle = 45
			} else if unitX > wayX {
				AllUnits[i].angle = 270
			} else if unitX < wayX {
				AllUnits[i].angle = 90
			} else if unitY > wayY {
				AllUnits[i].angle = 0
			} else if unitY < wayY {
				AllUnits[i].angle = 180
			}
			if unitX > wayX {
				AllUnits[i].posX -= float64(velocity)
			}
			if unitX < wayX {
				AllUnits[i].posX += float64(velocity)
			}
			if unitY > wayY {
				AllUnits[i].posY -= float64(velocity)
			}
			if unitY < wayY {
				AllUnits[i].posY += float64(velocity)
			}
		}
	}
}

// Логика стрельбы юнита

func UnitStrikeLogic() {
	for i := 0; i < len(AllUnits); i++ {
		for j := 0; j < len(AllUnits); j++ {
			if AllUnits[i].rechargeStatus >= AllUnits[i].rechargeRateSeconds {
				if (AllUnits[i].Enemy == false && AllUnits[j].Enemy == true) || (AllUnits[i].Enemy == true && AllUnits[j].Enemy == false) {
					if AllUnits[i].posChunkX-AllUnits[i].fireRange <= AllUnits[j].posChunkX && AllUnits[i].posChunkX+AllUnits[i].fireRange >= AllUnits[j].posChunkX {
						if AllUnits[i].posChunkY-AllUnits[i].fireRange <= AllUnits[j].posChunkY && AllUnits[i].posChunkY+AllUnits[i].fireRange >= AllUnits[j].posChunkY {
							shootAngle := 0
							unitX := AllUnits[i].posX
							unitY := AllUnits[i].posY
							wayX := AllUnits[j].posX
							wayY := AllUnits[j].posY
							if unitX > wayX && unitY > wayY {
								AllUnits[i].angle = -45
								shootAngle = -45
							} else if unitX < wayX && unitY < wayY {
								AllUnits[i].angle = 135
								shootAngle = 135
							} else if unitX > wayX && unitY < wayY {
								AllUnits[i].angle = -135
								shootAngle = -135
							} else if unitX < wayX && unitY > wayY {
								AllUnits[i].angle = 45
								shootAngle = 45
							} else if unitX > wayX {
								AllUnits[i].angle = 270
								shootAngle = 270
							} else if unitX < wayX {
								AllUnits[i].angle = 90
								shootAngle = 90
							} else if unitY > wayY {
								AllUnits[i].angle = 0
								shootAngle = 0
							} else if unitY < wayY {
								AllUnits[i].angle = 180
								shootAngle = 180
							}
							AllUnits[i].wayX = AllUnits[i].posX
							AllUnits[i].wayY = AllUnits[i].posY
							CreateAnimation(AllUnits[i].hitAnimIndex, wayX+float64(Random(-20, 20)), wayY+float64(Random(-20, 20)), shootAngle)
							CreateAnimation(AllUnits[i].shootAnimIndex, unitX, unitY, shootAngle)
							AllUnits[j].hp = CountDamage(AllUnits[i].damage, AllUnits[i].penetration, AllUnits[j].hp, AllUnits[j].armor)
							AllUnits[i].rechargeStatus = 0
						}
					}
				}
			}
		}
	}
}
