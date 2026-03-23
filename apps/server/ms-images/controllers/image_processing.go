package controllers

import (
	"fmt"
	"github.com/HugoSmits86/nativewebp"
	"github.com/disintegration/imaging"
	"github.com/samber/lo"
	"image"
	"log"
	"ms-images/forms"
	"ms-images/models"
	"ms-images/pkg/dotenv"
	"ms-images/utils"
	"os"
)

func ProcessResizedImage(image image.Image, filename string, fileExt string) {
	log.Print("ProcessResizedImage start")

	src := imaging.Resize(
		image,
		image.Bounds().Dx()/2,
		image.Bounds().Dy()/2,
		imaging.CatmullRom,
	)

	err := saveWebp(fmt.Sprintf(dotenv.GetEnv("SERVER_IMAGES")+"/%s/%v", filename, "resized"+fileExt), src)
	if err != nil {
		fmt.Printf("Error: Fail to save resized image: %v\n", err)
	}

	log.Print("ProcessResizedImage end")
}

func ProcessThumbnailImage(image image.Image, filename string, fileExt string, orientation string) {

	log.Print("ProcessThumbnailImage start")
	filenameThumbnail := "thumbnail" + fileExt

	src := imaging.Resize(
		image,
		lo.Ternary(utils.IsHorizontal(orientation), 150, 0),
		lo.Ternary(utils.IsHorizontal(orientation), 0, 150),
		imaging.CatmullRom,
	)
	err := saveWebp(fmt.Sprintf(dotenv.GetEnv("SERVER_IMAGES")+"/%s/%v", filename, filenameThumbnail), src)
	if err != nil {
		fmt.Printf("Error: Fail to save thumbnail image: %v\n", err)
	}

	filenameThumbnailZoom := "thumbnailZoom" + fileExt
	src = imaging.Resize(
		image,
		lo.Ternary(utils.IsHorizontal(orientation), 500, 0),
		lo.Ternary(utils.IsHorizontal(orientation), 0, 500),
		imaging.CatmullRom,
	)

	err = saveWebp(fmt.Sprintf(dotenv.GetEnv("SERVER_IMAGES")+"/%s/%v", filename, filenameThumbnailZoom), src)
	if err != nil {
		fmt.Printf("Error: Fail to save thumbnail zoom image: %v\n", err)
	}

	log.Print("ProcessThumbnailImage end")

}

func ProcessLogoImage(image image.Image, filename string, fileExt string, orientation string) {
	log.Print("ProcessLogoImage start")
	filenameThumbnail := "logo" + fileExt

	width := 250
	height := 0

	if utils.IsSquare(orientation) {
		width = 160
		height = 160
	} else if utils.IsVertical(orientation) {
		width = 0
		height = 250
	}

	src := imaging.Resize(
		image,
		width,
		height,
		imaging.CatmullRom,
	)
	err := saveWebp(fmt.Sprintf(dotenv.GetEnv("SERVER_IMAGES")+"/%s/%v", filename, filenameThumbnail), src)

	if err != nil {
		fmt.Printf("Error: Fail to save logo image: %v\n", err)
	}

	log.Print("ProcessLogoImage end")

}

func ProcessProfilePictureImage(image image.Image, filename string, fileExt string, orientation string) {
	log.Print("ProcessProfilePictureImage start")
	filenameProfilePicture := "profilePicture" + fileExt

	width := 400
	height := 0

	if utils.IsSquare(orientation) {
		width = 400
		height = 400
	} else if utils.IsVertical(orientation) {
		width = 0
		height = 400
	}

	src := imaging.Resize(
		image,
		width,
		height,
		imaging.CatmullRom,
	)
	err := saveWebp(fmt.Sprintf(dotenv.GetEnv("SERVER_IMAGES")+"/%s/%v", filename, filenameProfilePicture), src)

	if err != nil {
		fmt.Printf("Error: Fail to save profile picture image: %v\n", err)
	}

	width = 150
	if utils.IsSquare(orientation) {
		width = 150
		height = 150
	} else if utils.IsVertical(orientation) {
		width = 0
		height = 150
	}

	filenameProfilePictureThumbnail := "profilePictureThumbnail" + fileExt
	src = imaging.Resize(
		image,
		width,
		height,
		imaging.CatmullRom,
	)
	err = saveWebp(fmt.Sprintf(dotenv.GetEnv("SERVER_IMAGES")+"/%s/%v", filename, filenameProfilePictureThumbnail), src)

	if err != nil {
		fmt.Printf("Error: Fail to save profile picture image: %v", err)
	}

	width = 32
	if utils.IsSquare(orientation) {
		width = 32
		height = 32
	} else if utils.IsVertical(orientation) {
		width = 0
		height = 32
	}

	filenameProfilePictureSmall := "profilePictureSmall" + fileExt
	src = imaging.Resize(
		image,
		width,
		height,
		imaging.CatmullRom,
	)
	err = saveWebp(fmt.Sprintf(dotenv.GetEnv("SERVER_IMAGES")+"/%s/%v", filename, filenameProfilePictureSmall), src)

	if err != nil {
		fmt.Printf("Error: Fail to save profile picture image: %v", err)
	}

	log.Print("ProcessProfilePictureImage end")
}

func FreeTransform(freeTransformParams forms.FreeTranform) (image.Image, error) {
	file := models.GetFile(freeTransformParams.Id)

	log.Printf("Resizing with width %v and height %v", freeTransformParams.Width, freeTransformParams.Height)

	imageOriginal, errOpen := imaging.Open("./" + dotenv.GetEnv("SERVER_IMAGES") + "/" + freeTransformParams.Id + "/original" + file.Extension)

	if errOpen != nil {
		fmt.Printf("Error: Fail to open image: %v\n", errOpen)
		return nil, errOpen
	}

	src := imageOriginal

	if freeTransformParams.Width != 0 || freeTransformParams.Height != 0 {
		src = imaging.Resize(
			imageOriginal,
			freeTransformParams.Width,
			freeTransformParams.Height,
			imaging.CatmullRom,
		)
	}

	log.Printf("Add blur: %v", freeTransformParams.Blur)
	if freeTransformParams.Blur {
		src = imaging.Blur(src, freeTransformParams.BlurValue)
	}

	return src, nil

}

func saveWebp(name string, img *image.NRGBA) error {
	f, err := os.Create(name)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	if err := nativewebp.Encode(f, img); err != nil {
		err := f.Close()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return err
		}
		return err
	}

	if err := f.Close(); err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	return nil
}
