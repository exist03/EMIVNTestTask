package daimyo

import (
	"EMIVNTestTask/internal/handlers/states"
	"EMIVNTestTask/internal/keyboards"
	"EMIVNTestTask/pkg/models/mysql"
	"database/sql"
	"fmt"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
	"log"
)

var (
	reportState              = states.InputSG.New("ReportState")
	onChosePeriodReportState = states.InputSG.New("onChosePeriodReportState")
	onBeginPeriodReportState = states.InputSG.New("onBeginPeriodReportState")
	onEndPeriodReportState   = states.InputSG.New("onEndPeriodReportState")
)

func initReportHandlers(manager *fsm.Manager, db *sql.DB) {

	//данные за смену
	manager.Bind(&keyboards.BtnDataShift, reportState, showSamurais(db))
	//запросить отчет
	manager.Bind(&keyboards.BtnGetReport, reportState, onChosePeriodReport)

	//за смену
	manager.Bind(&keyboards.BtnReportShift, onChosePeriodReportState, reportShift(db))
	//за период
	manager.Bind(&keyboards.BtnReportPeriod, onChosePeriodReportState, onBeginPeriodReport)
	manager.Bind(tele.OnText, onBeginPeriodReportState, onEndPeriodReport)
	manager.Bind(tele.OnText, onEndPeriodReportState, reportPeriod(db))
}

func showCards(db *sql.DB) fsm.Handler {
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

func onChosePeriodReport(c tele.Context, state fsm.FSMContext) error {
	go state.Set(onChosePeriodReportState)
	return c.Send("выберите", keyboards.DaimyoReportPeriodKB())
}

func reportShift(db *sql.DB) fsm.Handler {
	daimyoModel := mysql.DaimyoModel{DB: db}
	return func(c tele.Context, state fsm.FSMContext) error {
		res, _ := daimyoModel.GetReportShift(c.Sender().Username)
		for k, v := range res {
			if v.Sber != v.SberCheck || v.Tink != v.TinkCheck {
				amountContoller := v.SberCheck + v.TinkCheck
				amountSamurai := v.Sber + v.Tink
				c.Send(fmt.Sprintf("%s\nВсего\n +"+
					"%.0f / %.0f / %.0f\n\n "+
					"сбер\n%.0f / %.0f / %.0f\n\n"+
					"тинькофф\n%.0f / %.0f / %.0f", k, amountSamurai, amountContoller, amountSamurai-amountContoller, v.Sber, v.SberCheck, v.Sber-v.SberCheck, v.Tink, v.TinkCheck, v.Tink-v.TinkCheck))
			} else {
				c.Send(fmt.Sprintf("расхождение по %s отсутствуют", k))
			}
		}
		return nil
	}
}

func onBeginPeriodReport(c tele.Context, state fsm.FSMContext) error {
	go state.Set(onBeginPeriodReportState)
	return c.Send("Введите дату начало периода в формате гггг-мм-дд")
}
func onEndPeriodReport(c tele.Context, state fsm.FSMContext) error {
	go state.Set(onEndPeriodReportState)
	state.Update("dateBegin", c.Text())
	return c.Send("Введите дату конца периода в формате гггг-мм-дд")
}

func reportPeriod(db *sql.DB) fsm.Handler {
	daimyoModel := mysql.DaimyoModel{DB: db}
	return func(c tele.Context, state fsm.FSMContext) error {
		dateBegin, _ := state.Get("dateBegin")
		dateEnd := c.Text()
		res, _ := daimyoModel.GetReportPeriod(c.Sender().Username, dateBegin, dateEnd)
		result := 0.0
		for _, v := range res {
			result += v.SberCheck - v.Sber
		}
		c.Send(fmt.Sprintf("%s\nОборот: %.0f\n0.0015 -> %.0f", c.Sender().Username, result, result*0.0015))
		for k, v := range res {
			c.Send(fmt.Sprintf("%s\n"+
				"%.0f / %.0f ", k, v.Sber, v.SberCheck))
		}
		go state.Set(BeginDaimyoState)
		return c.Send("Конец отчета", keyboards.DaimyoKB())
	}
}

//func onCreateSamurai(c tele.Context, state fsm.FSMContext) error {
//	go state.Set(onCreateSamuraiUsernameState)
//	return c.Send("Введите username самурая")
//}
//func onCreateSamuraiUsername(c tele.Context, state fsm.FSMContext) error {
//	go state.Update("samuraiUsername", c.Text())
//	go state.Set(onCreateSamuraiNicknameState)
//	return c.Send("Введите nickname самурая")
//}
//func onCreateSamuraiNickname(db *sql.DB) fsm.Handler {
//	return func(c tele.Context, state fsm.FSMContext) error {
//		defer state.Set(hierarchyState)
//		samuraiUsername, err := state.Get("samuraiUsername")
//		samuraiModel := mysql.SamuraiModel{DB: db}
//		if err != nil {
//			log.Print(err)
//			return c.Send("Возникла ошибка")
//		}
//		samurai := users.Samurai{
//			Owner:            c.Sender().Username,
//			TelegramUsername: samuraiUsername,
//			Nickname:         c.Text(),
//		}
//		err = samuraiModel.Insert(samurai)
//		if err != nil {
//			log.Println(err)
//			return c.Send("Произошла ошбика", keyboards.DaimyoHierarchyKB())
//		}
//		return c.Send("Данные записаны", keyboards.DaimyoHierarchyKB())
//	}
//}
