package transports

import (
	"net/http"

	"github.com/haunt98/togo/internal/services/usecases"
)

type TaskTransport struct {
	taskUseCase usecases.TaskUseCase
}

func (tp *TaskTransport) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {

}
