package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	CountLines := flag.Int("n", -1, "Выводит первое КОЛИЧЕСТВО строк файла")
	CountBytes := flag.Int("c", -1, "Выводит первое КОЛИЧЕСТВО байт файла")
	isVerbose := flag.Bool("v", false, "Выводит название(шапку) перед каждым файлом")

	flag.Parse()
	if *CountLines < 0 && *CountBytes < 0 {
		*CountLines = 10
	}

	files := flag.Args()
	for _, fileName := range files {
		var output string

		if *isVerbose {
			output += fmt.Sprintf("==> %s <==\n", fileName)
		}

		file, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("head: %s: %v\n", fileName, err)
			continue
		}
		defer file.Close()

		if *CountLines > 0 && *CountBytes > 0 {
			lines := scanLines(file, *CountLines)
			bytes := scanBytes(file, *CountBytes)
			if len(lines) >= len(bytes) {
				output += lines
			} else {
				output += bytes
			}
		} else if *CountLines > 0 {
			output += scanLines(file, *CountLines)
		} else if *CountBytes > 0 {
			output += scanBytes(file, *CountBytes)
		}
		fmt.Println(output)
	}
}

func scanLines(file *os.File, lines int) string {
	var output string

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan() && i < lines; i++ {
		output += fmt.Sprintln(scanner.Text())
	}

	return output
}

func scanBytes(file *os.File, bytes int) string {
	var output string

	buffer := make([]byte, bytes)
	byteRead, _ := file.Read(buffer)
	output = string(buffer[:byteRead])

	return output
}
