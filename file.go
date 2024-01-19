package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	flag.Parse()

	files := flag.Args()

	for _, fileName := range files {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("file: %s: %v\n", fileName, err)
			continue
		}
		defer file.Close()

		buffer := make([]byte, 512)
		file.Read(buffer)
		fmt.Println(fileName + ": ", http.DetectContentType(buffer))
	}

}
