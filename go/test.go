package main

import "fmt"

func main() {
	ch := make(chan bool)
	d := <-ch
	go func() {
		ch <- true 
	}()
	fmt.Println(d)
}