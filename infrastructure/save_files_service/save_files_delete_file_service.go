package savefilesservice

import "log"

func (c *SaveFileService) DeleteImage(folder, path string) error {
	log.Printf("folder %s path%s", folder, path)
	return nil
}
