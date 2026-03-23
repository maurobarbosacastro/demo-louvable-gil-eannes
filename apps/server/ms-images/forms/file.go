package forms

import "mime/multipart"

type CreateFileEntityForm struct {
	Name       string   `form:"name" json:"name" binding:"required"`
	Size       int      `form:"size" json:"size" binding:"required"`
	Dimensions string   `form:"dimensions" json:"dimensions" binding:"required"`
	FileType   string   `form:"fileType" json:"fileType"`
	Extension  string   `form:"extension" json:"extension"`
	Alt        string   `form:"alt" json:"alt"`
	Types      []string `form:"types" json:"types" binding:"required"`
}

type CreateFileForm struct {
	File       multipart.File        `form:"file" json:"file" binding:"required"`
	Name       string                `form:"name" json:"name" binding:"required"`
	Extension  string                `form:"extension" json:"extension" binding:"required"`
	Size       int64                 `form:"size" json:"size" binding:"required"`
	Alt        string                `form:"alt" json:"alt" binding:"required"`
	Types      []string              `form:"types" json:"types" binding:"required"`
	FileHeader *multipart.FileHeader `form:"fileHeader" json:"fileHeader"`
}

type JsonRequest struct {
	FileName string `json:"filename"`
	Alt      string `json:"alt"`
	Types    string `json:"types"`
	File     string `json:"file"`
}
