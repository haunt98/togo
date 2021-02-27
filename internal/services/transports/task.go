package transports

import "net/http"

type TaskTransport struct {
}

func (tp *TaskTransport) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
}
