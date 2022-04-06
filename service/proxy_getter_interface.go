package service

type ProxyGetterInterface interface {
	GetContentHtml(requestUrl string) string
	ParseHtml(body string) [][]string
	GetUrlList() []string
}
