package urls

import (
	"github.com/gorilla/mux"
	"net/http"
)

func shortURLsServer(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.URL.Path
	fullUrl := GetUrlDBInstance().URLs[shortUrl].(string)

	http.Redirect(w,r,fullUrl,302)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w,r,"https://www.youtube.com/watch?v=dQw4w9WgXcQ",302)
}

func SetUpURLs() http.Handler {
	r := mux.NewRouter()

	for k, _ := range GetUrlDBInstance().URLs {
		r.HandleFunc(k, shortURLsServer).Methods(http.MethodGet)
	}

	r.NotFoundHandler = http.HandlerFunc(NotFound)

	return r
}
