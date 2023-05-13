package handlers

import (
	"encoding/json"
	"fmt"
	"iam/pkg/httphelpers"
	"iam/src/core/ports"
	"net/http"
)

type UsersHandler struct {
	usersService ports.UsersService
}

func NewUsersHandler(usersService ports.UsersService) *UsersHandler {
	return &UsersHandler{
		usersService: usersService,
	}
}

func (h *UsersHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var args ports.RegisterArgs
	err := json.NewDecoder(r.Body).Decode(&args)
	fmt.Println("===========", args)
	if err != nil {
		httphelpers.WriteError(http.StatusInternalServerError, "error", err.Error())(w, r)
		return
	}
	answer, err := h.usersService.Register(args)
	if err != nil {
		httphelpers.WriteError(http.StatusInternalServerError, "error", err.Error())(w, r)
		return
	}
	httphelpers.WriteSuccess(http.StatusOK, "User created successfully", answer)(w, r)
}

func (h *UsersHandler) WhoAmI(w http.ResponseWriter, r *http.Request) {
	httphelpers.WriteSuccess(http.StatusOK, "Foobar", struct{}{})(w, r)
}
