package admin

import (
	"EMIVNTestTask/internal/handlers/shogun"
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
	onCreateShogunUsernameState    = states.InputSG.New("onCreateShogunUsernameState")
	onCreateShogunNicknameState    = states.InputSG.New("onCreateShogunNicknameState")
	onCreateDaimyoUsernameState    = states.InputSG.New("onCreateDaimyoUsernameState")
	onCreateDaimyoNicknameState    = states.InputSG.New("onCreateDaimyoNicknameState")
	onCreateSamuraiUsernameState   = states.InputSG.New("onCreateSamuraiUsernameState")
	onCreateSamuraiNicknameState   = states.InputSG.New("onCreateSamuraiNicknameState")
	onCreateCollectorUsernameState = states.InputSG.New("onCreateCollectorUsernameState")
	onCreateCollectorNicknameState = states.InputSG.New("onCreateCollectorNicknameState")
)

func initCreatingHandlers(manager *fsm.Manager, db *sql.DB) {
	//create card
	manager.Bind(&keyboards.BtnCreateCard, onCreateState, shogun.OnCreateCard)
	manager.Bind(tele.OnText, shogun.OnInputCardCreateState, shogun.OnInputCardCreate)
	manager.Bind(tele.OnText, shogun.OnInputBankState, shogun.OnInputBank)
	manager.Bind(tele.OnText, shogun.OnInputLimitState, shogun.OnInputLimit(db))
	//create shogun
	manager.Bind(&keyboards.BtnShogun, onCreateState, onCreateShogun)
	manager.Bind(tele.OnText, onCreateShogunUsernameState, onCreateShogunUsername)
	manager.Bind(tele.OnText, onCreateShogunNicknameState, onCreateShogunNickname(db))
	//create daimyo
	manager.Bind(&keyboards.BtnDaimyo, onCreateState, onCreateDaimyo)
	manager.Bind(tele.OnText, onCreateDaimyoUsernameState, onCreateDaimyoUsername)
	manager.Bind(tele.OnText, onCreateDaimyoNicknameState, onCreateDaimyonNickname(db))
	//create samurai
	manager.Bind(&keyboards.BtnSamurai, onCreateState, onCreateSamurai)
	manager.Bind(tele.OnText, onCreateSamuraiUsernameState, onCreateSamuraiUsername)
	manager.Bind(tele.OnText, onCreateSamuraiNicknameState, onCreateSamuraiNickname(db))
	//collector
	manager.Bind(&keyboards.BtnCollector, onCreateState, onCreateCollector)
	manager.Bind(tele.OnText, onCreateCollectorUsernameState, onCreateCollectorUsername)
	manager.Bind(tele.OnText, onCreateCollectorNicknameState, onCreateCollectorNickname(db))
}
func onCreating(c tele.Context, state fsm.FSMContext) error {
	go state.Set(onCreateState)
	return c.Send("Выберите", keyboards.AdminCreateKB())
}

// shogun
func onCreateShogun(c tele.Context, state fsm.FSMContext) error {
	go state.Set(onCreateShogunUsernameState)
	return c.Send("Введите username сегуна")
}
func onCreateShogunUsername(c tele.Context, state fsm.FSMContext) error {
	go state.Update("shogunUsername", c.Text())
	go state.Set(onCreateShogunNicknameState)
	return c.Send("Введите nickname сегуна")
}
func onCreateShogunNickname(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		defer state.Set(beginAdminState)
		shogunUsername, err := state.Get("shogunUsername")
		shogunModel := mysql.ShogunModel{DB: db}
		if err != nil {
			log.Print(err)
			return c.Send("Возникла ошибка")
		}
		shogun := users.Shogun{
			TelegramUsername: shogunUsername,
			Nickname:         c.Text(),
		}
		err = shogunModel.Insert(shogun)
		if err != nil {
			return c.Send("Произошла ошбика", keyboards.AdminKB())
		}
		return c.Send("Сегун создан", keyboards.AdminKB())
	}
}

// daimyo
func onCreateDaimyo(c tele.Context, state fsm.FSMContext) error {
	go state.Set(onCreateDaimyoUsernameState)
	return c.Send("Введите username дайме")
}
func onCreateDaimyoUsername(c tele.Context, state fsm.FSMContext) error {
	go state.Update("daimyoUsername", c.Text())
	go state.Set(onCreateDaimyoNicknameState)
	return c.Send("Введите nickname дайме")
}
func onCreateDaimyonNickname(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		defer state.Set(beginAdminState)
		daimyoUsername, err := state.Get("daimyoUsername")
		daimyoModel := mysql.DaimyoModel{DB: db}
		if err != nil {
			log.Print(err)
			return c.Send("Возникла ошибка")
		}
		daimyo := users.Daimyo{
			Owner:            "1",
			TelegramUsername: daimyoUsername,
			Nickname:         c.Text(),
		}
		err = daimyoModel.Insert(daimyo)
		if err != nil {
			log.Println(err)
			return c.Send("Произошла ошбика", keyboards.AdminKB())
		}
		return c.Send("Дайме создан", keyboards.AdminKB())
	}
}

// samurai
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
		defer state.Set(beginAdminState)
		samuraiUsername, err := state.Get("samuraiUsername")
		samuraiModel := mysql.SamuraiModel{DB: db}
		if err != nil {
			log.Print(err)
			return c.Send("Возникла ошибка")
		}
		samurai := users.Samurai{
			Owner:            "1",
			TelegramUsername: samuraiUsername,
			Nickname:         c.Text(),
		}
		err = samuraiModel.Insert(samurai)
		if err != nil {
			log.Println(err)
			return c.Send("Произошла ошбика", keyboards.AdminKB())
		}
		return c.Send("Самурай создан", keyboards.AdminKB())
	}
}

// collector
func onCreateCollector(c tele.Context, state fsm.FSMContext) error {
	go state.Set(onCreateCollectorUsernameState)
	return c.Send("Введите username инкассатора")
}
func onCreateCollectorUsername(c tele.Context, state fsm.FSMContext) error {
	go state.Update("collectorUsername", c.Text())
	go state.Set(onCreateCollectorNicknameState)
	return c.Send("Введите nickname инкассатора")
}
func onCreateCollectorNickname(db *sql.DB) fsm.Handler {
	return func(c tele.Context, state fsm.FSMContext) error {
		defer state.Set(beginAdminState)
		collectorUsername, err := state.Get("collectorUsername")
		collectorModel := mysql.CollectorModel{DB: db}
		if err != nil {
			log.Print(err)
			return c.Send("Возникла ошибка")
		}
		collector := users.Collector{
			TelegramUsername: collectorUsername,
			Nickname:         c.Text(),
		}
		err = collectorModel.Insert(collector)
		if err != nil {
			return c.Send("Произошла ошбика", keyboards.AdminKB())
		}
		return c.Send("Инкассатор создан", keyboards.AdminKB())
	}
}
