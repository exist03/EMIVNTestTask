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
	return fmt.Sprintf("\n%s\nВладелец: %s\nBank: %s\nОстаток: %s\nЛимит: %s", c.ID, c.Owner, c.BankInfo, c.Balance, c.LimitInfo)
}
