package handlers

import (
	"net/http"
	"time"

	"github.com/eduardor2m/task-manager/src/api/handlers/dto/request"
	"github.com/eduardor2m/task-manager/src/api/handlers/dto/response"
	"github.com/eduardor2m/task-manager/src/core/domain/user"
	"github.com/eduardor2m/task-manager/src/core/interfaces/primary"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandlers struct {
	service primary.UserManager
}

func (instance UserHandlers) SignUp(context echo.Context) error {
	var userRequest request.UserDTO

	err := context.Bind(&userRequest)

	if err != nil {
		message := response.TaskMessage{
			Message: err.Error(),
		}
		return context.JSON(400, message)
	}

	currentTime := time.Now()

	userInstance, err := user.NewBuilder().WithID(uuid.New()).WithUsername(userRequest.Username).WithEmail(userRequest.Email).WithPassword(userRequest.Password).WithCreatedAt(&currentTime).WithUpdatedAt(&currentTime).Build()

	if err != nil {
		message := response.TaskMessage{
			Message: err.Error(),
		}
		return context.JSON(400, message)
	}

	userID, err := instance.service.SignUp(*userInstance)

	if err != nil {
		message := response.TaskMessage{
			Message: err.Error(),
		}
		return context.JSON(400, message)
	}

	idJson := response.UserID{
		ID: *userID,
	}

	context.JSON(http.StatusCreated, idJson)

	return nil
}

func (instance UserHandlers) SignIn(context echo.Context) error {
	var userRequest request.UserDTO

	err := context.Bind(&userRequest)

	if err != nil {
		message := response.TaskMessage{
			Message: err.Error(),
		}
		return context.JSON(400, message)
	}

	token, err := instance.service.SignIn(userRequest.Email, userRequest.Password)

	if err != nil {
		message := response.TaskMessage{
			Message: err.Error(),
		}

		return context.JSON(400, message)

	}

	tokenJson := response.UserToken{
		Token: *token,
	}

	context.JSON(200, tokenJson)

	return nil
}

func (instance UserHandlers) DeleteUserByEmail(context echo.Context) error {
	var userRequest request.UserDTO

	err := context.Bind(&userRequest)

	if err != nil {
		message := response.TaskMessage{
			Message: err.Error(),
		}

		return context.JSON(400, message)
	}

	err = instance.service.DeleteUserByEmail(userRequest.Email)

	if err != nil {
		message := response.TaskMessage{
			Message: err.Error(),
		}

		return context.JSON(400, message)
	}

	message := response.TaskMessage{
		Message: "User deleted successfully",
	}

	return context.JSON(200, message)
}

func NewUserHandlers(service primary.UserManager) *UserHandlers {
	return &UserHandlers{
		service: service,
	}
}
