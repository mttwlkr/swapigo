package model

// IndexPageResponse is the response to a request for a page
type IndexPageResponse struct {
	PageNumber   int
	NextPage     string
	PreviousPage string
	Cards        []SubCard
	PageTitle    string
}

// DetailPageResponse is the response to a request for an individual user
type DetailPageResponse struct {
	PageTitle   string
	MainCard    MainCard
	Cards1      []SubCard
	Cards1Title string
	Cards2      []SubCard
	Cards2Title string
	Cards3      []SubCard
	Cards3Title string
}
