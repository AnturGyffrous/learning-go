package main

import (
	"embed"
	"fmt"
	"os"
)

//go:embed *_rights.txt
var rights embed.FS

func main() {
	if len(os.Args) == 1 {
		fmt.Println("missing required language")
		os.Exit(1)
	}
	data, err := rights.ReadFile(os.Args[1] + "_rights.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}
