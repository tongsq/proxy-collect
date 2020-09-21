package service

import "io"

type GetProxyInterface interface {
	GetContentHtml(requestUrl string) io.ReadCloser
	ParseHtml(body io.ReadCloser) [][]string
	GetUrlList() []string
}
