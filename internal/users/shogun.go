package users

import (
	"fmt"
)

type Shogun struct {
	DaimyoList       map[string]Daimyo
	TelegramUsername string
	Nickname         string
}

func (s *Shogun) ShowDaimyoList() {
	for _, v := range s.DaimyoList {
		fmt.Print(v.Nickname)
	}
}

func (s *Shogun) CheckDaimyoSamurais(daimyo *Daimyo) map[string]Samurai {
	return daimyo.SamuraiList
}

func (s *Shogun) CreateCard(info string) Card {
	return Card{BankInfo: info, LimitInfo: 2000000, Balance: 2000000}
}

func (s *Shogun) SetLimit(card *Card, limit float64) {
	card.LimitInfo = limit
}

func (s *Shogun) SetCardOwner(card *Card, owner string) {
	card.Owner = owner
}

func (s *Shogun) CheckDaimyoInfo(daimyo *Daimyo) string {
	var mainTurnOver float64
	for _, v := range daimyo.SamuraiList {
		mainTurnOver += v.TurnOver
	}
	//SAMURAI LIST IF NEED
	//sList := ""
	//for _, v := range daimyo.SamuraiList {
	//	sList += fmt.Sprintf("Nickname: %s\nTG Username: %s\nTurnover: %f\nOwner: %s\n", v.Nickname, v.TelegramUsername, v.TurnOver, v.Owner.Nickname)
	//}
	res := fmt.Sprintf("TG Username: %s\nUsername: %s\nCards: %v\nTurnover: %f", daimyo.TelegramUsername, daimyo.Nickname, daimyo.CardList, mainTurnOver)
	return res
}

func (s *Shogun) CheckSamuraiInfo(daimyo *Daimyo, nickName string) float64 {
	return daimyo.CheckSamuraiTurnOver(daimyo.SamuraiList[nickName])
}
