package handlers

import (
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
		context.JSON(400, err)
	}

	currentTime := time.Now()

	userInstance, err := user.NewBuilder().WithID(uuid.New()).WithEmail(userRequest.Email).WithPassword(userRequest.Password).WithCreatedAt(&currentTime).WithUpdatedAt(&currentTime).Build()

	if err != nil {
		context.JSON(400, err)
	}

	userID, _ := instance.service.SignUp(*userInstance)

	idJson := response.ID{
		ID: *userID,
	}

	context.JSON(200, idJson)

	return nil
}

func (instance UserHandlers) SignIn(context echo.Context) error {
	var userRequest request.UserDTO

	err := context.Bind(&userRequest)

	if err != nil {
		context.JSON(400, err)
	}

	token, _ := instance.service.SignIn(userRequest.Email, userRequest.Password)

	tokenJson := response.Token{
		Token: *token,
	}

	context.JSON(200, tokenJson)

	return nil
}

func NewUserHandlers(service primary.UserManager) *UserHandlers {
	return &UserHandlers{
		service: service,
	}
}
