package Task_1

type Person struct {
	Name    string
	Age     uint8
	Address string
}

type BankAccount struct {
	Number  int
	Balance int
	Owner   string
}

type BankClients struct {
	Client map[string]*BankAccount
}
