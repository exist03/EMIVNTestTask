package users

import "fmt"

type Daimyo struct {
	CardList         []Card             //by nickname
	SamuraiList      map[string]Samurai //same
	Owner            Shogun
	TelegramUsername string
	Nickname         string
}

func (d *Daimyo) FillRemainingFunds(card *Card, sum float64) {
	card.FundBalance = sum
}

func (d *Daimyo) ShowSamuraiList() {
	for _, v := range d.SamuraiList {
		fmt.Print(v.Nickname)
	}
}

func (d *Daimyo) ShowCardList() {
	for _, v := range d.CardList {
		fmt.Print(v)
	}
}

//TODO
// rebuild
//func (d *Daimyo) CreateApplication(cardList []Card) []Card {
//	return cardList
//}

func (d *Daimyo) CheckSamuraiTurnOver(samurai Samurai) float64 {
	return samurai.TurnOver
}

func (d *Daimyo) SetOwner(shogun *Shogun) {
	d.Owner = *shogun
}
