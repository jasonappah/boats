package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/pocketbase/pocketbase/cmd"
	"boats/lib"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	
	pocketBaseApp := lib.InitPocketbase(true)
	pocketBaseApp.Bootstrap()
	serveCmd := cmd.NewServeCommand(pocketBaseApp, true)

	go func() {
		serveCmd.Execute()
		// TODO: need to figure out how to gracefully shutdown the server on wails shutdown
	}()


	// Create application with options
	err := wails.Run(&options.App{
		Title:  "boats",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
