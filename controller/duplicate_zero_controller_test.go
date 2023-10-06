package controller

import (
	"github.com/dipay/internal/validations"
	"github.com/dipay/internal/validations/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func Test_duplicateZeroController_DuplicateZero(t *testing.T) {
	type fields struct {
		ContextTimeOut int
		Validator      validations.IValidator
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		mockValidator []validations.ValidationError
		wantErr       bool
	}{
		{
			name: "error bind duplicate zero",
			args: args{
				c: failedInitialEcho(),
			},
			wantErr: false,
		},
		{
			name: "error validator",
			args: args{
				c: successInitialEcho(),
			},
			mockValidator: []validations.ValidationError{
				{
					Field: "n",
					Error: "error n is empty",
				},
			},
			wantErr: false,
		},
		{
			name: "error validator",
			args: args{
				c: successInitialEcho(),
			},
			mockValidator: nil,
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &duplicateZeroController{}
			mockValidator := new(mocks.IValidator)
			mockValidator.On("Struct", mock.Anything).Return(tt.mockValidator)
			d.Validator = mockValidator

			if err := d.DuplicateZero(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("DuplicateZero() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewDuplicateZeroController(t *testing.T) {
	type args struct {
		contextTimeOut int
		validator      validations.IValidator
	}
	tests := []struct {
		name string
		args args
		want IDuplicateZeroController
	}{
		{
			name: "initial duplicate zero controller",
			args: args{},
			want: &duplicateZeroController{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDuplicateZeroController(tt.args.contextTimeOut, tt.args.validator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDuplicateZeroController() = %v, want %v", got, tt.want)
			}
		})
	}
}
