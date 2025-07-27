package main

import "fmt"

func generate(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func consume(ch chan int) {
	for num := range ch {
		fmt.Println("Получено число:", num)
	}
}

func main() {
	ch := make(chan int)
	go generate(ch)
	consume(ch)
}
