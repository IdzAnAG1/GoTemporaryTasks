package Temp_1

import (
	"SIdekick_01/TempTasks/Temp_1/constatns"
	"bufio"
	"fmt"
	"os"
	"time"
)

func (g *Game) Action_1() {
	scanner := bufio.NewScanner(os.Stdin)
	reader := bufio.NewReader(os.Stdin)
	var choice string
	fmt.Println("Выберете город в который собираетесь отправиться")
	g.World.ShowLocations()

	reader.ReadString('\n')
	scanner.Scan()
	choice = scanner.Text()

	switch choice {

	case "Гавань":
		g.Player.Location = g.World.Locations["harbor"]
	case "Родной Город":
		g.Player.Location = g.World.Locations["city"]
	case "Пустыня":
		g.Player.Location = g.World.Locations["desert"]
	case "Пещера":
		g.Player.Location = g.World.Locations["cave"]
	case "Полоса Препятствий":
		g.Player.Location = g.World.Locations["obstacle_course"]
	case "Гробница":
		g.Player.Location = g.World.Locations["tomb"]
	default:
		fmt.Println("Нет таких локаций")
	}
	// Продумать способ обработки неверного ввода локации
	fmt.Printf("\n\t Вы прибыли в %s \n", choice)
	fmt.Println()
}

func (g *Game) Action_2() {
	fmt.Println("Список врагов")

	for _, enm := range g.Player.Location.Enemies {
		fmt.Printf("%s : {"+
			"\n\t Здороовье : %d ,"+
			"\n\t Урон : %d,"+
			"\n}\n", enm.Name, enm.HeatPoint, enm.Attack)
		fmt.Println()
	}
}
func Action_3() {
	fmt.Println("Третье действие")
}
func Action_4() {
	fmt.Println("Четвертое действие")
}
func Default() {
	fmt.Println("Нет таких действий")
	for i := 0; i <= 5; i++ {
		if i != 5 {
			fmt.Println(".")
			time.Sleep(1 * time.Second)
		}
	}

}
func (g *Game) Update() {
	fmt.Println(constatns.MENU)
	var numOfAction int
	fmt.Scan(&numOfAction)
	fmt.Scan()
	switch numOfAction {
	case 1:
		g.Action_1()

	case 2:
		g.Action_2()
	case 3:
		Action_3()
	case 4:
		Action_4()
	case 5:
		g.Running = false
	default:
		Default()
	}
}

func setName() string {
	var name string
	fmt.Print("Введите никнейм игрока: ")
	fmt.Scan(&name)
	fmt.Println()
	if name == "" {
		return setName()
	} else {
		return name
	}
}
func (w *World) ShowLocations() {
	for _, val := range w.Locations {
		fmt.Println(" ==>", val.Name)
	}
	fmt.Println()
}

func SetPlayer() Player {
	return Player{
		Name:      setName(),
		HeatPoint: 100,
		Inventory: nil,
		Location:  nil,
	}
}
