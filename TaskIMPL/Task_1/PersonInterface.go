package Task_1

import (
	"errors"
	"fmt"
)

type AboutPerson interface {
	// Установка для значений структуры Person
	SetPersonName(name string)
	SetPersonAge(age uint8)
	SetPersonAddress(address string)
	// Получение значений структуры Person
	GetPersonName() string
	GetPersonAge() uint8
	GetPersonAddress() string
	// String - метод для вывода информации
	String()
}

// Реализация методов
func (p *Person) SetPersonName(name string) {
	p.Name = name
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
func (p *Person) SetPersonAddress(address string) {
	p.Address = address
}
func (p *Person) GetPersonName() string {
	return p.Name
}
func (p *Person) GetPersonAge() uint8 {
	return p.Age
}
func (p *Person) GetPersonAddress() string {
	return p.Address
}

func (p *Person) String() string {
	return fmt.Sprintf("Карточка Person : {\n\t Имя : %s \n\t Возраст : %d \n\t Адресс : %s \n }", p.Name, p.Age, p.Address)
}
