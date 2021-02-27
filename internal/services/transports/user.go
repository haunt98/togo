package transports

import "net/http"

type UserTransport struct {
}

func (up *UserTransport) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
}
