package main

import (
	"crypto/internal/service"
	"log"
	"os"
)

func main() {

	err := service.GetExchangeServiceFor(os.Getenv("exchange")).GetPriceDetails()
	if err != nil {
		log.Panic("error while fetching detail", err)
	}

	log.Println("successfully fetched details")
}
