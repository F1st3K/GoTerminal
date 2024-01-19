package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

// Функция exit инициирует процесс завершения программы mate-terminal
func main() {
	cmd := exec.Command("pidof", os.Getenv("SHELL"))
	output, error := cmd.Output()
	if error != nil {
		fmt.Println(error)
	}
	strData := string(output)
	str := strings.TrimSpace(strData)
	intValue, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}

	// Завершение процесса mate-terminal с использованием идентификатора процесса
	er := syscall.Kill(intValue, syscall.SIGKILL)
	if er != nil {
		panic(er.Error())
	}
}
