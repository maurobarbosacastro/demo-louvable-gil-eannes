package files

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"mime/multipart"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/pkg/dotenv"
	"os"
	"path/filepath"
)

func HandleFile(fileHeader *multipart.FileHeader, file *models.File) (*string, map[string]string) {
	// Open the uploaded file
	src, err := fileHeader.Open()
	if err != nil {
		return nil, map[string]string{"error": "Could not open file"}
	}
	defer src.Close()

	fileUuid, errMap := createFile(fileHeader, src, file)

	if errMap != nil {
		return nil, errMap
	}

	return fileUuid, nil
}

func createFile(fileHeader *multipart.FileHeader, src multipart.File, file *models.File) (*string, map[string]string) {
	// Create directory path using UUID
	pathToFile := dotenv.GetEnv("PATH_FILE")
	baseDir := filepath.Join(pathToFile, file.Uuid.String())
	if err := os.MkdirAll(baseDir, os.ModePerm); err != nil {
		return nil, map[string]string{"error": "Could not create directory"}
	}

	// Create the full file path
	filePath := filepath.Join(baseDir, fileHeader.Filename)

	// Create the destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return nil, map[string]string{"error": "Could not create file"}
	}
	defer dst.Close()

	// Copy the file content
	if _, err := io.Copy(dst, src); err != nil {
		// Cleanup on error
		os.RemoveAll(baseDir)
		return nil, map[string]string{"error": "Could not save file content"}
	}

	uuidStr := file.Uuid.String()
	return &uuidStr, nil
}

func CreateCSVFile(header []string, data [][]string, fileName string) (*bytes.Buffer, error) {
	fmt.Printf("START service.CreatingCSVFile  - %v\n", fileName)

	// Create a buffer to write the CSV to
	buf := new(bytes.Buffer)
	writer := csv.NewWriter(buf)
	if err := writer.Write(header); err != nil {
		fmt.Printf("Error writing header - %v\n", err)
		return nil, err
	}

	// Write the data on the CSV file
	if err := writer.WriteAll(data); err != nil {
		fmt.Printf("Error writing row - %v\n", err)
		return nil, err
	}
	fmt.Printf("ALL data was save on the CSV file - %v\n", fileName)

	// Flush the writer to ensure all data is written to the buffer
	writer.Flush()
	if err := writer.Error(); err != nil {
		return nil, err
	}

	fmt.Printf("END service.CreatingCSVFile - %v\n", fileName)
	return buf, nil

}
