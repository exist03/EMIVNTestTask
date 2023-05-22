package keyboards

import tele "gopkg.in/telebot.v3"

var (
	daimyoMenu           = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnReport            = tele.Btn{Text: "Отчет"}
	BtnGetReport         = tele.Btn{Text: "Запросить отчет"}
	BtnReportShift       = tele.Btn{Text: "За смену"}
	BtnReportPeriod      = tele.Btn{Text: "За период"}
	BtnDataShift         = tele.Btn{Text: "Ввести данные за смену"}
	BtnHierarchy         = tele.Btn{Text: "Иерархия"}
	BtnShowSamurais      = tele.Btn{Text: "Список подчиненных"}
	BtnCardLimit         = tele.Btn{Text: "Лимит по карте"}
	BtnShowCards         = tele.Btn{Text: "Карты💳"}
	BtnRemainingFunds    = tele.Btn{Text: "Ввести остаток на карте"}
	BtnCreateApplication = tele.Btn{Text: "Запросить пополнение"}
)

func DaimyoKB() *tele.ReplyMarkup {
	daimyoMenu.Reply(
		daimyoMenu.Row(BtnHierarchy, BtnCardLimit),
		daimyoMenu.Row(BtnReport, BtnCreateApplication),
		daimyoMenu.Row(BtnCancel))
	return daimyoMenu
}
func DaimyoReportKB() *tele.ReplyMarkup {
	daimyoMenu.Reply(
		daimyoMenu.Row(BtnGetReport, BtnDataShift),
		daimyoMenu.Row(BtnCancel))
	return daimyoMenu
}
func DaimyoReportPeriodKB() *tele.ReplyMarkup {
	daimyoMenu.Reply(
		daimyoMenu.Row(BtnReportShift),
		daimyoMenu.Row(BtnReportPeriod),
		daimyoMenu.Row(BtnCancel))
	return daimyoMenu
}
func DaimyoHierarchyKB() *tele.ReplyMarkup {
	daimyoMenu.Reply(
		daimyoMenu.Row(BtnShowSamurais),
		daimyoMenu.Row(BtnSamurai),
		daimyoMenu.Row(BtnCancel))
	return daimyoMenu
}
