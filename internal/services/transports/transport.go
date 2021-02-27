package transports

import (
	"log"
	"net/http"
)

type Transport struct {
	taskTransport *TaskTransport
	userTransport *UserTransport
}

func (t *Transport) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)

	switch req.URL.Path {
	case "/login":
		return t.TaskTransport.ServeHTTP(rsp, req)
	case "/tasks":
		return t.userTransport.ServeHTTP(rsp, req)
	default:
		// TODO: return unimplement
		return
	}
}
