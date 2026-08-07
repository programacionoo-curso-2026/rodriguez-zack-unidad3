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
	orders := generateOrders(5)

	fmt.Printf("Órdenes generadas: %d\n", len(orders))
}
func generateOrders(count int) []*Order {
	orders := make([]*Order, count)

	for i := 0; i < count; i++ {
		orders[i] = &Order{
			ID:     i + 1,
			Status: "Pendiente",
		}
	}

	return orders
}
