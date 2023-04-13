package withdraw

type Command struct {
	UserID string
	From   string
	To     string
	Amount float64
}
