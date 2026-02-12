package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	length, err := fileLen(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(length)
}

func fileLen(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	var length = 0
	buffer := make([]byte, 2048)
	for {
		count, err := file.Read(buffer)
		length += count
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			return length, nil
		}
	}
}
