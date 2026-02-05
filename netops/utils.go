package netops

import (
	"dietpizza/hymn/types"
	"fmt"
	"mime"
	"net/url"
	"path"
	"strings"
)

func GetRangeHeader(byte_range types.ByteRange) types.RangeHeaderInfo {
	return types.RangeHeaderInfo{
		Key:   "Range",
		Value: fmt.Sprintf("bytes=%d-%d", byte_range.Start, byte_range.End),
	}
}

func GetFileName(url string, contentDisposition string) string {
	if name, ok := ExtractFilename(url); ok {
		return name
	}
	if name, ok := ParseContentDispositionFilename(contentDisposition); ok {
		return name
	}

	return ""
}

func ParseContentDispositionFilename(headerValue string) (string, bool) {
	if headerValue == "" {
		return "", false
	}

	_, params, err := mime.ParseMediaType(headerValue)
	if err != nil {
		return "", false
	}

	if name, ok := params["filename"]; ok && name != "" {
		return name, true
	}

	return "", false
}

func ExtractFilename(urlStr string) (string, bool) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", false
	}

	filename := path.Base(u.Path)
	if filename == "." || filename == "/" {
		return "", false
	}

	filename = strings.Split(filename, "?")[0]

	return filename, true
}

var DefaultChunkSize = int64(25 * 1024 * 1024)

func GetChunkRanges(content_length int64) []types.ByteRange {
	var chunk_ranges []types.ByteRange

	offset := int64(0)

	for offset < content_length {
		end := Clamp(offset+DefaultChunkSize, content_length)
		chunk := types.ByteRange{Start: offset, End: end}
		offset = chunk.End + 1

		fmt.Println(chunk)

		chunk_ranges = append(chunk_ranges, chunk)
	}

	return chunk_ranges
}

func Clamp(value int64, max int64) int64 {
	if value > max {
		return max
	}
	return value
}
