package tests

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func getToken() string {
	return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2ODgyMzgzMTgsInVzZXJfaWQiOiJmNThmNmI2MC1jZWYyLTQxZGQtYWY4Yi1jYjkyMWYzOGE0MmEifQ.I3ubIZnXXIj2tezFrCHnINzHZHtKO4lA4ZBkqqh8lHI"
}

func TestWhenListAllTasks(t *testing.T) {
	requestUrl := "http://localhost:9090/api/task"
	clientRequest, err := http.NewRequest(http.MethodGet, requestUrl, nil)

	if err != nil {
		t.Error(err)
	}

	token := getToken()

	clientRequest.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

	serverResponse, err := http.DefaultClient.Do(clientRequest)

	if err != nil {
		t.Error(err)
	}

	serverData, err := io.ReadAll(serverResponse.Body)

	if err != nil {
		t.Error(err)
	}

	t.Log(string(serverData))

	assert.Equal(t, http.StatusOK, serverResponse.StatusCode)

}

func TestWhenCreateTask(t *testing.T) {
	requestUrl := "http://localhost:9090/api/task"
	requestBodyJson := `{
		"title": "Teste",
		"description": "Teste",
		"status": false
	}`

	requestBodyIo := io.Reader(strings.NewReader(requestBodyJson))

	clientRequest, err := http.NewRequest(http.MethodPost, requestUrl, requestBodyIo)

	if err != nil {
		t.Error(err)
	}

	token := getToken()

	clientRequest.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	clientRequest.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	serverResponse, err := http.DefaultClient.Do(clientRequest)

	if err != nil {
		t.Error(err)
	}

	serverData, err := io.ReadAll(serverResponse.Body)

	if err != nil {
		t.Error(err)
	}

	t.Log(string(serverData))

	assert.Equal(t, http.StatusCreated, serverResponse.StatusCode)

}

func TestWhenUpdateStatusTask(t *testing.T) {
	requestUrl := "http://localhost:9090/api/task/d5c7913e-896d-41ce-a295-2dd45b5f827b/status"

	clientRequest, err := http.NewRequest(http.MethodPut, requestUrl, nil)

	if err != nil {
		t.Error(err)
	}

	token := getToken()

	clientRequest.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	clientRequest.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	serverResponse, err := http.DefaultClient.Do(clientRequest)

	if err != nil {
		t.Error(err)
	}

	serverData, err := io.ReadAll(serverResponse.Body)

	if err != nil {
		t.Error(err)
	}

	t.Log(string(serverData))

	assert.Equal(t, http.StatusOK, serverResponse.StatusCode)
}

func TestWhenDeleteTask(t *testing.T) {
	requestUrl := "http://localhost:9090/api/task/5ea1da69-9d85-4c45-8c44-e631a7ec2192"

	clientRequest, err := http.NewRequest(http.MethodDelete, requestUrl, nil)

	if err != nil {
		t.Error(err)
	}

	token := getToken()

	clientRequest.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

	serverResponse, err := http.DefaultClient.Do(clientRequest)

	if err != nil {
		t.Error(err)
	}

	serverData, err := io.ReadAll(serverResponse.Body)

	if err != nil {
		t.Error(err)
	}

	t.Log(string(serverData))

	assert.Equal(t, http.StatusOK, serverResponse.StatusCode)

}

func TestWhenDeleteAllTasks(t *testing.T) {
	requestUrl := "http://localhost:9090/api/task"

	clientRequest, err := http.NewRequest(http.MethodDelete, requestUrl, nil)

	if err != nil {
		t.Error(err)
	}

	token := getToken()

	clientRequest.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

	serverResponse, err := http.DefaultClient.Do(clientRequest)

	if err != nil {
		t.Error(err)
	}

	serverData, err := io.ReadAll(serverResponse.Body)

	if err != nil {
		t.Error(err)
	}

	t.Log(string(serverData))

	assert.Equal(t, http.StatusOK, serverResponse.StatusCode)
}

// func TestWhenUpdateTask(t *testing.T) {
// 	requestUrl := "http://localhost:9090/api/task/5ea1da69-9d85-4c45-8c44-e631a7ec2192"
// 	requestBodyJson := `{
// 		"title": "Teste",
// 		"description": "Teste",
// 		"status": true
// 	}`

// 	requestBodyIo := io.Reader(strings.NewReader(requestBodyJson))

// 	clientRequest, err := http.NewRequest(http.MethodPut, requestUrl, requestBodyIo)

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	token := getToken()

// 	clientRequest.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
// 	clientRequest.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 	serverResponse, err := http.DefaultClient.Do(clientRequest)

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	serverData, err := io.ReadAll(serverResponse.Body)

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	t.Log(string(serverData))

// 	assert.Equal(t, http.StatusOK, serverResponse.StatusCode)

// }
