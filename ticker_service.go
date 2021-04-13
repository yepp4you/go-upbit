package upbit

import (
	"context"
	"encoding/json"

	"github.com/yepp4you/go-upbit/common"
)

type ListPricesService struct {
	c       *Client
	markets string
}

func (s *ListPricesService) Markets(markets string) *ListPricesService {
	s.markets = markets
	return s
}

// Do send request
func (s *ListPricesService) Do(ctx context.Context, opts ...RequestOption) (res []*MarketPrice, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/v1/ticker",
	}
	r.setParam("markets", s.markets)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*MarketPrice{}, err
	}
	data = common.ToJSONList(data)
	res = make([]*MarketPrice, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarketPrice{}, err
	}
	return res, nil
}

// SymbolPrice define symbol and price pair
type MarketPrice struct {
	Market             string  `json:"market"`                // "KRW-BTC",
	TradeDate          string  `json:"trade_date"`            // : "20180418",
	TradeTime          string  `json:"trade_time"`            // : "102340",
	TradeDateKst       string  `json:"trade_date_kst"`        // : "20180418",
	TradeTimeKst       string  `json:"trade_time_kst"`        // : "192340",
	TradeTimestamp     int64   `json:"trade_timestamp"`       // : 1524047020000,
	OpeningPrice       float64 `json:"opening_price"`         // : 8450000,
	HighPrice          float64 `json:"high_price"`            // : 8679000,
	LowPrice           float64 `json:"low_price"`             // : 8445000,
	TradePrice         float64 `json:"trade_price"`           // : 8621000,
	PrevClosingPrice   float64 `json:"prev_closing_price"`    // : 8450000,
	Change             string  `json:"change"`                // : "RISE",
	ChangePrice        float64 `json:"change_price"`          // : 171000,
	ChangeRate         float64 `json:"change_rate"`           // : 0.0202366864,
	SignedChangePrice  float64 `json:"signed_change_price"`   // : 171000,
	SignedChangeRate   float64 `json:"signed_change_rate"`    // : 0.0202366864,
	TradeVolume        float64 `json:"trade_volume"`          // : 0.02467802,
	AccTradePrice      float64 `json:"acc_trade_price"`       // : 108024804862.58254,
	AccTradePrice24H   float64 `json:"acc_trade_price_24h"`   // : 232702901371.09309,
	AccTradeVolume     float64 `json:"acc_trade_volume"`      // : 12603.53386105,
	AccTradeVolume24H  float64 `json:"acc_trade_volume_24h"`  // : 27181.31137002,
	Highest52WeekPrice float64 `json:"highest_52_week_price"` // : 28885000,
	Highest52WeekDate  string  `json:"highest_52_week_date"`  // `: "2018-01-06",
	Lowest52WeekPrice  float64 `json:"lowest_52_week_price"`  // : 4175000,
	Lowest52WeekDate   string  `json:"lowest_52_week_date"`   // : "2017-09-25",
	Timestamp          int64   `json:"timestamp"`             // : 1524047026072
}
