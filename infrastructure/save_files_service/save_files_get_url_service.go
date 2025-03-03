package savefilesservice

import (
	jsierralibs "github.com/jSierraB3991/jsierra-libs"
)

func (c *SaveFileService) GetImageUrlTemp(imagePaths []string, folder string) ([]jsierralibs.ImageNameTemp, error) {
	var result []jsierralibs.ImageNameTemp

	for _, v := range imagePaths {
		result = append(result, jsierralibs.ImageNameTemp{
			UrlImageTemp: v,
			ImageName:    v,
		})
	}

	return result, nil
}
