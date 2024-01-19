package main

import (
	"errors"
	"flag"
	"os"
	"time"
)

func main() {
	isNoCreated := flag.Bool("c", false, "Не создавать файлы")

	flag.Parse()

	files := flag.Args()

	for _, fileName := range files {
		err := os.Chtimes(fileName, time.Now(), time.Now())

		if errors.Is(err, os.ErrNotExist) && *isNoCreated == false {
			os.Create(fileName)
		}

	}
}
