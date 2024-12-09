package EveryDays

import (
	"encoding/json"
	"fmt"
	"os"
)

type test struct {
	Description string `json:"Description"`
}

var (
	num_1 = []test{
		{Description: "Первый JSON вывод"},
		{Description: "Второй JSON вывод"},
		{Description: "Третий JSON вывод"},
	}
	num_2 = []test{
		{Description: "Четвертый JSON вывод"},
		{Description: "Пятый JSON вывод"},
		{Description: "Шестой JSON вывод"},
	}
)

func Sunday_01_12_2024() {
	temp := num_1
	var err error = nil

	for i := 0; i < 5; i++ {
		err, temp = writeWithAppendJSON("EveryDays/ResultsToFile/Result_5.txt", temp)
		if err != nil {
			_ = fmt.Errorf("Ошибка вышла %w ", err)
		}
		fmt.Println(len(temp))
	}

}

func writeToFile(filename, data string) error {
	return os.WriteFile(filename, []byte(data), 0644)
}

func writeToFile_2(pathToFile, data string) error {
	file, err := os.Create(pathToFile)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(data)
	return err
}

func writeToFile_3(filename string) error {
	data, err := json.MarshalIndent(num_1, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func writeWithAppend(filename, data string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	return err
}

func writeWithAppendJSON(filename string, newTests []test) (error, []test) {
	var tests []test
	data, err := os.ReadFile(filename)
	if err != nil && !os.IsNotExist(err) {
		return err, nil
	}
	if len(data) > 0 {
		if err := json.Unmarshal(data, &tests); err != nil {
			return err, nil
		}
	}

	for _, nTest := range newTests {
		tests = append(tests, nTest)
	}

	updatedData, err := json.MarshalIndent(tests, "", "  ")

	if err != nil {
		return err, nil
	}

	return os.WriteFile(filename, updatedData, 0644), tests
}
