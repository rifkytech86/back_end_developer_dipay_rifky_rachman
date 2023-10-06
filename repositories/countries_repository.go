package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dipay/pkg/httpClient"
)

type countriesRepository struct {
	clientHttp      httpClient.IClientHttp
	exAPIGetCountry string
}

//go:generate mockery --name ICountriesRepository
type ICountriesRepository interface {
	GetCountries(ctx context.Context) ([]CountriesPayload, error)
}

func NewCountriesRepository(client httpClient.IClientHttp, exSpiGetCountry string) ICountriesRepository {
	return &countriesRepository{
		clientHttp:      client,
		exAPIGetCountry: exSpiGetCountry,
	}
}

func (co *countriesRepository) GetCountries(ctx context.Context) ([]CountriesPayload, error) {
	responseBytes, err := co.clientHttp.Get(co.exAPIGetCountry)
	//responseBytes, err := co.clientHttp.Get("https://gist.githubusercontent.com/herysepty/ba286b815417363bfbcc472a5197edd0/raw/aed8ce8f5154208f9fe7f7b04195e05de5f81fda/coutries.json")
	if err != nil {
		return nil, err
	}
	var countriesPayload []CountriesPayload
	err = json.Unmarshal(responseBytes, &countriesPayload)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	return countriesPayload, nil
}
