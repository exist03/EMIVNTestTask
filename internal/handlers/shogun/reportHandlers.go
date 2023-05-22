package shogun

import (
	"EMIVNTestTask/internal/handlers/states"
	"EMIVNTestTask/internal/keyboards"
	"EMIVNTestTask/pkg/models/mysql"
	"database/sql"
	"fmt"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
	"log"
	"math"
	"time"
)

var (
	reportState              = states.InputSG.New("ReportState")
	onChosePeriodReportState = states.InputSG.New("onChosePeriodReportState")
	onBeginPeriodReportState = states.InputSG.New("onBeginPeriodReportState")
	onEndPeriodReportState   = states.InputSG.New("onEndPeriodReportState")
)

func initReportHandlers(manager *fsm.Manager, db *sql.DB) {
	//за смену
	manager.Bind(&keyboards.BtnReportShift, reportState, reportShift(db))
	//за период
	manager.Bind(&keyboards.BtnReportPeriod, reportState, onBeginPeriodReport)
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
	return c.Send("выберите", keyboards.ShogunReportPeriodKB())
}

func reportShift(db *sql.DB) fsm.Handler {
	shogunModel := mysql.ShogunModel{DB: db}
	return func(c tele.Context, state fsm.FSMContext) error {
		res, _ := shogunModel.GetReportShift(c.Sender().Username)
		response := fmt.Sprintf("Отчет за %s\n\n", time.Now().Add(-24*time.Hour).Format("2006-01-02"))
		for k, v := range res {
			if v.TurnoverSamurai+v.AdditionSum+v.CardSum != v.RemainingFunds {
				difference := v.TurnoverSamurai + v.AdditionSum + v.CardSum - v.RemainingFunds
				response += fmt.Sprintf("%s: -%.0f\n", k, math.Abs(difference))
				//c.Send(fmt.Sprintf("%s: -%.0f", k, math.Abs(difference)))
			} else {
				response += fmt.Sprintf("%s: OK\n", k)
				//c.Send(fmt.Sprintf("%s: OK", k))
			}
		}
		go state.Set(BeginShogunState)
		return c.Send(response, keyboards.ShogunKB())
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
		go state.Set(BeginShogunState)
		return c.Send("Конец отчета", keyboards.DaimyoKB())
	}
}
