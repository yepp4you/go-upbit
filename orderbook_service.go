package upbit

import (
	"context"
	"encoding/json"

	"github.com/yepp4you/go-upbit/common"
)

// KlinesService list klines
type ListOrderBooksService struct {
	c       *Client
	markets string
}

func (s *ListOrderBooksService) Markets(markets string) *ListOrderBooksService {
	s.markets = markets
	return s
}

// Do send request
func (s *ListOrderBooksService) Do(ctx context.Context, opts ...RequestOption) (res []*Orderbook, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/v1/orderbook",
	}

	r.setParam("markets", s.markets)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	res = make([]*Orderbook, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Kline define kline info
type Orderbook struct {
	Market         string          `json:"market"`         // "KRW-BTC",
	Timestamp      int64           `json:"timestamp"`      //: 1524046594584,
	TotalAskSize   float64         `json:"total_ask_size"` //: 8.83621228,
	TotalBitSize   float64         `json:"total_bid_size"` //: 2.43976741,
	OrderbookUnits []OrderbookUnit `json:"orderbook_units,omitempty"`
}

type OrderbookUnit struct {
	AskPrice float64 `json:"ask_price"` //: 6956000,
	BidPrice float64 `json:"bid_price"` //: 6954000,
	AskSize  float64 `json:"ask_size"`  //: 0.24078656,
	BidSize  float64 `json:"bid_size"`  //: 0.00718341
}
