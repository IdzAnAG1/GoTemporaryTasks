package Temp_1

import "SIdekick_01/TempTasks/Temp_1/Enemies"

type Player struct {
	Name      string
	HeatPoint uint8
	Inventory []string
	Location  *Location
}

type Location struct {
	Name        string
	Description string
	Enemies     Enemies.Enemies
}

type World struct {
	Locations map[string]*Location
}

type Game struct {
	Player  *Player
	World   *World
	Running bool
}
