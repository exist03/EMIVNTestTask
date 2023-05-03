package keyboards

import tele "gopkg.in/telebot.v3"

var (
	startMenu    = &tele.ReplyMarkup{ResizeKeyboard: true, RemoveKeyboard: true}
	BtnAdmin     = tele.Btn{Text: "Админ👑"}
	BtnShogun    = tele.Btn{Text: "Сегун"}
	BtnDaimyo    = tele.Btn{Text: "Дайме"}
	BtnSamurai   = tele.Btn{Text: "Самурай🥷"}
	BtnCollector = tele.Btn{Text: "Инкассатор💵"}
	BtnCancel    = tele.Btn{Text: "❌ Отмена"}
)

func StartKB() *tele.ReplyMarkup {
	startMenu.Reply(
		startMenu.Row(BtnAdmin, BtnShogun, BtnDaimyo),
		startMenu.Row(BtnSamurai, BtnCollector))
	return startMenu
}
