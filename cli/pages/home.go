package pages

import (
	"github.com/rivo/tview"
	"os"
	"os/signal"
	"syscall"
)

type Page func(app *tview.Application, pages *tview.Pages) tview.Primitive

var pages = make(map[string]Page)

func GetPages() map[string]Page {
	pages["welcome"] = WelcomePage
	pages["create_new_chat"] = CreateNewChatPage
	pages["join_chat"] = JoinChatPage
	pages["chat"] = ChatPage

	return pages
}

func GetPage(title string) Page {
	return GetPages()[title]
}

func HomePage() {
	app := tview.NewApplication()
	pages := tview.NewPages()

	for name, page := range GetPages() {
		if name != "chat" {
			pages.AddPage(name, page(app, pages), true, name == "welcome")
		}
	}

	// Signal channel to handle OS interrupts
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChan
		os.Exit(0)
	}()

	RenderPage(app, pages)
}
