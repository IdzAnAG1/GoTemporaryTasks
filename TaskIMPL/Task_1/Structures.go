package Task_1

type Person struct {
	Name    string
	Gender  bool
	Age     uint8
	Address string
}

type BankAccount struct {
	Number  int
	Balance int
	Owner   Person
}

type BankClients struct {
	Client map[string]*BankAccount
}
