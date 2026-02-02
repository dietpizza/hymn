package fileops

import (
	"dietpizza/hymn/types"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

var ChunkDir = "./temp"

func CreateFileWithParentDirs(path string) bool {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return false
	}

	file, err := os.Create(path)
	if err != nil {
		return false
	}
	defer file.Close()

	return true
}

func GetChunkFilePath(file_name string, byte_range types.ByteRange) (string, error) {
	chunk_file_name := fmt.Sprintf("%d-%d", byte_range.Start, byte_range.End)
	chunk_path := path.Join(ChunkDir, file_name, chunk_file_name)

	if ok := CreateFileWithParentDirs(chunk_path); !ok {
		return "", errors.New("error creating chunk file")
	}

	return chunk_path, nil
}
