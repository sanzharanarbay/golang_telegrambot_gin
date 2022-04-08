package models

// Album detail model
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// AlbumList is the response of list user endpoint
type AlbumList struct {
	Albums []Album `json:"albums"`
}