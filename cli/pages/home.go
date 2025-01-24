package pages

import (
	"github.com/Hossara/linkin-chat/cli/components"
	"github.com/rivo/tview"
	"os"
	"os/signal"
	"syscall"
)

func HomePage(server, token string) {
	app := tview.NewApplication()
	pages := tview.NewPages()

	pages.AddPage("welcome", components.WelcomePage(app), true, true)

	// Signal channel to handle OS interrupts
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChan
		app.Stop()
	}()

	RenderPage(app, pages)
}

func NavigateTo(pages *tview.Pages, pageName string, page tview.Primitive) {
	if pages.HasPage(pageName) {
		pages.SwitchToPage(pageName)
	} else {
		pages.AddPage(pageName, page, true, true)
	}
}
