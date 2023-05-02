package users

import "fmt"

type Card struct {
	ID        string
	Owner     string
	BankInfo  string
	LimitInfo float64
	Balance   float64
}

func (c Card) String() string {
	return fmt.Sprintf("\n%s\nOwner: %s\nBank: %s\nRemain funds: %.2f\nDaily limits: %.2f\n", c.ID, c.Owner, c.BankInfo, c.Balance, c.LimitInfo)
}
