package service

type ProxyGetterInterface interface {
	GetContentHtml(requestUrl string) string
	/**
	result format:
	[
		["ip", "port", "protocol", "user", "password"]
	]
	*/
	ParseHtml(body string) [][]string
	GetUrlList() []string
}
