package upbit

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type websocketServiceTestSuite struct {
	baseTestSuite
	origWsServe func(*WsConfig, WsHandler, ErrHandler) (chan struct{}, chan struct{}, error)
	serveCount  int
}

func TestWebsocketService(t *testing.T) {
	suite.Run(t, new(websocketServiceTestSuite))
}

func (s *websocketServiceTestSuite) SetupTest() {
	s.origWsServe = wsServe
}

func (s *websocketServiceTestSuite) TearDownTest() {
	wsServe = s.origWsServe
	s.serveCount = 0
}

func (s *websocketServiceTestSuite) mockWsServe(data []byte, err error) {
	wsServe = func(cfg *WsConfig, handler WsHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, innerErr error) {
		s.serveCount++
		doneC = make(chan struct{})
		stopC = make(chan struct{})
		go func() {
			<-stopC
			close(doneC)
		}()
		handler(data)
		if err != nil {
			errHandler(err)
		}
		return doneC, stopC, nil
	}
}

func (s *websocketServiceTestSuite) assertWsServe(count ...int) {
	e := 1
	if len(count) > 0 {
		e = count[0]
	}
	s.r().Equal(e, s.serveCount)
}

func (s *websocketServiceTestSuite) TestTickerServe() {
	endpoint := GetWsEndpoint()
	code := "KRW-EOS"
	ticket := "test1234"
	typ := "ticker"
	codes := []string{"KRW-EOS"}
	cfg := NewWsConfig(endpoint, code, ticket, typ, codes)
	doneC, stopC, err := WsTickerServe(cfg, func(event *WsTicker) {
		e := &WsTicker{}
		fmt.Printf("%+v \n", event)
		s.assertWsTickerEqual(e, event)
		vp := event.AccBidVolume / event.AccAskVolume * 100
		fmt.Printf("%+v \n", vp)
	}, func(err error) {
		fmt.Printf("err => %+v \n", err)
		// s.r().EqualError(err, fakeErrMsg)
	})

	if err != nil {
		fmt.Printf("err => %+v \n", err)
	}

	time.Sleep(time.Second * 5)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsTickerEqual(e, a *WsTicker) {
	// r := s.r()
}

func (s *websocketServiceTestSuite) TestWsTradeServe() {
	endpoint := GetWsEndpoint()
	code := "KRW-BTC"
	ticket := "test12345"
	typ := "trade"
	codes := []string{"KRW-BTC"}
	cfg := NewWsConfig(endpoint, code, ticket, typ, codes)
	doneC, stopC, err := WsTradeServe(cfg, func(event *WsTrade) {
		fmt.Printf("%+v\n", event)
		e := &WsTrade{}
		s.assertWsAggTradeEqual(e, event)
	}, func(err error) {
		fmt.Printf("%+v\n", err)
		//s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	time.Sleep(time.Second * 5)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsAggTradeEqual(e, a *WsTrade) {
	/*
		r := s.r()
		r.Equal(e.Event, a.Event, "Event")
		r.Equal(e.Time, a.Time, "Time")
		r.Equal(e.Symbol, a.Symbol, "Symbol")
		r.Equal(e.AggTradeID, a.AggTradeID, "AggTradeID")
		r.Equal(e.Price, a.Price, "Price")
		r.Equal(e.Quantity, a.Quantity, "Quantity")
		r.Equal(e.FirstBreakdownTradeID, a.FirstBreakdownTradeID, "FirstBreakdownTradeID")
		r.Equal(e.LastBreakdownTradeID, a.LastBreakdownTradeID, "LastBreakdownTradeID")
		r.Equal(e.TradeTime, a.TradeTime, "TradeTime")
		r.Equal(e.IsBuyerMaker, a.IsBuyerMaker, "IsBuyerMaker")
	*/
}

func (s *websocketServiceTestSuite) TestWsOrderBookServe() {
	endpoint := GetWsEndpoint()
	code := "KRW-BTC"
	ticket := "test12345"
	typ := "orderbook"
	codes := []string{"KRW-BTC"}
	cfg := NewWsConfig(endpoint, code, ticket, typ, codes)
	doneC, stopC, err := WsOrderBookServe(cfg, func(event *WsOrderBook) {
		fmt.Printf("%+v\n", event)
		e := &WsOrderBook{}
		s.assertWsOrderBookEqual(e, event)
	}, func(err error) {
		fmt.Printf("err => %+v \n", err)
	})

	if err != nil {
		fmt.Printf("err => %+v \n", err)
	}

	time.Sleep(time.Second * 5)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsOrderBookEqual(e, a *WsOrderBook) {
	/*
		r := s.r()
		r.Equal(e.Event, a.Event, "Event")
		r.Equal(e.Time, a.Time, "Time")
		r.Equal(e.Symbol, a.Symbol, "Symbol")
		r.Equal(e.AggTradeID, a.AggTradeID, "AggTradeID")
		r.Equal(e.Price, a.Price, "Price")
		r.Equal(e.Quantity, a.Quantity, "Quantity")
		r.Equal(e.FirstBreakdownTradeID, a.FirstBreakdownTradeID, "FirstBreakdownTradeID")
		r.Equal(e.LastBreakdownTradeID, a.LastBreakdownTradeID, "LastBreakdownTradeID")
		r.Equal(e.TradeTime, a.TradeTime, "TradeTime")
		r.Equal(e.IsBuyerMaker, a.IsBuyerMaker, "IsBuyerMaker")
	*/
}
