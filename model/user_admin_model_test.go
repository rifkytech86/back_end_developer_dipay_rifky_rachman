package model

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"testing"
	"time"
)

func TestUserAdmin_GetTableName(t *testing.T) {
	type fields struct {
		ID        primitive.ObjectID
		UserName  string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get table",
			want: "admins",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserAdmin{
				ID:        tt.fields.ID,
				UserName:  tt.fields.UserName,
				Password:  tt.fields.Password,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if got := u.GetTableName(); got != tt.want {
				t.Errorf("GetTableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserAdmin(t *testing.T) {
	tests := []struct {
		name string
		want IUserAdmin
	}{
		{
			name: "initial model",
			want: &UserAdmin{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserAdmin(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserAdmin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserAdmin_IsValidPassword(t *testing.T) {
	type fields struct {
		ID        primitive.ObjectID
		UserName  string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	type args struct {
		userAdminPassword string
		userReqPassword   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "is valid password",
			args: args{
				userAdminPassword: "$2a$10$qoahj5NG6GzK.acOk0jn6ugflySm4dW.vVsjLMYg1pZJnAx74KsiS",
				userReqPassword:   "ZXCasdqwe123!",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserAdmin{
				ID:        tt.fields.ID,
				UserName:  tt.fields.UserName,
				Password:  tt.fields.Password,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if err := u.IsValidPassword(tt.args.userAdminPassword, tt.args.userReqPassword); (err != nil) != tt.wantErr {
				t.Errorf("IsValidPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserAdmin_EncryptedPassword(t *testing.T) {
	type fields struct {
		ID        primitive.ObjectID
		UserName  string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	type args struct {
		userAdminPassword string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantHashed string
		wantErr    bool
	}{
		{
			name:    "encrypted success",
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserAdmin{
				ID:        tt.fields.ID,
				UserName:  tt.fields.UserName,
				Password:  tt.fields.Password,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			_, err := u.EncryptedPassword(tt.args.userAdminPassword)
			assert.NoError(t, err)
		})
	}
}
