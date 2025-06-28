package handler

import (
	"fmt"
	"net/http"
	"url-shortener/api"
	"url-shortener/internal/service"
)

func CreateHandler(svc *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		originalURL := r.FormValue("url")
		if originalURL == "" {
			http.Error(w, "Missing URL parameter", http.StatusBadRequest)
			return
		}

		resp, err := svc.CreateLink(r.Context(), &api.CreateLinkRequest{
			OriginalUrl: originalURL,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, resp.ShortUrl)
	}
}
