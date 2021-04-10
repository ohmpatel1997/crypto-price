package service

import (
	"crypto/internal/common"
	"log"
)

type ExchangeService interface {
	GetPriceDetails() error
}

func GetExchangeServiceFor(exc string) ExchangeService {
	switch exc {
	case common.ExchangeBinance:
		return newBinanceService()
	default:
		log.Panic("invalid exchange")
	}

	return nil
}
