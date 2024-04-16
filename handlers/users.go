package handlers

import (
	"encoding/json"
	"github.com/ODawah/url-shortener/models"
	"github.com/ODawah/url-shortener/persistence"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"strconv"
)

func CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		log.Println(user)
		err = user.CreateUser(persistence.DB)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		render.JSON(w, r, user)
	}
}

func FindUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			render.Status(r, 400)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		if id <= 0 {
			render.Status(r, 400)
			render.JSON(w, r, map[string]string{"error": "id must be greater than 0"})
			return
		}
		user, err := models.FindUserByID(id, persistence.DB)
		if err != nil {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		render.JSON(w, r, user)
	}
}
