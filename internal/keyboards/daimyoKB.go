package keyboards

import tele "gopkg.in/telebot.v3"

var (
	daimyoMenu      = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnShowSamurais = shogunMenu.Text("daimyo samurais")
	BtnShowCards    = shogunMenu.Text("daimyo cards")
)

func DaimyoKB() *tele.ReplyMarkup {
	daimyoMenu.Reply(
		daimyoMenu.Row(BtnShowSamurais),
		daimyoMenu.Row(BtnShowCards))
	return daimyoMenu
}
