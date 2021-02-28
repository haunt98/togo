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
		newReq, ok := t.userTransport.ValidateToken(rsp, req)
		if !ok {
			return
		}

		switch req.Method {
		case http.MethodGet:
			t.taskTransport.List(rsp, newReq)
			return
		case http.MethodPost:
			t.taskTransport.Add(rsp, newReq)
			return
		default:
			// TODO: return unimplement
		}
	default:
		// TODO: return unimplement
		return
	}
}
