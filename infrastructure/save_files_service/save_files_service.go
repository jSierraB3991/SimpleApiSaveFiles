package savefilesservice

import savefilesclient "github.com/jSierraB3991/SimpleApiSaveFiles/infrastructure/save_files_client"

type SaveFileService struct {
	randomLetters int
	client        *savefilesclient.SaveFilesClient
}

func NewSaveFileService(lettersToAddToFiles int, apiVersion, apiToken string) *SaveFileService {
	return &SaveFileService{
		randomLetters: lettersToAddToFiles,
		client:        savefilesclient.NewSaveFilesClient("https://savefiles.org/api", apiVersion, apiToken),
	}
}
