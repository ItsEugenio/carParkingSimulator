package main

import "concurrent-parking/views"

func main() {
	mainView := views.NewMainView()
	mainView.Run()
}
