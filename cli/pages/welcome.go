package pages

import (
	"fmt"
	"github.com/Hossara/linkin-chat/cli/components"
	"github.com/Hossara/linkin-chat/cli/services"
	"github.com/Hossara/linkin-chat/cli/types"
	"github.com/Hossara/linkin-chat/pkg/utils"
	"github.com/rivo/tview"
	"github.com/spf13/viper"
	"log"
)

func confirmJoinAction(app *tview.Application, pages *tview.Pages) func(name string, code string) {
	return func(name string, code string) {
		modal := tview.NewModal().
			SetText(fmt.Sprintf("What would you like to do with [black]%s?", name)).
			AddButtons([]string{"Join", "Remove", "Cancel"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				switch buttonLabel {
				case "Join":
					viper.Set("chat.code", code)
					app.SetRoot(pages, true)
					RemoveAndNavigate(pages, app, "chat")
				case "Remove":
					fmt.Println("Removing", name)
				default:
					app.SetRoot(pages, true)
				}
			})

		app.SetRoot(modal, true)
	}
}

func getChats() ([]types.ResponseChatRoom, error) {
	chats, err := services.GetAllChats()

	if err != nil {
		return nil, err
	}

	return chats, nil
}

func WelcomePage(app *tview.Application, pages *tview.Pages) tview.Primitive {
	username := viper.GetString("login.username")
	var root *tview.Flex

	welcomeText := tview.NewTextView().
		SetText("Welcome to Linkin Chat").
		SetDynamicColors(true)

	// Logged-in user text and logout button
	userInfo := tview.NewTextView().
		SetText(fmt.Sprintf("You are logged in as [green]@%s", username)).
		SetDynamicColors(true)

	logoutBtn := tview.NewButton("Logout").SetSelectedFunc(func() {
		viper.Set("login.username", "")
		viper.Set("login.password", "")
		err := viper.WriteConfig()

		if err != nil {
			app.SetRoot(components.ErrorModal("Error writing config file",
				func(buttonIndex int, buttonLabel string) {
					app.SetRoot(root, true)
				}), true)
			return
		}
		app.Stop()
	})

	closeBtn := tview.NewButton("Exit").SetSelectedFunc(func() {
		app.Stop()
	})

	logoutFlex := tview.NewFlex().
		SetDirection(tview.FlexColumnCSS).
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRowCSS).
			AddItem(userInfo, 0, 1, false).
			AddItem(tview.NewFlex().
				AddItem(closeBtn, 10, 0, false).
				AddItem(nil, 2, 0, false).
				AddItem(logoutBtn, 10, 0, false),
				22, 0, false),
			1, 0, false).
		AddItem(nil, 1, 0, false)

	confirmActionModal := confirmJoinAction(app, pages)

	// Chatrooms list
	chatroomsList := tview.NewList()
	var chatList []types.ResponseChatRoom

	var err error
	chatList, err = getChats()

	if err != nil {
		viper.Set("login.username", "")
		viper.Set("login.password", "")
		log.Fatal(err)
	}

	for _, room := range chatList {
		chatroomsList.AddItem(room.Title, "[Join] [Remove]", 0, func() {
			confirmActionModal(room.Title, room.Code)
		})
	}

	chatroomsList.SetBorder(true).SetTitle(
		utils.IfThenElse(len(chatList) == 0, "  You have not chat!  ", "  Your Chatrooms  ").(string),
	)

	// Buttons to create and join chatrooms
	createButton := tview.NewButton("Create New Chatroom").SetSelectedFunc(func() {
		RemoveAndNavigate(pages, app, "create_new_chat")
	})

	joinButton := tview.NewButton("Join Another Chatroom").SetSelectedFunc(func() {
		RemoveAndNavigate(pages, app, "join_chat")
	})

	chatroomButtons := tview.NewFlex().SetDirection(tview.FlexRowCSS).
		AddItem(createButton, 0, 1, false).
		AddItem(nil, 1, 0, false).
		AddItem(joinButton, 0, 1, false)

	// Layout
	root = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(nil, 1, 0, false).
		AddItem(welcomeText, 1, 0, false).
		AddItem(logoutFlex, 2, 0, false).
		AddItem(chatroomsList, 0, 1, true).
		AddItem(chatroomButtons, 3, 0, false)

	root.SetInputCapture(CaptureFocus(app, 2, []tview.Primitive{
		closeBtn,
		logoutBtn,
		chatroomsList,
		createButton,
		joinButton,
	}))

	return root
}
