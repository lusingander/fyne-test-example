package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	mainApp := app.New()
	w := mainApp.NewWindow("Sample")
	w.Resize(fyne.NewSize(200, 30))

	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			widget.NewButton("Counter", func() { newCounterWindow(mainApp).Show() }),
		),
	)
	w.ShowAndRun()
}
