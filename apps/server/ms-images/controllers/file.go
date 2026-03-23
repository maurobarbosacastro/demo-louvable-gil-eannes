package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"image"
	"io"
	"log"
	"ms-images/forms"
	"ms-images/models"
	"ms-images/pkg/dotenv"
	"ms-images/utils"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const webpExtension = ".webp"

func GetFile(id string) models.File {
	return models.GetFile(id)
}

func processFileTypes(types []string) []string {
	processedTypes := []string{"original"}
	lo.ForEach(types, func(item string, index int) {
		switch strings.ToLower(strings.TrimSpace(item)) {
		case "resized":
			//resize resizes
			processedTypes = append(processedTypes, "resized")
			break
		case "thumbnail":
			//resize thumbnail and zoom
			processedTypes = append(processedTypes, "thumbnail", "thumbnailZoom")
			break
		case "logo":
			//resize logo
			processedTypes = append(processedTypes, "logo")
			break
		case "profilepicture":
			//resize profile picture
			processedTypes = append(processedTypes, "profilePicture", "profilePictureThumbnail", "profilePictureSmall")
			break
		}
	})
	return processedTypes
}

func CreateFile(fileForm forms.CreateFileForm, c *gin.Context) (int, any, error) {
	log.Print("------------------------")
	log.Print("Create File init")
	log.Print(fileForm)

	imageFile, _, err := image.Decode(fileForm.File)
	if err != nil {
		return http.StatusInternalServerError, nil, errors.New(fmt.Sprintf("Error decoding image: %v", err))
	}

	fileDB := models.CreateFile(forms.CreateFileEntityForm{
		Name:       fileForm.Name,
		Size:       int(fileForm.Size),
		Dimensions: fmt.Sprintf("%vx%v", imageFile.Bounds().Dx(), imageFile.Bounds().Dy()),
		Extension:  fileForm.Extension,
		Alt:        fileForm.Alt,
		Types:      processFileTypes(fileForm.Types),
	})

	log.Print(dotenv.GetEnv("SERVER_IMAGES") + "/" + fileDB.ID.String())
	errDir := os.MkdirAll(dotenv.GetEnv("SERVER_IMAGES")+"/"+fileDB.ID.String(), os.ModePerm)

	if errDir != nil {
		return http.StatusInternalServerError, nil, errors.New(fmt.Sprintf("Error creating directory: %v", errDir))
	}

	orientation := utils.GetOrientation(imageFile)
	log.Printf("Image orientation: %v", orientation)

	generatedImages := gin.H{
		"id":       fileDB.ID,
		"original": os.Getenv("MS_IMAGES_SERVER_PUBLIC_URL") + fileDB.ID.String() + "/original" + fileForm.Extension,
	}

	errSaving := c.SaveUploadedFile(fileForm.FileHeader, fmt.Sprintf(dotenv.GetEnv("SERVER_IMAGES")+"/%s/%v", fileDB.ID.String(), "original"+fileForm.Extension))
	if errSaving != nil {
		return http.StatusInternalServerError, nil, errSaving
	}

	// Resize according to types given by user
	// original will always be saved
	// thumbnail will always generate thumbnailZoom too
	lo.ForEach(fileForm.Types, func(item string, index int) {
		switch strings.ToLower(strings.TrimSpace(item)) {
		case "resized":
			//resize resizes
			go ProcessResizedImage(imageFile, fileDB.ID.String(), webpExtension)
			generatedImages = lo.Assign(
				generatedImages,
				gin.H{"resized": os.Getenv("MS_IMAGES_SERVER_PUBLIC_URL") + fileDB.ID.String() + "/resized" + webpExtension},
			)
			break
		case "thumbnail":
			//resize thumbnail and zoom
			go ProcessThumbnailImage(imageFile, fileDB.ID.String(), webpExtension, orientation)
			generatedImages = lo.Assign(
				generatedImages,
				gin.H{
					"thumbnail":     os.Getenv("MS_IMAGES_SERVER_PUBLIC_URL") + fileDB.ID.String() + "/thumbnail" + webpExtension,
					"thumbnailZoom": os.Getenv("MS_IMAGES_SERVER_PUBLIC_URL") + fileDB.ID.String() + "/thumbnailZoom" + webpExtension,
				},
			)
			break
		case "logo":
			//resize logo
			go ProcessLogoImage(imageFile, fileDB.ID.String(), webpExtension, orientation)
			generatedImages = lo.Assign(
				generatedImages,
				gin.H{"logo": os.Getenv("MS_IMAGES_SERVER_PUBLIC_URL") + fileDB.ID.String() + "/logo" + webpExtension},
			)
			break
		case "profilepicture":
			//resize profile picture
			go ProcessProfilePictureImage(imageFile, fileDB.ID.String(), webpExtension, orientation)
			generatedImages = lo.Assign(
				generatedImages,
				gin.H{
					"profilePicture":          os.Getenv("MS_IMAGES_SERVER_PUBLIC_URL") + fileDB.ID.String() + "/profilePicture" + webpExtension,
					"profilePictureThumbnail": os.Getenv("MS_IMAGES_SERVER_PUBLIC_URL") + fileDB.ID.String() + "/profilePictureThumbnail" + webpExtension,
					"profilePictureSmall":     os.Getenv("MS_IMAGES_SERVER_PUBLIC_URL") + fileDB.ID.String() + "/profilePictureSmall" + webpExtension,
				},
			)
			break
		}
	})

	return http.StatusCreated, generatedImages, nil

}

