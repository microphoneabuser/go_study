package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const domain = "https://omdbapi.com/"
const apikey = "?apikey=9e824146&"

func main() {
	name := strings.ReplaceAll(os.Args[1], " ", "+")
	resp, err := http.Get(domain + apikey + "t=" + name)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Println(resp.Status)
	}
	var film Film
	if err := json.NewDecoder(resp.Body).Decode(&film); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fmt.Printf("Title: %s\nYear: %s\nRated: %s\nReleased: %s\n", film.Title, film.Year, film.Rated, film.Released)
	fmt.Printf("Runtime: %s\nGenre: %s\nDirector: %s\n", film.Runtime, film.Genre, film.Director)
	fmt.Printf("Writer: %s\nActors: %s\nPlot: %s\nLanguage: %s\n", film.Writer, film.Actors, film.Plot, film.Language)
	fmt.Printf("Country: %s\nAwards: %s\nPoster: %s\nRatings: %s\n", film.Country, film.Awards, film.Poster, film.Ratings)
	fmt.Printf("Metascore: %s\nIMDBRating: %s\nIMDBVote: %s\nIMDBID: %s\n", film.Metascore, film.IMDBRating, film.IMDBVote, film.IMDBID)
	fmt.Printf("Type: %s\nDVD: %s\nBoxOffice: %s\n", film.Type, film.DVD, film.BoxOffice)
	fmt.Printf("Production: %s\nWebsite: %s\nResponse: %s\n", film.Production, film.Website, film.Response)
}

type Film struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Ratings    []Ratings
	Metascore  string
	IMDBRating string `json:"imdbRating"`
	IMDBVote   string `json:"imdbVote"`
	IMDBID     string `json:"imdbID"`
	Type       string
	DVD        string
	BoxOffice  string
	Production string
	Website    string
	Response   string
}

type Ratings struct {
	Source string
	Value  string
}
