package postgres

import (
	"context"
	"fmt"

	"github.com/eduardor2m/task-manager/src/infra/postgres/bridge"

	"github.com/eduardor2m/task-manager/src/core/domain/user"
	"github.com/eduardor2m/task-manager/src/core/interfaces/repository"
	"github.com/google/uuid"
)

var _ repository.UserLoader = &UserSQLiteRepository{}

type UserSQLiteRepository struct {
	connectorManager
}

func (instance UserSQLiteRepository) SignUp(userInstance user.User) (*uuid.UUID, error) {
	conn, err := instance.connectorManager.getConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	ctx := context.Background()

	queries := bridge.New(conn)

	userFormated := bridge.SignupParams{
		ID:        userInstance.ID(),
		Username:  userInstance.Username(),
		Email:     userInstance.Email(),
		Password:  userInstance.Password(),
		CreatedAt: *userInstance.CreatedAt(),
		UpdatedAt: *userInstance.UpdatedAt(),
	}

	userDB, err := queries.Signup(ctx, userFormated)

	fmt.Println(err)

	if err != nil {
		return nil, err
	}

	idLastInsert := userDB.ID

	return &idLastInsert, nil

}

func (instance UserSQLiteRepository) SignIn(email string, password string) (*string, error) {
	conn, err := instance.connectorManager.getConnection()
	if err != nil {
		return nil, err
	}
	defer instance.connectorManager.closeConnection(conn)

	ctx := context.Background()

	queries := bridge.New(conn)

	loginParams := bridge.SigninParams{
		Email:    email,
		Password: password,
	}

	userDB, err := queries.Signin(ctx, loginParams)

	if err != nil {
		return nil, err
	}

	idDB := userDB.ID

	idDBString := idDB.String()

	return &idDBString, nil
}

func NewUserSQLiteRepository(connectorManager connectorManager) UserSQLiteRepository {
	return UserSQLiteRepository{connectorManager: connectorManager}
}
