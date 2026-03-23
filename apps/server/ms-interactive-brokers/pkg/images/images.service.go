package images

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"ms-interactive-brokers/pkg/dotenv"
	"ms-interactive-brokers/pkg/http_client"
	"ms-interactive-brokers/pkg/utils"
	"net/http"
)

func UploadImage(file *multipart.FileHeader, fileName string, types string) (*string, map[string]string) {
	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return nil, map[string]string{"error": "Could not open file"}
	}
	defer src.Close()

	responseUuid, errMap := createImage(file, src, fileName, types)
	if errMap != nil {
		return nil, errMap
	}

	return responseUuid, nil
}

func createImage(file *multipart.FileHeader, src multipart.File, fileName string, types string) (*string, map[string]string) {
	// Create a new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a form file field
	part, err := writer.CreateFormFile("file", file.Filename)
	if err != nil {
		return nil, map[string]string{"error": "Could not create form file"}
	}

	// Copy file contents to the form file field
	_, err = io.Copy(part, src)
	if err != nil {
		return nil, map[string]string{"error": "Could not copy file contents"}
	}

	// Add any additional form fields if needed
	writer.WriteField("fileName", fileName)
	writer.WriteField("types", types)

	// Close the multipart writer
	err = writer.Close()
	if err != nil {
		return nil, map[string]string{"error": "Could not close multipart writer"}
	}

	msImagesUrl := dotenv.GetEnv("MS_IMAGES_URL")
	// Create the HTTP request to send the file
	req, err := http.NewRequest("POST", msImagesUrl+"api/image/", body)
	if err != nil {
		return nil, map[string]string{"error": "Could not create HTTP request"}
	}

	// Set the content type to multipart/form-data
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, map[string]string{"error": "Could not send file to destination"}
	}
	defer resp.Body.Close()

	// Read the response
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, map[string]string{"error": "Could not read response"}
	}

	var result map[string]interface{}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, map[string]string{"error": "Could not unmarshal JSON: " + err.Error()}
	}

	return utils.StringPointer(result["id"].(string)), nil
}

func GetLogoUrl(uuid string) string {
	url := dotenv.GetEnv("MS_IMAGES_URL")
	return url + "images/" + uuid + "/logo.webp"
}

func DeleteImage(id string) error {
	if id == "" {
		return nil
	}

	httpClient := &http_client.HttpClient{HttpClient: &http.Client{}}
	msImage := dotenv.GetEnv("MS_IMAGES_URL")
	_, err := httpClient.Delete(msImage+"api/image/"+id, nil)

	if err != nil {
		return err
	}

	return nil
}
