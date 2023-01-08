package main

import (
	askiied "ascii-art-web/makeAscii"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	fmt.Println("Link -->   " + "http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Data struct {
	InitialWord  string
	Str          string
	Format       string
	ErrorCode    int
	ErrorMessage string
}

func errorHandler(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	_, err := template.ParseFiles("templates/error.html")
	if err != nil {
		fmt.Fprintf(w, "Internal server error!")
		return
	}
	if status == http.StatusNotFound {
		renderTemplate(w, "error", Data{ErrorCode: status, ErrorMessage: "Page not found!"})
	}
	if status == http.StatusBadRequest {
		renderTemplate(w, "error", Data{ErrorCode: status, ErrorMessage: "Bad request!"})
	}
	if status == http.StatusInternalServerError {
		renderTemplate(w, "error", Data{ErrorCode: status, ErrorMessage: "Internal Server Error!"})
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" && r.Method == "GET" {
		renderTemplate(w, "home", Data{InitialWord: "", Format: "standard"})
		return
	}
	if r.URL.Path == "/ascii-art" && r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		str := r.PostForm["str"]
		if len(str[0]) == 0 {
			renderTemplate(w, "home", Data{InitialWord: "", Format: "standard"})
			return
		}
		format := r.PostForm["format"]
		res, err := askiied.Askiied(str[0], format[0])
		if err != 0 {
			errorHandler(w, err)
			return
		}
		renderTemplate(w, "home", Data{InitialWord: str[0], Str: res, Format: format[0]})
		return
	}
	if r.URL.Path == "/ascii-art" || r.URL.Path == "/" {
		errorHandler(w, 400)
		return
	}
	errorHandler(w, 404)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data Data) {
	t, err := template.ParseFiles("templates/" + tmpl + ".html")
	if err != nil {
		errorHandler(w, 500)
		return
	}
	t.Execute(w, data)
}
