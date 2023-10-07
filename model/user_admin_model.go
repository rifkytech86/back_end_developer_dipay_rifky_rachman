package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserAdmin struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserName  string             `bson:"username,omitempty"`
	Password  string             `bson:"password,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

//go:generate mockery --name IUserAdmin
type IUserAdmin interface {
	GetTableName() string
	IsValidPassword(userAdminPassword string, userReqPassword string) error
	EncryptedPassword(userAdminPassword string) (hashed string, err error)
}

func NewUserAdmin() IUserAdmin {
	return &UserAdmin{}
}

func (u *UserAdmin) GetTableName() string {
	return "admins"
}

func (u *UserAdmin) IsValidPassword(userAdminPassword string, userReqPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userAdminPassword), []byte(userReqPassword))
	if err != nil {
		return err
	}
	return nil
}

func (u *UserAdmin) EncryptedPassword(userAdminPassword string) (hashed string, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userAdminPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
