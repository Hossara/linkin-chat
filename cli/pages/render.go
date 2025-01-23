package pages

import (
	"github.com/rivo/tview"
)

func RenderPage(app *tview.Application, view tview.Primitive) {
	if err := app.SetRoot(view, true).Run(); err != nil {
		panic(err)
	}
}
