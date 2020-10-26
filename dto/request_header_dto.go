package dto

type RequestHeaderDto struct {
	Referer                 string
	UserAgent               string
	Host                    string
	UpgradeInsecureRequests string
	Accept                  string
	AcceptEncoding          string
	AcceptLanguage          string
	SecFetchDest            string
	SecFetchMode            string
}
