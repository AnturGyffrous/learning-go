package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg1 sync.WaitGroup
	wg1.Add(2)
	go func() {
		defer wg1.Done()
		for n := 3; n <= 30; n += 3 {
			ch <- n
		}
	}()
	go func() {
		defer wg1.Done()
		for n := 5; n < 50; n += 5 {
			ch <- n
		}
	}()
	go func() {
		wg1.Wait()
		close(ch)
	}()
	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for n := range ch {
			fmt.Println(n)
		}
	}()
	wg2.Wait()
}
