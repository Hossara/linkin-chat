package components

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/viper"
)

func confirmAction(app *tview.Application) func(root *tview.Flex, chatroomName string) {
	return func(root *tview.Flex, chatroomName string) {
		modal := tview.NewModal().
			SetText(fmt.Sprintf("What would you like to do with %s?", chatroomName)).
			AddButtons([]string{"Join", "Remove", "Cancel"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				switch buttonLabel {
				case "Join":
					fmt.Println("Joining", chatroomName)
				case "Remove":
					fmt.Println("Removing", chatroomName)
				}
				app.SetRoot(root, true) // Return to the list
			})

		app.SetRoot(modal, true)
	}
}

func WelcomePage(app *tview.Application) *tview.Flex {
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
		viper.Reset()
		err := viper.WriteConfig()

		if err != nil {
			ErrorModal("Error writing config file", nil)
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

	confirmActionModal := confirmAction(app)

	// Chatrooms list
	chatroomsList := tview.NewList().
		AddItem("Chatroom 1", "[Join] [Remove]", 0, func() {
			fmt.Println("Joining Chatroom 1")
		}).
		AddItem("Chatroom 2", "[Join] [Remove]", 0, func() {
			fmt.Println("Joining Chatroom 1")
		}).
		AddItem("Chatroom 3", "[Join] [Remove]", 0, func() {
			confirmActionModal(root, "Chatroom 3")
		})
	chatroomsList.SetBorder(true).SetTitle("  Your Chatrooms  ")

	// Buttons to create and join chatrooms
	createButton := tview.NewButton("Create New Chatroom").SetSelectedFunc(func() {
		// Action for creating a new chatroom
		fmt.Println("Create new chatroom clicked")
	})
	joinButton := tview.NewButton("Join Another Chatroom").SetSelectedFunc(func() {
		// Action for joining another chatroom
		fmt.Println("Join another chatroom clicked")
	})
	chatroomButtons := tview.NewFlex().SetDirection(tview.FlexRowCSS).
		AddItem(createButton, 0, 1, false).
		AddItem(nil, 1, 0, false).
		AddItem(joinButton, 0, 1, false)

	focusList := []tview.Primitive{
		closeBtn,
		logoutBtn,
		chatroomsList,
		createButton,
		joinButton,
	}

	// Manage focus index
	currentFocus := 2
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab: // Move to the next focusable item
			currentFocus = (currentFocus + 1) % len(focusList)
			app.SetFocus(focusList[currentFocus])
			return nil
		case tcell.KeyBacktab: // Move to the previous focusable item
			currentFocus = (currentFocus - 1 + len(focusList)) % len(focusList)
			app.SetFocus(focusList[currentFocus])
			return nil
		default:
		}
		return event
	})

	// Layout
	root = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(nil, 1, 0, false).
		AddItem(welcomeText, 1, 0, false).
		AddItem(logoutFlex, 2, 0, false).
		AddItem(chatroomsList, 0, 1, true).
		AddItem(chatroomButtons, 3, 0, false)

	return root
}
