package model

type Movie struct {
	ImgURL         string
	Name           string
	AvailableTimes []string
	Rating         string
	Languages      []string
	IMAX           bool
}
