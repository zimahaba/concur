package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var Apis ApiConfig

type ApiClient interface {
	SetRemoteCurrencies(conversion *Conversion, currencies string) error
}

func GetApiClient() (ApiClient, error) {
	if Apis.Active == "currencyapi" {
		return CurrencyAPI{}, nil
	} else if Apis.Active == "exchangerateapi" {
		return ExchangeRateAPI{}, nil
	}
	return nil, fmt.Errorf("unable to get api client.")
}

type CurrencyAPI struct{}

func (c CurrencyAPI) SetRemoteCurrencies(conversion *Conversion, currencies string) error {
	url := fmt.Sprintf(Apis.Available[Apis.Active].Url, conversion.BaseCurrency, currencies)
	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", Apis.Available[Apis.Active].Apikey)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("could not connect to currency api.")
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&conversion)
	if err != nil {
		return fmt.Errorf("could not read currency response body.")
	}

	return nil
}

type ExchangeRateAPI struct{}

func (c ExchangeRateAPI) SetRemoteCurrencies(conversion *Conversion, currencies string) error {

	if strings.Contains(currencies, conversion.BaseCurrency) == false {
		currencies = currencies + "," + conversion.BaseCurrency
	}

	url := fmt.Sprintf(Apis.Available[Apis.Active].Url, Apis.Available[Apis.Active].Apikey, currencies)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("could not connect to exchange rate api.")
	}
	defer resp.Body.Close()

	exchangeRate := ExchangeRate{}
	err = json.NewDecoder(resp.Body).Decode(&exchangeRate)
	if err != nil {
		return fmt.Errorf("could not read currency response body.")
	}

	baseCurrencyRate := exchangeRate.Rates[conversion.BaseCurrency]
	for k, v := range exchangeRate.Rates {
		if k != conversion.BaseCurrency {
			conversion.Data[k] = Currency{Code: k, Value: baseCurrencyRate / v}
		}
	}

	return nil
}
