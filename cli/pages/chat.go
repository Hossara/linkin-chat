package pages

import (
	"github.com/rivo/tview"
	"github.com/spf13/viper"
)

func ChatPage(app *tview.Application, pages *tview.Pages) tview.Primitive {

	root := tview.NewFlex().SetDirection(tview.FlexRow)

	chatCode := viper.GetString("chat.code")

	if chatCode == "" {
		RemoveAndNavigate(pages, app, "welcome")
	}

	return root.AddItem(tview.NewTextView().SetText(chatCode), 0, 1, false)
}
