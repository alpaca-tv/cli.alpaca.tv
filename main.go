package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/alpaca-tv/alpclib"
)

type Arguments struct {
	Search     string
	Series     bool
	Voicecover string
	Quality    string
	Season     int
	Episode    int
}

func ParseArguments() *Arguments {
	// Override usage
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "alpc")
		flag.PrintDefaults()
	}
	search := flag.String("search", "", "Search query")
	series := flag.Bool("series", false, "Search for series")
	voicecover := flag.String("voicecover", "", "Voicecover filter")
	quality := flag.String("quality", "", "Quality filter")
	season := flag.Int("season", 1, "Series season filter (0 for all)")
	episode := flag.Int("episode", 1, "Series episode filter (0 for all)")
	flag.Parse()
	return &Arguments{
		Search:     *search,
		Series:     *series,
		Voicecover: *voicecover,
		Quality:    *quality,
		Season:     *season,
		Episode:    *episode,
	}
}

func main() {
	args := ParseArguments()
	r := alpclib.Rezka{}
	if !args.Series {
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
	} else {
		serieslist, err := r.ListSeries(&alpclib.ListParameters{
			Search: args.Search,
		})
		if err != nil {
			panic(err)
		}
		series, err := r.GetSeries(serieslist[0].ID, args.Season, args.Episode)
		fmt.Println("Name:", series.Name)
		fmt.Println("Poster:", series.PosterURL)
		fmt.Println("Description:", series.Description)
		fmt.Println("Rating:", series.Rating)
		fmt.Println("Country:", series.Country)
		fmt.Println("Sources:")
		for _, source := range series.Sources {
			if source.URL == "" {
				continue
			}
			if args.Voicecover != "" && !strings.Contains(source.Voicecover, args.Voicecover) {
				continue
			}
			if args.Quality != "" && !strings.Contains(source.Quality, args.Quality) {
				continue
			}
			fmt.Println("-", fmt.Sprintf("s%v:e%v", source.Season, source.Episode), source.Voicecover, source.Quality, source.URL)
		}
	}
}
