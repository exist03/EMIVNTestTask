package handlers

import (
	"EMIVNTestTask/internal/keyboards"
	"EMIVNTestTask/pkg/models/mysql"
	"database/sql"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
	"log"
	"time"
)

//	case "get_samurai_info": //admin get_samurai_info [samuraiID]
//		if len(command) != 2 {
//			return "Wrong message"
//		}
//		samuraiID := command[1]
//		samuraiModel := mysql.SamuraiModel{DB: db}
//		return samuraiModel.Get(samuraiID)

var (
	onAdditionalShogunState          = InputSG.New("onAdditionalShogunState")
	onAdditionalDaimyoState          = InputSG.New("onAdditionalDaimyoState")
	onAdditionalSamuraiUsernameState = InputSG.New("onAdditionalSamuraiUsernameState")
	onAdditionalSamuraiResState      = InputSG.New("onAdditionalSamuraiResState")
)

func initAdditionalHandlers(manager *fsm.Manager, db *sql.DB) {
	//shogun
	manager.Bind(&keyboards.BtnShogun, onAdditionalState, onAdditionalShogun)
	manager.Bind(tele.OnText, onAdditionalShogunState, onAdditionalShogunRes(db))
	//daimyo
	manager.Bind(&keyboards.BtnDaimyo, onAdditionalState, onAdditionalDaimyo)
	manager.Bind(tele.OnText, onAdditionalDaimyoState, onAdditionalDaimyoRes(db))
	//samurai
	manager.Bind(&keyboards.BtnSamurai, onAdditionalState, onAdditionalSamurai)
	manager.Bind(tele.OnText, onAdditionalSamuraiUsernameState, onAdditionalSamuraiUsername)
	manager.Bind(tele.OnText, onAdditionalSamuraiResState, onAdditionalSamuraiRes(db))
}

func onAdditional(c tele.Context, state fsm.FSMContext) error {
	go state.Set(onAdditionalState)
	return c.Send("Выберите", keyboards.AdminAdditionalKB())
}

// shogun
func onAdditionalShogun(c tele.Context, state fsm.FSMContext) error {
	go state.Set(onAdditionalShogunState)
	return c.Send("Введите username сегуна")
}
func onAdditionalShogunRes(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		defer state.Set(beginAdminState)
		shogunModel := mysql.ShogunModel{DB: db}
		return c.Send(shogunModel.Get(c.Text()), keyboards.AdminKB())
	}
}

// daimyo
func onAdditionalDaimyo(c tele.Context, state fsm.FSMContext) error {
	go state.Set(onAdditionalDaimyoState)
	return c.Send("Введите username дайме")
}
func onAdditionalDaimyoRes(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		defer state.Set(beginAdminState)
		daimyoModel := mysql.DaimyoModel{DB: db}
		return c.Send(daimyoModel.Get(c.Text()), keyboards.AdminKB())
	}
}

// samurai
func onAdditionalSamurai(c tele.Context, state fsm.FSMContext) error {
	go state.Set(onAdditionalSamuraiUsernameState)
	return c.Send("Введите username дайме")
}
func onAdditionalSamuraiUsername(c tele.Context, state fsm.FSMContext) error {
	go state.Set(onAdditionalSamuraiResState)
	go state.Update("samuraiUsername", c.Text())
	return c.Send("Введите дату в формате YYYY-MM-DD")
}
func onAdditionalSamuraiRes(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		defer state.Set(beginAdminState)
		report := mysql.ReportModel{DB: db}
		samurai, err := state.Get("samuraiUsername")
		date, err := time.Parse("2006-01-02", c.Text())
		if err != nil {
			log.Print(err)
			return c.Send("Возникла ошибка")
		}
		return c.Send(report.Samurais(samurai, date), keyboards.AdminKB())
	}
}

//report := mysql.ReportModel{DB: db}
//		time, _ := time2.Parse("2006-01-02", command[1])
//		return report.Samurais(samuraiID, time)
