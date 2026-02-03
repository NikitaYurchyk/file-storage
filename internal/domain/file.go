package domain

import (
	"context"
	"io"
	"time"

	"github.com/google/uuid"
)


type File struct {
	ID           uuid.UUID `json:"id" gorm:"type:text;primaryKey"`
	CreatedAt    int64     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    int64     `json:"updated_at" gorm:"autoUpdateTime"`
	Filename     string    `json:"filename" gorm:"type:text;not null"`
	Size         int64     `json:"size" gorm:"type:integer"`
	ThumbnailURL string    `json:"thumbnail_url" gorm:"type:text"`
	FileURL      string    `json:"file_url" gorm:"type:text;not null"`
	FileType     string    `json:"file_type" gorm:"type:text"`
	AspectRatio  string    `json:"aspect_ratio,omitempty" gorm:"type:text"`
	Title        string    `json:"title" gorm:"type:text"`
	Description  string    `json:"description" gorm:"type:text"`
}

func (File) TableName() string {
	return "files"
}


type FileStorage interface {
	Upload(ctx context.Context, key string, reader io.Reader, contentType string) (string, error)
	Download(ctx context.Context, key string) (io.ReadCloser, error)
	Delete(ctx context.Context, key string) error
	GetSignedURL(ctx context.Context, key string, filename string, expiry time.Duration) (string, error)
}

type FileRepo interface{
	GetFile(id string) (File, error)
	ListFiles() ([]File, error)
	UploadFile(data File) (File, error)
	DeleteFile(id string) error
}


