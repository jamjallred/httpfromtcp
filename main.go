package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Printf("Error opening file: %w", err)
		return
	}

	fileBytes := make([]byte, 8)

	for {
		bytesNum, err := file.Read(fileBytes)
		if err == io.EOF {
			break
		}
		fmt.Printf("read: %s\n", fileBytes[0:bytesNum])
	}
}
