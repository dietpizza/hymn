package netops

import (
	"mime"
	"net/url"
	"path"
	"strings"
)

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
