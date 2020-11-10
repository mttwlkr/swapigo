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
	MainCard  model.MainCard
	SubCard1  []model.SubCard
	SubCard2  []model.SubCard
	SubCard3  []model.SubCard
}
