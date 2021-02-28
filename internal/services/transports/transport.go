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
		t.userTransport.Login(rsp, req)
		return
	case "/tasks":
		t.taskTransport.ServeHTTP(rsp, req)
		return
	default:
		// TODO: return unimplement
		return
	}
}
