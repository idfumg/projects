package main

// go get -u github.com/inancgumus/screen

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

type placeholder [5]string

func getDigits() []placeholder {
	return []placeholder{
		{
			"▉▉▉",
			"▉ ▉",
			"▉ ▉",
			"▉ ▉",
			"▉▉▉",
		},
		{
			"▉▉ ",
			" ▉ ",
			" ▉ ",
			" ▉ ",
			"▉▉▉",
		},
		{
			"▉▉▉",
			"  ▉",
			"▉▉▉",
			"▉  ",
			"▉▉▉",
		},
		{
			"▉▉▉",
			"  ▉",
			"▉▉▉",
			"  ▉",
			"▉▉▉",
		},
		{
			"▉ ▉",
			"▉ ▉",
			"▉▉▉",
			"  ▉",
			"  ▉",
		},
		{
			"▉▉▉",
			"▉  ",
			"▉▉▉",
			"  ▉",
			"▉▉▉",
		},
		{
			"▉▉▉",
			"▉  ",
			"▉▉▉",
			"▉ ▉",
			"▉▉▉",
		},
		{
			"▉▉▉",
			"  ▉",
			"  ▉",
			"  ▉",
			"  ▉",
		},
		{
			"▉▉▉",
			"▉ ▉",
			"▉▉▉",
			"▉ ▉",
			"▉▉▉",
		},
		{
			"▉▉▉",
			"▉ ▉",
			"▉▉▉",
			"  ▉",
			"▉▉▉",
		},
	}
}

func getColon() placeholder {
	return placeholder{
		"   ",
		" ▦ ",
		"   ",
		" ▦ ",
		"   ",
	}
}

func getEmptyColon() placeholder {
	return placeholder{
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
	}
}

func getActualColon(seconds int) placeholder {
	if seconds%2 == 0 {
		return getEmptyColon()
	}
	return getColon()
}

func getClock() []placeholder {
	digits := getDigits()
	now := time.Now()
	hours, minutes, seconds := now.Hour(), now.Minute(), now.Second()

	return []placeholder{
		digits[hours/10], digits[hours%10],
		getActualColon(seconds),
		digits[minutes/10], digits[minutes%10],
		getActualColon(seconds),
		digits[seconds/10], digits[seconds%10],
	}
}

func printClock(clock []placeholder) {
	for line := range clock[0] {
		for index := range clock {
			fmt.Print(clock[index][line], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	for {
		screen.Clear()
		screen.MoveTopLeft()
		printClock(getClock())
		time.Sleep(time.Second)
	}
}
