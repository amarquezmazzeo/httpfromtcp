package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// Initialize read byte buffer and open file

	file, err := os.Open("./messages.txt")
	if err != nil {
		log.Fatalf("error opening the file: %v\n", err)
	}

	byteStream := make([]byte, 8)
	line := ""

	// Read every 8 bytes of the file
	for {
		bytesRead, err := file.Read(byteStream)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("error reading the file: %v\n", err)
		}
		line = line + string(byteStream[:bytesRead])
		// fmt.Println(line)
		lines := strings.Split(line, "\n")
		// fmt.Println(lines)
		if len(lines) == 2 {
			fmt.Printf("read: %s\n", lines[0])
			line = lines[1]
		}
		// fmt.Printf("read: %s\n", byteStream[:bytesReady])
	}

	// Close file
	err = file.Close()
	if err != nil {
		log.Fatalf("error closing the file: %v\n", err)
	}
}
