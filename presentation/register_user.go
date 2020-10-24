package presentation

import (
	"errors"
	"jweb-notifier/usecase"
	"net/http"

	"github.com/go-chi/render"
)

func registerUser(w http.ResponseWriter, r *http.Request) {
	data := &registerUserRequest{}

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
	resp := &registerUserResponse{}
	render.Render(w, r, resp)
}

type registerUserRequest struct {
	id       string `json:"id"`
	password string `json:"password"`
}

func (u *registerUserRequest) Bind(r *http.Request) error {
	if u.id == "" {
		return errors.New("missing required id fields.")
	}
	if u.password == "" {
		return errors.New("missing required password fields.")
	}

	return nil
}

type registerUserResponse struct{}

func (u *registerUserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
