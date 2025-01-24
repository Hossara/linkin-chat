package pages

import (
	"fmt"
	"github.com/Hossara/linkin-chat/cli/components"
	"github.com/Hossara/linkin-chat/cli/services"
	"github.com/rivo/tview"
)

func createChat(title string) (string, error) {
	code, err := services.CreateNewChat(title)

	if err != nil {
		return "", err
	}

	return code, nil
}

func CreateNewChatPage(app *tview.Application, pages *tview.Pages) tview.Primitive {
	loadingText := tview.NewTextView().
		SetText("Loading...").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	// Create new chat form
	form := tview.NewForm()
	var chatTitle string

	// Title input field
	form.AddInputField("Title", "", 30, func(text string, lastChar rune) bool {
		return len(text) <= 255
	}, func(text string) {
		chatTitle = text
	})

	// Create button
	form.AddButton("Create", func() {
		if len(chatTitle) < 3 || len(chatTitle) > 255 {
			// Show an error message if the title is too short
			form.GetFormItem(0).(*tview.InputField).SetText("")

			app.SetRoot(components.ErrorModal("Title must be between 2 and 255 characters.",
				func(buttonIndex int, buttonLabel string) {
					app.SetRoot(form, true)
				}), true)
			return
		}
		// Show loading text and trigger function
		app.SetRoot(loadingText, true)

		code, err := createChat(chatTitle)

		if err != nil {
			app.SetRoot(components.ErrorModal(err.Error(),
				func(buttonIndex int, buttonLabel string) {
					app.SetRoot(pages, true)
					NavigateTo(pages, "welcome")
				}), true)
			return
		}

		app.SetRoot(components.ErrorModal(
			fmt.Sprintf("Chat created successfully!\nShare this code with your friends: %s", code),
			func(buttonIndex int, buttonLabel string) {
				app.SetRoot(pages, true)
			}), true)

		RemoveAndNavigate(pages, app, "welcome")
	})

	// Cancel button to exit
	form.AddButton("Cancel", func() {
		NavigateTo(pages, "welcome")
	})

	form.SetBorder(true).
		SetTitle("  Create New Chat  ").
		SetTitleAlign(tview.AlignLeft)

	return form
}
