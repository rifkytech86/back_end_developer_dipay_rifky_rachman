package repositories

import (
	"context"
	"errors"
	"github.com/dipay/internal"
	"github.com/dipay/pkg/httpClient"
	"github.com/dipay/pkg/httpClient/mocks"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func Test_countriesRepository_GetCountries(t *testing.T) {
	type fields struct {
		clientHttp      httpClient.IClientHttp
		exAPIGetCountry string
	}
	type mockResponse struct {
		res []byte
		err error
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		want         []CountriesPayload
		wantErr      bool
		mockResponse mockResponse
	}{
		{
			name: "error get countries",
			args: args{
				ctx: context.TODO(),
			},
			mockResponse: mockResponse{
				err: errors.New(internal.ErrorInternalServer.String()),
				res: []byte{},
			},
			wantErr: true,
		},
		{
			name: "error format response from api",
			args: args{
				ctx: context.TODO(),
			},
			mockResponse: mockResponse{
				err: nil,
				res: []byte(`{"id": 123, "name: "John"}`),
			},
			wantErr: true,
		},
		{
			name: "happy flow call api",
			args: args{
				ctx: context.TODO(),
			},
			mockResponse: mockResponse{
				err: nil,
				res: []byte(`[{"name":"Afghanistan","topLevelDomain":[".af"],"alpha2Code":"AF","alpha3Code":"AFG","callingCodes":["93"],"capital":"Kabul","altSpellings":["AF","Afġānistān"],"region":"Asia","subregion":"Southern Asia","population":27657145,"latlng":[33,65],"demonym":"Afghan","area":652230,"gini":27.8,"timezones":["UTC+04:30"],"borders":["IRN","PAK","TKM","UZB","TJK","CHN"],"nativeName":"افغانستان","numericCode":"004","currencies":[{"code":"AFN","name":"Afghan afghani","symbol":"؋"}],"languages":[{"iso639_1":"ps","iso639_2":"pus","name":"Pashto","nativeName":"پښتو"},{"iso639_1":"uz","iso639_2":"uzb","name":"Uzbek","nativeName":"Oʻzbek"},{"iso639_1":"tk","iso639_2":"tuk","name":"Turkmen","nativeName":"Türkmen"}],"translations":{"de":"Afghanistan","es":"Afganistán","fr":"Afghanistan","ja":"アフガニスタン","it":"Afghanistan","br":"Afeganistão","pt":"Afeganistão","nl":"Afghanistan","hr":"Afganistan","fa":"افغانستان"},"flag":"https://restcountries.eu/data/afg.svg","regionalBlocs":[{"acronym":"SAARC","name":"South Asian Association for Regional Cooperation","otherAcronyms":[],"otherNames":[]}],"cioc":"AFG"}]`),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			co := &countriesRepository{}
			co.exAPIGetCountry = "/tester"
			mockClient := new(mocks.IClientHttp)
			mockClient.On("Get", mock.Anything).Return(tt.mockResponse.res, tt.mockResponse.err)
			co.clientHttp = mockClient
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

func TestNewCountriesRepository(t *testing.T) {
	type args struct {
		client          httpClient.IClientHttp
		exSpiGetCountry string
	}
	tests := []struct {
		name string
		args args
		want ICountriesRepository
	}{
		{
			name: "initial countries",
			args: args{},
			want: &countriesRepository{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCountriesRepository(tt.args.client, tt.args.exSpiGetCountry); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCountriesRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
