package upbit

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yepp4you/go-upbit/common"
)

// KlinesService list klines
type KlinesService struct {
	c        *Client
	market   string
	to       *string
	count    *int
	interval *string
	unit     *string
}

func (s *KlinesService) Market(market string) *KlinesService {
	s.market = market
	return s
}

func (s *KlinesService) To(to string) *KlinesService {
	s.to = &to
	return s
}

func (s *KlinesService) Count(count int) *KlinesService {
	s.count = &count
	return s
}

func (s *KlinesService) Unit(unit string) *KlinesService {
	s.unit = &unit
	return s
}

func (s *KlinesService) Interval(interval string) *KlinesService {
	s.interval = &interval
	return s
}

// Do send request
func (s *KlinesService) Do(ctx context.Context, opts ...RequestOption) (res []*Kline, err error) {
	endpoint := "/v1/candles"
	if *s.interval == "minutes" {
		endpoint = fmt.Sprintf("%s/%s/%s", endpoint, *s.interval, *s.unit)
	} else {
		endpoint = fmt.Sprintf("%s/%s", endpoint, *s.interval)
	}
	r := &request{
		method:   "GET",
		endpoint: endpoint,
	}

	r.setParam("market", s.market)
	if s.to != nil {
		r.setParam("to", *s.to)
	}
	if s.count != nil {
		r.setParam("count", *s.count)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	res = make([]*Kline, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Kline define kline info
type Kline struct {
	Market               string  `json:"market"`                  // "KRW-BTC",
	CandleDateTimeUtc    string  `json:"candle_date_time_utc"`    //"2018-04-18T10:16:00",
	CandleDateTimeKst    string  `json:"candle_date_time_kst"`    //: "2018-04-18T19:16:00",
	OpenPrice            float64 `json:"opening_price"`           //: 8615000,
	HighPrice            float64 `json:"high_price"`              //: 8618000,
	LowPrice             float64 `json:"low_price"`               //: 8611000,
	TradePrice           float64 `json:"trade_price"`             //: 8616000,
	Timestamp            int64   `json:"timestamp"`               //: 1524046594584,
	CandleAccTradePrice  float64 `json:"candle_acc_trade_price"`  //: 60018891.90054,
	CandleAccTradeVolume float64 `json:"candle_acc_trade_volume"` //: 6.96780929,
	Unit                 int     `json:"unit"`                    //: 1
}
