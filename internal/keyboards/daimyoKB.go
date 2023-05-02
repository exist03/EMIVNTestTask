package keyboards

import tele "gopkg.in/telebot.v3"

var (
	daimyoMenu           = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnShowSamurais      = tele.Btn{Text: "Подчиненные"}
	BtnShowCards         = tele.Btn{Text: "Карты💳"}
	BtnRemainingFunds    = tele.Btn{Text: "Ввести остаток на карте"}
	BtnCreateApplication = tele.Btn{Text: "Заявка на пополнение"}
)

//daimyo set [cardID] [balance] - Остаток на карте
//- daimyo application [cardID] [value] - Создание заявки на пополнение карты[cardID] до суммы [value]

func DaimyoKB() *tele.ReplyMarkup {
	daimyoMenu.Reply(
		daimyoMenu.Row(BtnShowSamurais, BtnShowCards),
		daimyoMenu.Row(BtnRemainingFunds, BtnCreateApplication))
	return daimyoMenu
}
