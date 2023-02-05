package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"acsii-art-web/functions"
)

type Errors struct {
	Status  int
	Message string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/ascii-art", result)
	mux.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	log.Println("Starting a web server on  http://127.0.0.1:8080")
	http.ListenAndServe(":8080", mux)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// 404
		Errorhandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		w.Header().Set("Get", http.MethodGet)
		// 405
		Errorhandler(w, http.StatusMethodNotAllowed)
		return
	}
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		Errorhandler(w, 500)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		Errorhandler(w, 500)
	}
}

func result(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		// 404
		Errorhandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// 405
		Errorhandler(w, http.StatusMethodNotAllowed)
		return
	}
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		Errorhandler(w, 500)
		return
	}
	message := r.FormValue("text")
	bannerName := r.FormValue("font")
	for _, v := range message {
		if (v < 32 || v > 126) && !(v == '\r' || v == '\n') {
			Errorhandler(w, 400)
			return
		}
	}
	if bannerName == "" {
		Errorhandler(w, 400)
		return
	}
	banner, err2 := functions.MakeMapOfBanner(bannerName)
	if err2 != nil {
		Errorhandler(w, http.StatusInternalServerError)
		return
	}
	word, err1 := functions.MakeGraphicWord(banner, message)
	if err1 != nil {
		Errorhandler(w, 400)
		return
	}
	err = t.Execute(w, word)
	if err != nil {
		Errorhandler(w, 500)
		return
	}
}

func Errorhandler(w http.ResponseWriter, status int) {
	var ErrResult Errors
	ErrResult.Status = status
	ErrResult.Message = http.StatusText(status)
	html, err := template.ParseFiles("templates/error.html")
	if err != nil {
		w.WriteHeader(status)
		fmt.Fprintf(w, "%s %d", ErrResult.Message, status)
		return
	}
	w.WriteHeader(status)
	err = html.Execute(w, ErrResult)
	if err != nil {
		w.WriteHeader(status)
		fmt.Fprintf(w, "%s %d", ErrResult.Message, status)
		return
	}
}
