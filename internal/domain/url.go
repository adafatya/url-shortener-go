package domain

import "time"

type URL struct {
	ID        string    `json:"id"`
	ShortCode string    `json:"short_code"`
	TargetURL string    `json:"target_url"`
	CreatedAt time.Time `json:"created_at"`
}
