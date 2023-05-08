package keyboards

import tele "gopkg.in/telebot.v3"

var (
	collectorMenu       = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnShowApplications = tele.Btn{Text: "Показать заявки"}
	BtnApplyApplication = tele.Btn{Text: "Выполнить заявку"}
)

func CollectorKB() *tele.ReplyMarkup {
	collectorMenu.Reply(
		collectorMenu.Row(BtnShowApplications),
		collectorMenu.Row(BtnApplyApplication))
	return collectorMenu
}
