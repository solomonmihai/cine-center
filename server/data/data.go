package data

import (
	"errors"
	"main/utils"
	"sort"
	"time"
)

type Film struct {
	Title    string `json:"title"`
	Date     int    `json:"date"`
	Price    string `json:"price"`
	Link     string `json:"link"`
	ImgURL   string `json:"img_url"`
	Location string `json:"location"`
}

// Cinema represents a cinema with a name, URL, and a list of films.
type Cinema struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Films []Film `json:"films"`
}

// Cinemas represents a map of cinema names to Cinema structs.
type Cinemas map[string]Cinema

const DATA_PATH = "../data.json"

var CinemaData Cinemas
var CinemaNames []string
var AllFilmsByDate []Film

func loadData() error {
	data, err := utils.LoadJSON[Cinemas](DATA_PATH)
	if err != nil {
		return err
	}

	CinemaData = data
	CinemaNames = getCinemaNames()
	AllFilmsByDate = getAllFilmsByDate()

	return nil
}

func getCinemaNames() []string {
	names := []string{}
	for _, cinema := range CinemaData {
		names = append(names, cinema.Name)
	}
	return names
}

func getAllFilmsByDate() []Film {
	films := []Film{}

	now := time.Now().Unix()

	for _, cinema := range CinemaData {
		for _, film := range cinema.Films {
			if int64(film.Date) > now {
				films = append(films, film)
			}
		}
	}

	sort.Slice(films, func(i, j int) bool {
		return films[i].Date < films[j].Date
	})

	return films
}

func GetCinemaData(name string) (Cinema, error) {
	cinema, exists := CinemaData[name]
	if !exists {
		return Cinema{}, errors.New("cinema does not exist")
	}

	return cinema, nil
}

func GetFilmsByDate(date string) ([]Film, error) {
	return nil, nil
}
