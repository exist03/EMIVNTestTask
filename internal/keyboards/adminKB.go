package keyboards

import (
	tele "gopkg.in/telebot.v3"
)

var (
	adminMenu  = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnCreate  = tele.Btn{Text: "Создать"}
	BtnConnect = tele.Btn{Text: "Привязать"}
	//create
	BtnCreateCard = tele.Btn{Text: "Создать карту💳"}
	//connect
	BtnAdditional     = tele.Btn{Text: "Подробнее"}
	BtnConnectCard    = tele.Btn{Text: "Привязять карту💳"}
	BtnConnectDaimyo  = tele.Btn{Text: "Дайме"}
	BtnConnectSamurai = tele.Btn{Text: "Самурая🥷"}
	//additional
)

func AdminKB() *tele.ReplyMarkup {
	adminMenu.Reply(
		adminMenu.Row(BtnCreate, BtnConnect, BtnAdditional),
		adminMenu.Row(BtnCancel))
	return adminMenu
}
func AdminCreateKB() *tele.ReplyMarkup {
	adminMenu.Reply(
		adminMenu.Row(BtnCreateCard, BtnShogun, BtnDaimyo),
		adminMenu.Row(BtnSamurai, BtnCollector),
		adminMenu.Row(BtnCancel))
	return adminMenu
}
func AdminConnectKB() *tele.ReplyMarkup {
	adminMenu.Reply(
		adminMenu.Row(BtnConnectCard, BtnDaimyo, BtnSamurai),
		adminMenu.Row(BtnCancel))
	return adminMenu
}

func AdminAdditionalKB() *tele.ReplyMarkup {
	adminMenu.Reply(
		adminMenu.Row(BtnConnectCard, BtnConnectDaimyo, BtnConnectSamurai),
		adminMenu.Row(BtnCancel))
	return adminMenu
}
