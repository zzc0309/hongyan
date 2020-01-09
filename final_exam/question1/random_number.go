package main

import (
	"math/rand"
	"time"
)

func random_num()int{
	rand.Seed(time.Now().Unix())

	return rand.Intn(5)
}