package users

import "fmt"

type Daimyo struct {
	CardList         []Card             //by nickname
	SamuraiList      map[string]Samurai //same
	Owner            string
	TelegramUsername string
	Nickname         string
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

//func (d *Daimyo) CheckSamuraiTurnOver(samurai Samurai) float64 {
//	return samurai.TurnOver
//}

func (d *Daimyo) SetOwner(owner string) {
	d.Owner = owner
}

func (d Daimyo) String() string {
	return fmt.Sprintf("TG Username: %s\nNickname: %s\nOwner: %s\n\n", d.TelegramUsername, d.Nickname, d.Owner)
}
