package fileservice

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	dto "file_storage/internal/DTO"

	"github.com/google/uuid"
)

type GenericStrategy struct{}

func NewGenericStrategy() *GenericStrategy {
	return &GenericStrategy{}
}

func (s *GenericStrategy) Process(ctx context.Context, input *dto.UploadFileRequestDTO) (*ProcessedFile, error) {
	ext := filepath.Ext(input.Filename)
	name := strings.TrimSuffix(input.Filename, ext)
	key := fmt.Sprintf("files/%s-%s%s", name, uuid.New().String()[:8], ext)
	return &ProcessedFile{
		Reader:      input.Reader,
		Key:         key,
		ContentType: input.ContentType,
		Size:        input.Size,
		Metadata: FileMetadata{
			Title:       input.Title,
			Description: input.Description,
		},
	}, nil
}
