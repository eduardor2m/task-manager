package tests

import (
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func deleteUserByEmail(email string) {
	requestUrl := "http://localhost:9090/api/user/delete"
	requestBodyJson := `{
		"email": "` + email + `"
	}`

	requestBodyIo := io.Reader(strings.NewReader(requestBodyJson))

	clientRequest, err := http.NewRequest(http.MethodDelete, requestUrl, requestBodyIo)

	if err != nil {
		panic(err)
	}

	token := getToken()

	clientRequest.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

	serverResponse, err := http.DefaultClient.Do(clientRequest)

	if err != nil {
		panic(err)
	}

	serverData, err := io.ReadAll(serverResponse.Body)

	if err != nil {
		panic(err)
	}

	os.Stdout.Write(serverData)
}

func TestWhenSignupUser(t *testing.T) {
	requestUrl := "http://localhost:9090/api/user/signup"
	requestBodyJson := `{
		"username": "test",
		"email": "test@gmail.com",
		"password": "123456"
	}`

	defer deleteUserByEmail("test@gmail.com")

	requestBodyIo := io.Reader(strings.NewReader(requestBodyJson))

	clientRequest, err := http.NewRequest(http.MethodPost, requestUrl, requestBodyIo)

	if err != nil {
		t.Error(err)
	}
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

func TestWhenLoginUser(t *testing.T) {
	requestUrl := "http://localhost:9090/api/user/signin"
	requestBodyJson := `{
		"email": "test@gmail.com",
		"password": "123456"
	}`

	defer deleteUserByEmail("test@gmail.com")

	requestBodyIo := io.Reader(strings.NewReader(requestBodyJson))

	clientRequest, err := http.NewRequest(http.MethodPost, requestUrl, requestBodyIo)

	if err != nil {
		t.Error(err)
	}

	clientRequest.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	serverResponse, err := http.DefaultClient.Do(clientRequest)

	if err != nil {
		t.Error(err)
	}

	serverData, err := io.ReadAll(serverResponse.Body)

	if err != nil {

		t.Error(err)
	}

	token := string(serverData)

	os.WriteFile("./data/token.json", []byte(token), 0666)

	t.Log(string(serverData))

	assert.Equal(t, http.StatusOK, serverResponse.StatusCode)

}
