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
	"golang.org/x/crypto/bcrypt"
	"time"
)

type useCaseUserAdmin struct {
	UserAdminRepository repositories.IUserAdminRepository
	JWT                 jwt.IJWTRSAToken
	Secret              string
}

//go:generate mockery --name IUseCaseUserAdmin
type IUseCaseUserAdmin interface {
	Login(context.Context, *api.LoginJSONBody) (string, error)
	Register(context.Context, *api.LoginJSONBody) (string, error)
}

func NewUseCaseUserAdmin(userAdminRepository repositories.IUserAdminRepository, jwt jwt.IJWTRSAToken, secret string) IUseCaseUserAdmin {
	return &useCaseUserAdmin{
		UserAdminRepository: userAdminRepository,
		JWT:                 jwt,
		Secret:              secret,
	}
}

func (u *useCaseUserAdmin) Login(ctx context.Context, req *api.LoginJSONBody) (string, error) {
	var userAdmin model.UserAdmin
	err := u.UserAdminRepository.Fetch(ctx, bson.M{"username": req.Username}, &userAdmin)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", mongo.ErrNoDocuments
		}
		return "", errors.New(internal.ErrorInternalServer.String())
	}

	err = bcrypt.CompareHashAndPassword([]byte(userAdmin.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New(internal.ErrorInternalGenerateToken.String())
	}

	token, err := u.JWT.GenerateToken(userAdmin.ID.String(), userAdmin.UserName)
	if err != nil {
		return "", errors.New(internal.ErrorInternalGenerateToken.String())
	}

	return token, nil
}

func (u *useCaseUserAdmin) Register(ctx context.Context, req *api.LoginJSONBody) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New(internal.ErrorInternalServer.String())
	}
	userAdmin := model.UserAdmin{
		UserName:  req.Username,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	lastInsertedID, err := u.UserAdminRepository.Create(ctx, userAdmin)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", mongo.ErrNoDocuments
		}
		return "", errors.New(internal.ErrorInternalServer.String())
	}

	token, err := u.JWT.GenerateToken(lastInsertedID.String(), req.Password)
	if err != nil {
		return "", errors.New(internal.ErrorInternalGenerateToken.String())
	}

	return token, nil
}
