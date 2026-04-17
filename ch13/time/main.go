package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.Parse("2006-01-02 15:04:05 MST", "1955-11-12 06:04:00 GMT")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.Format(time.RFC1123))
}
