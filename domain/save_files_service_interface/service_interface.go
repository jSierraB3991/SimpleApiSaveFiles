package savefilesserviceinterface

import (
	"bytes"
	"mime/multipart"

	savefilesresponse "github.com/jSierraB3991/SimpleApiSaveFiles/infrastructure/save_files_response"
	jsierralibs "github.com/jSierraB3991/jsierra-libs"
)

type SaveFilesServiceeInterface interface {
	UploadImage(file *multipart.FileHeader, folder string) (string, error)
	DeleteImage(folder, path string) error
	GetImageUrlTemp(imagePaths []string, folder string) ([]jsierralibs.ImageNameTemp, error)
	GetInfoFolderService(folderName string) (*savefilesresponse.EntryResponse, error)
	GetImageByUrl(imageUrl string) (*bytes.Buffer, error)
}
