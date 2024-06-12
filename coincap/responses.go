package coincap

import "fmt"

type assetsResponse struct {
	Assets    []Asset `json:"data"`
	Timestamp int64   `json:"timestamp"`
}

type assetResponse struct {
	Asset     Asset `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

type Asset struct {
	ID           string `json:"id"`
	Rank         string `json:""`
	Symbol       string `json:""`
	Name         string `json:""`
	Supply       string `json:""`
	MaxSupply    string `json:""`
	MarketCapUSD string `json:""`
	VolumeUSD24h string `json:""`
	PriceUSD     string `json:"priceUSD"`
}

func (d Asset) Info() string {
	return fmt.Sprintf("[ID] %s | [Rank] %s | [Symbol] %s | [Name] %s | [Price] %s", d.ID, d.Rank, d.Symbol, d.Name, d.PriceUSD)
}
