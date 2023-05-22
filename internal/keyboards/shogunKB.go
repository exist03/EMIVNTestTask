package keyboards

import tele "gopkg.in/telebot.v3"

var (
	shogunMenu            = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnShowDaimyos        = tele.Btn{Text: "Показать дайме"}
	BtnShowDaimyoSamurais = tele.Btn{Text: "Показать самураев"}
)

func ShogunKB() *tele.ReplyMarkup {
	shogunMenu.Reply(
		shogunMenu.Row(BtnGetReport),
		shogunMenu.Row(BtnShowDaimyos, BtnShowDaimyoSamurais),
		shogunMenu.Row(BtnCreateCard, BtnConnectCard),
		shogunMenu.Row(BtnCancel))
	return shogunMenu
}

func ShogunReportPeriodKB() *tele.ReplyMarkup {
	daimyoMenu.Reply(
		daimyoMenu.Row(BtnReportShift),
		daimyoMenu.Row(BtnReportPeriod),
		daimyoMenu.Row(BtnCancel))
	return daimyoMenu
}
