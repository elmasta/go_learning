package main

import (
	"io"
	"os"
)

func main() {
	if len(os.Args) >= 2 {
		for i, v := range os.Args {
			if i > 0 {
				content, err := os.Open(v)
				if err != nil {
					os.Stdout.Write([]byte("ERROR: open " + v + ": no such file or directory\n"))
					os.Exit(1)
				} else {
					contSize, _ := content.Stat()
					arr := make([]byte, contSize.Size())
					content.Read(arr)
					os.Stdout.Write(arr)
					content.Close()
				}
			}
		}
	} else {
		bytes, _ := io.ReadAll(os.Stdin)
		os.Stdout.Write(bytes)
	}
}
