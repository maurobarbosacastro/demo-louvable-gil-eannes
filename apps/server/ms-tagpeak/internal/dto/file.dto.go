package dto

type CreateFileDto struct {
	FileName  string `json:"fileName,omitempty"`
	Extension string `json:"extension,omitempty"`
}
