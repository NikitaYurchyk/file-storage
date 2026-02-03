package filerepository

import (
	"errors"

	"file_storage/internal/apperror"
	"file_storage/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetFile(id string) (domain.File, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return domain.File{}, apperror.ErrInvalidFileID
	}

	var file domain.File
	result := r.db.First(&file, "id = ?", uid.String())
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.File{}, apperror.ErrFileNotFound
		}
		return domain.File{}, apperror.WrapInternal(result.Error)
	}

	return file, nil
}

func (r *Repository) ListFiles() ([]domain.File, error) {
	var files []domain.File
	result := r.db.Order("created_at DESC").Find(&files)
	if result.Error != nil {
		return nil, apperror.WrapInternal(result.Error)
	}
	return files, nil
}

func (r *Repository) UploadFile(data domain.File) (domain.File, error) {
	if data.ID == uuid.Nil {
		data.ID = uuid.New()
	}

	result := r.db.Create(&data)
	if result.Error != nil {
		return domain.File{}, apperror.WrapInternal(result.Error)
	}

	return data, nil
}

func (r *Repository) DeleteFile(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return apperror.ErrInvalidFileID
	}

	result := r.db.Delete(&domain.File{}, "id = ?", uid.String())
	if result.Error != nil {
		return apperror.WrapInternal(result.Error)
	}

	if result.RowsAffected == 0 {
		return apperror.ErrFileNotFound
	}

	return nil
}
