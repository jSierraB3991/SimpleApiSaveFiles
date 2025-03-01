package savefilesserviceinterface

import (
	"mime/multipart"

	savefilesresponse "github.com/jSierraB3991/SimpleApiSaveFiles/infrastructure/save_files_response"
)

type SaveFilesServiceeInterface interface {
	UploadImage(file *multipart.FileHeader, folder string) (string, error)
	DeleteImage(folder, path string) error
	GetImageUrlTemp(imagePaths []string, folder string) ([]interface{}, error)
	GetInfoFolderService(folderName string) (*savefilesresponse.EntryResponse, error)
}
