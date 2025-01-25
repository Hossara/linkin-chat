package pages

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func RenderPage(app *tview.Application, view tview.Primitive) {
	if err := app.SetRoot(view, true).Run(); err != nil {
		panic(err)
	}
}

func NavigateTo(pages *tview.Pages, pageName string) {
	if pages.HasPage(pageName) {
		pages.SwitchToPage(pageName)
	}
}

func RemoveAndNavigate(pages *tview.Pages, app *tview.Application, name string) {
	if pages.HasPage(name) {
		pages.RemovePage(name)
	}

	pages.AddAndSwitchToPage(name, GetPage(name)(app, pages), true)
}

type CaptureType func(event *tcell.EventKey) *tcell.EventKey

func CaptureFocus(app *tview.Application, currentFocus int, focusList []tview.Primitive) CaptureType {
	return func(event *tcell.EventKey) *tcell.EventKey {
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
	}
}
