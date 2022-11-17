package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Переменный анимаций

var (
	AllAnimations   []Animation
	AnimationFrames [][]*ebiten.Image
)

// Конструкция анимаций

type Animation struct {
	id            int16           // Айди анимации
	size          float64         // Размер текстуры анимации
	posX, posY    float64         // Позиция
	angle         int             // Угол поворота
	nowFrame      *ebiten.Image   // Текущий кадр
	frames        []*ebiten.Image // Все кадры анимации
	nowFrameIndex int8            // Текущий кадр индекс
	delay         float32         // Задержка между кадрами
	nowDelay      float32         // Задержка между кадрами
}

// Инициализация анимаций

func InisFrames() {
	t1, _, _ := ebitenutil.NewImageFromFile("textures/animations/unit_strike/explosion_unit_anim_1.png")
	t2, _, _ := ebitenutil.NewImageFromFile("textures/animations/unit_strike/explosion_unit_anim_2.png")
	t3, _, _ := ebitenutil.NewImageFromFile("textures/animations/unit_strike/explosion_unit_anim_3.png")
	t4, _, _ := ebitenutil.NewImageFromFile("textures/animations/unit_strike/explosion_unit_anim_4.png")
	t5, _, _ := ebitenutil.NewImageFromFile("textures/animations/unit_strike/explosion_unit_anim_5.png")
	t6, _, _ := ebitenutil.NewImageFromFile("textures/animations/unit_strike/explosion_unit_anim_6.png")
	t7, _, _ := ebitenutil.NewImageFromFile("textures/animations/unit_strike/explosion_unit_anim_7.png")
	t8, _, _ := ebitenutil.NewImageFromFile("textures/animations/unit_strike/explosion_unit_anim_8.png")
	t9, _, _ := ebitenutil.NewImageFromFile("textures/animations/unit_strike/explosion_unit_anim_9.png")
	t10, _, _ := ebitenutil.NewImageFromFile("textures/animations/unit_strike/explosion_unit_anim_10.png")
	AnimationFrames = append(AnimationFrames, nil)
	AnimationFrames[0] = append(AnimationFrames[0], t1)
	AnimationFrames[0] = append(AnimationFrames[0], t2)
	AnimationFrames[0] = append(AnimationFrames[0], t3)
	AnimationFrames[0] = append(AnimationFrames[0], t4)
	AnimationFrames[0] = append(AnimationFrames[0], t5)
	AnimationFrames[0] = append(AnimationFrames[0], t6)
	AnimationFrames[0] = append(AnimationFrames[0], t7)
	AnimationFrames[0] = append(AnimationFrames[0], t8)
	AnimationFrames[0] = append(AnimationFrames[0], t9)
	AnimationFrames[0] = append(AnimationFrames[0], t10)
	t1, _, _ = ebitenutil.NewImageFromFile("textures/animations/unit_shoot/shoot_frame_T141_1.png")
	t2, _, _ = ebitenutil.NewImageFromFile("textures/animations/unit_shoot/shoot_frame_T141_2.png")
	t3, _, _ = ebitenutil.NewImageFromFile("textures/animations/unit_shoot/shoot_frame_T141_3.png")
	t4, _, _ = ebitenutil.NewImageFromFile("textures/animations/unit_shoot/shoot_frame_T141_4.png")
	t5, _, _ = ebitenutil.NewImageFromFile("textures/animations/unit_shoot/shoot_frame_T141_5.png")
	AnimationFrames = append(AnimationFrames, nil)
	AnimationFrames[1] = append(AnimationFrames[1], t1)
	AnimationFrames[1] = append(AnimationFrames[1], t2)
	AnimationFrames[1] = append(AnimationFrames[1], t3)
	AnimationFrames[1] = append(AnimationFrames[1], t4)
	AnimationFrames[1] = append(AnimationFrames[1], t5)
	t1, _, _ = ebitenutil.NewImageFromFile("textures/animations/unit_shoot/shoot_frame_desertfox1.png")
	t2, _, _ = ebitenutil.NewImageFromFile("textures/animations/unit_shoot/shoot_frame_desertfox2.png")
	t3, _, _ = ebitenutil.NewImageFromFile("textures/animations/unit_shoot/shoot_frame_desertfox3.png")
	t4, _, _ = ebitenutil.NewImageFromFile("textures/animations/unit_shoot/shoot_frame_desertfox4.png")
	t5, _, _ = ebitenutil.NewImageFromFile("textures/animations/unit_shoot/shoot_frame_desertfox5.png")
	AnimationFrames = append(AnimationFrames, nil)
	AnimationFrames[2] = append(AnimationFrames[2], t1)
	AnimationFrames[2] = append(AnimationFrames[2], t2)
	AnimationFrames[2] = append(AnimationFrames[2], t3)
	AnimationFrames[2] = append(AnimationFrames[2], t4)
	AnimationFrames[2] = append(AnimationFrames[2], t5)

}

// Создание анимации

func CreateAnimation(id int16, posX float64, posY float64, angle int) {
	if id == 0 {
		AllAnimations = append(AllAnimations, Animation{0, 50, posX, posY, angle,
			AnimationFrames[0][0], AnimationFrames[0], 0, 3, 0})
	}
	if id == 1 {
		AllAnimations = append(AllAnimations, Animation{1, 50, posX, posY, angle,
			AnimationFrames[1][0], AnimationFrames[1], 0, 5, 0})
	}
	if id == 2 {
		AllAnimations = append(AllAnimations, Animation{2, 50, posX, posY, angle,
			AnimationFrames[2][0], AnimationFrames[2], 0, 5, 0})
	}
}

// Логика анимаций

func AnimationLogic() {
	for i := 0; i < len(AllAnimations); i++ {
		if AllAnimations[i].nowDelay < AllAnimations[i].delay {
			AllAnimations[i].nowDelay += 1
		} else {
			if int(AllAnimations[i].nowFrameIndex) >= len(AllAnimations[i].frames) {
				AllAnimations = RemoveAnimation(i, AllAnimations)
				break
			}
			AllAnimations[i].nowDelay = 0
			AllAnimations[i].nowFrame = AllAnimations[i].frames[AllAnimations[i].nowFrameIndex]
			AllAnimations[i].nowFrameIndex += 1
		}
	}
}
