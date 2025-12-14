package main

import (
	"log"
	"my_app_name/cmd/app/handler"
	"my_app_name/cmd/app/middleware"
	"my_app_name/internal"
	"my_app_name/internal/store/repository"

	baseHandler "github.com/socle-lab/pkg/http/handler"
	baseMiddleware "github.com/socle-lab/pkg/http/middleware"
)

func initApp() *application {
	core, err := internal.Boot("app")
	if err != nil {
		log.Fatal(err)
	}

	myMiddleware := &middleware.Middleware{
		Middleware: baseMiddleware.Middleware{
			Core: core,
		},
	}

	myHandlers := &handler.Handler{
		Handler: baseHandler.Handler{
			Core: core,
		},
	}

	app := &application{
		Core:       core,
		Handler:    myHandlers,
		Middleware: myMiddleware,
	}

	store := repository.NewStore(app.Core.DB.Pool)

	app.Middleware.Store = store
	myHandlers.Store = store

	app.Core.Routes.Mount("/", app.routes())
	return app
}
