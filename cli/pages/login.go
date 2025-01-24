package pages

import (
	"fmt"
	"github.com/Hossara/linkin-chat/cli/components"
	"github.com/Hossara/linkin-chat/cli/constants"
	"github.com/Hossara/linkin-chat/cli/services"
	"github.com/rivo/tview"
	"github.com/spf13/viper"
)

func LoginPage(server string) {
	app := tview.NewApplication()

	textView := tview.NewTextView().
		SetText(constants.Logo).
		SetDynamicColors(true)

	var form *tview.Form
	var flex *tview.Flex

	form = tview.NewForm().
		AddInputField("Username", "", 20, tview.InputFieldMaxLength(74), nil).
		AddInputField("Password", "", 20, tview.InputFieldMaxLength(74), nil).
		AddButton("Login", func() {
			usernameInput := form.GetFormItemByLabel("Username").(*tview.InputField)
			username := usernameInput.GetText()

			password := form.GetFormItemByLabel("Password").(*tview.InputField).GetText()

			showModal := func(message string) {
				app.SetRoot(components.ErrorModal(message,
					func(buttonIndex int, buttonLabel string) {
						app.SetRoot(flex, true).SetFocus(usernameInput)
					}),
					false,
				)
			}

			if len(username) > 4 && len(username) < 74 && len(password) > 4 && len(password) < 74 {
				token, err := services.Login(username, password, server)

				if err != nil {
					showModal(err.Error())
					return
				}

				viper.Set("login.username", username)
				viper.Set("login.password", password)
				err = viper.WriteConfig()

				if err != nil {
					showModal(fmt.Sprintf("Cannot save files to configuration file! Please restart the app. Error: %v", err))
					return
				}

				viper.Set("login.token", token)

				app.Stop()
			} else {
				showModal("Username and password must be between 4 and 74 characters")
			}
		})

	flex = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(textView, 7, 1, false).
		AddItem(form, 35, 1, true)

	RenderPage(app, flex)
}
