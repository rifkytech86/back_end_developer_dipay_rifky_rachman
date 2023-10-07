package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"testing"
	"time"
)

func TestCompanies_GetTableName(t *testing.T) {
	type fields struct {
		ID              primitive.ObjectID
		CompanyName     string
		TelephoneNumber string
		Address         string
		IsActive        bool
		CreatedAt       time.Time
		UpdatedAt       time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get table",
			want: "companies",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Companies{
				ID:              tt.fields.ID,
				CompanyName:     tt.fields.CompanyName,
				TelephoneNumber: tt.fields.TelephoneNumber,
				Address:         tt.fields.Address,
				IsActive:        tt.fields.IsActive,
				CreatedAt:       tt.fields.CreatedAt,
				UpdatedAt:       tt.fields.UpdatedAt,
			}
			if got := u.GetTableName(); got != tt.want {
				t.Errorf("GetTableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCompanies(t *testing.T) {
	tests := []struct {
		name string
		want ICompanies
	}{
		{
			name: "initial companies",
			want: &Companies{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCompanies(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCompanies() = %v, want %v", got, tt.want)
			}
		})
	}
}
