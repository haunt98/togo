package transports

import (
	"errors"
	"net/http"

	"github.com/haunt98/togo/internal/services/usecases"
)

const (
	createdDateField = "created_date"
)

type TaskTransport struct {
	taskUseCase usecases.TaskUseCase
}

func (t *TaskTransport) List(rsp http.ResponseWriter, req *http.Request) {
	userID, err := getUserIDFromCtx(req.Context)
	if err != nil {
		makeJSONResponse(rsp, http.StatusInternalServerError, nil, err)
		return
	}

	createdDate := req.FormValue(createdDateField)
	if createdDate == "" {
		makeJSONResponse(rsp, http.StatusBadRequest, nil, errors.New("some errors here"))
		return
	}

	tasks, err := t.taskUseCase.List(req.Context, userID, createdDate)
	if err != nil {
		makeJSONResponse(rsp, http.StatusInternalServerError, nil, err)
		return
	}

	makeJSONResponse(rsp, http.StatusOK, tasks, nil)
}

func (t *TaskTransport) Add(rsp http.ResponseWriter, req *http.Request) {
}
