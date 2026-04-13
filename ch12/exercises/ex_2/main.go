package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for n := 3; n <= 30; n += 3 {
			ch1 <- n
		}
		close(ch1)
	}()
	go func() {
		for n := 5; n < 50; n += 5 {
			ch2 <- n
		}
		close(ch2)
	}()
	done := 2
	for done > 0 {
		select {
		case n, ok := <-ch1:
			if !ok {
				ch1 = nil
				done--
				break
			}
			fmt.Println("ch1 - ", n)
		case n, ok := <-ch2:
			if !ok {
				ch2 = nil
				done--
				break
			}
			fmt.Println("ch2 - ", n)
		}
	}
}
