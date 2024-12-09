package Example_2

import (
	"encoding/json"
	"fmt"
	"os"
)

func (lib *Library) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(lib, "", " ")
	if err != nil {
		return fmt.Errorf("Ошибка сериализации данных: %w ", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Errorf("Возникла ошибка при записи в файл %w ", err)
	}

	fmt.Println("Данные успешно записаны и сохранены в файл:", filename)
	return nil
}
