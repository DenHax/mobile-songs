package models

type Song struct {
	id     int
	Name   string
	Group  string
	Lyrics string
}

type UpdateSong struct {
	Name   string
	Group  string
	Lyrics string
}
