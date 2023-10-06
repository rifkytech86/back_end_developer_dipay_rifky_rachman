package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Companies struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	CompanyName     string             `bson:"company_name,omitempty"`
	TelephoneNumber string             `bson:"telephone_number,omitempty"`
	Address         string             `bson:"address,omitempty"`
	IsActive        bool               `bson:"is_active"`
	CreatedAt       time.Time          `bson:"created_at,omitempty"`
	UpdatedAt       time.Time          `bson:"updated_at,omitempty"`
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
