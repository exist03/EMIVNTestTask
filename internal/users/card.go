package users

import "fmt"

type Card struct {
	ID        interface{}
	Owner     string
	BankInfo  interface{}
	LimitInfo string
	Balance   string
}

func (c Card) String() string {
	return fmt.Sprintf("\n%s\nOwner: %s\nBank: %s\nRemain funds: %.2f\nDaily limits: %.2f\n", c.ID, c.Owner, c.BankInfo, c.Balance, c.LimitInfo)
}
