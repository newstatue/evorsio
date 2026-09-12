package main

import "github.com/wailsapp/wails/v3/pkg/application"

var windowOpt = application.WebviewWindowOptions{
	Title:  "Evorsio",
	Width:  1000,
	Height: 618,
	Mac: application.MacWindow{
		InvisibleTitleBarHeight: 50,
		Backdrop:                application.MacBackdropTranslucent,
		TitleBar:                application.MacTitleBarHiddenInset,
	},
	Frameless: true,
	URL:       "/",
}
