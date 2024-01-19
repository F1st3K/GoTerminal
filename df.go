package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	humanReadable := flag.Bool("h", false, "Отображение размеров в человеко-читаемом формате")
	inodes := flag.Bool("i", false, "Отображение информации об нодах вместо использования блоков")
	fileSystemType := flag.String("t", "", "Ограничение списка файловых систем по указанному типу")

	flag.Parse()

	displayDiskFree(*humanReadable, *inodes, *fileSystemType)
}

func displayDiskFree(humanReadable, inodes bool, fileSystemType string) {
	command := "df"
	args := []string{"-P"} // Используем формат вывода POSIX для упрощения обработки данных

	if humanReadable {
		args = append(args, "-h")
	}

	if inodes {
		args = append(args, "-i")
	}

	if fileSystemType != "" {
		args = append(args, "-t", fileSystemType)
	}

	output := getFreeDisk(command, args...)
	fmt.Println(output)
}

func getFreeDisk(command string, args ...string) string {
	cmdOutput, err := exec.Command(command, args...).Output()
	if err != nil {
		return fmt.Sprintf("Ошибка выполнения команды %s: %s", command, err)
	}
	return strings.TrimSpace(string(cmdOutput))
}