func CreateFileBase64(file []byte, name string, extension string, size int64, alt string, types []string, c *gin.Context) (int, any, error) {
	log.Print("------------------------")
	log.Print("Create File init")
	log.Print(name)
	log.Print(extension)
	log.Print(file)

	log.Printf("--------------------")
	imageFile, _, err := image.Decode(bytes.NewReader(file))
	log.Printf("--------------------")
	log.Printf("%s %d %s %s %s", name, size, extension, alt, types)

	if err != nil {
		return http.StatusInternalServerError, nil, errors.New(fmt.Sprintf("Error decoding image: %v", err))
	}

	fileDB := models.CreateFile(forms.CreateFileEntityForm{
		Name:       name,
		Size:       int(size),
		Dimensions: fmt.Sprintf("%vx%v", imageFile.Bounds().Dx(), imageFile.Bounds().Dy()),
		Extension:  extension,
		Alt:        alt,
		Types:      append(types, []string{"original"}...),
	})

	log.Print(dotenv.GetEnv("SERVER_IMAGES") + "/" + fileDB.ID.String())
	errDir := os.MkdirAll(dotenv.GetEnv("SERVER_IMAGES")+"/"+fileDB.ID.String(), os.ModePerm)

	if errDir != nil {
		return http.StatusInternalServerError, nil, errors.New(fmt.Sprintf("Error creating directory: %v", errDir))
	}

	orientation := utils.GetOrientation(imageFile)
	log.Printf("Image orientation: %v", orientation)

	generatedImages := gin.H{
		"id":       fileDB.ID,
		"original": os.Getenv("MS_IMAGES_SERVER_PUBLIC_URL") + fileDB.ID.String() + "/original" + extension,
	}

	errSaving := CustomSave(file, fmt.Sprintf(dotenv.GetEnv("SERVER_IMAGES")+"/%s/%v", fileDB.ID.String(), "original"+extension), c)

	if errSaving != nil {
		return http.StatusInternalServerError, nil, errSaving
	}

	// Resize according to types given by user
	// original will always be saved
	// thumbnail will always generate thumbnailZoom too
	lo.ForEach(types, func(item string, index int) {
		switch strings.ToLower(strings.TrimSpace(item)) {
		case "resized":
			//resize resizes
			go ProcessResizedImage(imageFile, fileDB.ID.String(), webpExtension)
			generatedImages = lo.Assign(
				generatedImages,
				gin.H{"resized": os.Getenv("MS_IMAGES_SERVER_PUBLIC_URL") + fileDB.ID.String() + "/resized" + webpExtension},
			)
			break
		case "thumbnail":
			//resize thumbnail and zoom
			go ProcessThumbnailImage(imageFile, fileDB.ID.String(), webpExtension, orientation)
			generatedImages = lo.Assign(
				generatedImages,
				gin.H{
					"thumbnail":     os.Getenv("MS_IMAGES_SERVER_PUBLIC_URL") + fileDB.ID.String() + "/thumbnail" + webpExtension,
					"thumbnailZoom": os.Getenv("MS_IMAGES_SERVER_PUBLIC_URL") + fileDB.ID.String() + "/thumbnailZoom" + webpExtension,
				},
			)
			break
		case "logo":
			//resize logo
			go ProcessLogoImage(imageFile, fileDB.ID.String(), webpExtension, orientation)
			generatedImages = lo.Assign(
				generatedImages,
				gin.H{"logo": os.Getenv("MS_IMAGES_SERVER_PUBLIC_URL") + fileDB.ID.String() + "/logo" + webpExtension},
			)
			break
		}
	})

	return http.StatusCreated, generatedImages, nil
}

// This function save the file without transform
func CreateSVGFile(file []byte, name string, extension string, size int64, alt string, types []string, c *gin.Context) (int, any, error) {

	log.Printf("Create SVG file")
	log.Printf("---------------")

	fileDB := models.CreateFile(forms.CreateFileEntityForm{
		Name:       name,
		Size:       int(size),
		Dimensions: fmt.Sprintf("svg file"),
		Extension:  extension,
		Alt:        alt,
		Types:      append(types, []string{"original"}...),
	})

	log.Print(dotenv.GetEnv("SERVER_IMAGES") + "/" + fileDB.ID.String())
	errDir := os.MkdirAll(dotenv.GetEnv("SERVER_IMAGES")+"/"+fileDB.ID.String(), os.ModePerm)

	if errDir != nil {
		return http.StatusInternalServerError, nil, errors.New(fmt.Sprintf("Error creating directory: %v", errDir))
	}

	generatedImages := gin.H{
		"id":       fileDB.ID,
		"original": os.Getenv("MS_IMAGES_SERVER_PUBLIC_URL") + fileDB.ID.String() + "/original" + extension,
	}

	errSaving := CustomSave(file, fmt.Sprintf(dotenv.GetEnv("SERVER_IMAGES")+"/%s/%v", fileDB.ID.String(), "original"+extension), c)

	if errSaving != nil {
		return http.StatusInternalServerError, nil, errSaving
	}

	return http.StatusCreated, generatedImages, nil
}

func CustomSave(text []byte, dst string, g *gin.Context) error {

	os.MkdirAll(filepath.Dir(dst), 0750)

	out, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, bytes.NewReader(text))
	return err
}

func DeleteFile(id string) error {
	file, err := models.DeleteFile(id)

	if err != nil {
		log.Printf("Error deleting file: %v", err)
		return err
	}
	log.Printf("Dleting file: %v", file.ID.String())
	return nil
}
