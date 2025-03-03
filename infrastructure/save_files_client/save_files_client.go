package savefilesclient

type SaveFilesClient struct {
	UrlBase      string
	ApiVersion   string
	TokenApi     string
	urlWithNoApi string
}

func NewSaveFilesClient(urlBase, urlWithNoApi, apiVersion, tokenApi string) *SaveFilesClient {
	return &SaveFilesClient{
		UrlBase:      urlBase,
		ApiVersion:   apiVersion,
		urlWithNoApi: urlWithNoApi,
		TokenApi:     tokenApi,
	}
}
