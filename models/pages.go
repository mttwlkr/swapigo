package model

// IndexPageResponse is the response to a request for a page
type IndexPageResponse struct {
	PageNumber   int
	NextPage     string
	PreviousPage string
	Cards        []SubCard
	PageTitle    string
}
