package usecase

import (
	"context"
	"errors"
	"github.com/dipay/api"
	"github.com/dipay/internal"
	"github.com/dipay/internal/jwt"
	"github.com/dipay/model"
	"github.com/dipay/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type useCaseUserAdmin struct {
	UserAdminRepository repositories.IUserAdminRepository
	JWT                 jwt.IJWTRSAToken
}

type IUseCaseUserAdmin interface {
	Login(context.Context, *api.LoginJSONBody) (string, error)
}

func NewUseCaseUserAdmin(userAdminRepository repositories.IUserAdminRepository, jwt jwt.IJWTRSAToken) IUseCaseUserAdmin {
	return &useCaseUserAdmin{
		UserAdminRepository: userAdminRepository,
		JWT:                 jwt,
	}
}

func (u *useCaseUserAdmin) Login(ctx context.Context, req *api.LoginJSONBody) (string, error) {
	var userAdmin model.UserAdmin
	err := u.UserAdminRepository.Fetch(ctx, bson.M{"username": req.Username, "password": req.Password}, &userAdmin)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", errors.New(internal.ErrorInvalidLogin.String())
		}
		return "", errors.New(internal.ErrorInternalServer.String())
	}

	token, err := u.JWT.GenerateToken(userAdmin.ID.String(), userAdmin.UserName)
	if err != nil {
		return "", errors.New(internal.ErrorInternalGenerateToken.String())
	}

	return token, nil
}
