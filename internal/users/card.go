package users

import "fmt"

type Card struct {
	Owner       Daimyo
	BankInfo    string
	LimitInfo   float64
	FundBalance float64
}

func (c *Card) String() string {
	return fmt.Sprintf("Owner: %s\nBank-emitter: %s\nRemain funds: %f\nDaily limits: %.2f", c.Owner, c.BankInfo, c.FundBalance, c.LimitInfo)
}

func (c *Card) SetOwner(daimyo *Daimyo) {
	c.Owner = *daimyo
}
