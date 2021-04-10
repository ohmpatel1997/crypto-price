package service

import (
	"crypto/internal/common"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type serviceBinance struct {
}

type Symbol struct {
	Symbol     string `json:"symbol"`
	BaseAsset  string `json:"baseAsset"`
	QuoteAsset string `json:"quoteAsset"`
}

type AllResponse struct {
	Symbols []Symbol `json:"symbols"`
}

func newBinanceService() *serviceBinance {
	return &serviceBinance{}
}

func (s *serviceBinance) GetPriceDetails() error {
	f, err := os.Create(os.Getenv("exchange") + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	url, err := common.GetUrlForExchange(os.Getenv("exchange"))("/api/v3/exchangeInfo", common.Params{})
	if err != nil {
		return err
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	allDetails := new(AllResponse)

	err = json.Unmarshal(body, allDetails)
	if err != nil {
		return err
	}

	for _, tck := range allDetails.Symbols {
		_, err = f.WriteString(fmt.Sprintf("%s/%s \n", tck.BaseAsset, tck.QuoteAsset))
		if err != nil {
			return err
		}
	}

	return nil
}
