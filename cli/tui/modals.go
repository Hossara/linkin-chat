package tui

import "github.com/rivo/tview"

func ErrorModal(message string, done func(buttonIndex int, buttonLabel string)) *tview.Modal {
	return tview.NewModal().
		SetText(message).
		AddButtons([]string{"OK"}).
		SetDoneFunc(done)
}
