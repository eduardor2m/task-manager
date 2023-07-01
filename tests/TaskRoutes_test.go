package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// func createUser() {
// 	requestUrl := "http://localhost:9090/api/user/signup"
// 	requestBodyJson := `{
// 		"username": "test",
// 		"email": "test@gmail.com",
// 		"password": "123456"
// 	}`

// 	requestBodyIo := io.Reader(strings.NewReader(requestBodyJson))

// 	clientRequest, err := http.NewRequest(http.MethodPost, requestUrl, requestBodyIo)

// 	if err != nil {
// 		panic(err)
// 	}
// 	clientRequest.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 	serverResponse, err := http.DefaultClient.Do(clientRequest)

// 	if err != nil {
// 		panic(err)
// 	}

// 	_, err = io.ReadAll(serverResponse.Body)

// 	if err != nil {
// 		panic(err)
// 	}

// }

// func getToken() string {
// 	createUser()
// 	defer deleteUserByEmail("test@gmail.com")
// 	requestUrl := "http://localhost:9090/api/user/signin"
// 	requestBodyJson := `{
// 		"email": "test@gmail.com",
// 		"password": "123456"
// 	}`

// 	requestBodyIo := io.Reader(strings.NewReader(requestBodyJson))

// 	clientRequest, err := http.NewRequest(http.MethodPost, requestUrl, requestBodyIo)

// 	if err != nil {
// 		panic(err)
// 	}

// 	clientRequest.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 	serverResponse, err := http.DefaultClient.Do(clientRequest)

// 	if err != nil {
// 		panic(err)
// 	}

// 	serverData, err := io.ReadAll(serverResponse.Body)

// 	if err != nil {
// 		panic(err)
// 	}

// 	token := struct {
// 		Token string `json:"token"`
// 	}{}

// 	err = json.Unmarshal(serverData, &token)

// 	if err != nil {
// 		panic(err)
// 	}

// 	return token.Token
// }

func getToken() string {
	file, err := os.Open("./data/token.json")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	fileData, err := io.ReadAll(file)

	if err != nil {
		panic(err)
	}

	token := struct {
		Token string `json:"token"`
	}{}

	err = json.Unmarshal(fileData, &token)

	if err != nil {
		panic(err)
	}

	return token.Token
}

func getTaskId() string {
	file, err := os.Open("./data/taskId.json")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	fileData, err := io.ReadAll(file)

	if err != nil {
		panic(err)
	}

	taskId := struct {
		Id string `json:"id"`
	}{}

	err = json.Unmarshal(fileData, &taskId)

	if err != nil {
		panic(err)
	}

	return taskId.Id
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

	os.WriteFile("./data/taskId.json", serverData, 0666)

	t.Log(string(serverData))

	assert.Equal(t, http.StatusCreated, serverResponse.StatusCode)

}

func TestWhenUpdateStatusTask(t *testing.T) {
	id := getTaskId()
	requestUrl := "http://localhost:9090/api/task/" + id + "/status"

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
