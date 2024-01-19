package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	isNumeric := flag.Bool("n", false, "Нумеровать все строки при выводе")
	isNonEmptyNumeric := flag.Bool("b", false, "Нумеровать непустые строки при выводе")
	isShowEnd := flag.Bool("E", false, "Показыватьь $  в конце каждой строки")

	flag.Parse()

	files := flag.Args()

	var countLine int64 = 0
	for _, fileName := range files {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("cat: %s: %v\n", fileName, err)
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			if *isNonEmptyNumeric {
				if strings.TrimSpace(line) != "" {
					countLine++
					line = fmt.Sprintf("%6s  %s", strconv.FormatInt(countLine, 10), line)
				}
			} else if *isNumeric {
				countLine++
				line = fmt.Sprintf("%6s  %s", strconv.FormatInt(countLine, 10), line)
			}
			if *isShowEnd {
				line = line + "$"
			}
			fmt.Println(line)
		}
	}
}
