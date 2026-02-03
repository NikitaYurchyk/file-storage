package video

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type AspectRatio string

const (
	Portrait  AspectRatio = "portrait"
	Landscape AspectRatio = "landscape"
)

func GetVideoAspectRatio(filePath string) (AspectRatio, error) {
	cmd := exec.Command("ffprobe",
		"-v", "error",
		"-print_format", "json",
		"-show_streams",
		filePath,
	)

	var buffer bytes.Buffer
	cmd.Stdout = &buffer

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("ffprobe failed: %w", err)
	}

	type stream struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	}

	type ffprobeOutput struct {
		Streams []stream `json:"streams"`
	}

	var output ffprobeOutput
	if err := json.Unmarshal(buffer.Bytes(), &output); err != nil {
		return "", fmt.Errorf("failed to parse ffprobe output: %w", err)
	}

	if len(output.Streams) == 0 {
		return "", fmt.Errorf("no streams found in video")
	}

	width := output.Streams[0].Width
	height := output.Streams[0].Height

	if width == 0 || height == 0 {
		return "", fmt.Errorf("invalid video dimensions: %dx%d", width, height)
	}

	if height > width {
		return Portrait, nil
	}
	return Landscape, nil
}

func ProcessForFastStart(filePath string) (string, error) {
	outputPath := filePath + ".processing.mp4"
	cmd := exec.Command("ffmpeg",
		"-i", filePath,
		"-c", "copy",
		"-movflags", "faststart",
		"-f", "mp4",
		"-y",
		outputPath,
	)

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("ffmpeg processing failed: %w", err)
	}

	return outputPath, nil
}
