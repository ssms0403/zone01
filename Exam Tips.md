package server

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/helper"
	"groupie-tracker/models"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"
)
type Artistlocation struct{
	Location models.Alllocation
	Artist []models.Artist
}

var ErrorMsgMap = map[int]Error{
	http.StatusBadRequest:          {http.StatusBadRequest, "Bad Request"},
	http.StatusNotFound:            {http.StatusNotFound, "Not Found"},
	http.StatusInternalServerError: {http.StatusInternalServerError, "Internal Server Error"},
	http.StatusMethodNotAllowed:    {http.StatusMethodNotAllowed, "Method Not Allowed"},
}

type Error struct {
	Code int
	Msg  string
}

func getArtists() ([]models.Artist, error) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var artists []models.Artist

	err = json.NewDecoder(response.Body).Decode(&artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func getArtistById(id int, artist *models.Artist) error {

	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", id)
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(artist)
	if artist.Id == 0 {
		err = fmt.Errorf("no artist with this ID: %d", id)
	}

	if err != nil {
		return err
	}

	return err
}

func getLocations(url string, location *models.Location) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(location)
	if err != nil {
		return err
	}

	return err
}

func getDates(url string, date *models.Dates) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(date)
	if err != nil {
		return err
	}

	return err
}

func getRelations(url string, relations *models.Relations) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	err = json.Unmarshal(body, &relations)
	if err != nil {
		return err
	}

	return err
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RenderErrorPage(http.StatusNotFound, w)
		return
	}
	if r.Method != http.MethodGet {
		RenderErrorPage(http.StatusMethodNotAllowed, w)
		return
	}
	artist, err := getArtists()
	 art ,err:=getAlllocation()

	if err != nil {
		RenderErrorPage(http.StatusInternalServerError, w)
		return
	}
	s := "templates/index.html"
	tmpl, err := template.ParseFiles(s)

	if err != nil {
		RenderErrorPage(http.StatusInternalServerError, w)
		return
	}
	var artloc =Artistlocation{
	art ,
	artist,
	}

	tmpl.Execute(w, artloc)
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artist" {
		RenderErrorPage(http.StatusNotFound, w)
		return
	}
	if r.Method != http.MethodGet {
		RenderErrorPage(http.StatusMethodNotAllowed, w)
		return
	}
	s := "templates/artist.html"
	tmpl, err := template.ParseFiles(s)
	if err != nil {
		RenderErrorPage(http.StatusInternalServerError, w)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if err != nil || id < 0 {
		RenderErrorPage(http.StatusBadRequest, w)
		return
	}
	var artist models.Artist

	err = getArtistById(id, &artist)

	if err != nil {
		RenderErrorPage(http.StatusNotFound, w)
		return
	}

	var artistRes models.ArtistResponse

	artistRes.Id = artist.Id
	artistRes.CreationDate = artist.CreationDate
	artistRes.FirstAlbum = helper.Form