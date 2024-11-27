package EveryDays

import (
	"fmt"
	"strings"
	"time"
)

func run() {

	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h")

	GameName("Grizzly")
	fmt.Print("\033[2J")
	ProgressBar()

	lastAnimation()
}

func GameName(name string) {
	for _, item := range name {
		fmt.Printf("%c", item)
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("\r")
}
func ProgressBar() {
	eq := "=========="
	for i := 0; i <= 10; i++ {
		temp1 := 10
		temp := i * 10
		fmt.Printf("[%s%s]%d%%", eq[:i], strings.Repeat(" ", temp1-i), temp)
		fmt.Printf("\r")
		time.Sleep(1 * time.Second)
	}
	fmt.Print("\033[2J")
}

func lastAnimation() {
	for i := 0; i <= 10; i++ {
		temp := 1
		var line string
		if i != 10 {
			line = "Запуск"
		} else {
			line = "Готово"
		}

		fmt.Printf("%s %s %s \r", strings.Repeat("<", temp+i), line, strings.Repeat(">", temp+i))

		time.Sleep(1 * time.Second)
	}
}
