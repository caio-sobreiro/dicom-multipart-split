package services

import (
	"bytes"
	"io"
)

/*
ParseMultipartFile reads a multipart file from the provided io.Reader and splits it into parts based on the specified boundary.
This is a rough implementation that reads the entire file into memory, which may not be suitable for very large files.
*/
func ParseMultipartFile(fp io.Reader, boundary string) ([]io.Reader, error) {
	// Read entire file and store it in a buffer
	data, err := io.ReadAll(fp)
	if err != nil {
		return nil, err
	}

	// Split the data by the boundary
	boundaryBytes := []byte("\r\n--" + boundary)
	parts := bytes.Split(data, boundaryBytes)
	if len(parts) < 2 {
		return nil, nil // No parts found
	}

	// Create a ReadCloser for each part
	var readers []io.Reader
	var part io.Reader
	for _, p := range parts[1:] { // Skip the first part which is before the boundary
		if len(p) == 0 {
			continue // Skip empty parts
		}
		part = bytes.NewReader(p)
		readers = append(readers, part)
	}

	return readers, nil
}

func RemoveHeaderLines(part io.Reader) (io.Reader, error) {
	// Read the part into a buffer
	data, err := io.ReadAll(part)
	if err != nil {
		return nil, err
	}

	// Find the end of the header (double CRLF)
	headerEnd := bytes.Index(data, []byte("\r\n\r\n"))
	if headerEnd == -1 {
		return bytes.NewReader(data), nil // No header found, return as is
	}

	// Return the part after the header
	return bytes.NewReader(data[headerEnd+4:]), nil
}
