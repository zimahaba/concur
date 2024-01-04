package pkg

import (
	"fmt"
	"strings"
)

type Conversion struct {
	BaseCurrency string
	Data         map[string]Currency `json:"data"`
}

type Currency struct {
	Code  string  `json:"code"`
	Value float32 `json:"value"`
}

type ExchangeRate struct {
	Rates map[string]float32 `json:"rates"`
}

type ApiConfig struct {
	Active    string         `yaml:"active"`
	Available map[string]Api `yaml:"available"`
}

func (apis ApiConfig) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("Active API: %s\n", apis.Active))
	sb.WriteString(fmt.Sprintln("Available APIS:"))
	for k, v := range apis.Available {
		sb.WriteString(fmt.Sprintf("- %s (%s)\n", v.Name, k))
	}
	return sb.String()
}

type Api struct {
	Name   string
	Url    string
	Apikey string
}
