package main

import (
	"fmt"
	"scraper/model"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	url := "https://cavea.ge"

	var movies []model.Movie

	c.OnHTML("div.movie", func(e *colly.HTMLElement) {
		var movie model.Movie

		movie.Name = e.ChildText("a h5")

		movie.ImgURL = e.ChildAttr("img.movie-avatar-soon", "src")

		e.ForEach("div.movie-sessions a", func(i int, elem *colly.HTMLElement) {
			movie.AvailableTimes = append(movie.AvailableTimes, elem.Text)
		})

		if e.DOM.Find("div.imax-logo").Length() > 0 {
			movie.IMAX = true
		}

		e.ForEach("div.rating a ul.ratings li", func(i int, elem *colly.HTMLElement) {
			text := strings.TrimSpace(elem.Text)
			style := strings.TrimSpace(elem.Attr("style"))

			if text != "" {
				if style != "" {
					movie.Languages = append(movie.Languages, text)
				} else {
					movie.Rating = text
				}
			}
		})

		movies = append(movies, movie)
	})

	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error:", err)
	}

	for i, movie := range movies {
		fmt.Printf("Movie #%d\n", i+1)
		fmt.Printf("Name: %s\n", movie.Name)
		fmt.Printf("Image URL: %s\n", movie.ImgURL)
		fmt.Printf("Available Times: %v\n", movie.AvailableTimes)
		fmt.Printf("Rating: %s\n", movie.Rating)
		fmt.Printf("Languages: %v\n", movie.Languages)
		fmt.Printf("IMAX: %v\n", movie.IMAX)
		fmt.Println("-----------------------------------------")
	}
}
