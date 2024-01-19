package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func main() {
	clearHistory := flag.Bool("c", false, "Очистить историю команд")

	flag.Parse()

	if *clearHistory {
		clearHistoryFile()
		return
	}

	user, _ := user.Current()

	filePath := filepath.Join(user.HomeDir, ".bash_history")
	file, _ := os.ReadFile(filePath)

	fileContent := strings.Split(string(file), "\n")

	for index, command := range fileContent {
		if command != "" {
			fmt.Printf("%5d  %s\n", index+1, command)
		}
	}
}

func clearHistoryFile() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
		return
	}

	filePath := filepath.Join(user.HomeDir, ".bash_history")
	err = os.WriteFile(filePath, []byte(""), 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
}
