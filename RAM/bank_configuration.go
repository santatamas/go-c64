package RAM

type BankConfiguration int

const (
	ROM BankConfiguration = iota + 1
	RAM
	IO
)
