package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	isUniversal := flag.Bool("u", false, "Выводит текущее Universal Time (UTC)")
	isRFC := flag.Bool("R", false, "Выводит время и дату в формате RFC 3339")

	flag.Parse()

	currentTime := time.Now()
	if *isUniversal {
		currentTime = currentTime.UTC()
	}

	if *isRFC {
		fmt.Println(currentTime.Format("Mon, _2 Jan 2006 15:04:05 -0700"))
	} else {
		fmt.Println(currentTime.Format("Mon Jan _2 03:04:05 PM MST 2006"))
	}
}
