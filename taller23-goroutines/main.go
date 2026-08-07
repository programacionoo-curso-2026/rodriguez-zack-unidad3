package main

import (
	"fmt"
	"sync"
)

type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}

var (
	totalUpdates int
	updateMutex  sync.Mutex
)

func main() {
	fmt.Println("Estructura Order creada")
}
