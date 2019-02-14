//origin code: https://github.com/vadv/gopher-lua-libs/tree/master/http
package http_interface

import (
	"net/http"
	"net/url"
)

type LuaHTTPClient interface {
	DoRequest(*http.Request) (*http.Response, error)
	PostFormRequest(string, url.Values) (*http.Response, error)
}
