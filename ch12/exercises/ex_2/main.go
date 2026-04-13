package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	quit := make(chan bool)
	go func() {
		for n := 3; n <= 30; n += 3 {
			ch1 <- n
		}
		quit <- true
	}()
	go func() {
		for n := 5; n < 50; n += 5 {
			ch2 <- n
		}
		quit <- true
	}()
	done := 0
	for {
		select {
		case <-quit:
			if done++; done > 1 {
				return
			}
		case a := <-ch1:
			fmt.Println(a)
		case b := <-ch2:
			fmt.Println(b)
		}
	}
}
