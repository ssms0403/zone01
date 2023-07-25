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
	response, 