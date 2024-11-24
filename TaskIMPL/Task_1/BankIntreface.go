package Task_1

type BankAccountInterface interface {
	SetNumber(number int)
	SetBalance()
	SetOwner(p Person)

	GetNumber() int
	GetBalance() int
	GetOwner() Person
}

type BankClientsManager interface {
	AddNewClient()
	RemoveClient()
	SortClientsByAge()
	SortClientByGender()
}
