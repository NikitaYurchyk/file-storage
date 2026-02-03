package apperror

import "fmt"

type AppError struct {
	Code    int    `json:"-"`
	Message string `json:"error"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func New(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

func Wrap(code int, message string, err error) *AppError {
	return &AppError{Code: code, Message: message, Err: err}
}

var (
	ErrUnsupportedFileType = New(400, "unsupported file type")
	ErrNoFileUploaded      = New(400, "no file uploaded")
	ErrInvalidFileID       = New(400, "invalid file id")
)

var (
	ErrFileNotFound = New(404, "file not found")
)

var (
	ErrInternal            = New(500, "internal server error")
	ErrTempFileCreate      = New(500, "couldn't create temp file")
	ErrTempFileWrite       = New(500, "couldn't write to temp file")
	ErrTempFileOpen        = New(500, "couldn't open processed file")
	ErrGetAspectRatio      = New(500, "couldn't get aspect ratio")
	ErrVideoProcessing     = New(500, "couldn't process video for fast start")
	ErrS3Upload            = New(500, "couldn't upload file to storage")
)

func WrapInternal(err error) *AppError {
	return Wrap(500, "internal server error", err)
}

func WrapTempFileCreate(err error) *AppError {
	return Wrap(500, "couldn't create temp file", err)
}

func WrapTempFileWrite(err error) *AppError {
	return Wrap(500, "couldn't write to temp file", err)
}

func WrapTempFileOpen(err error) *AppError {
	return Wrap(500, "couldn't open processed file", err)
}

func WrapGetAspectRatio(err error) *AppError {
	return Wrap(500, "couldn't get aspect ratio", err)
}

func WrapVideoProcessing(err error) *AppError {
	return Wrap(500, "couldn't process video for fast start", err)
}

func WrapS3Upload(err error) *AppError {
	return Wrap(500, "couldn't upload file to storage", err)
}
