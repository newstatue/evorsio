package main

import "github.com/wailsapp/wails/v3/pkg/application"

var WindowsWindow = application.WindowsWindow{
	CustomTheme: application.ThemeSettings{
		LightModeActive: &application.WindowTheme{
			TitleBarColour:  application.NewRGBPtr(225, 225, 225),
			TitleTextColour: application.NewRGBPtr(38, 38, 38),
			BorderColour:    application.NewRGBPtr(195, 195, 195),
		},
		LightModeInactive: &application.WindowTheme{
			TitleBarColour:  application.NewRGBPtr(225, 225, 225),
			TitleTextColour: application.NewRGBPtr(38, 38, 38),
			BorderColour:    application.NewRGBPtr(195, 195, 195),
		},
		DarkModeActive: &application.WindowTheme{
			TitleBarColour:  application.NewRGBPtr(32, 32, 32),
			TitleTextColour: application.NewRGBPtr(235, 235, 235),
			BorderColour:    application.NewRGBPtr(55, 55, 55),
		},
		DarkModeInactive: &application.WindowTheme{
			TitleBarColour:  application.NewRGBPtr(32, 32, 32),
			TitleTextColour: application.NewRGBPtr(235, 235, 235),
			BorderColour:    application.NewRGBPtr(55, 55, 55),
		},
	},
}

var MacWindow = application.MacWindow{
	InvisibleTitleBarHeight: 50,
	Backdrop:                application.MacBackdropTranslucent,
	TitleBar:                application.MacTitleBarHiddenInset,
}
