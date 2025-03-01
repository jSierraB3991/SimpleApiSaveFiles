package savefilesresponse

type FileResponse struct {
	FileEntry EntryResponse `json:"fileEntry"`
	Status    string        `json:"status"`
}
