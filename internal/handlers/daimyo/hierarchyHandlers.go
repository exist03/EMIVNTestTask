package daimyo

import (
	"EMIVNTestTask/internal/handlers/states"
	"EMIVNTestTask/internal/keyboards"
	"EMIVNTestTask/internal/users"
	"EMIVNTestTask/pkg/models/mysql"
	"database/sql"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
	"log"
)

var (
	hierarchyState               = states.InputSG.New("HierarchyState")
	onCreateSamuraiUsernameState = states.InputSG.New("onDaimyoCreatingSamuraiUsernameState")
	onCreateSamuraiNicknameState = states.InputSG.New("onDaimyoCreatingSamuraiNicknameState")
)

func initHierarchyHandlers(manager *fsm.Manager, db *sql.DB) {

	//show
	manager.Bind(&keyboards.BtnShowSamurais, hierarchyState, showSamurais(db))
	//create
	manager.Bind(&keyboards.BtnSamurai, hierarchyState, onCreateSamurai)
	manager.Bind(tele.OnText, onCreateSamuraiUsernameState, onCreateSamuraiUsername)
	manager.Bind(tele.OnText, onCreateSamuraiNicknameState, onCreateSamuraiNickname(db))
}

func showSamurais(db *sql.DB) fsm.Handler {
	samuraiModel := mysql.SamuraiModel{DB: db}
	return func(c tele.Context, state fsm.FSMContext) error {
		res, _ := samuraiModel.GetList(c.Sender().Username)
		for _, v := range res {
			err := c.Send(v)
			if err != nil {
				log.Println(err)
				return c.Send(err)
			}
		}
		return nil
	}
}

func onCreateSamurai(c tele.Context, state fsm.FSMContext) error {
	go state.Set(onCreateSamuraiUsernameState)
	return c.Send("Введите username самурая")
}
func onCreateSamuraiUsername(c tele.Context, state fsm.FSMContext) error {
	go state.Update("samuraiUsername", c.Text())
	go state.Set(onCreateSamuraiNicknameState)
	return c.Send("Введите nickname самурая")
}
func onCreateSamuraiNickname(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		defer state.Set(hierarchyState)
		samuraiUsername, err := state.Get("samuraiUsername")
		samuraiModel := mysql.SamuraiModel{DB: db}
		if err != nil {
			log.Print(err)
			return c.Send("Возникла ошибка")
		}
		samurai := users.Samurai{
			Owner:            c.Sender().Username,
			TelegramUsername: samuraiUsername,
			Nickname:         c.Text(),
		}
		err = samuraiModel.Insert(samurai)
		if err != nil {
			log.Println(err)
			return c.Send("Произошла ошбика", keyboards.DaimyoHierarchyKB())
		}
		return c.Send("Данные записаны", keyboards.DaimyoHierarchyKB())
	}
}
