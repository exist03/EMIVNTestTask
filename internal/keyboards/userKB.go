package keyboards

import tele "gopkg.in/telebot.v3"

var (
	startMenu    = &tele.ReplyMarkup{ResizeKeyboard: true, RemoveKeyboard: true}
	BtnAdmin     = tele.Btn{Text: "/admin"}
	BtnShogun    = tele.Btn{Text: "/shogun"}
	BtnDaimyo    = tele.Btn{Text: "/daimyo"}
	BtnSamurai   = tele.Btn{Text: "/samurai"}
	BtnCollector = tele.Btn{Text: "/collector"}
	BtnCancel    = tele.Btn{Text: "Назад"}
)

func StartKB() *tele.ReplyMarkup {
	startMenu.Reply(
		startMenu.Row(BtnAdmin, BtnShogun, BtnDaimyo),
		startMenu.Row(BtnSamurai, BtnCollector))
	return startMenu
}
