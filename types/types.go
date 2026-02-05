package types

type ByteRange struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

type RangeHeaderInfo struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type RemoteFileMetadata struct {
	Name                string `json:"name"`
	Size                int64  `json:"size"`
	SupportsRangeHeader bool   `json:"supports_range_header"`
}
