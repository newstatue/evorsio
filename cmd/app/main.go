package main

import (
	"log"
	"time"

	"github.com/newstatue/evorsio"
	"github.com/newstatue/evorsio/internal/drive"
	"github.com/wailsapp/wails/v3/pkg/application"
)

func init() {
	application.RegisterEvent[string]("time")
}

func main() {
	app := application.New(application.Options{
		Name:        "app",
		Description: "A demo of using raw HTML & CSS",
		Services: []application.Service{
			application.NewService(&drive.GreetService{}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(evorsio.Assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:  "Window 1",
		Width:  1000,
		Height: 618,
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(6, 7, 15),
		URL:              "/",
	})

	go func() {
		for {
			now := time.Now().Format(time.RFC1123)
			app.Event.Emit("time", now)
			time.Sleep(time.Second)
		}
	}()

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
