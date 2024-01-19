package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	isVerbose := flag.Bool("v", false, "Поясняет производимые действия")
	isRecursive := flag.Bool("r", false, "Копирует все подкаталоги рекурсивно")
	isForce := flag.Bool("f", false, "Игнорирует возникающие исключения")

	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("cp: skipped source operand")
		return
	} else if len(flag.Args()) < 2 {
		fmt.Println("cp: skipped purpose operand")
		return
	}

	args := flag.Args()
	sources := args[:len(args)-1]
	purpose := args[len(args)-1]

	for _, source := range sources {
		copyEntries(source, purpose, *isVerbose, *isRecursive, *isForce)
	}
}

func getNamePath(path string) string {
	paths := strings.Split(path, "/")
	return ("/" + paths[len(paths)-1])
}

func copyEntries(source string, purpose string, isVerbose bool, isRecursive bool, isForce bool) {
	info, err := os.Stat(source)
	if err != nil && isForce == false {
		fmt.Printf("cp: %s: %v\n", source, err)
		return
	}

	if info.IsDir() {
		if isRecursive == false {
			fmt.Printf("cp: no such flag -r; skipped dir '%s'\n", source)
			return
		}

		copyDir(source, purpose, isVerbose, isForce)
	} else {
		copyFile(source, purpose, isVerbose, isForce)
	}
}

func copyDir(source string, purpose string, isVerbose bool, isForce bool) {
	purpose += getNamePath(source)
	os.MkdirAll(purpose, os.ModePerm)

	if isVerbose {
		fmt.Printf("'%s' -> '%s'\n", source, purpose)
	}

	dirs, err := os.ReadDir(source)
	if err != nil && isForce == false {
		fmt.Printf("cp: %s: %v\n", source, err)
		return
	}

	for _, entry := range dirs {
		copyEntries(source+"/"+entry.Name(), purpose, isVerbose, true, isForce)
	}
}

func copyFile(source string, purpose string, isVerbose bool, isForce bool) {
	info, err := os.Stat(purpose)
	if err == nil && info.IsDir() {
		purpose += getNamePath(source)
	}

	sourceFile, err := os.Open(source)
	if err != nil && isForce == false {
		fmt.Printf("cp: %s: %v\n", source, err)
		return
	}
	defer sourceFile.Close()

	purposeFile, err := os.Create(purpose)
	if err != nil && isForce == false {
		fmt.Printf("pcp: %s: %v\n", purpose, err)
		return
	}
	defer purposeFile.Close()

	_, err = io.Copy(purposeFile, sourceFile)
	if err != nil && isForce == true {
		fmt.Printf("cp: %v\n", err)
	}
	if isVerbose {
		fmt.Printf("'%s' -> '%s'\n", source, purpose)
	}
}
