package savefilesclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *SaveFilesClient) Get(result interface{}, uri string) error {

	req, err := http.NewRequest("GET", c.UrlBase+"/"+c.ApiVersion+"/"+uri, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.TokenApi)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New(resp.Status)
}

func (c *SaveFilesClient) GetImage(imageUrl string, result *bytes.Buffer) error {
	req, err := http.NewRequest("GET", c.urlWithNoApi+"/"+imageUrl, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.TokenApi)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if _, err := io.Copy(result, resp.Body); err != nil {
		return err
	}

	return nil
}
