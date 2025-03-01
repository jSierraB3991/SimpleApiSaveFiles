package savefilesservice

import (
	savefileslibs "github.com/jSierraB3991/SimpleApiSaveFiles/domain/save_files_libs"
	savefilesresponse "github.com/jSierraB3991/SimpleApiSaveFiles/infrastructure/save_files_response"
)

func (c *SaveFileService) GetInfoFolderService(folderName string) (*savefilesresponse.EntryResponse, error) {
	var pagginationEntry savefilesresponse.PaginationEntry
	err := c.client.Get(&pagginationEntry, savefileslibs.ENTRIES_URI_CONST+folderName)
	if err != nil {
		return nil, err
	}

	var result savefilesresponse.EntryResponse
	for _, v := range pagginationEntry.Data {
		if v.Type == savefileslibs.FOLDER_CONST {
			result = v
		}
	}
	return &result, nil
}
