package httpObj

import (
	"strings"
)

const (
	GET        = "GET"
	POST       = "POST"
	RootDomain = "http://demo.dxbim.com:16001/"
	Path       = "/"
)

type Request interface {
	GetUrl() string
	GetHttpMethod() string
	SetHttpMethod(method string)
}
type Response interface {
	ParseErrorFromHttpResponse(body []byte) error
}

type BaseRequest struct {
	httpMethod string
	domain     string
	path       string
	params     map[string]string
	formParams map[string]string
}
type BaseResponse struct {
}

func (r *BaseRequest) Init() *BaseRequest {
	r.domain = ""
	r.path = Path
	r.params = make(map[string]string)
	r.formParams = make(map[string]string)
	return r
}
func (r *BaseResponse) ParseErrorFromHttpResponse(body []byte) error {
	return nil
}
func (r *BaseRequest) SetHttpMethod(method string) {
	switch strings.ToUpper(method) {
	case POST:
		r.httpMethod = POST
	case GET:
		r.httpMethod = GET
	default:
		r.httpMethod = GET
	}
}
