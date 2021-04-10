package common

import (
	"os"
)

const (
	baseBinanceDevelopmentURL = "api.binance.com"
)

//BuildAbsoluteForApp builds an absolute path from protocol to params
func BuildAbsoluteForBinance(path string, urlParams Params) (string, error) {
	return BuildAbsolute(path, getBaseBinanceURLForENV(), urlParams)
}

func getBaseBinanceURLForENV() string {
	baseURL := protocol

	env := os.Getenv("API_ENV")

	switch env {
	case EnvDev:
		baseURL = baseURL + baseBinanceDevelopmentURL
	default:
		baseURL = baseURL + baseBinanceDevelopmentURL
	}

	return baseURL
}
