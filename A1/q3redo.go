package main

import (
	"fmt"
	"sync"
	"time"
)

func Multiple(x int, m int) bool {
	return x%m == 0
}

func RandomGenerator(wg *sync.WaitGroup, stop <-chan bool, m int) <-chan int {
	intStream := make(chan int)
	go func() {
		defer wg.Done()
		defer close(intStream)
		for {
			select {
			case <-stop:
				return
			case n := <-time.After(time.Millisecond * 100):
				if Multiple(n.Nanosecond(), m) {
					intStream <- n.Nanosecond()
				}
			}
		}
	}()
	return intStream
}

func main() {
	var wg sync.WaitGroup
	stop := make(chan bool)

	wg.Add(3)
	ch5 := RandomGenerator(&wg, stop, 5)
	ch13 := RandomGenerator(&wg, stop, 13)
	ch97 := RandomGenerator(&wg, stop, 97)

	count5, count13, count97 := 0, 0, 0

	for i := 0; i < 100; i++ {
		select {
		case n := <-ch5:
			count5++
			fmt.Printf("%d is a multiple of 5\n", n)
		case n := <-ch13:
			count13++
			fmt.Printf("%d is a multiple of 13\n", n)
		case n := <-ch97:
			count97++
			fmt.Printf("%d is a multiple of 97\n", n)
		}
	}

	close(stop)
	wg.Wait()

	fmt.Printf("Total multiples of 5: %d\n", count5)
	fmt.Printf("Total multiples of 13: %d\n", count13)
	fmt.Printf("Total multiples of 97: %d\n", count97)
}
