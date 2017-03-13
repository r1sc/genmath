package main

import (
	"math/rand"
)

func chance(max int) bool {
	return rand.Intn(max) == 1
}
