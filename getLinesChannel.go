package main

import (
	"io"
	"log"
	"strings"
)

func getLinesChannel(file io.ReadCloser) <-chan string {
	ch := make(chan string)

	byteStream := make([]byte, 8)
	line := ""

	go func() {
		for {
			bytesRead, err := file.Read(byteStream)
			if err != nil {
				if line != "" {
					ch <- line
					line = ""
				}
				if err == io.EOF {
					break
				}
				log.Fatalf("error reading the file: %v\n", err)
				break
			}
			strRead := string(byteStream[:bytesRead])
			parts := strings.Split(strRead, "\n")
			for i := 0; i < len(parts)-1; i++ {
				ch <- (line + parts[i])
				line = ""
			}
			line += parts[len(parts)-1]
		}
		close(ch)
		err := file.Close()
		if err != nil {
			log.Fatalf("error closing the file: %v\n", err)
		}
	}()

	return ch
}
