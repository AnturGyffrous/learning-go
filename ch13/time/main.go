package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.Parse("2006-01-02 15:04:05 MST", "2015-10-21 16:29:00 PDT")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(t.Location(), t.Format(time.RFC1123))

	loc, err := time.LoadLocation("Europe/London")
	if err != nil {
		fmt.Println(err)
		return
	}

	t2 := t.In(loc)

	fmt.Println(t2.Location(), t2.Format("Monday, 2 January 2006, 15:04:05 MST (-07:00)"))
}
