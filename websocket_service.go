package upbit

import (
	"encoding/json"
	"time"
)

// Endpoints
const (
	baseWsMainURL = "wss://api.upbit.com/websocket/v1"
)

var (
	// WebsocketTimeout is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketTimeout = time.Second * 60
	// WebsocketKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketKeepalive = false
)

// getWsEndpoint return the base endpoint of the WS according the UseTestnet flag
func GetWsEndpoint() string {
	return baseWsMainURL
}

type WsTickerHandler func(event *WsTicker)

// WsTickerServe serve websocket kline handler with a symbol and interval like 15m, 30s
func WsTickerServe(cfg *WsConfig, handler WsTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	// endpoint := getWsEndpoint()
	//cfg := NewWsConfig(endpoint, code)
	wsHandler := func(message []byte) {
		event := new(WsTicker)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// WsTicker define websocket kline
type WsTicker struct {
	Type                string  `json:"type"`                 // "ticker"
	Code                string  `json:"code"`                 // market cde "KRW-BTC"
	OpeningPrice        float64 `json:"opening_price"`        // 8450000,
	HighPrice           float64 `json:"high_price"`           // 8679000,
	LowPrice            float64 `json:"low_price"`            // 8445000,
	TradePrice          float64 `json:"trade_price"`          // 8621000
	PrevClosingPrice    float64 `json:"prev_closing_price"`   // 8450000,
	Change              string  `json:"change"`               // "RISE",
	ChangePrice         float64 `json:"change_price"`         // 171000,
	SignedChangePrice   float64 `json:"signed_change_price"`  // 171000,
	ChangeRate          float64 `json:"change_rate"`          // 0.0202366864,
	SignedChangeRate    float64 `json:"signed_change_rate"`   // 0.0202366864,
	TradeVolume         float64 `json:"trade_volume"`         // 0.02467802,
	AccTradePrice       float64 `json:"acc_trade_price"`      // 108024804862.58254,
	AccTradePrice24h    float64 `json:"acc_trade_price_24h"`  // 232702901371.09309,
	AccTradeVolume      float64 `json:"acc_trade_volume"`     // 12603.53386105,
	AccTradeVolume24h   float64 `json:"acc_trade_volume_24h"` // 27181.31137002,
	TradeDate           string  `json:"trade_date"`           // "20180418",
	TradeTime           string  `json:"trade_time"`           // "102340",
	TradeDateKst        string  `json:"trade_date_kst"`       // "20180418",
	TradeTimeKst        string  `json:"trade_time_kst"`       // "192340",
	TradeTimestamp      int64   `json:"trade_timestamp"`      // 1524047020000,
	AskBid              string  `json:"ask_bid"`
	AccAskVolume        float64 `json:"acc_ask_volume"`        // 27181.31137002,
	AccBidVolume        float64 `json:"acc_bid_volume"`        // 27181.31137002,
	Hightest52WeekPrice float64 `json:"highest_52_week_price"` // 28885000,
	Highest52WeekDate   string  `json:"highest_52_week_date"`  // "2018-01-06",
	Lowest52WeekPrice   float64 `json:"lowest_52_week_price"`  // 4175000,
	Lowest52WeekDate    string  `json:"lowest_52_week_date"`   // "2017-09-25",
	TradeStatus         string  `json:"trade_status"`          // 거래상태,
	MarketStatus        string  `json:"market_status"`         // 거래상태,
	MarketStatForIos    string  `json:"market_state_for_ios"`  // 거래상태,
	IsTradingSuspended  bool    `json:"is_trading_suspended"`  // 거래정지 여부,
	DelistingDate       string  `json:"delisting_date"`
	MarketWarning       string  `json:"market_warning"`
	Timestamp           int64   `json:"timestamp"` // 1524047026072,
	StreamType          string  `json:"stream_type"`
}

// WsAggTradeHandler handle websocket aggregate trade event
type WsTradeHandler func(event *WsTrade)

// WsAggTradeServe serve websocket aggregate handler with a symbol
func WsTradeServe(cfg *WsConfig, handler WsTradeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	// endpoint := getWsEndpoint()
	// cfg := newWsConfig(endpoint, code)
	wsHandler := func(message []byte) {
		event := new(WsTrade)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

type WsTrade struct {
	Type             string  `json:"type"`               // "trade",
	Code             string  `json:"code"`               //"KRW-BTC",
	Timestamp        int64   `json:"timestamp"`          //1617846849357,
	TradeDate        string  `json:"trade_date"`         //"2021-04-08",
	TradeTime        string  `json:"trade_time"`         //"01:54:09",
	TradeTimestamp   int64   `json:"trade_timestamp"`    //1617846849000,
	TradePrice       float64 `json:"trade_price"`        //72898000.0,
	TradeVolume      float64 `json:"trade_volume"`       //0.0046,
	AskBid           string  `json:"ask_bid"`            //"BID",
	PrevClosingPrice float64 `json:"prev_closing_price"` //72850000.00000000,
	Change           string  `json:"change"`             //"RISE",
	ChangePrice      float64 `json:"change_price"`       //48000.00000000,
	SequentialId     int64   `json:"sequential_id"`      //1617846849000002,
	StreamType       string  `json:"stream_type"`        //"REALTIME"
}

type WsOrderBook struct {
	Market         string  `json:"market"`         // "KRW-BTC",
	Timestamp      int64   `json:"timestamp"`      //: 1524047680880,
	TotalAskSize   float64 `json:"total_ask_size"` // 11.2909676,
	TotalBidSize   float64 `json:"total_bid_size"` // 33.92373073,
	OrderBookUnits []Unit  `json:"orderbook_units"`
}

type Unit struct {
	AskPrice float64 `json:"ask_price"`
	BidPrice float64 `json:"bid_price"`
	AskSize  float64 `json:"ask_size"`
	BidSize  float64 `json:"bid_size"`
}

// WsBookTickerHandler handle websocket that pushes updates to the best bid or ask price or quantity in real-time for a specified symbol.
type WsOrderBookHandler func(event *WsOrderBook)

// WsBookTickerServe serve websocket that pushes updates to the best bid or ask price or quantity in real-time for a specified symbol.
func WsOrderBookServe(cfg *WsConfig, handler WsOrderBookHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	// endpoint := getWsEndpoint()
	// cfg := newWsConfig(endpoint, code)
	wsHandler := func(message []byte) {
		event := new(WsOrderBook)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}
