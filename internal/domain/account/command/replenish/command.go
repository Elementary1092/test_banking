package replenish

type Command struct {
	AccountNumber string
	FromCard      string
	Amount        float64
	Currency      string
}
