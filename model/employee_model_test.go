package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"testing"
	"time"
)

func TestEmployees_GetTableName(t *testing.T) {
	type fields struct {
		ID          primitive.ObjectID
		Name        string
		Email       string
		PhoneNumber string
		JobTitle    string
		CompanyID   string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get table",
			want: "employees",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Employees{
				ID:          tt.fields.ID,
				Name:        tt.fields.Name,
				Email:       tt.fields.Email,
				PhoneNumber: tt.fields.PhoneNumber,
				JobTitle:    tt.fields.JobTitle,
				CompanyID:   tt.fields.CompanyID,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
			}
			if got := u.GetTableName(); got != tt.want {
				t.Errorf("GetTableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEmployees(t *testing.T) {
	tests := []struct {
		name string
		want IEmployees
	}{
		{
			name: "initial modal",
			want: &Employees{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmployees(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmployees() = %v, want %v", got, tt.want)
			}
		})
	}
}
