package savefilesclient

import (
	"encoding/json"
	"errors"
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
