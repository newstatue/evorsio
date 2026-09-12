package main

import "github.com/wailsapp/wails/v3/pkg/application"

var windowOpt = application.WebviewWindowOptions{
	Title:  "Evorsio",
	Width:  1000,
	Height: 618,
	Windows: application.WindowsWindow{
		BackdropType:                      application.Mica,
		DisableFramelessWindowDecorations: false,
		NonClientRegionSupport:            true,
		WebView2CompositionHosting:        true,
	},
	BackgroundType: application.BackgroundTypeTranslucent,
	Frameless:      true,
	URL:            "/",
}
