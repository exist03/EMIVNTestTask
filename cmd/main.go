package main

import (
	"EMIVNTestTask/internal/handlers"
	"EMIVNTestTask/pkg/DB"
	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
)

// func main() {
// samura1 := users.CreateSamurai("Samurai1_u", "Samurai1_n")
// samura2 := users.CreateSamurai("Samurai2_u", "Samurai2_n")
//
//	sho := users.Shogun{
//		TelegramUsername: "shogun_u",
//		Nickname:         "shogun_n",
//	}
//
// card1 := sho.CreateCard("infoCard1")
// card2 := sho.CreateCard("infoCard2")
// samuraMap := make(map[string]users.Samurai)
// samuraMap[samura1.Nickname] = samura1
// samuraMap[samura2.Nickname] = samura2
//
//	dai1 := users.Daimyo{
//		CardList:         []users.Card{card1, card2},
//		SamuraiList:      samuraMap,
//		Owner:            sho,
//		TelegramUsername: "daimyo1_u",
//		Nickname:         "daimyo1_n",
//	}
//
//	dai2 := users.Daimyo{
//		CardList:         []users.Card{card1, card2},
//		SamuraiList:      samuraMap,
//		Owner:            sho,
//		TelegramUsername: "daimyo1_u",
//		Nickname:         "daimyo1_n",
//	}
//
// daiMap := make(map[string]users.Daimyo)
// daiMap[dai1.Nickname] = dai1
// daiMap[dai2.Nickname] = dai2
// sho.DaimyoList = daiMap
// temp := sho.CheckDaimyoInfo(&dai2)
// fmt.Print(temp)
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	pref := tele.Settings{
		Token: os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{
			Timeout: 5 * time.Second,
		},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		panic("can`t create bot")
	}
	db, err := DB.OpenDB("quest:quest@/EMIVN")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("db is open")
	defer db.Close()
	handlers.InitHanders(b, db)
	//m := mysql.CardModel{DB: db}
	//b.Handle(tele.OnText, func(c tele.Context) error {
	//	temp, err := m.GetList("nickname")
	//	if err != nil {
	//		return err
	//	}
	//	return c.Send(temp)
	//})

	b.Start()
}
