package keyboards

import tele "gopkg.in/telebot.v3"

var (
	shogunMenu     = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnShowDaimyos = shogunMenu.Text("shogun daimyos")
)

func ShogunKB() *tele.ReplyMarkup {
	shogunMenu.Reply(
		shogunMenu.Row(BtnShowDaimyos))
	return shogunMenu
}
