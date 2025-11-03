package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"

	shortenurl "github.com/cjosue15/url-shortener/internal"
	"github.com/cjosue15/url-shortener/internal/db"
)

type PageData struct {
	Url string
}

type IndexData struct {
	Data  []shortenurl.ShortUrl
	Error bool
}

type Data struct {
	Url string `json:"url"`
}

var pagaData = PageData{
	Url: "",
}

func shortenURLHandler(w http.ResponseWriter, r *http.Request, s *shortenurl.ShortenUrl) {
	url := r.FormValue("url")
	short, err := s.CreateShortUrl(url)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating short URL"))
		return
	}

	pagaData.Url = "http://localhost:8080/" + *short

	http.Redirect(w, r, "/shorten", http.StatusSeeOther)
}

func main() {
	database := db.Connect()
	database.AutoMigrate(&shortenurl.ShortUrl{})

	s := shortenurl.NewShortenUrl(database)

	tmpl := template.Must(template.New("").ParseGlob("./web/templates/*"))
	router := http.NewServeMux()
	fs := http.FileServer(http.Dir("./web/styles/"))
	router.Handle("/web/styles/", http.StripPrefix("/web/styles/", fs))

	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		data, err := s.GetAllUrls()
		tmpl.ExecuteTemplate(w, "index.html", IndexData{
			Data:  data,
			Error: err != nil,
		})
	})

	router.HandleFunc("GET /{code}", func(w http.ResponseWriter, r *http.Request) {
		code := r.PathValue("code")
		url, err := s.GetOriginalUrl(code)

		if err != nil {

			if err.Error() == "URL not found" {
				http.Redirect(w, r, "/not-found", http.StatusSeeOther)

			}

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Intenal Server Error"))
			return
		}

		http.Redirect(w, r, *url, http.StatusSeeOther)
	})

	router.HandleFunc("GET /shorten", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "shorten.html", PageData{
			Url: pagaData.Url,
		})
	})

	router.HandleFunc("GET /not-found", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "404.html", nil)
	})

	router.HandleFunc("POST /api/shorten", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("API shorten URL requested")
		shortenURLHandler(w, r, s)
	})

	srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Starting website at localhost:8080")

	err := srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Println("An error occured:", err)
	}
}
