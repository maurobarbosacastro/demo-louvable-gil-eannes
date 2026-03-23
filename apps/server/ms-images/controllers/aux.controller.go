package controllers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"image"
	"io"
	"io/ioutil"
	"ms-images/forms"
	"ms-images/models"
	"ms-images/pkg/dotenv"
	"ms-images/utils"
	"net/http"
	"os"
	"strings"
)

var url = "https://img.logo.dev/"
var urlExtra = "?token=pk_RjCwIdHdS2m9TOt7Nb7PhQ&size=277&format=png&retina=true"

func GetLogo(c *gin.Context) {
	nameSearch := c.Query("name")

	// Download the file
	resp, err := http.Get(url + nameSearch + urlExtra)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fileDB := models.CreateFile(forms.CreateFileEntityForm{
		Name:       nameSearch,
		Size:       0,
		Dimensions: "514x514",
		Extension:  ".png",
		Alt:        "",
		Types:      []string{"original", "logo"},
	})

	generatedImages, err := saveAndTransformImage(fileDB, c, body, []string{"logo"})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, generatedImages)
	return
}

func GetImageFromUrl(c *gin.Context) {
	fmt.Printf("\nGetImageFromUrl\n")
	var bodyRequest map[string]string
	if err := c.ShouldBindJSON(&bodyRequest); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	fileTypes := strings.Split(bodyRequest["types"], ",")

	url := bodyRequest["url"]

	// Download the file
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fileDB := models.CreateFile(forms.CreateFileEntityForm{
		Name:      bodyRequest["name"],
		Size:      0,
		Extension: ".png",
		Alt:       "",
		Types:     append(fileTypes, []string{"original"}...),
	})

	generatedImages, err := saveAndTransformImage(fileDB, c, body, fileTypes)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, generatedImages)
	return
}

func saveAndTransformImage(fileDB models.File, c *gin.Context, imageBody []byte, fileTypes []string) (*gin.H, error) {

	publicPath := dotenv.GetEnv("SERVER_IMAGES")

	errDir := os.MkdirAll(publicPath+"/"+fileDB.ID.String(), os.ModePerm)

	if errDir != nil {
		fmt.Printf("Error: %v\n", errDir)
		return nil, errDir
	}

	generatedImages := gin.H{
		"id":       fileDB.ID,
		"original": os.Getenv("MS_IMAGES_SERVER_PUBLIC_URL") + fileDB.ID.String() + "/original" + ".png",
	}

	file, err := os.Create(publicPath + "/" + fileDB.ID.String() + "/original.png")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}
	defer file.Close()

	_, err = io.Copy(file, bytes.NewBuffer(imageBody))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	imageFile, _, err := image.Decode(bytes.NewBuffer(imageBody))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	orientation := utils.GetOrientation(imageFile)

	lo.ForEach(fileTypes, func(fileType string, index int) {
		switch strings.ToLower(fileType) {
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

	return &generatedImages, nil
}
