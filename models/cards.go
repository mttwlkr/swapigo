package model

// MainCard is the main focus of the page
type MainCard struct {
	Title, SubTitle, Body1, Body2, Body3, Body4, Body5 string
}

// SubCard describes cards lower on the page
type SubCard struct {
	Title, SubTitle, SubTitle2, Body, URL, Tags string
}
