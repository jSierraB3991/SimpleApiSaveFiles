package savefilesservice

import (
	"bytes"
	"io"
	"mime/multipart"

	savefileslibs "github.com/jSierraB3991/SimpleApiSaveFiles/domain/save_files_libs"
	savefilesresponse "github.com/jSierraB3991/SimpleApiSaveFiles/infrastructure/save_files_response"
	jsierralibs "github.com/jSierraB3991/jsierra-libs"
)

func (c *SaveFileService) UploadImage(file *multipart.FileHeader, folder string) (string, error) {

	folderInSaveFiles, err := c.GetInfoFolderService(folder)
	if err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}

	defer src.Close()
	cadena, err := jsierralibs.GenerateRandomCode(c.randomLetters, jsierralibs.Lowercase)
	if err != nil {
		return "", err
	}

	body := bytes.Buffer{}
	writer := multipart.NewWriter(&body)

	// Adjunta el archivo
	part, _ := writer.CreateFormFile("file", cadena+file.Filename)
	_, _ = io.Copy(part, src)

	// Agrega los otros campos
	_ = writer.WriteField("parentId", jsierralibs.GetStringNumberFor(folderInSaveFiles.ID))
	_ = writer.WriteField("relativePath", "")

	// Finaliza el writer
	err = writer.Close()
	if err != nil {
		return "", err
	}

	contentType := writer.FormDataContentType()
	var result savefilesresponse.FileResponse
	err = c.client.Post(&result, savefileslibs.UPLOAD_FILE_URI_CONST, contentType, &body)
	if err != nil {
		return "", err
	}
	return *result.FileEntry.URL, nil
}
