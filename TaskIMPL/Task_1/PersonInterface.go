package Task_1

import (
	"errors"
	"fmt"
)

type AboutPerson interface {

	// SetPersonName - Установка имени для Персоны
	SetPersonName(name string)
	// SetPersonAge - Установка возраста Персоны
	SetPersonAge(age uint8)
	// SetPersonAddress - Установка адреса Персоны
	SetPersonAddress(address string)

	// GetPersonName - Получение Имени Персоны
	GetPersonName() string
	// GetPersonAge - Получение Возраста Персоны
	GetPersonAge() uint8
	// GetPersonAddress - Получение Адреса Персоны
	GetPersonAddress() string
	// String - метод для вывода информации
	String()
}

/* Реализация методов */

/* Сеттеры */

func (p *Person) SetPersonName(name string) error {
	if p.Name == "" {
		return errors.New("имя не может быть пустим")
	}
	p.Name = name
	return nil
}
func (p *Person) SetPersonAge(age uint8) error {
	if age > 105 {
		return errors.New("вероятнее всего вы ошиблись при указании возраста")
	} else if age < 14 {
		return errors.New("вы не проходите по возротному параметру")

	}
	p.Age = age
	return nil
}
func (p *Person) SetPersonAddress(address string) error {
	if p.Address == "" {
		return errors.New("адрес не может быть пустым ")
	}
	p.Address = address
	return nil
}

/* Геттеры */

func (p *Person) GetPersonName() string {
	return p.Name
}
func (p *Person) GetPersonAge() uint8 {
	return p.Age
}
func (p *Person) GetPersonAddress() string {
	return p.Address
}

/* Вывод в консоль */

func (p *Person) String() string {
	return fmt.Sprintf("Карточка Person : {\n\t Имя : %s \n\t Возраст : %d \n\t Адресс : %s \n }", p.Name, p.Age, p.Address)
}
