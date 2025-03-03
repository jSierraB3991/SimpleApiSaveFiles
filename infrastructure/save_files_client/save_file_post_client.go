package savefilesclient

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

func (c *SaveFilesClient) Post(result interface{}, uri, contentTye string, body io.Reader) error {

	req, err := http.NewRequest("POST", c.UrlBase+"/"+c.ApiVersion+"/"+uri, body)
	if err != nil {
		return err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", contentTye)
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
	} else {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error al leer la respuesta:", err)
			return err
		}

		// Imprimir el cuerpo de la respuesta
		log.Println("Respuesta:", string(body))
	}
	return errors.New(resp.Status)
}
