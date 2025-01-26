package models

type URL struct {
	ID        int    `json:"id"`
	ShortURL  string `json:"short_url"`
	Original  string `json:"original"`
	CreatedAt string `json:"created_at"`
	ExpiresAt string `json:"expires_at"`
}
