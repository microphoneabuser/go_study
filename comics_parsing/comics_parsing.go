package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const domain = "https://xkcd.com"
const jsonName = "info.0.json"

func main() {
	i := 1
	var comics Comics
	for {
		query := fmt.Sprintf("%s/%d/%s", domain, i, jsonName)
		var comic Comic
		if err := ReadComic(query, &comic); err != nil {
			fmt.Errorf("search query failed. %s", err)
			break
		}
		comics.Comics = append(comics.Comics, &comic)
		fmt.Println(comic.Num)
		i++
	}
	if err := WriteComics("comics.json", &comics); err != nil {
		fmt.Errorf("writing json file failed. %s", err)
	} else {
		fmt.Printf("Done.")
	}
}

func ReadComic(query string, comic *Comic) error {
	resp, err := http.Get(query)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return err
	}
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		resp.Body.Close()
		return err
	}
	resp.Body.Close()
	return nil
}

func WriteComics(path string, comics *Comics) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	if err := json.NewEncoder(file).Encode(comics); err != nil {
		return err
	}
	file.Close()
	return nil
}

type Comics struct {
	Comics []*Comic `json:"comics"`
}

type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"self_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}
