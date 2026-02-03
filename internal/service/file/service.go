package fileservice

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	dto "file_storage/internal/DTO"
	"file_storage/internal/apperror"
	"file_storage/internal/config"
	"file_storage/internal/domain"
)

type Service struct {
	repo       domain.FileRepo
	storage    domain.FileStorage
	cfg        *config.Config
	strategies map[string]UploadStrategy
	fallback   UploadStrategy
}

func NewService(repo domain.FileRepo, storage domain.FileStorage, cfg *config.Config) *Service {
	return &Service{
		repo:    repo,
		storage: storage,
		cfg:     cfg,
		strategies: map[string]UploadStrategy{
			"video/mp4": NewVideoStrategy(),
		},
		fallback: NewGenericStrategy(),
	}
}

func (s *Service) Upload(ctx context.Context, input *dto.UploadFileRequestDTO) (*dto.UploadFileResponseDTO, error) {
	strategy := s.getStrategy(input.ContentType)

	processed, err := strategy.Process(ctx, input)
	if err != nil {
		return nil, err
	}

	if closer, ok := processed.Reader.(interface{ Close() error }); ok {
		defer closer.Close()
	}

	url, err := s.storage.Upload(ctx, processed.Key, processed.Reader, processed.ContentType)
	if err != nil {
		return nil, apperror.WrapS3Upload(err)
	}

	fullURL := fmt.Sprintf("https://%s/%s", s.cfg.CloudFrontDistributionDomain, processed.Key)
	if url != "" {
		fullURL = url
	}

	file := domain.File{
		Filename:    input.Filename,
		FileURL:     fullURL,
		FileType:    processed.ContentType,
		Size:        input.Size,
		AspectRatio: processed.Metadata.AspectRatio,
		Title:       processed.Metadata.Title,
		Description: processed.Metadata.Description,
	}

	savedFile, err := s.repo.UploadFile(file)
	if err != nil {
		return nil, apperror.WrapInternal(err)
	}

	return &dto.UploadFileResponseDTO{
		ID:          savedFile.ID.String(),
		Filename:    savedFile.Filename,
		ContentType: savedFile.FileType,
		Size:        processed.Size,
		URL:         fullURL,
	}, nil
}

const defaultSignedURLExpiry = 1 * time.Hour

func (s *Service) GetFile(ctx context.Context, id string) (*dto.GetFileResponseDTO, error) {
	file, err := s.repo.GetFile(id)
	if err != nil {
		return nil, err
	}

	key := extractS3Key(file.FileURL)
	signedURL, err := s.storage.GetSignedURL(ctx, key, file.Filename, defaultSignedURLExpiry)
	if err != nil {
		return nil, apperror.WrapInternal(err)
	}

	expiresAt := time.Now().Add(defaultSignedURLExpiry)

	return &dto.GetFileResponseDTO{
		ID:          file.ID.String(),
		Filename:    file.Filename,
		ContentType: file.FileType,
		Size:        file.Size,
		SignedURL:   signedURL,
		AspectRatio: file.AspectRatio,
		Title:       file.Title,
		Description: file.Description,
		CreatedAt:   file.CreatedAt,
		ExpiresAt:   expiresAt,
	}, nil
}

func (s *Service) ListFiles(ctx context.Context) ([]domain.File, error) {
	return s.repo.ListFiles()
}

func (s *Service) DeleteFile(ctx context.Context, id string) error {
	return s.repo.DeleteFile(id)
}

func (s *Service) getStrategy(contentType string) UploadStrategy {
	if strategy, ok := s.strategies[contentType]; ok {
		return strategy
	}
	return s.fallback
}

func extractS3Key(fileURL string) string {
	parsed, err := url.Parse(fileURL)
	if err != nil {
		return fileURL
	}
	return strings.TrimPrefix(parsed.Path, "/")
}
