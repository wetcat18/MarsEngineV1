package MarsEngine

// Удаление игрового объекта из массива

func removeObject(i int, mass []gameObject) []gameObject {
	return append(mass[:i], mass[i+1:]...)
}

// Удаление буллиана из массива

func removeBool(i int, mass []bool) []bool {
	return append(mass[:i], mass[i+1:]...)
}

// Удаление стринга из массива

func removeString(i int, mass []string) []string {
	return append(mass[:i], mass[i+1:]...)
}
