package controller

import (
	"errors"
	"github.com/dipay/internal"
	"github.com/dipay/repositories"
	"github.com/dipay/usecase"
	"github.com/dipay/usecase/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func Test_countriesController_GetDataCountries(t *testing.T) {
	type fields struct {
		CountriesUseCase usecase.ICountriesUseCase
		ContextTimeOut   int
	}
	type mockGetCountries struct {
		data []repositories.CountriesPayload
		err  error
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantErr          bool
		mockGetCountries mockGetCountries
	}{
		{
			name: "error get countries",
			args: args{
				c: successInitialEcho(),
			},
			mockGetCountries: mockGetCountries{
				err: errors.New(internal.ErrorInternalServer.String()),
			},
			wantErr: true,
		},
		{
			name: "Happy flow",
			args: args{
				c: successInitialEcho(),
			},
			mockGetCountries: mockGetCountries{
				data: []repositories.CountriesPayload{
					{
						Name: "tester",
					},
				},
				err: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			co := &countriesController{}
			mockCountriesUseCase := new(mocks.ICountriesUseCase)
			mockCountriesUseCase.On("GetCountries", mock.Anything).Return(tt.mockGetCountries.data, tt.mockGetCountries.err)
			co.CountriesUseCase = mockCountriesUseCase

			if err := co.GetDataCountries(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetDataCountries() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewCountriesController(t *testing.T) {
	type args struct {
		countriesUseCase usecase.ICountriesUseCase
		contextTimeOut   int
	}
	tests := []struct {
		name string
		args args
		want ICountriesController
	}{
		{
			name: "initial countries controller",
			args: args{},
			want: &countriesController{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCountriesController(tt.args.countriesUseCase, tt.args.contextTimeOut); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCountriesController() = %v, want %v", got, tt.want)
			}
		})
	}
}
