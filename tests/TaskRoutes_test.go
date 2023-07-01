package tests

import (
	"io"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestWhenListAllTasks(t *testing.T) {
	requestUrl := "http://localhost:9090/api/task"
	clientRequest, err := http.NewRequest(http.MethodGet, requestUrl, nil)

	if err != nil {
		t.Error(err)
	}

	token := ""

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
