// This package provides a command-line interface (CLI) tool for
// splitting a multipart file into individual parts by the provided boundary.

package main

import (
	"dicom-multipart-splitter/services"
	"fmt"
	"io"
	"os"
)

func main() {
	// Check if the correct number of arguments is provided
	if len(os.Args) != 3 {
		fmt.Println("Usage: multipart-split <file-path> <boundary>")
		os.Exit(1)
	}

	fmt.Println("Starting multipart file splitter...")
	fmt.Println("File path:", os.Args[1])
	fmt.Println("Boundary:", os.Args[2])

	filePath := os.Args[1]
	boundary := os.Args[2]

	// Open the file
	fp, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err.Error())
		os.Exit(1)
	}
	defer fp.Close()

	fmt.Println("File opened successfully")
	fmt.Println("Parsing file...")

	// Parse the multipart file
	parts, err := services.ParseMultipartFile(fp, boundary)
	if err != nil {
		fmt.Println("Error parsing multipart file:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Number of parts found:", len(parts))

	// remove headers from each part
	for i, part := range parts {
		fmt.Println("Removing header from part", i+1)
		updated, err := services.RemoveHeaderLines(part)
		if err != nil {
			fmt.Println("Error removing header from part:", err.Error())
			continue
		}
		parts[i] = updated
	}

	// save parts to files
	for i, part := range parts {
		fmt.Println("Saving part", i+1)
		partFileName := fmt.Sprintf("part_%d.dcm", i+1)
		partFile, err := os.Create(partFileName)
		if err != nil {
			fmt.Println("Error creating part file:", err.Error())
			continue
		}
		defer partFile.Close()

		if _, err := io.Copy(partFile, part); err != nil {
			fmt.Println("Error writing part to file:", err.Error())
			continue
		}

		fmt.Println("Part saved to:", partFileName)
	}
}
