package main

import (
	"dietpizza/hymn/netops"
	"fmt"
)

var url = "https://fi.arch.niranjan.co/iso/2026.02.01/archlinux-bootstrap-x86_64.tar.zst"

func main() {
	metadata, err := netops.GetFileMetadata(url)
	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println("Content-Length", metadata.Size)

	netops.GetChunkRanges(metadata.Size)
	// fmt.Println("Chunks", chunk_ranges)
}
