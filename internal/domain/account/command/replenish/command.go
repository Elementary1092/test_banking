package replenish

type Command struct {
	UserID        string
	AccountNumber string
	FromCard      string
	Amount        float64
	Currency      string
}
