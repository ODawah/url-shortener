package handlers

import (
	"encoding/json"
	"github.com/ODawah/url-shortener/kafka"
	"github.com/ODawah/url-shortener/models"
	"github.com/ODawah/url-shortener/persistence"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

func EncodeURL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var url models.URL
		err := json.NewDecoder(r.Body).Decode(&url)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		if url.Original == "" {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		err = url.CreateURL(persistence.DB)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		render.JSON(w, r, url)
	}
}

func Redirect() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortURL := chi.URLParam(r, "url")

		if len(shortURL) != 7 {
			render.Status(r, 400)
			render.JSON(w, r, map[string]string{"error": "encoded url must be 7 characters length"})
			return
		}
		url, err := models.GetOriginalURL(shortURL, persistence.DB)
		if err != nil {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		go func() {
			RequestData := ExtractRequestData(r)
			if RequestData.IP == "" {
				log.Println("Error from Extract Request functions")
				return
			}
			err = kafka.ProduceMessage(url.Short, RequestData)
			if err != nil {
				log.Println("error Producing message")
				log.Println(err)
				return
			}
			log.Println("Producer Sent the message")
		}()

		render.JSON(w, r, url)
		return
	}
}

func ExtractRequestData(r *http.Request) models.RequestData {
	ip := r.RemoteAddr

	host := r.Host

	userAgent := r.UserAgent()

	return models.RequestData{IP: ip, Host: host, Browser: userAgent}
}
