package usecase

import (
	"context"
	"errors"
	"github.com/dipay/internal"
	"github.com/dipay/repositories"
	"github.com/dipay/repositories/mocks"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func Test_countriesUseCase_GetCountries(t *testing.T) {
	type fields struct {
		CountriesRepository repositories.ICountriesRepository
	}
	type args struct {
		ctx context.Context
	}
	type mockGetCountry struct {
		data []repositories.CountriesPayload
		err  error
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		want           []repositories.CountriesPayload
		wantErr        bool
		mockGetCountry mockGetCountry
	}{
		{
			name: "error get country",
			args: args{},
			mockGetCountry: mockGetCountry{
				data: []repositories.CountriesPayload{},
				err:  errors.New(internal.ErrorInternalServer.String()),
			},
			wantErr: true,
		},
		{
			name: "Happy flow get country",
			args: args{},
			mockGetCountry: mockGetCountry{
				data: []repositories.CountriesPayload{},
				err:  nil,
			},
			wantErr: false,
			want:    []repositories.CountriesPayload{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			co := &countriesUseCase{}
			mockRepo := new(mocks.ICountriesRepository)
			mockRepo.On("GetCountries", mock.Anything).Return(tt.mockGetCountry.data, tt.mockGetCountry.err)
			co.CountriesRepository = mockRepo

			got, err := co.GetCountries(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCountries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCountries() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCountriesUseCase(t *testing.T) {
	type args struct {
		countriesRepository repositories.ICountriesRepository
	}
	tests := []struct {
		name string
		args args
		want ICountriesUseCase
	}{
		{
			name: "initial new countries repository",
			args: args{},
			want: &countriesUseCase{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCountriesUseCase(tt.args.countriesRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCountriesUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
