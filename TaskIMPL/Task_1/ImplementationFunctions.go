package Task_1

import "errors"

func (ba *BankAccount) SetNumber(number int) error {
	if number == 0 {
		return errors.New("Не может быть такого номера ")
	}
	ba.Number = number
	return nil
}

func (ba *BankAccount) SetBalance() {
	ba.Number = 0
}

func (ba *BankAccount) SetOwner(p Person) error {
	if p.Age < 14 {
		return errors.New("Не проходит по возрастной категории ")
	}
	ba.Owner = p
	return nil
}

func (ba *BankAccount) GetNumber() int {
	return ba.Number
}

func (ba *BankAccount) GetBalance() int {
	return ba.Balance
}

func (ba *BankAccount) GetOwner() Person {
	return ba.Owner
}

/* Bank Clients Manager */
/*
func (bck *BankClients) AddNewClient(variant Person) {
	bck.Client = append(bck.Client, variant)
}
*/
// Ожидает реализации
