package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
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
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	orders := generateOrders(20)

	wg.Add(1)

	go func() {
		defer wg.Done()

		for _, order := range orders {
			updateOrderStatus(order)
		}
	}()

	wg.Wait()
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
func updateOrderStatus(order *Order) {
	order.mu.Lock()

	status := []string{
		"Procesando",
		"Despachando",
		"Entregado",
	}[rand.Intn(3)]

	order.Status = status

	fmt.Printf("Orden %d -> %s\n",
		order.ID,
		status)

	order.mu.Unlock()
}
