package netops

import (
	"dietpizza/hymn/fileops"
	"dietpizza/hymn/types"
	"fmt"
	"net/http"
)

func GetRangeHeaderValue(byte_range types.ByteRange) types.RangeHeaderInfo {
	return types.RangeHeaderInfo{
		Key:   "Range",
		Value: fmt.Sprintf("bytes=%d-%d", byte_range.Start, byte_range.End),
	}
}

func DownloadChunk(url string, file_name string, byte_range types.ByteRange) error {
	range_header := GetRangeHeaderValue(byte_range)
	chunk_path, err := fileops.GetChunkFilePath(file_name, byte_range)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add(range_header.Key, range_header.Value)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
