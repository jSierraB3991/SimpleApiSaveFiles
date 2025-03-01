package savefilesclient

type SaveFilesClient struct {
	UrlBase    string
	ApiVersion string
	TokenApi   string
}

func NewSaveFilesClient(urlBase, apiVersion, tokenApi string) *SaveFilesClient {
	return &SaveFilesClient{
		UrlBase:    urlBase,
		ApiVersion: apiVersion,
		TokenApi:   tokenApi,
	}
}
