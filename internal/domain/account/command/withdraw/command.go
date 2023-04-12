package withdraw

type Command struct {
	From   string
	To     string
	Amount float64
}
