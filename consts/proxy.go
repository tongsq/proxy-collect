package consts

const STATUS_NO int8 = 0
const STATUS_YES int8 = 1
const STATUS_RECHECK int8 = 2

const PROTO_HTTP = "http"
const PROTO_HTTPS = "https"
const PROTO_SOCKS4 = "socks4"
const PROTO_SOCKS4A = "socks4a"
const PROTO_SOCKS5 = "socks5"
const PROTO_SS = "ss"

// protocols
var PROTO_LIST = []string{
	PROTO_HTTP,
	PROTO_SOCKS5,
	PROTO_HTTPS,
	PROTO_SOCKS4,
	PROTO_SOCKS4A,
	PROTO_SS,
}
