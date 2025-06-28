package handler

import (
	"net/http"
	"url-shortener/api"
	"url-shortener/internal/service"
)

func GetHandler(svc *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		shortURL := r.URL.Path[len("/get/"):]
		if len(shortURL) != 10 {
			http.Error(w, "Invalid short URL", http.StatusBadRequest)
			return
		}

		resp, err := svc.GetLink(r.Context(), &api.GetLinkRequest{
			ShortUrl: shortURL,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Redirect(w, r, resp.OriginalUrl, http.StatusFound)
	}
}
