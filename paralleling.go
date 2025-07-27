package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Горутина %d начала работу\n", id)
	// Можно добавить time.Sleep для симуляции работы
	fmt.Printf("Горутина %d завершила работу\n", id)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("Все горутины завершены")
}
