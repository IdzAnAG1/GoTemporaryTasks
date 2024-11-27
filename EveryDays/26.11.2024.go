package EveryDays

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

func run_1() {
	temp := []string{"Привет", "Как дела", "Что делать"}
	for _, tiem := range temp {
		Animated(tiem)
	}
	fmt.Println()
}

func Animated(word string) {
	Wordlen := utf8.RuneCountInString(word)
	for i := 0; i <= len(word); i++ {
		fmt.Printf("\r%s", word[:i])
		time.Sleep(time.Microsecond * 100000)
	}

	for i := Wordlen; i >= 0; i-- {
		current := []rune(word)[:i]
		underscores := strings.Repeat(" ", Wordlen-i)
		fmt.Printf("\r%s%s", string(current), underscores)
		time.Sleep(time.Microsecond * 100000)
	}
}
