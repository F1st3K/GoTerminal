package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	linesFlag := flag.Bool("l", false, "Отображение количества строк")
	wordsFlag := flag.Bool("w", false, "Отображение количества слов")
	bytesFlag := flag.Bool("c", false, "Отображение количества байт")

	flag.Parse()

	filename := flag.Arg(0)

	if filename == "" {
		fmt.Println("wc: enter first operand, file name")
		return
	}

	// Если не указаны ключи, установим все ключи для выполнения
	if !*linesFlag && !*wordsFlag && !*bytesFlag {
		*linesFlag, *wordsFlag, *bytesFlag = true, true, true
	}

	displayWordCount(*linesFlag, *wordsFlag, *bytesFlag, filename)
}

func displayWordCount(countLines, countWords, countBytes bool, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("wc: %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount, wordCount, byteCount := 0, 0, 0

	for scanner.Scan() {
		lineCount++
		words := strings.Fields(scanner.Text())
		wordCount += len(words)
		byteCount += len(scanner.Text()) + 1
	}

	if countLines {
		fmt.Printf("wc: count lines: %d\n", lineCount)
	}

	if countWords {
		fmt.Printf("wc: count words: %d\n", wordCount)
	}

	if countBytes {
		fmt.Printf("wc: count bytes: %d\n", byteCount)
	}
}
