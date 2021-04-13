package upbit

import (
	"context"
	"encoding/json"
)

// GetAccountService get account info
type GetAccountService struct {
	c *Client
}

// Do send request
func (s *GetAccountService) Do(ctx context.Context, opts ...RequestOption) ([]Account, error) {
	r := &request{
		method:   "GET",
		endpoint: "/v1/accounts",
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := make([]Account, 0, 256)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Account define account info
type Account struct {
	Currency            string `json:"currency,omitempty"`               // "currency":"KRW",
	Balance             string `json:"balance,omitempty"`                // "1000000.0",
	Locked              string `json:"locked,omitempty"`                 // "0.0",
	AvgBuyPrice         string `json:"avg_buy_price,omitempty"`          // "0",
	AvgBuyPriceModified bool   `json:"avg_buy_price_modified,omitempty"` // false,
	UnitCurrency        string `json:"unit_currency,omitempty"`          // "KRW"
}
