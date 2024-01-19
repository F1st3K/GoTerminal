package main

import (
	"flag"
	"fmt"
)

func main() {
	isScroll := flag.Bool("x", false, "Проскролить вместо очистки")
	flag.Parse()

	if *isScroll {
		fmt.Print("\033[H\033[2J")
	} else {
		fmt.Print("\x1bc")
	}

}
