package main

import (
	"SIdekick_01/TaskIMPL/Task_1"
	"fmt"
)

func main() {
	var Egor Task_1.Person

	Egor.SetPersonName("Егор")
	Egor.SetPersonAddress("МО, г.Балашиха, пр-кт Ленина 80")
	err := Egor.SetPersonAge(11)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Egor.String())
}
