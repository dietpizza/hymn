package types

import "fmt"

type WriteCounter struct {
	Total   int64
	Written int64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Written += int64(n)

	fmt.Printf("downloaded %d of %d\n", wc.Written, wc.Total)

	return n, nil
}
