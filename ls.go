package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	isAll := flag.Bool("a", false, "Включить скрытые флаги")
	isLong := flag.Bool("l", false, "Показать подробную информацию")
	isReverse := flag.Bool("r", false, "Показать в обратном порядке")
	isHuman := flag.Bool("h", false, "Печатать размеры в удобном для еловека виде")

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		args = append(args, ".")
	}

	for _, arg := range args {

		info, err := os.Stat(arg)
		if err != nil {
			fmt.Printf("%s: %v\n", arg, err)
			continue
		}
		if !info.IsDir() {
			fmt.Printf("%s: Не является директорией\n", arg)
			continue
		}

		files, err := os.ReadDir(arg)
		if err != nil {
			fmt.Printf("%s: %v\n", arg, err)
			continue
		}

		var entries []os.DirEntry

		for _, file := range files {
			if strings.HasPrefix(file.Name(), ".") && *isAll == false {
				continue
			}
			entries = append(entries, file)
		}

		if len(args) > 1 {
			fmt.Println(arg + ":")
		}

		sort.Slice(entries, func(i, j int) bool {
			return *isReverse
		})
		for _, entry := range entries {
			if len(args) > 1 {
				fmt.Print("\t")
			}
			PrintFile(entry, *isLong, *isHuman)
		}
	}
}

func PrintFile(entry os.DirEntry, isLong bool, isHuman bool) {
	ls := entry.Name()
	if entry.IsDir() {
		ls = "./" + ls
	}
	if isLong {
		info, _ := entry.Info()
		size := strconv.FormatInt(info.Size(), 10)
		if isHuman {
			size = strconv.FormatInt(info.Size()/1024, 10) + "MB"
		}
		ls = size + "\t" + info.ModTime().Format("Jan _2 03:15 ") + ls
	}
	fmt.Println(ls)
}
