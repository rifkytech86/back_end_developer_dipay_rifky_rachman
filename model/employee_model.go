package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Employees struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Email       string             `bson:"email,omitempty"`
	PhoneNumber string             `bson:"phone_number,omitempty"`
	JobTitle    string             `bson:"jobtitle,omitempty"`
	CompanyID   string             `bson:"company_id,omitempty"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

type IEmployees interface {
	GetTableName() string
}

func NewEmployees() IEmployees {
	return &Employees{}
}

func (u *Employees) GetTableName() string {
	return "employees"
}
