package main

import (
	"my_app_name/cmd/api/handler"
	"my_app_name/cmd/api/middleware"
	"os"
	"os/signal"
	"sync"
	"syscall"

	socle "github.com/socle-lab/core"
)

type application struct {
	Core       *socle.Socle
	Handler    *handler.Handler
	Middleware *middleware.Middleware
	wg         sync.WaitGroup
}

func main() {
	app := initApp()
	go app.listenForShutdown()
	err := app.Core.ListenAndServe()
	app.Core.Log.ErrorLog.Println(err)
}

func (a *application) shutdown() {
	// put any clean up tasks here

	// block until the WaitGroup is empty
	a.wg.Wait()
}

func (app *application) listenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit

	app.Core.Log.InfoLog.Println("Received signal", s.String())
	app.shutdown()

	os.Exit(0)
}
