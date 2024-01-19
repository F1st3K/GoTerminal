package main

import (
	"flag"
	"fmt"
	"syscall"
)

func main() {
	isAll := flag.Bool("a", false, "Вся информация")
	isNName := flag.Bool("n", false, "Имя ядра")
	isSName := flag.Bool("s", false, "Имя пк в сети")
	isKRelease := flag.Bool("r", false, "Релиз ядра")

	flag.Parse()

	if !*isAll && !*isNName && !*isSName && !*isKRelease {
		*isSName = true
	}

	var name syscall.Utsname
	syscall.Uname(&name)

	if *isSName || *isAll {
		fmt.Print(bytesToString(name.Sysname[:]) + " ")
	}
	if *isNName || *isAll {
		fmt.Print(bytesToString(name.Nodename[:]) + " ")
	}
	if *isKRelease || *isAll {
		fmt.Print(bytesToString(name.Release[:]) + " ")
	}

	fmt.Println()
}

func bytesToString(b []int8) string {
	var s []byte
	for _, v := range b {
		if v == 0 {
			break
		}
		s = append(s, byte(v))
	}
	return string(s)
}
