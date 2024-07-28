package main

import (
	"boats/lib"
	"context"
	"io"
	"net/http"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/cmd"
	"github.com/pocketbase/pocketbase/core"
)

// App struct
type App struct {
	ctx        context.Context
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

func (a *App) Shutdown(ctx context.Context) {
	// TODO: this still doesn't run when i expect it to. why
	println("Shutting down FN")
	a.pocketbase.OnTerminate().Trigger(&core.TerminateEvent{
		App:       a.pocketbase.App,
		IsRestart: false,
	})
	result := a.pocketbase.App.ResetBootstrapState()
	println(result)
	println("Shutdown complete")
}

type ScholarshipURL struct {
	Id      string `db:"id"`
	Url     string `db:"url"`
	created string `db:"created"`
	updated string `db:"updated"`
}

func (a *App) FetchScholarshipHTML() error {
	urls := []ScholarshipURL{}

	err := a.pocketbase.Dao().DB().
		NewQuery("SELECT * FROM scholarship_urls EXCEPT SELECT * FROM scholarship_urls JOIN scholarship_html ON scholarship_urls.id = scholarship_html.scholarship_id").All(&urls)

	if err != nil {
		println(err.Error())
		return err
	}

	httpClient := http.Client{}

	for _, url := range urls {
		resp, err := httpClient.Get(url.Url)
		if err != nil {
			println(err.Error())
			return err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			println(err.Error())
			continue
		}

		// TODO: later we can use a html parser to extract the needed text to save space and make prompts smaller too. we can also minify the html - we should keep it as html so that we still have knowledge of text hierarchy
		_, err = a.pocketbase.Dao().DB().
			NewQuery("UPDATE scholarship_html SET html = {:body} WHERE scholarship_id = {:scholarshipId}").Bind(dbx.Params{
			"body":          string(body),
			"scholarshipId": url.Id,
		}).Execute()

		if err != nil {
			println(err.Error())
			return err
		}
	}

	return nil
}
