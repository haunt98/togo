package transports

import (
	"net/http"

	"github.com/haunt98/togo/internal/services/usecases"
)

type TaskTransport struct {
	taskUseCase usecases.TaskUseCase
}

func (t *TaskTransport) List(rsp http.ResponseWriter, req *http.Request) {
}

func (t *TaskTransport) Add(rsp http.ResponseWriter, req *http.Request) {
}
