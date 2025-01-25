package pages

import (
	"github.com/Hossara/linkin-chat/cli/components"
	"github.com/rivo/tview"
	"github.com/spf13/viper"
)

func JoinChatPage(app *tview.Application, pages *tview.Pages) tview.Primitive {
	// Create new chat form
	form := tview.NewForm()
	var code string

	// Title input field
	form.AddInputField("Code", "", 30, func(text string, lastChar rune) bool {
		return len(text) <= 8
	}, func(text string) {
		code = text
	})

	// Create button
	form.AddButton("Join", func() {
		if len(code) != 8 {
			// Show an error message if the title is too short
			form.GetFormItem(0).(*tview.InputField).SetText("")

			app.SetRoot(components.ErrorModal("Code must be 8 characters long",
				func(buttonIndex int, buttonLabel string) {
					app.SetRoot(form, true)
					form.SetFocus(0)
				}), true)
			return
		}

		app.SetRoot(pages, true)
		viper.Set("chat.code", code)
		RemoveAndNavigate(pages, app, "chat")
	})

	// Cancel button to exit
	form.AddButton("Back To Home", func() {
		NavigateTo(pages, "welcome")
	})

	form.SetBorder(true).
		SetTitle("  Join Chat  ").
		SetTitleAlign(tview.AlignLeft)

	return form
}
