package netops

import (
	"dietpizza/hymn/fileops"
	"dietpizza/hymn/types"
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadChunk(url string, file_name string, byte_range types.ByteRange) error {
	range_header := GetRangeHeader(byte_range)
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

	isSuccessfulRequest := resp.StatusCode == http.StatusPartialContent || resp.StatusCode == http.StatusOK
	if !isSuccessfulRequest {
		return fmt.Errorf("bad request - status code: %d", resp.StatusCode)
	}

	wc := &types.WriteCounter{Total: resp.ContentLength}
	file, err := os.Open(chunk_path)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, io.TeeReader(resp.Body, wc))
	if err != nil {
		return err
	}

	return nil
}
