package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Companies struct {
	ID              primitive.ObjectID `bson:"_id"`
	CompanyName     string             `bson:"company_name"`
	TelephoneNumber string             `bson:"telephone_number"`
	Address         string             `bson:"address"`
	IsActive        bool               `bson:"is_active"`
	CreatedAt       time.Time          `bson:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at"`
}

//go:generate mockery --name ICompanies
type ICompanies interface {
	GetTableName() string
}

func NewCompanies() IUserAdmin {
	return &Companies{}
}

func (u *Companies) GetTableName() string {
	return "companies"
}
