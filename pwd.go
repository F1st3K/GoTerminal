package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	isValue := flag.Bool("L", false, "Показать значение $PWD")
	isPhysical := flag.Bool("P", false, "Показать физическую директорию, без символических ссылок")

	flag.Parse()

	pwd, _ := os.Getwd()
	if *isValue {
		fmt.Println(pwd)
		return
	}
	if *isPhysical {
		pwd, _ = filepath.Abs(filepath.Dir("."))
	}

	fmt.Println(pwd)
}
