package pages

import (
	"fmt"
	"github.com/Hossara/linkin-chat/cli/types"
	"github.com/Hossara/linkin-chat/pkg/utils"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/viper"
	"time"
)

func getUsersList() []types.ChatRoomUser {
	return []types.ChatRoomUser{
		{
			Username: "hossara",
			Role:     0,
		},
		{
			Username: "mmd",
			Role:     1,
		},
		{
			Username: "asqar",
			Role:     1,
		},
	}
}

func getMessagesList() []types.Message {
	return []types.Message{
		{
			Sender: types.ChatRoomUser{
				Username: "hossara",
				Role:     0,
			},
			Content:   "Hello World!",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "mmd",
				Role:     1,
			},
			Content:   "Hiiiiiiiiiiiiiiiiiiiiiiii",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
		{
			Sender: types.ChatRoomUser{
				Username: "asar",
				Role:     1,
			},
			Content:   "dfdfdfd",
			CreatedAt: time.Now(),
		},
	}
}

func ChatPage(app *tview.Application, pages *tview.Pages) tview.Primitive {
	chatCode := viper.GetString("chat.code")
	//username := viper.GetString("login.username")

	var users []types.ChatRoomUser
	var messages []types.Message

	if chatCode == "" {
		RemoveAndNavigate(pages, app, "welcome")
	}

	root := tview.NewFlex()

	// User List on the Right
	userList := tview.NewList()
	userList.
		ShowSecondaryText(false).
		SetBorder(true).SetTitle("Users")
	userList.SetSelectedBackgroundColor(tcell.ColorBlack)
	userList.SetSelectedTextColor(tcell.ColorWhite)

	// Message List
	messageList := tview.NewList()
	messageList.SetBorder(true).SetTitle("  Messages  ")
	messageList.SetSelectedBackgroundColor(tcell.NewHexColor(0x121212))

	getUsers := func() {
		users = getUsersList()

		for _, user := range users {
			userList.AddItem(fmt.Sprintf(
				"%s%s", user.Username, utils.IfThenElse(
					user.Role == 0, " [green](Admin)", "",
				),
			), "", 0, nil)
		}
	}

	getMessages := func() {
		messages = getMessagesList()

		for _, message := range messages {
			messageList.AddItem(fmt.Sprintf(
				" %s%s [white]%s     [#d3d3d3](%s)",
				utils.IfThenElse(message.Sender.Role == 0, "[red]", "[green]"),
				message.Sender.Username,
				message.Content,
				message.CreatedAt.Format(time.RFC3339),
			), "", 0, nil)
		}

		messageList.SetCurrentItem(messageList.GetItemCount() - 1)
	}

	getUsers()
	getMessages()

	// Chat Room on the Left
	chatRoom := tview.NewFlex()

	chatRoom.
		SetDirection(tview.FlexRow).
		SetBorder(true).
		SetTitle("  Chat Room  ")

	// Username Display
	usernameTxt := tview.NewTextView().
		SetText(fmt.Sprintf(" Username: [green]@%s", viper.GetString("login.username"))).
		SetDynamicColors(true)

	exitBtn := tview.NewButton("Exit Chat")
	deleteBtn := tview.NewButton("Delete Chat")
	copyJoinBtn := tview.NewButton("Copy Join Code")

	header := tview.NewFlex().
		SetDirection(tview.FlexColumnCSS).
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRowCSS).
			AddItem(usernameTxt, 0, 1, false).
			AddItem(tview.NewFlex().
				AddItem(exitBtn, 13, 0, false).
				AddItem(nil, 2, 0, false).
				AddItem(deleteBtn, 15, 0, false).
				AddItem(nil, 2, 0, false).
				AddItem(copyJoinBtn, 18, 0, false).
				AddItem(nil, 2, 0, false),
				52, 0, false),
			1, 0, false).
		AddItem(nil, 1, 0, false)

	// Message Input and Send Button
	messageInput := tview.NewInputField()
	messageInput.
		SetFieldTextColor(tcell.ColorWhite).
		SetFieldBackgroundColor(tcell.ColorBlack)

	messageInputPhStyle := tcell.Style{}.Background(tcell.ColorBlack)

	messageInput.
		SetPlaceholder("Type a message...").
		SetPlaceholderStyle(messageInputPhStyle).
		SetTitleAlign(tview.AlignLeft).SetBorderPadding(1, 1, 2, 2)

	sendButton := tview.NewButton("Send")

	inputRow := tview.NewFlex().
		AddItem(messageInput, 0, 1, true).
		AddItem(sendButton, 20, 0, false)

	// Assemble Chat Room
	chatRoom.
		AddItem(nil, 1, 0, false).
		AddItem(header, 1, 0, false).
		AddItem(nil, 1, 0, false).
		AddItem(messageList, 0, 1, false).
		AddItem(inputRow, 3, 0, false)

	root.SetInputCapture(CaptureFocus(app, 0, []tview.Primitive{
		exitBtn,
		deleteBtn,
		copyJoinBtn,
		messageList,
		messageInput,
		sendButton,
	}))

	// Main Layout
	root.
		AddItem(chatRoom, 0, 3, false).
		AddItem(userList, 30, 1, false)

	return root
}
