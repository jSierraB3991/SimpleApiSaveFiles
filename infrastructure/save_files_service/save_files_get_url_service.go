package savefilesservice

import (
	"bytes"

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

func (c *SaveFileService) GetImageByUrl(imageUrl string) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)

	err := c.client.GetImage(imageUrl, buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
