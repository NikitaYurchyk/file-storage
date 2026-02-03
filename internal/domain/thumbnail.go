package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Thumbnail struct {
	ID           uuid.UUID `json:"id" gorm:"type:text;primaryKey"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	ThumbnailURL string    `json:"thumbnail_url" gorm:"type:text;not null"`
	FileID       uuid.UUID `json:"file_id" gorm:"type:text;not null;index"`
}

func (Thumbnail) TableName() string {
	return "thumbnails"
}

type ThumbnailRepo interface {
	Create(ctx context.Context, thumbnail *Thumbnail) error
	GetByID(ctx context.Context, id uuid.UUID) (*Thumbnail, error)
	GetByFileID(ctx context.Context, fileID uuid.UUID) ([]Thumbnail, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
