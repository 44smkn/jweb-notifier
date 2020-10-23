package presentation

import (
	"jweb-notifier/usecase"
	"net/http"

	"github.com/go-chi/render"
)

func registerUser(w http.ResponseWriter, r *http.Request) {
	data := struct {
		id       string `json:"id"`
		password string `json:"password"`
	}{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, errInvalidRequest(err))
		return
	}

	param := &usecase.RegisterUserParam{
		Id:       data.id,
		Password: data.password,
	}

	if err := usecase.RegisterUser(param); err != nil {
		render.Render(w, r, errInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
	resp := &struct{}{}
	render.Render(w, r, resp)
}
