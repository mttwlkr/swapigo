package model

// Film is a struct of a Star Wars movie
type Film struct {
	Title        string   `json:"title"`
	EpisodeID    int64    `json:"episode_id"`
	OpeningCrawl string   `json:"opening_crawl"`
	Director     string   `json:"director"`
	Producer     string   `json:"producer"`
	Characters   []string `json:"characters"`
	Planets      []string `json:"planets"`
	Starships    []string `json:"starships"`
	Vehicles     []string `json:"vehicles"`
	Species      []string `json:"species"`
	Created      string   `json:"created"`
	Edited       string   `json:"edited"`
	URL          string   `json:"url"`
}

// func getFilms(urls []string, wg *sync.WaitGroup) (films Film[], error) {
// 	for _, url := range urls {
// 		var f Film

// 		// if err = Get(url, &f); err != nil {
// 			// return
// 		// }
// 		films = append(films, f)
// 	}
// 	return
// }
