package keyboards

import tele "gopkg.in/telebot.v3"

var (
	samuraiMenu = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnTurnover = tele.Btn{Text: "Ввести данные на конец смены"}
)

func SamuraiKB() *tele.ReplyMarkup {
	samuraiMenu.Reply(
		daimyoMenu.Row(BtnTurnover),
		daimyoMenu.Row(BtnCancel))
	return samuraiMenu
}
