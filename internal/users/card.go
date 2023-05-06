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
	return fmt.Sprintf("\n%s\nOwner: %s\nBank: %s\nRemain funds: %s\nDaily limits: %s\n", c.ID, c.Owner, c.BankInfo, c.Balance, c.LimitInfo)
}
