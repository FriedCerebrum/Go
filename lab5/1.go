package main

import (
	"fmt"
	"sync"
)

func count(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Println(num * num)
	}
}

func main() {

	ch := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go count(ch, &wg)

	for i := 1; i <= 5; i++ {
		ch <- i
	}

	close(ch)

	wg.Wait()
}
