package savefilesservice

import (
	savefileslibs "github.com/jSierraB3991/SimpleApiSaveFiles/domain/save_files_libs"
	savefilesresponse "github.com/jSierraB3991/SimpleApiSaveFiles/infrastructure/save_files_response"
	jsierralibs "github.com/jSierraB3991/jsierra-libs"
)

func (c *SaveFileService) GetImageUrlTemp(imagePaths []string, folder string) ([]jsierralibs.ImageNameTemp, error) {
	var result []jsierralibs.ImageNameTemp

	for _, v := range imagePaths {
		var resultTemp savefilesresponse.PaginationEntry
		err := c.client.Get(&resultTemp, savefileslibs.SEARCH_ENTRIES+v)
		if err != nil {
			return nil, err
		}
		result = append(result, jsierralibs.ImageNameTemp{
			UrlImageTemp: *resultTemp.Data[0].URL,
			ImageName:    v,
		})
	}

	return result, nil
}
