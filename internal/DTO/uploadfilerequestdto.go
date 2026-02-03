package dto

import "io"


type UploadFileRequestDTO struct {
    Reader      io.Reader
    Filename    string
    ContentType string
    Size        int64
    // Optional metadata from form
    Title       string
    Description string
}