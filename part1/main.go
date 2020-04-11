package main

import (
	"github.com/goki/gi/gi"
	"github.com/goki/gi/gimain"
)

func main() {
	gimain.Main(func() {
		mainrun() // calls the function that makes the GUI
	})
}

func mainrun() {
	width := 1024 // pixel sizes of screen
	height := 768 // pixel sizes of screen

	win := gi.NewMainWindow("nameOfYourGame", "Name of Your Game", width, height) // makes the window

	vp := win.WinViewport2D() // makes the vieport in the window
	updt := vp.UpdateStart()  // starts the update loop on the viewport

	mfr := win.SetMainFrame()                   // main frame in the window
	gi.AddNewLabel(mfr, "text", "Hello, World") // add a new piece of text to the main frame called text with text hello world

	vp.UpdateEndNoSig(updt) // end the update loop

	win.StartEventLoop() // starts the event loop that constantly runs
}
