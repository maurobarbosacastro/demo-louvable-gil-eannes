package models

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	db "ms-images/db"
	"ms-images/forms"
	"ms-images/pkg/dotenv"
	"os"
	"time"
)

type File struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name       string
	Size       int `gorm:"default:1"`
	Dimensions string
	FileTypes  []FileType `gorm:"many2many:file_file_types;foreignKey:id"`
	Extension  string
	Alt        string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func InitFileModel() {
	log.Print("Init File Model & Data")
	localDB := db.GetDB()

	log.Print("Creating Entity File")
	err := localDB.AutoMigrate(&File{})

	if err != nil {
		log.Panicf("Error creating model file %v", err)
		return
	}

	log.Print("Entity Created")

	log.Print("Creating Default Data")

	log.Print("Default Data Created")
	log.Print("-----------------------------")
}

func CreateFile(createForm forms.CreateFileEntityForm) File {
	getDb := db.GetDB()

	fileTypes, _ := GetFileTypes(createForm.Types)

	file := File{
		Name:       createForm.Name,
		Size:       createForm.Size,
		Dimensions: createForm.Dimensions,
		FileTypes:  fileTypes,
		Extension:  createForm.Extension,
		Alt:        createForm.Alt,
	}

	getDb.Create(&file)

	return file
}

func GetFile(id string) File {
	getDb := db.GetDB()

	var file File
	uuidId, _ := uuid.Parse(id)
	getDb.Preload("FileTypes").Model(&File{}).Where("id = ?", uuidId).Find(&file)

	return file

}

func DeleteFile(id string) (File, error) {
	getDb := db.GetDB()

	var file File

	// Parse the UUID and handle any potential error
	uuidId, err := uuid.Parse(id)
	if err != nil {
		return file, fmt.Errorf("invalid UUID: %v", err)
	}

	// Find the file record based on the id
	if err := getDb.Preload("FileTypes").Where("id = ?", uuidId).First(&file).Error; err != nil {
		return file, fmt.Errorf("file not found: %v", err)
	}

	// Delete the file record in the database
	if err := getDb.Delete(&file).Error; err != nil {
		return file, fmt.Errorf("failed to delete file: %v", err)
	}

	err = DeleteFolderFromStorage(file.ID.String()) // Assuming file.ID is a UUID or some unique identifier
	if err != nil {
		log.Printf("Error deleting folder from storage: %v", err)
	}
	// You would delete the file from storage here (assuming external file storage is used)

	// Return the deleted file
	return file, nil
}

func DeleteFolderFromStorage(fileID string) error {

	// Build the folder path
	folderPath := dotenv.GetEnv("SERVER_IMAGES") + "/" + fileID

	// Check if the folder exists
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// Folder doesn't exist
		return nil
	}

	// Delete the folder and its contents
	err := os.RemoveAll(folderPath)
	if err != nil {
		return err
	}

	return nil
}
