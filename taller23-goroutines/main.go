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

	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()

			for _, order := range orders {
				updateOrderStatus(order)
			}
		}()
	}

	wg.Wait()

	reportOrderStatus(orders)

	fmt.Println("\nTodas las operaciones completadas.")
	fmt.Printf("Total Actualizaciones: %d\n", totalUpdates)
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

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	status := []string{
		"Procesando",
		"Despachando",
		"Entregado",
	}[rand.Intn(3)]

	order.Status = status

	fmt.Printf("Actualizando orden %d con estado: %s\n",
		order.ID, status)

	order.mu.Unlock()

	updateMutex.Lock()
	defer updateMutex.Unlock()

	currentUpdates := totalUpdates
	time.Sleep(5 * time.Millisecond)
	totalUpdates = currentUpdates + 1
}
