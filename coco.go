package coco

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Coco Framework
type Coco struct {
	httpRouter              *httprouter.Router
	Request                 *http.Request
	ResponseWriter          http.ResponseWriter
	Params                  httprouter.Params
	AdditionResponseHeaders map[string]string
}

func NewCoco() *Coco {
	c := Coco{
		httpRouter:              httprouter.New(),
		AdditionResponseHeaders: make(map[string]string),
	}
	return &c
}

//ServeHTTP 实现了该方法的对象，可以作为 http.ListenAndServe 的 handler
// http server 在监听到请求之后会把请求交给 ServeHTTP 处理，主要处理方式就是路由并分发给相应的 handler
func (c *Coco) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//c.Request = r
	//c.ResponseWriter = w
	//handler := c.Router.data[r.RequestURI]
	//handler(c)
	for key, val := range c.AdditionResponseHeaders {
		w.Header().Set(key, val)
	}
	c.httpRouter.ServeHTTP(w, r)
}

//func (c *Coco) AddRouter(method string, path string, f func(coco *Coco)) {
//	c.Router.data[path] = f
//}

func (c *Coco) AddRouter(method string, path string, h Handler) {
	c.httpRouter.Handler(method, path, h)
}

func (c *Coco) Run(host string, port int) error {
	addr := fmt.Sprintf("%s:%d", host, port)
	return http.ListenAndServe(addr, c)
}
