package main

import (
	"context"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/cmd"
	"boats/lib"
)

// App struct
type App struct {
	ctx context.Context
	pocketbase *pocketbase.PocketBase
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	a.pocketbase = lib.InitPocketbase(true)
	a.pocketbase.Bootstrap()
	
	go func() {
		serveCmd := cmd.NewServeCommand(a.pocketbase, true)
		serveCmd.Execute()
	}()
}

func (a *App) shutdown(ctx context.Context) {
	// TODO: this still doesn't run when i expect it to. why
	println("Shutting down FN")
	a.pocketbase.OnTerminate().Trigger(&core.TerminateEvent{
		App: a.pocketbase.App,
		IsRestart: false,
	})
	result := a.pocketbase.App.ResetBootstrapState() 
	println(result) 
	println("Shutdown complete")
}
