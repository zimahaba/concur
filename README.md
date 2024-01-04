# Concur
Concur is a command line currency converter.

## Build / Install
Binary currently not available. To build the application run `make` in the root directory. A binary will be created in the /bin directory.
### Prerequisites
- [Go](https://go.dev/): version 1.21.5.

## Run
`./concur [command] [ARGS] [FLAGS]`

- Print version:
`./concur version`
- Print settings:
`./concur config show`
- Configure api and apikey:
`./concur config -a api -k apikey`
- Print available currencies:
`./concur currencies`
- Convert currency:
`./concur convert [ARGS] [FLAGS]`

## Configuration
Concur uses two different currency apis, so it is required to create an account and an apikey in the select api.
- Currency API: https://currencyapi.com/
- Exchange Rate API: https://exchangeratesapi.io/

To set the api in Concur run the `config` command passing the api name and apikey:

- `./concur config -a currencyapi -k [apikey]`
- `./concur config -a exchangerateapi -k [apikey]`

The first time Concur runs a `config.yaml` file will be created in the applicaton directory. You can also configure the settings by manually updating the file.

## Convert
After configuring the api, you can convert currencies by running the `convert` command passing the value you want to convert, the base currency and the currencies to be converted.

The first argument of the convert can be the value you want to convert. If no value is informed, Concur uses 1.

The first argument after the value must be the base currency you want to use in the conversion.

The next arguments are the list of currencies you want the value to be converted to.

- Example 1: `./concur convert 2 usd brl` will convert 2 USD to BRL.
- Example 2: `./concur convert eur usd gbp brl` will convert 1 EUR to USD, GBP and BRL.

At least 2 arguments is required to run the `convert` command.

### Cache
Concur uses sqlite database to cache the conversion rates. The cache will be setup the first time Concur runs.

To use the cache pass the `-c` flag to the `convert` command.
- `./concur convert 5 usd brl ars gbp -c`

Concur will use the cache only for already used currencies, so even if the cache flag is passed the currency api configured will be used in case a currency hasn't been used before.
