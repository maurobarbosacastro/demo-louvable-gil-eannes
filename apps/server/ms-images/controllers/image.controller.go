package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"image/jpeg"
	"io/ioutil"
	"log"
	"ms-images/forms"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

func TransformImage(c *gin.Context) {
	width, _ := strconv.Atoi(c.Query("width"))
	height, _ := strconv.Atoi(c.Query("height"))
	data := forms.FreeTranform{Id: c.Param("id"), Width: width, Height: height}

	data.Blur = c.Query("blur") != ""
	if data.Blur {
		data.BlurValue, _ = strconv.ParseFloat(c.Query("blur"), 64)
	}

	img, errTransform := FreeTransform(data)
	if errTransform != nil {
		c.IndentedJSON(http.StatusInternalServerError, errTransform)
	}

	buf := new(bytes.Buffer)
	_ = jpeg.Encode(buf, img, nil)

	_, err := c.Writer.Write(buf.Bytes())
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	return
}

func CreateImage(c *gin.Context) {
	if c.Request.ContentLength > (11 * 1024 * 1024) {
		c.IndentedJSON(http.StatusRequestEntityTooLarge, "File too big, max is 10M")
		return
	}

	file, header, errFormFile := c.Request.FormFile("file")
	fileName := c.Request.FormValue("filename")
	alt := c.Request.FormValue("alt")
	types := c.Request.FormValue("types")

	fileExt := filepath.Ext(header.Filename)
	allowedExtensions := []string{"jpeg", "jpg", "png", "webp"}

	if errFormFile != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Sprintf("Form File file error : %s", errFormFile.Error()))
		return
	}

	if fileName == "" {
		log.Print("No filename provided, using default filename: " + header.Filename)
		fileName = header.Filename
	} else {
		log.Print("Filename provided, filename: " + fileName + fileExt)
	}

	log.Printf("Extension: %v | Allowed extensions: %v", strings.Replace(strings.ToLower(fileExt), ".", "", 1), allowedExtensions)
	if !lo.Contains(allowedExtensions, strings.Replace(strings.ToLower(fileExt), ".", "", 1)) {
		c.IndentedJSON(http.StatusBadRequest, "Extension not allowed")
		return
	}

	createFile, response, err := CreateFile(
		forms.CreateFileForm{
			File: file,
			Alt:  alt,
			Name: fileName,
			Types: lo.Map(strings.Split(types, ","), func(item string, index int) string {
				return strings.TrimSpace(item)
			}),
			Size:       c.Request.ContentLength,
			Extension:  fileExt,
			FileHeader: header,
		},
		c)

	if err != nil {
		log.Print(err)
		c.IndentedJSON(createFile, err)
		return
	}

	c.IndentedJSON(createFile, response)
}

func HandleJsonRequest(c *gin.Context) {
	var request forms.JsonRequest

	// Read the request body
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Parse the JSON request
	err = json.Unmarshal(body, &request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Decode the base64-encoded file
	decodedFile, err := base64.StdEncoding.DecodeString(request.File)
	log.Printf("File decoded...")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Handle the other fields (FileName, Alt, Types) as needed
	fileName := request.FileName
	alt := request.Alt
	types := request.Types
	// Access the decoded file as needed
	// You can use the 'decodedFile' byte slice directly or write it to a file if required

	// Rest of your code
	// ...

	log.Printf("Filename is -> %s", fileName)

	fileExt := filepath.Ext(fileName)
	allowedExtensions := []string{"jpeg", "jpg", "png", "webp", "svg"}

	if fileName == "" {
		log.Print("No filename provided, using default filename")
		fileName = "default_filename" + fileExt
	} else {
		log.Print("Filename provided: " + fileName)
	}

	log.Printf("Extension: %v | Allowed extensions: %v", strings.Replace(strings.ToLower(fileExt), ".", "", 1), allowedExtensions)

	//If the image saved is SVG the ms will save the file and not transform it.
	if fileExt == ".svg" {
		createFile, response, err := CreateSVGFile(
			decodedFile,
			fileName,
			fileExt,
			c.Request.ContentLength,
			alt,
			strings.Split(types, ","),
			c)

		if err != nil {
			log.Print(err)
			c.IndentedJSON(createFile, err)
			return
		}

		c.IndentedJSON(createFile, response)

		return
	}

	createFile, response, err := CreateFileBase64(
		decodedFile,
		fileName,
		fileExt,
		c.Request.ContentLength,
		alt,
		strings.Split(types, ","),
		c)

	if err != nil {
		log.Print(err)
		c.IndentedJSON(createFile, err)
		return
	}

	c.IndentedJSON(createFile, response)
}
