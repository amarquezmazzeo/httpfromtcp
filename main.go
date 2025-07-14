package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Initialize read byte buffer and open file

	file, err := os.Open("./messages.txt")
	if err != nil {
		log.Fatalf("error opening the file: %v\n", err)
	}

	linesChannel := getLinesChannel(file)

	for line := range linesChannel {
		fmt.Printf("read: %s\n", line)
	}
}
