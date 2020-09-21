package service

type GetProxyInterface interface {
	GetContentHtml(requestUrl string) string
	ParseHtml(body string) [][]string
	GetUrlList() []string
}
