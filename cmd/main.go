package main

import (
	"EMIVNTestTask/internal/users"
	"fmt"
)

func main() {
	samura1 := users.CreateSamurai("Samurai1_u", "Samurai1_n")
	samura2 := users.CreateSamurai("Samurai2_u", "Samurai2_n")
	sho := users.Shogun{
		TelegramUsername: "shogun_u",
		Nickname:         "shogun_n",
	}
	card1 := sho.CreateCard("infoCard1")
	card2 := sho.CreateCard("infoCard2")
	samuraMap := make(map[string]users.Samurai)
	samuraMap[samura1.Nickname] = samura1
	samuraMap[samura2.Nickname] = samura2
	dai1 := users.Daimyo{
		CardList:         []users.Card{card1, card2},
		SamuraiList:      samuraMap,
		Owner:            sho,
		TelegramUsername: "daimyo1_u",
		Nickname:         "daimyo1_n",
	}
	dai2 := users.Daimyo{
		CardList:         []users.Card{card1, card2},
		SamuraiList:      samuraMap,
		Owner:            sho,
		TelegramUsername: "daimyo1_u",
		Nickname:         "daimyo1_n",
	}
	daiMap := make(map[string]users.Daimyo)
	daiMap[dai1.Nickname] = dai1
	daiMap[dai2.Nickname] = dai2
	sho.DaimyoList = daiMap
	temp := sho.CheckDaimyoInfo(&dai2)
	fmt.Print(temp)
}
