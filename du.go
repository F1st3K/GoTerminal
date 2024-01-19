package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	humanReadable := flag.Bool("h", false, "Отображение размеров в человеко-читаемом формате")
	total := flag.Bool("c", false, "Отображение итоговой информации")

	flag.Parse()

	displayDiskUsage(flag.Args(), *humanReadable, *total)
}

func displayDiskUsage(args []string, humanReadable bool, total bool) {
	command := "du"

	if humanReadable {
		args = append(args, "-h")
	}

	if total {
		args = append(args, "-c")
	}

	output := getUsageDisk(command, args...)
	fmt.Println(output)
}

func getUsageDisk(command string, args ...string) string {
	cmdOutput, err := exec.Command(command, args...).Output()
	if err != nil {
		return fmt.Sprintf("Ошибка выполнения команды %s: %s", command, err)
	}
	return strings.TrimSpace(string(cmdOutput))
}
