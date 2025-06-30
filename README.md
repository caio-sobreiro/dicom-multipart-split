# DICOM Multipart Splitter
This package provides a command-line interface (CLI) tool for splitting a multipart file into individual parts by the provided boundary.

# Build
```bash
go build -o multipart-split cli/main.go
```

# Usage
```bash
./multipart-split [file] [boundary]
```

- File is the path to the file to be split
- Boundary is the string that splits the file (header boundary in the HTTP response)
