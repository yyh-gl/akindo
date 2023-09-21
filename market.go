package main

import "net/http"

type Instrument string

const (
	InstrumentEURUSD Instrument = "EUR_USD"
)

type Order struct{}

type Market interface {
	CreateOrder(units int, instrument Instrument) error
	GetOrder(instrument Instrument) (Order, error)
}

type OANDAMarket struct {
	httpClient  *http.Client
	accessToken string
	accountID   string
}

func (m OANDAMarket) CreateOrder(units int, instrument Instrument) error {
	return nil
}

func (m OANDAMarket) GetOrder(instrument Instrument) (Order, error) {
	return Order{}, nil
}

var _ Market = &OANDAMarket{}

func NewOANDAMarket(accessToken, accountID string) *OANDAMarket {
	return &OANDAMarket{
		httpClient:  http.DefaultClient,
		accessToken: accessToken,
		accountID:   accountID,
	}
}
