package http

import (
	"net/http"

	"github.com/parnurzeal/gorequest"
)

// NewRequest 新建请求对象，默认useragent 为 chrome 75.0, 数据类型 json
func NewRequest() *Request {
	r := &Request{
		http:    gorequest.New(),
		headers: make(map[string]string),
	}
	r.Type("json")
	r.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")

	return r
}

// Request 请求结构
type Request struct {
	http      *gorequest.SuperAgent
	Req       *gorequest.Request
	headers   map[string]string
	userAgent string
	querys    []interface{}
	// mdls
}

// Type 请求提交方式，默认json
func (r *Request) Type(name string) *Request {
	r.http = r.http.Type(name)
	return r
}

// UserAgent 设置请求 user-agent，默认是 chrome 75.0
func (r *Request) UserAgent(name string) *Request {
	r.http = r.http.Set("User-Agent", name)
	return r
}

// Cookie 设置请求 Cookie
func (r *Request) Cookie(c *http.Cookie) *Request {
	r.http = r.http.AddCookie(c)
	return r
}

// Header 设置请求 Header
func (r *Request) Header(key, val string) *Request {
	r.http = r.http.Set(key, val)
	return r
}

// Proxy 设置请求代理
func (r *Request) Proxy(url string) *Request {
	r.http = r.http.Proxy(url)
	return r
}

// Query 增加查询参数
func (r *Request) Query(query interface{}) *Request {
	r.querys = append(r.querys, query)
	return r
}

// Do 发出请求，method 请求方法，url 请求地址， query 查询参数，body 请求数据，file 文件对象/地址
func (r *Request) Do(method, url string, args ...interface{}) (*Response, error) {
	var query interface{} // 查询参数
	var body interface{}  // body 数据
	var file interface{}  // 发送文件

	// get query & body
	for i, v := range args {
		switch i {
		case 0:
			query = v
		case 1:
			body = v
		case 2:
			file = v
		}

	}

	// set mthod url
	r.http = r.http.CustomMethod(method, url)

	// set query string
	if query != nil {
		r.http = r.http.Query(query)
	}
	for _, q := range r.querys {
		r.http = r.http.Query(q)
	}

	// set body
	if body != nil {
		r.http = r.http.Send(body)
	}

	if file != nil {
		r.Type("multipart")
		r.http = r.http.SendFile(file)
	}

	res, resBody, errs := r.http.EndBytes()

	response := &Response{
		Request: r,
		Raw:     &res,
		Body:    resBody,
		Errs:    errs,
	}

	return response, response.Err()
}

// Head 发起 head 请求
func (r *Request) Head(url string, query interface{}) (*Response, error) {
	return r.Do("HEAD", url, query, nil, nil)
}

// Get 发起 get 请求， query 查询参数
func (r *Request) Get(url string, query interface{}) (*Response, error) {
	return r.Do("GET", url, query, nil, nil)
}

// Post 发起 post 请求，body 是请求带的参数，可使用json字符串或者结构体
func (r *Request) Post(url string, body interface{}) (*Response, error) {
	return r.Do("POST", url, nil, body, nil)
}

// Put 发起 put 请求，body 是请求带的参数，可使用json字符串或者结构体
func (r *Request) Put(url string, body interface{}) (*Response, error) {
	return r.Do("PUT", url, nil, body, nil)
}

// Del 发起 delete 请求，body 是请求带的参数，可使用json字符串或者结构体
func (r *Request) Del(url string, body interface{}) (*Response, error) {
	return r.Do("DELETE", url, nil, body, nil)
}

// Patch 发起 patch 请求，body 是请求带的参数，可使用json字符串或者结构体
func (r *Request) Patch(url string, body interface{}) (*Response, error) {
	return r.Do("PATCH", url, nil, body, nil)
}

// Options 发起 options 请求，query 查询参数
func (r *Request) Options(url string, query interface{}) (*Response, error) {
	return r.Do("OPTIONS", url, query, nil, nil)
}

// PostFile 发起 post 请求上传文件，将使用表单提交，file 是文件地址或者文件流， body 是请求带的参数，可使用json字符串或者结构体
func (r *Request) PostFile(url string, file interface{}, body interface{}) (*Response, error) {
	return r.Do("PUT", url, nil, body, file)
}

// PutFile 发起 put 请求上传文件，将使用表单提交，file 是文件地址或者文件流， body 是请求带的参数，可使用json字符串或者结构体
func (r *Request) PutFile(url string, file interface{}, body interface{}) (*Response, error) {
	return r.Do("PUT", url, nil, body, file)
}
