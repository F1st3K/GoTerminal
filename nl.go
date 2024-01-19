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
	flag.Parse()

	files := flag.Args()

	var countLine int64 = 0
	for _, fileName := range files {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("nl: %s: %v\n", fileName, err)
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			if strings.TrimSpace(line) != "" {
				countLine++
				line = fmt.Sprintf("%6s  %s", strconv.FormatInt(countLine, 10), line)
			}
			fmt.Println(line)
		}
	}
}
