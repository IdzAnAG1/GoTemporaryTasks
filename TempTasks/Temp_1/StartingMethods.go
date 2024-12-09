package Temp_1

import (
	"SIdekick_01/TempTasks/Temp_1/Enemies"
	"fmt"
)

func (g *Game) Start() {
	fmt.Println("Добро пожаловать в игру ...")
	g.Running = true
	err := Enemies.LoadCityEnemiesFromFile()
	if err != nil {
		_ = fmt.Errorf("Ошибка чтения из файла %v ", err)
	}
	for g.Running {
		g.Update()
	}
}
