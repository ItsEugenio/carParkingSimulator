package views

import (
	"concurrent-parking/scenes"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type MainView struct{}

func NewMainView() *MainView {
	return &MainView{}
}

func (v *MainView) Run() {
	myApp := app.New()
	window := myApp.NewWindow("Parking")
	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(400, 800))

	mainScene := scenes.NewMainScene(window)
	mainScene.Show()
	go mainScene.Run()
	window.ShowAndRun()
}
