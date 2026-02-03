package fileservice

import (
	"context"
	dto "file_storage/internal/DTO"
	"io"
)

type UploadStrategy interface {
	Process(ctx context.Context, input *dto.UploadFileRequestDTO) (*ProcessedFile, error)
}

type ProcessedFile struct {
	Reader      io.Reader
	Key         string
	ContentType string
	Size        int64
	Metadata    FileMetadata
}

type FileMetadata struct {
	AspectRatio string
	Title       string
	Description string
}
