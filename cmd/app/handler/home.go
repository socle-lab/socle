package handler

import (
	"my_app_name/cmd/app/view"
	"my_app_name/cmd/app/viewmodel"
	"net/http"

	"github.com/socle-lab/render"
)

func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {

	var err error

	vm := viewmodel.NewIndexViewModel("Modules de recherche", nil)

	// err = h.render(w, r, render.PageOptions{
	// 	Data: views.Home(vm),
	// })

	err = h.Render(w, r, render.PageOptions{
		ComponentFunc: view.Home,
		ViewModel:     &vm,
		Data:          nil,
	})
	if err != nil {
		h.Log("error", err)
	}
}
