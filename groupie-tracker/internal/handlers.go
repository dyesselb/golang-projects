package internal

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)



func(h *Router) Home(w http.ResponseWriter, r *http.Request) {
	allArtists, err := GetArtists()
	if err != nil {
		h.Errorhandler(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	if r.URL.Path != root {
		h.Errorhandler(w, NotFound, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		h.Errorhandler(w, MethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}
	temp, err := template.New("").ParseFiles("templates/index.html", "templates/artist.html")
	if err != nil {
		h.Errorhandler(w, InternalServerError, http.StatusInternalServerError)
		return
	}

	err = temp.ExecuteTemplate(w, "index", allArtists)
	if err != nil {
		h.Errorhandler(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	log.Printf("%d %s\n", http.StatusOK, OK)
}

func (h *Router) Search(w http.ResponseWriter, r *http.Request) {
	allArtists, err := GetArtists()
	if err != nil {
		h.Errorhandler(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodGet {
		h.Errorhandler(w, MethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}
	text := r.FormValue("text")
	if len(text) == 0 {
		h.Errorhandler(w, BadRequest, http.StatusBadRequest)
		return
	}
	result, err := FindData(strings.TrimSpace(text), allArtists)
	if err != nil {
		h.Errorhandler(w, BadRequest, http.StatusBadRequest)
		return
	}
	html, err := template.New("").ParseFiles("templates/index.html", "templates/artist.html")
	if err != nil {
		h.Errorhandler(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	err = html.ExecuteTemplate(w, "index", result)
	if err != nil {
		h.Errorhandler(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	log.Printf("%d %s\n", http.StatusOK, OK)
}

func FindData(text string, allData []Artist) ([]Artist, error) {
	var result []Artist

	text = strings.ToLower(text)
	for _, v := range allData {
		if strings.Contains(strings.ToLower(v.Name), text) {
			result = append(result, v)
			continue
		}
		if strings.Contains(v.FirstAlbum, text) {
			if Check(v.Id, result) {
				result = append(result, v)
				continue
			}
		}
		if strings.Contains(strconv.Itoa(v.CreationDate), text) {
			if Check(v.Id, result) {
				result = append(result, v)
				continue
			}
		}
		for _, member := range v.Members {
			if strings.Contains(strings.ToLower(member), text) {
				if Check(v.Id, result) {
					result = append(result, v)
				}
			}
		}
		for key := range v.Relations {
			if strings.Contains(strings.ToLower(key), text) {
				if Check(v.Id, result) {
					result = append(result, v)
				}
			}
		}
	}
	if result == nil {
		myError := errors.New(BadRequest)
		return nil, myError
	}
	
	return result, nil
}

func Check(id int, Data []Artist) bool {
	for _, m := range Data {
		if m.Id == id {
			return false
		}
	}
	return true
}

func (h *Router) trackerHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != trackerURL {
		h.Errorhandler(w, NotFound, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		h.Errorhandler(w, MethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	artistLoc, err := GetArtistLocation(id)
	if err != nil {
		h.Errorhandler(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	artistData, err := GetOneArtist(id)
	if err != nil {
		h.Errorhandler(w, NotFound, http.StatusNotFound)
		return
	}
	if artistData.Id == 0 {
		h.Errorhandler(w, NotFound, http.StatusNotFound)
		return
	}
	result := Result{
		Singer:   artistData,
		Relation: artistLoc,
		Text:     "",
		Type:     "",
	}
	if len(artistData.Members) == 1 {
		result.Text = fmt.Sprintf("started proffesional career in %d and released his first independent album in %s.", artistData.CreationDate, artistData.FirstAlbum)
		result.Type = "Real name:"
	} else {
		result.Text = fmt.Sprintf("has been created in %d. The band produced its first album in %s.", artistData.CreationDate, artistData.FirstAlbum)
		result.Type = "Members:"
	}
	html, err := template.New("").ParseFiles("templates/index.html", "templates/tracker.html")
	if err != nil {
		h.Errorhandler(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	err = html.ExecuteTemplate(w, "index", result)
	if err != nil {
		h.Errorhandler(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	log.Printf("%d %s\n", http.StatusOK, OK)
}

func (h *Router) Errorhandler(w http.ResponseWriter, err string, status int) {
	html, err1 := template.New("").ParseFiles("templates/index.html", "templates/error.html")
	if err1 != nil {
		log.Println(err1)
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	err = fmt.Sprintf("%d %s", status, err)
	err2 := html.ExecuteTemplate(w, "index", err)
	if err2 != nil {
		log.Println(err2)
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	log.Println(err)
}