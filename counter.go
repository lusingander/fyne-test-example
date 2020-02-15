package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type counter struct {
	*widget.Label
	*widget.Button
	count int
}

func newCounter() *counter {
	c := &counter{}
	c.Label = widget.NewLabel(fmt.Sprintf("Count: %d", c.count))
	c.Button = widget.NewButton("Click!", c.countUp)
	return c
}

func (c *counter) countUp() {
	c.count++
	c.Label.SetText(fmt.Sprintf("Count: %d", c.count))
}

type counterWindow struct {
	fyne.Window
	*counter
}

func newCounterWindow(app fyne.App) *counterWindow {
	counter := newCounter()
	w := app.NewWindow("Counter")
	w.Resize(fyne.NewSize(200, 30))
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			counter.Button, counter.Label,
		),
	)
	return &counterWindow{w, counter}
}
