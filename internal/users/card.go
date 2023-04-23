package users

import "fmt"

type Card struct {
	ID        int
	Owner     string
	BankInfo  string
	LimitInfo float64
	Balance   float64
}

func (c Card) String() string {
	return fmt.Sprintf("\nOwner: %s\nBank-emitter: %s\nRemain funds: %.2f\nDaily limits: %.2f", c.Owner, c.BankInfo, c.Balance, c.LimitInfo)
}
