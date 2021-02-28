package transports

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/haunt98/togo/internal/services/usecases"
	"github.com/haunt98/togo/internal/token"
)

const (
	userIDField   = "user_id"
	passwordField = "password"

	authorizationHeader = "Authorization"
)

var (
	incorrectLoginError          = errors.New("incorrect user_id/pwd")
	incorerectAuthorizationError = errors.New("incorerect authorization")
)

type UserTransport struct {
	userUseCase    usecases.UserUseCase
	tokenGenerator token.Generator
}

// Validate userID/password and create token
func (t *UserTransport) Login(rsp http.ResponseWriter, req *http.Request) {
	userID := req.FormValue(userIDField)
	password := req.FormValue(passwordField)
	if userID == "" || password == "" {
		makeJSONResponse(rsp, http.StatusUnauthorized, nil, incorrectLoginError)
		return
	}

	isValid := t.userUseCase.Validate(req.Context(), userID, password)
	if !isValid {
		makeJSONResponse(rsp, http.StatusUnauthorized, nil, incorrectLoginError)
		return
	}

	token, err := t.tokenGenerator.CreateToken(userID)
	if err != nil {
		log.Printf("failed to create token: %s", err)
		makeJSONResponse(rsp, http.StatusInternalServerError, nil, incorrectLoginError)
	}

	makeJSONResponse(rsp, http.StatusOK, token, nil)
}

// Validate token and save userID if valid
func (t *UserTransport) ValidateToken(rsp http.ResponseWriter, req *http.Request) (*http.Request, bool) {
	token := req.Header.Get(authorizationHeader)
	if token == "" {
		makeJSONResponse(rsp, http.StatusUnauthorized, nil, incorerectAuthorizationError)
		return nil, false
	}

	userID, err := t.tokenGenerator.ValidateToken(token)
	if err != nil {
		log.Printf("failed to validate token: %s", err)
		makeJSONResponse(rsp, http.StatusUnauthorized, nil, incorerectAuthorizationError)
		return nil, false
	}

	// Inject userID to ctx
	req = req.WithContext(context.WithValue(req.Context(), userIDField, userID))
	return req, true
}
