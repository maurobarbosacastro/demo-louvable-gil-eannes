package images

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"ms-tagpeak/internal/response_object"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/http_client"
	"ms-tagpeak/pkg/utils"
	"net/http"
)

func UploadImage(file *multipart.FileHeader, fileName string, types string) (*string, map[string]string) {
	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, map[string]string{"error": "Could not open file"}
	}
	defer src.Close()

	responseUuid, errMap := createImage(file, src, fileName, types)
	if errMap != nil {
		fmt.Printf("Error: %v\n", errMap)
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
		fmt.Printf("Error: %v\n", err)
		return nil, map[string]string{"error": "Could not create form file"}
	}

	// Copy file contents to the form file field
	_, err = io.Copy(part, src)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, map[string]string{"error": "Could not copy file contents"}
	}

	// Add any additional form fields if needed
	writer.WriteField("fileName", fileName)
	writer.WriteField("types", types)

	// Close the multipart writer
	err = writer.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, map[string]string{"error": "Could not close multipart writer"}
	}

	msImagesUrl := dotenv.GetEnv("MS_IMAGES_URL")
	// Create the HTTP request to send the file
	req, err := http.NewRequest("POST", msImagesUrl+"api/image/", body)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, map[string]string{"error": "Could not create HTTP request"}
	}

	token := http_client.GetInternalKeycloakToken()
	// Set the content type to multipart/form-data
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+token)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, map[string]string{"error": "Could not send file to destination"}
	}
	defer resp.Body.Close()

	// Read the response
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, map[string]string{"error": "Could not read response"}
	}

	var result map[string]interface{}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, map[string]string{"error": "Could not unmarshal JSON: " + err.Error()}
	}

	return utils.StringPointer(result["id"].(string)), nil
}

func GetLogoUrl(uuid string) string {
	url := dotenv.GetEnv("MS_IMAGES_URL")
	return url + "images/" + uuid + "/logo.webp"
}

func GetLogoFromDomain(domain string) (*response_object.AuxLogo, error) {
	httpClient := &http_client.InternalHttpClient{InternalHttpClient: &http.Client{}}
	msImage := dotenv.GetEnv("MS_IMAGES_URL")

	logo := &response_object.AuxLogo{}
	_, err := httpClient.Get(msImage+"aux/ai-logo?name="+domain, nil, logo)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	return logo, nil
}

func GetLogoFromUrl(url string, name string) (*response_object.AuxLogo, error) {

	httpClient := &http_client.InternalHttpClient{InternalHttpClient: &http.Client{}}

	msImage := dotenv.GetEnv("MS_IMAGES_URL")

	logo := &response_object.AuxLogo{}
	body := map[string]string{"url": url, "types": "logo", "name": name}
	_, err := httpClient.Post(msImage+"aux/url", nil, body, logo)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	return logo, nil

}

func GetBannerFromUrl(url string, name string) (*response_object.AuxLogo, error) {

	httpClient := &http_client.InternalHttpClient{InternalHttpClient: &http.Client{}}
	msImage := dotenv.GetEnv("MS_IMAGES_URL")

	logo := &response_object.AuxLogo{}
	body := map[string]string{"url": url, "types": "resized", "name": name}
	_, err := httpClient.Post(msImage+"aux/url", nil, body, logo)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	return logo, nil

}

func DeleteImage(id string) error {
	if id == "" {
		return nil
	}

	httpClient := &http_client.InternalHttpClient{InternalHttpClient: &http.Client{}}
	msImage := dotenv.GetEnv("MS_IMAGES_URL")
	_, err := httpClient.Delete(msImage+"api/image/"+id, nil)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	return nil
}

func GetProfilePictureFromUrl(url string, name string) (*response_object.AuxProfilePicture, error) {
	httpClient := &http_client.InternalHttpClient{InternalHttpClient: &http.Client{}}
	msImage := dotenv.GetEnv("MS_IMAGES_URL")

	profilePicture := &response_object.AuxProfilePicture{}
	body := map[string]string{"url": url, "types": "profilePicture", "name": name}
	_, err := httpClient.Post(msImage+"aux/url", nil, body, profilePicture)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	return profilePicture, nil

}
