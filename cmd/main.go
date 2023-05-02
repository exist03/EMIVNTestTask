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

func main() {
	err := godotenv.Load()
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
	handlers.InitHandlers(b, db)
	b.Start()
}
