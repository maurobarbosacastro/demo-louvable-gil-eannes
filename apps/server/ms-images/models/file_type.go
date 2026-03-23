package models

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	db "ms-images/db"
	"time"
)

// FileTypes are going to be: Thumbnail, logo, blog, banner

type FileType struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name      string
	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func InitFileTypeModel() {
	log.Print("Init File Type Model & Data")
	localDB := db.GetDB()

	log.Print("Creating Entity FileType")
	err := localDB.AutoMigrate(&FileType{})

	if err != nil {
		log.Panicf("Error creating model file type %v", err)
		return
	}

	log.Print("Entity Created")

	log.Print("Creating Default Data")
	localDB.FirstOrCreate(&FileType{Name: "original"}, FileType{Name: "original"})
	localDB.FirstOrCreate(&FileType{Name: "resized"}, FileType{Name: "resized"})
	localDB.FirstOrCreate(&FileType{Name: "thumbnail"}, FileType{Name: "thumbnail"})
	localDB.FirstOrCreate(&FileType{Name: "thumbnailZoom"}, FileType{Name: "thumbnailZoom"})
	localDB.FirstOrCreate(&FileType{Name: "logo"}, FileType{Name: "logo"})
	localDB.FirstOrCreate(&FileType{Name: "profilePicture"}, FileType{Name: "profilePicture"})
	localDB.FirstOrCreate(&FileType{Name: "profilePictureThumbnail"}, FileType{Name: "profilePictureThumbnail"})
	localDB.FirstOrCreate(&FileType{Name: "profilePictureSmall"}, FileType{Name: "profilePictureSmall"})
	log.Print("Default Data Created")
	log.Print("-----------------------------")

}

func GetFileTypes(types []string) ([]FileType, error) {
	localDB := db.GetDB()
	var typesResult []FileType

	localDB.Where("name IN ?", types).Find(&typesResult)

	if types == nil {
		errorMsg := fmt.Sprintf("Error retrieving %v file types", types)
		log.Print(errorMsg)
		return nil, errors.New(errorMsg)
	}

	return typesResult, nil
}
