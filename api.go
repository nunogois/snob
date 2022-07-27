package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type Rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

type Result struct {
	Title      string   `json:"Title"`
	Year       string   `json:"Year"`
	Rated      string   `json:"Rated"`
	Released   string   `json:"Released"`
	Runtime    string   `json:"Runtime"`
	Genre      string   `json:"Genre"`
	Director   string   `json:"Director"`
	Writer     string   `json:"Writer"`
	Actors     string   `json:"Actors"`
	Plot       string   `json:"Plot"`
	Language   string   `json:"Language"`
	Country    string   `json:"Country"`
	Awards     string   `json:"Awards"`
	Poster     string   `json:"Poster"`
	Ratings    []Rating `json:"Ratings"`
	Metascore  string   `json:"Metascore"`
	ImdbRating string   `json:"imdbRating"`
	ImdbVotes  string   `json:"imdbVotes"`
	ImdbID     string   `json:"imdbID"`
	Type       string   `json:"Type"`
	Dvd        string   `json:"DVD"`
	BoxOffice  string   `json:"BoxOffice"`
	Production string   `json:"Production"`
	Website    string   `json:"Website"`
}

var req = &http.Client{Timeout: 10 * time.Second}

func search(query string, result *Result, key string) error {
	if key == "" {
		key = getKey()
	}
	if key == "" {
		return errors.New("No API key set. Request one at https://www.omdbapi.com/apikey.aspx and then set it with `snob -k YOUR_KEY`.")
	}

	res, err := req.Get("https://www.omdbapi.com/?apikey=" + key + "&t=" + query)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(&result)
}
