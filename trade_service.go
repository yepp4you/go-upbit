package upbit

import (
	"context"
	"encoding/json"

	"github.com/yepp4you/go-upbit/common"
)

// KlinesService list klines
type ListTradesService struct {
	c       *Client
	market  string
	to      *string
	count   *int
	cursor  *int
	daysAgo *int
}

func (s *ListTradesService) Market(market string) *ListTradesService {
	s.market = market
	return s
}

func (s *ListTradesService) To(to string) *ListTradesService {
	s.to = &to
	return s
}

func (s *ListTradesService) Count(count int) *ListTradesService {
	s.count = &count
	return s
}

func (s *ListTradesService) Cursor(cursor int) *ListTradesService {
	s.cursor = &cursor
	return s
}

func (s *ListTradesService) DaysAgo(daysAgo int) *ListTradesService {
	s.daysAgo = &daysAgo
	return s
}

// Do send request
func (s *ListTradesService) Do(ctx context.Context, opts ...RequestOption) (res []*Trade, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/v1/trades/ticks",
	}

	r.setParam("market", s.market)
	if s.to != nil {
		r.setParam("to", *s.to)
	}
	if s.count != nil {
		r.setParam("count", *s.count)
	}
	if s.cursor != nil {
		r.setParam("cursor", *s.cursor)
	}
	if s.daysAgo != nil {
		r.setParam("daysago", *s.daysAgo)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	res = make([]*Trade, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Kline define kline info
type Trade struct {
	Market           string  `json:"market"`             //: "KRW-BTC",
	TradeDateUtc     string  `json:"trade_date_utc"`     //: //: "2018-04-18",
	TradeTimeUtc     string  `json:"trade_time_utc"`     //:: "10:19:58",
	Timestamp        int64   `json:"timestamp"`          //:: 1524046798000,
	TradePrice       float64 `json:"trade_price"`        //:: 8616000,
	TradeVolume      float64 `json:"trade_volume"`       //:: 0.03060688,
	PrevClosingPrice float64 `json:"prev_closing_price"` //:: 8450000,
	ChanePrice       float64 `json:"chane_price"`        //:: 166000,
	AskBid           string  `json:"ask_bid"`            //:: "ASK"
}
