package Enemies

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Enemy struct {
	Name      string
	HeatPoint int
	Attack    int
}

type Enemies map[string]Enemy

var (
	filepath    string  = "C:\\Users\\SusanoNOMikoto\\Desktop\\Studies\\Go\\goLangProjects\\SIdekick_01\\TempTasks\\Temp_1\\Configures\\CityEnemies.gob"
	CityEnemies Enemies = map[string]Enemy{}
)

func InsertInMap() {
	var newEnemy Enemy
	fmt.Scanln()
	fmt.Print("Enemy.Name :")
	fmt.Scanln(&newEnemy.Name)
	fmt.Print("Enemy.HeatPoint :")
	fmt.Scanln(&newEnemy.HeatPoint)
	fmt.Print("Enemy.Attack) :")
	fmt.Scanln(&newEnemy.Attack)

	CityEnemies[newEnemy.Name] = newEnemy
}

func SaveEnemiesInFile() error {
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("Ошибка создания файла %v ", err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)

	if err := encoder.Encode(CityEnemies); err != nil {
		return fmt.Errorf("Ошибка записи данных в файд %v ", err)
	}
	if FileExists(filepath) {
		fmt.Println("Данные записаны в файл")
	}
	return nil
}

func LoadCityEnemiesFromFile() error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("Ошибка при открытии файла "+
			"конфигурации CityEnemies %v ", err)
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	return decoder.Decode(&CityEnemies)
}
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
