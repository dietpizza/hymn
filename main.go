package main

import (
	"dietpizza/hymn/netops"
	"fmt"
)

var url = "https://in-mirror.garudalinux.org/archlinux/iso/2026.01.01/archlinux-bootstrap-x86_64.tar.zst"

func main() {
	metadata, err := netops.GetFileMetadata(url)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("Metadata", metadata)
}
