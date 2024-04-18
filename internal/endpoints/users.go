package endpoints

import (
	"github.com/go-chi/render"
	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	render.Status(r, 200)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) GetUserByCFP(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {}
