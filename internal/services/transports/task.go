package transports

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/haunt98/togo/internal/services/usecases"
	"github.com/haunt98/togo/internal/storages"
)

type TaskTransport struct {
	taskUseCase usecases.TaskUseCase
}

func (t *TaskTransport) ListTasks(rsp http.ResponseWriter, req *http.Request) {
	userID, err := getUserIDFromCtx(req.Context())
	if err != nil {
		makeJSONResponse(rsp, http.StatusInternalServerError, nil, err)
		return
	}

	createdDate := req.FormValue(createdDateField)
	if createdDate == "" {
		makeJSONResponse(rsp, http.StatusBadRequest, nil, errors.New("some errors here"))
		return
	}

	tasks, err := t.taskUseCase.ListTasks(req.Context(), userID, createdDate)
	if err != nil {
		makeJSONResponse(rsp, http.StatusInternalServerError, nil, err)
		return
	}

	makeJSONResponse(rsp, http.StatusOK, tasks, nil)
}

func (t *TaskTransport) AddTask(rsp http.ResponseWriter, req *http.Request) {
	t := &storages.Task{}
	if err := json.NewDecoder(req.Body).Decode(t); err != nil {
		makeJSONResponse(rsp, http.StatusInternalServerError, nil, err)
		return
	}
	defer req.Body.Close()

	userID, err := getUserIDFromCtx(req.Context())
	if err != nil {
		makeJSONResponse(rsp, http.StatusInternalServerError, nil, err)
		return
	}
}
