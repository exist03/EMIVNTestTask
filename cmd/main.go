package main

import (
	"EMIVNTestTask/internal/handlers"
	"EMIVNTestTask/pkg/DB"
	"github.com/joho/godotenv"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	"github.com/vitaliy-ukiru/fsm-telebot/storages/memory"
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
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

	b.Use(middleware.AutoRespond())
	bot := b.Group()
	storage := memory.NewStorage()
	defer storage.Close()
	manager := fsm.NewManager(bot, storage)

	db, err := DB.OpenDB("quest:quest@/EMIVN")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("db is open")
	defer db.Close()
	handlers.InitHandlers(bot, db, manager)
	b.Start()
}
