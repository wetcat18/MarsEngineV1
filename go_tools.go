package main

import (
	"math/rand"
)

func Random(min, max int) int {
	return rand.Intn((max+1)-min) + min
}
