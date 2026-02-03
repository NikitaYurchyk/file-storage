package fileservice

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	dto "file_storage/internal/DTO"
	"file_storage/internal/apperror"
	"file_storage/pkg/video"

	"github.com/google/uuid"
)

type VideoStrategy struct{}

func NewVideoStrategy() *VideoStrategy {
	return &VideoStrategy{}
}

func (s *VideoStrategy) Process(ctx context.Context, input *dto.UploadFileRequestDTO) (*ProcessedFile, error) {
	tmpFile, err := os.CreateTemp("", "upload-*.mp4")
	if err != nil {
		return nil, apperror.WrapTempFileCreate(err)
	}
	tmpPath := tmpFile.Name()
	defer os.Remove(tmpPath)

	if _, err := io.Copy(tmpFile, input.Reader); err != nil {
		tmpFile.Close()
		return nil, apperror.WrapTempFileWrite(err)
	}
	tmpFile.Close()

	aspectRatio, err := video.GetVideoAspectRatio(tmpPath)
	if err != nil {
		return nil, apperror.WrapGetAspectRatio(err)
	}

	processedPath, err := video.ProcessForFastStart(tmpPath)
	if err != nil {
		return nil, apperror.WrapVideoProcessing(err)
	}
	defer os.Remove(processedPath)

	processedFile, err := os.Open(processedPath)
	if err != nil {
		return nil, apperror.WrapTempFileOpen(err)
	}

	stat, err := processedFile.Stat()
	if err != nil {
		processedFile.Close()
		return nil, apperror.WrapInternal(err)
	}

	ext := filepath.Ext(input.Filename)
	if ext == "" {
		ext = ".mp4"
	}
	name := strings.TrimSuffix(input.Filename, ext)
	key := fmt.Sprintf("%s/%s-%s%s", aspectRatio, name, uuid.New().String()[:8], ext)

	return &ProcessedFile{
		Reader:      processedFile,
		Key:         key,
		ContentType: "video/mp4",
		Size:        stat.Size(),
		Metadata: FileMetadata{
			AspectRatio: string(aspectRatio),
			Title:       input.Title,
			Description: input.Description,
		},
	}, nil
}
