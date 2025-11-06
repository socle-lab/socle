package main

import (
	"log"
	"my_app_name/cmd/web/handler"
	"my_app_name/cmd/web/middleware"
	"my_app_name/internal"
	"my_app_name/internal/store"
)

func initApp() *application {
	core, err := internal.Boot("web")
	if err != nil {
		log.Fatal(err)
	}

	myMiddleware := &middleware.Middleware{
		Core: core,
	}

	myHandlers := &handler.Handler{
		Core: core,
	}

	app := &application{
		Core:       core,
		Handler:    myHandlers,
		Middleware: myMiddleware,
	}

	app.Core.Routes.Mount("/", app.routes())
	store := store.NewStore(app.Core.DB.Pool)

	app.Middleware.Store = store
	myHandlers.Store = store
	return app
}
