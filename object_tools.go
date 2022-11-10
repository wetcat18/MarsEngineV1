package MarsEngine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

// Конструкция игрового объекта

type gameObject struct {
	name          string
	texture       *ebiten.Image
	posX          float64
	posY          float64
	angle         int
	demonstration bool
}

// Найти объект в массиве всех объектов по имени

func findObject(name string) int {
	for i := 0; i < len(allObjects); i++ {
		if allObjects[i].name == name {
			return i
		}
	}
	return -1
}

// Создать игровой объект

func CreateObject(name string, texture string, posX float64, posY float64, angle int, demonstration bool) {
	img, _, err := ebitenutil.NewImageFromFile(texture)
	if err != nil {
		log.Fatal(err)
	}
	object := gameObject{name, img, posX, posY, angle, demonstration}
	allObjects = append(allObjects, object)
}

// Вернуть координаты объекта

func ObjectPosX(name string) float64 {
	index := findObject(name)
	return allObjects[index].posX
}

func ObjectPosY(name string) float64 {
	index := findObject(name)
	return allObjects[index].posY
}

// Вернуть отображаемость объекта

func ObjectDemo(name string) bool {
	index := findObject(name)
	return allObjects[index].demonstration
}

// Вернуть размер текстуры объекта

func ObjectSize(name string) (x, y int) {
	index := findObject(name)
	return allObjects[index].texture.Size()
}

// Вернуть угол поворота объекта

func ObjectAngle(name string) (angle int) {
	index := findObject(name)
	return allObjects[index].angle
}

// Изменить координаты объекта

func SetObjectPos(name string, posX float64, posY float64) {
	index := findObject(name)
	editObject := gameObject{allObjects[index].name, allObjects[index].texture, posX, posY, allObjects[index].angle, allObjects[index].demonstration}
	activeQueue = append(activeQueue, editObject)
}

// Изменить текстуру объекта

func SetObjectTexture(name string, texture string) {
	index := findObject(name)
	img, _, err := ebitenutil.NewImageFromFile(texture)
	if err != nil {
		log.Fatal(err)
	}
	editObject := gameObject{allObjects[index].name, img, allObjects[index].posX, allObjects[index].posY, allObjects[index].angle, allObjects[index].demonstration}
	activeQueue = append(activeQueue, editObject)
}

// Изменить видимость объекта

func SetObjectDemo(name string, demo bool) {
	index := findObject(name)
	editObject := gameObject{allObjects[index].name, allObjects[index].texture, allObjects[index].posX, allObjects[index].posY, allObjects[index].angle, demo}
	activeQueue = append(activeQueue, editObject)
}

// Подвинуть объект

func MoveObject(name string, posX float64, posY float64) {
	index := findObject(name)
	editObject := gameObject{allObjects[index].name, allObjects[index].texture, posX + allObjects[index].posX, posY + allObjects[index].posY, allObjects[index].angle, allObjects[index].demonstration}
	activeQueue = append(activeQueue, editObject)
}

// Изменить поворот объекта

func SetObjectAngle(name string, angle int) {
	index := findObject(name)
	editObject := gameObject{allObjects[index].name, allObjects[index].texture, allObjects[index].posX, allObjects[index].posY, angle, allObjects[index].demonstration}
	activeQueue = append(activeQueue, editObject)
}

// Повернуть объект на заданный градус

func RotateObject(name string, angle int) {
	index := findObject(name)
	editObject := gameObject{allObjects[index].name, allObjects[index].texture, allObjects[index].posX, allObjects[index].posY, allObjects[index].angle + angle, allObjects[index].demonstration}
	activeQueue = append(activeQueue, editObject)
}

// Проверить находится ли какой-то объект на указанных координатах, вернет имя этого объекта

func CheckObjectsOnCor(x int, y int) string {
	if len(allObjects) != 0 {
		for i := 0; i < len(allObjects); i++ {
			ObjectName := allObjects[i].name
			ObjectSizeX, ObjectSizeY := ObjectSize(ObjectName)
			UpLeftX, UpLeftY := int(allObjects[i].posX), int(allObjects[i].posY)
			DownRightX, DownRightY := UpLeftX+ObjectSizeX, UpLeftY+ObjectSizeY
			//DownLeftX, DownLeftY := UpLeftX, UpLeftY + ObjectSizeY
			//UpRightX, UpRightY := UpLeftX + ObjectSizeX, UpLeftY
			if x >= UpLeftX && x <= DownRightX {
				if y >= UpLeftY && y <= DownRightY {
					return ObjectName
				}
			}
		}
	}
	return "nil"
}
