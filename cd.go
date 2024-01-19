package main

/*
#include <unistd.h>
*/
import "C"

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	flag.Bool("L", false, "Принудительно следовать символьным ссылкам")
	flag.Bool("P", false, "Использовать физическую структуру каталогов, не следуя символическим ссылкам")

	flag.Parse()

	var directory string
	//если аргументы отсутствуют, перейти в домашний каталог пользователя
	countArgs := len(flag.Args())
	if len(flag.Args()) == 1 {
		directory = flag.Arg(0)
	}
	if countArgs == 0 {
		directory, _ = os.UserHomeDir()
	}
	if countArgs > 1 {
		fmt.Println("cd: too many arguments")
		return
	}

	//Поменяем текущий католог с помощью С
	_, err := C.chdir(C.CString(directory))
	if err != nil {
		log.Fatal(err)
	}

	//Получаем переменную окружения SHELL и запускаем оболочку
	shell := os.Getenv("SHELL")
	cmd := exec.Command(shell)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
