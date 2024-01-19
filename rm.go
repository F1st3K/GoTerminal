package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	isForce := flag.Bool("f", false, "Инорировать несуществующие файлы и аргументы")
	isRecursive := flag.Bool("r", false, "Рекурсивно удалять каталоги и их содержимое")
	isVerbose := flag.Bool("v", false, "Пояснять производимые действия")

	flag.Parse()

	files := flag.Args()

	for _, fileName := range files {
		remove(fileName, *isForce, *isRecursive, *isVerbose)
	}
}

func remove(file string, isForce bool, isRecursive bool, isVerbose bool) {
	var err error

	if isRecursive {
		dirs, _ := os.ReadDir(file)
		for _, f := range dirs {
			remove(file+"/"+f.Name(), isForce, isRecursive, isVerbose)
		}
	}

	err = os.Remove(file)
	if err != nil {
		if isForce == false {
			fmt.Printf("rm: failed to remove %s: %v\n", file, err)
		}
		return
	}

	if isVerbose {
		fmt.Printf("rm: removing '%s'\n", file)
	}
}
