package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	isByte := flag.Bool("b", false, "Показать информацию в байтах")
	isMByte := flag.Bool("mega", false, "Показать информацию в мегабайтах")
	isGByte := flag.Bool("giga", false, "Показать информацию в гигабайтах")

	flag.Parse()

	if *isByte {
		res(1024, true)
	} else if *isMByte {
		res(1024, false)
	} else if *isGByte {
		res(1024*1024, false)
	} else {
		res(1, false)
	}
}

func res(ch int, key_b bool) {
	var memTotal int
	var memFree int
	var memAvailable int
	var memBuffers int
	var memCashed int
	var memShared int
	var swapTotal int
	var swapFree int

	f, e := os.Open("/proc/meminfo")
	if e != nil {
		panic(e)
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	for s.Scan() {

		if strings.Contains(s.Text(), "MemTotal") {
			res, err := strconv.Atoi(strings.TrimSpace(strings.Split(strings.Split(s.Text(), ":")[1], "k")[0]))
			if err != nil {
				log.Fatal(err)
			}
			memTotal = res
		}
		if strings.Contains(s.Text(), "MemFree") {
			res, err := strconv.Atoi(strings.TrimSpace(strings.Split(strings.Split(s.Text(), ":")[1], "k")[0]))
			if err != nil {
				log.Fatal(err)
			}
			memFree = res
		}
		if strings.Contains(s.Text(), "MemAvailable") {
			res, err := strconv.Atoi(strings.TrimSpace(strings.Split(strings.Split(s.Text(), ":")[1], "k")[0]))
			if err != nil {
				log.Fatal(err)
			}
			memAvailable = res
		}
		if strings.Contains(s.Text(), "Buffers") {
			res, err := strconv.Atoi(strings.TrimSpace(strings.Split(strings.Split(s.Text(), ":")[1], "k")[0]))
			if err != nil {
				log.Fatal(err)
			}
			memBuffers = res
		}
		if strings.Contains(s.Text(), "Cached") && strings.Contains(s.Text(), "SwapCached") == false {
			res, err := strconv.Atoi(strings.TrimSpace(strings.Split(strings.Split(s.Text(), ":")[1], "k")[0]))
			if err != nil {
				log.Fatal(err)
			}
			memCashed = res
		}
		if strings.Contains(s.Text(), "Shmem:") {
			res, err := strconv.Atoi(strings.TrimSpace(strings.Split(strings.Split(s.Text(), ":")[1], "k")[0]))
			if err != nil {
				log.Fatal(err)
			}
			memShared = res
		}
		if strings.Contains(s.Text(), "SwapTotal") {
			res, err := strconv.Atoi(strings.TrimSpace(strings.Split(strings.Split(s.Text(), ":")[1], "k")[0]))
			if err != nil {
				log.Fatal(err)
			}
			swapTotal = res
		}
		if strings.Contains(s.Text(), "SwapFree") {
			res, err := strconv.Atoi(strings.TrimSpace(strings.Split(strings.Split(s.Text(), ":")[1], "k")[0]))
			if err != nil {
				log.Fatal(err)
			}
			swapFree = res
		}
	}
	if key_b {
		memTotal = memTotal * ch
		memFree = memFree * ch
		memAvailable = memAvailable * ch
		memBuffers = memBuffers * ch
		memCashed = memCashed * ch
		memShared = memShared * ch
		swapTotal = swapTotal * ch
		swapFree = swapFree * ch
	} else {
		memTotal = memTotal / ch
		memFree = memFree / ch
		memAvailable = memAvailable / ch
		memBuffers = memBuffers / ch
		memCashed = memCashed / ch
		memShared = memShared / ch
		swapTotal = swapTotal / ch
		swapFree = swapFree / ch
	}

	fmt.Printf("        %12s%12s%12s%12s%12s%12s\n",
		"total", "used", "free", "shared", "buff/cache", "avaliable")
	fmt.Printf("Mem:    %12s%12s%12s%12s%12s%12s\n",
		strconv.Itoa(memTotal), strconv.Itoa(memTotal-memFree),
		strconv.Itoa(memFree), strconv.Itoa(memShared),
		strconv.Itoa(memBuffers+memCashed), strconv.Itoa(memAvailable))
	fmt.Printf("Swap:   %12s%12s%12s%12s%12s%12s\n",
		strconv.Itoa(swapTotal), strconv.Itoa(swapTotal-swapFree),
		strconv.Itoa(swapFree), "", "", "")
}
