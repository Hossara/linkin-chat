package pages

import (
	"fmt"
	"github.com/Hossara/linkin-chat/cli/constants"
	"github.com/Hossara/linkin-chat/cli/services"
	"github.com/Hossara/linkin-chat/cli/tui"
	"github.com/rivo/tview"
)

func RegisterPage(username, password, server string) {
	app := tview.NewApplication()

	textView := tview.NewTextView().
		SetText(constants.Logo).
		SetDynamicColors(true)

	var form *tview.Form
	var flex *tview.Flex

	form = tview.NewForm().
		AddInputField("Username", username, 20, tview.InputFieldMaxLength(74), nil).
		AddInputField("Password", password, 20, tview.InputFieldMaxLength(74), nil).
		AddInputField("First Name", username, 20, tview.InputFieldMaxLength(74), nil).
		AddInputField("Last Name", password, 20, tview.InputFieldMaxLength(74), nil).
		AddButton("Register", func() {
			usernameInput := form.GetFormItemByLabel("Username").(*tview.InputField)

			username = usernameInput.GetText()
			password := form.GetFormItemByLabel("Password").(*tview.InputField).GetText()
			firstName := form.GetFormItemByLabel("First Name").(*tview.InputField).GetText()
			lastName := form.GetFormItemByLabel("Last Name").(*tview.InputField).GetText()

			showModal := func(message string) {
				app.SetRoot(tui.ErrorModal(message,
					func(buttonIndex int, buttonLabel string) {
						app.SetRoot(flex, true).SetFocus(usernameInput)
					}),
					false,
				)
			}

			if len(firstName) > 100 || len(lastName) > 100 || len(firstName) < 2 || len(lastName) < 2 {
				showModal("FirstName and LastName must be between 2 and 100 characters")
				return
			}

			if len(username) > 74 || len(username) > 74 || len(password) < 4 || len(password) < 4 {
				showModal("Username and password must be between 4 and 74 characters")
				return
			}

			token, err := services.Register(username, password, firstName, lastName, server)

			if err != nil {
				showModal(err.Error())
				return
			}

			if token == "" {
				showModal("Register uncompleted: No token returned!")
				return
			}

			app.SetRoot(tui.ErrorModal(
				fmt.Sprintf("User @%s registered successfuly! Now try login to system using join command. Use --help flag for more info.", username),
				func(buttonIndex int, buttonLabel string) {
					app.Stop()
				}),
				false,
			)
		})

	flex = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(textView, 7, 1, false).
		AddItem(form, 35, 1, true)

	RenderPage(app, flex)
}
