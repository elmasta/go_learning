package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		content, _ := os.Open(os.Args[1])
		arr := make([]byte, 14)
		content.Read(arr)
		fmt.Println(string(arr))
		content.Close()
	} else if len(os.Args) == 1 {
		fmt.Println("File name missing")
	} else {
		fmt.Println("Too many arguments")
	}
}
