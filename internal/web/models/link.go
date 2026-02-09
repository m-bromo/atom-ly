package models

type ShortenPayload struct {
	Url string `json:"url" validate:"url"`
}

type ShortenResponse struct {
	ShortLink string `json:"short_link"`
}
