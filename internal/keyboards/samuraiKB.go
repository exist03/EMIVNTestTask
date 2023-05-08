package keyboards

import tele "gopkg.in/telebot.v3"

var (
	samuraiMenu = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnTurnover = tele.Btn{Text: "Оборот"}
)

func SamuraiKB() *tele.ReplyMarkup {
	samuraiMenu.Reply(
		daimyoMenu.Row(BtnTurnover))
	return samuraiMenu
}
