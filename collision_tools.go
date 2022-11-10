package MarsEngine

// Добавить коллизию объекту

func AddCollision(name string) {
	collisionGroup = append(collisionGroup, name)
}

// Проверить коллизию объектов, вернет два объекта которые пересекаются и сторону с которой они пересекаются

func CheckCollision() (name1, name2 string) {
	allObjectsLen := len(collisionGroup)
	if allObjectsLen != 0 {
		for i := 0; i < allObjectsLen; i++ {
			index := findObject(collisionGroup[i])
			ObjectName1 := collisionGroup[i]
			ObjectSizeX1, ObjectSizeY1 := ObjectSize(ObjectName1)
			UpLeftX1, UpLeftY1 := int(allObjects[index].posX), int(allObjects[index].posY)
			DownRightX1, DownRightY1 := UpLeftX1+ObjectSizeX1, UpLeftY1+ObjectSizeY1
			DownLeftX1, DownLeftY1 := UpLeftX1, UpLeftY1+ObjectSizeY1
			UpRightX1, UpRightY1 := UpLeftX1+ObjectSizeX1, UpLeftY1
			for j := 0; j < allObjectsLen; j++ {
				jndex := findObject(collisionGroup[j])
				ObjectName2 := collisionGroup[j]
				if ObjectName1 != ObjectName2 {
					ObjectSizeX2, ObjectSizeY2 := ObjectSize(ObjectName2)
					UpLeftX2, UpLeftY2 := int(allObjects[jndex].posX), int(allObjects[jndex].posY)
					DownRightX2, DownRightY2 := UpLeftX2+ObjectSizeX2, UpLeftY2+ObjectSizeY2
					// Пересечение с верхним левым углом
					if UpLeftX1 >= UpLeftX2 && UpLeftX1 <= DownRightX2 {
						if UpLeftY1 >= UpLeftY2 && UpLeftY1 <= DownRightY2 {
							return ObjectName1, ObjectName2
						}
					}
					// Пересечение с нижним правым углом
					if DownRightX1 >= UpLeftX2 && DownRightX1 <= DownRightX2 {
						if DownRightY1 >= UpLeftY2 && DownRightY1 <= DownRightY2 {
							return ObjectName1, ObjectName2
						}
					}
					// Пересечение с нижним левым углом
					if DownLeftX1 >= UpLeftX2 && DownLeftX1 <= DownRightX2 {
						if DownLeftY1 >= UpLeftY2 && DownLeftY1 <= DownRightY2 {
							return ObjectName1, ObjectName2
						}
					}
					// Пересечение с верхним правым углом
					if UpRightX1 >= UpLeftX2 && UpRightX1 <= DownRightX2 {
						if UpRightY1 >= UpLeftY2 && UpRightY1 <= DownRightY2 {
							return ObjectName1, ObjectName2
						}
					}
				}
			}
		}
	}
	return "nil", "nil"
}
