 package dto

import "time"

type UploadFileResponseDTO struct {
	ID          string    `json:"id"`
	Filename    string    `json:"filename"`
	ContentType string    `json:"content_type"`
	Size        int64     `json:"size"`
	URL         string    `json:"url,omitempty"`
	UploadedAt  time.Time `json:"uploaded_at"`
}