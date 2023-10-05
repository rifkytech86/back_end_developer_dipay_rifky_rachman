package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserAdmin struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserName  string             `bson:"username,omitempty"`
	Password  string             `bson:"password,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type IUserAdmin interface {
	GetTableName() string
}

func NewUserAdmin() IUserAdmin {
	return &UserAdmin{}
}

func (u *UserAdmin) GetTableName() string {
	return "admins"
}
