package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/alpaca-tv/alpclib"
)

type Arguments struct {
	Search string
	Type   string
}

func ParseArguments() *Arguments {
	// Override usage
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "resotranslate")
		flag.PrintDefaults()
	}
	search := flag.String("search", "", "Search query")
	stype := flag.String("type", "film", "Search type (film,series)")
	flag.Parse()
	return &Arguments{
		Search: *search,
		Type:   *stype,
	}
}

func main() {
	args := ParseArguments()
	r := alpclib.Rezka{}
	if args.Type == "film" {
		films, err := r.ListFilms(&alpclib.ListParameters{
			Search: args.Search,
		})
		if err != nil {
			panic(err)
		}
		film, err := r.GetFilm(films[0].ID)
		if err != nil {
			panic(err)
		}
		fmt.Println("Name:", film.Name)
		fmt.Println("Poster:", film.PosterURL)
		fmt.Println("Description:", film.Description)
		fmt.Println("Year:", film.Year)
		fmt.Println("Rating:", film.Rating)
		fmt.Println("Country:", film.Country)
		fmt.Println("Genres:", strings.Join(film.Genres, ","))
		fmt.Println("Sources:")
		for _, source := range film.Sources {
			fmt.Println("-", source.Voicecover, source.Quality, source.URL)
		}
	}
}
