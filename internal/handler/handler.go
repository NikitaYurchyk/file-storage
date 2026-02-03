package handler

import (
	dto "file_storage/internal/DTO"
	"file_storage/internal/apperror"
	fileservice "file_storage/internal/service/file"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *fileservice.Service
}

func NewHandler(service *fileservice.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetFile(c *gin.Context) {
	id := c.Param("id")

	result, err := h.service.GetFile(c.Request.Context(), id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, result)
}

func (h *Handler) ListFiles(c *gin.Context) {
	files, err := h.service.ListFiles(c.Request.Context())
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, files)
}

func (h *Handler) UploadFile(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "No file uploaded"})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to open file"})
		return
	}
	defer file.Close()

	uploadInput := &dto.UploadFileRequestDTO{
		Reader:      file,
		Filename:    fileHeader.Filename,
		ContentType: fileHeader.Header.Get("Content-Type"),
		Size:        fileHeader.Size,
		Title:       c.PostForm("title"),
		Description: c.PostForm("description"),
	}

	result, err := h.service.Upload(c.Request.Context(), uploadInput)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(201, result)
}

func (h *Handler) DeleteFile(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteFile(c.Request.Context(), id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(204, nil)
}

func (h *Handler) UploadThumbnail(c *gin.Context) {
	// TODO: implement
	c.JSON(501, gin.H{"error": "Not implemented"})
}

func handleError(c *gin.Context, err error) {
	if appErr, ok := err.(*apperror.AppError); ok {
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}
	c.JSON(500, gin.H{"error": "Internal server error"})
}