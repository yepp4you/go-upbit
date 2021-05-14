package upbit

import (
	"context"
	"encoding/json"

	"github.com/yepp4you/go-upbit/common"
)

// KlinesService list klines
type MarketService struct {
	c         *Client
	isDetails *bool
}

func (s *MarketService) IsDetails(isDetails bool) *MarketService {
	s.isDetails = &isDetails
	return s
}

// Do send request
func (s *MarketService) Do(ctx context.Context, opts ...RequestOption) (res []*Market, err error) {
	endpoint := "/v1/market/all"
	r := &request{
		method:   "GET",
		endpoint: endpoint,
	}

	if s.isDetails != nil {
		r.setParam("isDetails", *s.isDetails)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	res = make([]*Market, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type Market struct {
	Market        string `json:"market"`         // "KRW-BTC",
	KoreanName    string `json:"korean_name"`    // "비트코인",
	EnglishName   string `json:"english_name"`   // "Bitcoin",
	MarketWarning string `json:"market_warning"` // NONE (해당 사항 없음), CAUTION(투자유의),
}
