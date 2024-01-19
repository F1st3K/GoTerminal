package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	isParent := flag.Bool("p", false, "Удалять родительские подкотологи если они пусты")
	isVerbose := flag.Bool("v", false, "Выводить информацию при удалении каталогов")

	flag.Parse()

	directories := flag.Args()
	if len(directories) == 0 {
		fmt.Println("rmdir: missing operand")
		return
	}

	for _, dir := range directories {
		removeDir(dir, *isParent, *isVerbose)
	}
}

func removeDir(dir string, isParent bool, isVerbose bool) {
	var err error

	status, _ := os.Stat(dir)

	if status.IsDir() == false {
		fmt.Printf("rmdir: failed to remove %s: %s\n", dir, "not a directory")
		return
	}

	err = os.Remove(dir)

	if err != nil {
		fmt.Printf("rmdir: failed to remove %s: %v\n", dir, err)
		return
	}

	if isVerbose {
		fmt.Printf("rmdir: removing directory '%s'\n", dir)
	}

	if index := strings.LastIndex(dir, "/"); index != -1 {
		parent := dir[:index]
		if isParent && err == nil {
			removeDir(parent, isParent, isVerbose)
		}
	}
}
