package Enemies

type Enemy struct {
	Name      string
	HeatPoint int
	Attack    int
}

type Enemies map[string]Enemy

var (
	CityEnemies Enemies = map[string]Enemy{
		"Clerc": {
			Name:      "Clerc",
			HeatPoint: 100,
			Attack:    5,
		},
		"Dog": {
			Name:      "Dog",
			HeatPoint: 60,
			Attack:    9,
		},
		"Rat": {
			Name:      "Rat",
			HeatPoint: 40,
			Attack:    2,
		},
		"Su-Shi Man": {
			Name:      "Su-Shi Man",
			HeatPoint: 100,
			Attack:    18,
		},
		"Chef": {
			Name:      "Chef",
			HeatPoint: 100,
			Attack:    25,
		},
	}
)
