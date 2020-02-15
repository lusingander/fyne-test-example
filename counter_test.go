package main

import (
	"testing"

	"fyne.io/fyne/test"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	c := newCounter()

	test.Tap(c.Button)
	test.Tap(c.Button)
	test.Tap(c.Button)

	assert.Equal(t, 3, c.count)
}

func TestNewCounterWindow(t *testing.T) {
	w := newCounterWindow(test.NewApp())

	assert.NotNil(t, w.Window)
	assert.NotNil(t, w.counter)
}
