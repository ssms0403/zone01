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