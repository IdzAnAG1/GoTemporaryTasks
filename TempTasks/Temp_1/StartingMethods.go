package Temp_1

import "fmt"

func (g *Game) Start() {
	fmt.Println("Добро пожаловать в игру ...")
	g.Running = true
	for g.Running {
		g.Update()
	}
}
