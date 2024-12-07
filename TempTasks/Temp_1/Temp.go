package Temp_1

import (
	"SIdekick_01/TempTasks/Temp_1/Enemies"
	"SIdekick_01/TempTasks/Temp_1/constatns"
)

func Gog() {

	player := SetPlayer()

	city := &Location{
		Name:        constatns.CITY,
		Description: constatns.CITY,
		Enemies:     Enemies.CityEnemies,
	}
	harbor := &Location{
		Name:        constatns.HARBOR,
		Description: constatns.HARBOR,
		Enemies:     nil,
	}
	desert := &Location{
		Name:        constatns.DESERT,
		Description: constatns.DESERT,
		Enemies:     nil,
	}
	cave := &Location{
		Name:        constatns.CAVE,
		Description: constatns.CAVE,
		Enemies:     nil,
	}

	obstacle_course := &Location{
		Name:        constatns.OBSTACLE_COURSE,
		Description: constatns.OBSTACLE_COURSE,
		Enemies:     nil,
	}
	tomb := &Location{
		Name:        constatns.TOMB,
		Description: constatns.TOMB,
		Enemies:     nil,
	}

	world := &World{
		Locations: map[string]*Location{
			"harbor":          harbor,
			"city":            city,
			"desert":          desert,
			"cave":            cave,
			"obstacle_course": obstacle_course,
			"tomb":            tomb,
		},
	}

	game := &Game{
		Player:  &player,
		World:   world,
		Running: false,
	}
	player.Location = city

	game.Start()
}
