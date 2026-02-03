package dto

import "time"

type GetFileResponseDTO struct {
	ID          string `json:"id"`
	Filename    string `json:"filename"`
	ContentType string `json:"content_type"`
	Size        int64  `json:"size"`
	SignedURL   string `json:"signed_url"`
	AspectRatio string `json:"aspect_ratio,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   int64  `json:"created_at"`
	ExpiresAt   time.Time `json:"expires_at"`
}
