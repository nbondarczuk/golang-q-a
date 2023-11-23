package model

type User struct {
	ID    int
	Name  string
	Email string
}

type Account struct {
	ID       int
	Name     string
	IBAN     string
	Currency string
}

type Balance struct {
	ID     int
	Name   string
	Amount int
}

type Transaction struct {
	ID       int
	Date     time.time
	Title    string
	FromIBAN string
	ToIBAN   string
}
