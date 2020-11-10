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
	PageTitle string
	MainCard  MainCard
	SubCard1  []SubCard
	SubCard2  []SubCard
	SubCard3  []SubCard
}
