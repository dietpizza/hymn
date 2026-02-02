package netops

import (
	"errors"
	"net/http"
	"strconv"
)

type RemoteFileMetadata struct {
	Name                string
	Size                int64
	SupportsRangeHeader bool
}

func GetFileMetadata(url string) (RemoteFileMetadata, error) {
	resp, err := http.Head(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return RemoteFileMetadata{}, errors.New("error getting file metadata")
	}
	defer resp.Body.Close()

	content_disposition := resp.Header.Get("Content-Disposition")
	file_name := GetFileName(url, content_disposition)

	content_length := resp.Header.Get("Content-Length")
	file_size, err := strconv.ParseInt(content_length, 10, 64)
	if err != nil {
		file_size = -1
	}

	var supports_range bool
	if file_size > 0 {
		accepts_ranges := resp.Header.Get("Accept-Ranges")
		supports_range = accepts_ranges == "bytes"
	}

	return RemoteFileMetadata{
		Name:                file_name,
		Size:                file_size,
		SupportsRangeHeader: supports_range,
	}, nil
}
