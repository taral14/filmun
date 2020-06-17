package entity

import "time"

type FilmImdb struct {
	Mark  float32 `json:"mark"`
	Votes int     `json:"votes"`
}

type FilmStatus int

type Film struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	Url            string     `json:"-"`
	OriginalName   string     `json:"original_name"`
	KpInternalId   int        `json:"-"`
	ImdbInternalId string     `json:"-"`
	Year           int        `json:"year"`
	CountSeasons   int        `json:"count_seasons"`
	IsSeries       int        `json:"is_series"`
	Imdb           FilmImdb   `json:"imdb"`
	ImageUrl       string     `json:"image"`
	KpImageUrl     string     `json:"-"`
	Description    string     `json:"description"`
	Status         FilmStatus `json:"status"`
	PremiereAt     time.Time  `json:"premiere_at"`
	Release720pAt  time.Time  `json:"release_720p_at"`
	Release1080pAt time.Time  `json:"release_1080p_at"`
	CreatedAt      time.Time  `json:"-"`
}

func (f *Film) GetNNN() string {
	return "nnn"
}
