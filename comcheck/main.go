package main

import (
	"fmt"
	"os"
)

func main() {
	for _, v := range os.Args {
		if v == "01" || v == "galaxy" || v == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
