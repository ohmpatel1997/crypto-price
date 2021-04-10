package common

import (
	"log"
	"net/url"
	goFilePath "path"
)

const (
	EnvDev = "dev"
)

const (
	ExchangeBinance = "binance"
)

const protocol = "https://"
const uprotocol = "http://"

//Params keys are the param names and the values are the param values
type Params map[string]string

//BuildAbsolute takes a path, baseURL and urlParams, returns a url as a string or an error
func BuildAbsolute(path string, baseURL string, urlParams Params) (string, error) {
	u, err := url.Parse(baseURL)

	if err != nil {
		return "", err
	}

	u.Path = goFilePath.Join(u.Path, path)

	query := u.Query()

	for key, value := range urlParams {
		query.Set(key, value)
	}

	u.RawQuery = query.Encode()

	return u.String(), nil
}

func GetUrlForExchange(exchange string) func(string, Params) (string, error) {

	switch exchange {
	case ExchangeBinance:
		return BuildAbsoluteForBinance
	default:
		log.Panic("invalid exchange")
	}

	return nil
}
