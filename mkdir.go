package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func main() {
	isParent := flag.Bool("p", false, "Создавать родительские подкотологи если они отсутствуют")
	isVerbose := flag.Bool("v", false, "Выводить информацию при создании каталогов")

	flag.Parse()

	directories := flag.Args()
	if len(directories) == 0 {
		fmt.Println("mkdir: missing operand")
		return
	}

	for _, dir := range directories {
		make(dir, *isParent, *isVerbose)
	}
}

func make(dir string, isParent bool, isVerbose bool) {
	var err error

	if index := strings.LastIndex(dir, "/"); index != -1 {
		parent := dir[:index]
		if isParent && err == nil {
			make(parent, isParent, isVerbose)
		}
	}

	err = os.Mkdir(dir, fs.ModePerm)

	if err != nil {
		fmt.Printf("mkdir: cannot create directory %s: %v\n", dir, err)
		return
	}

	if isVerbose {
		fmt.Printf("mkdir: created directory '%s'\n", dir)
	}

}
