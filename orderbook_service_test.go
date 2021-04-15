package upbit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type orderBookServiceTestSuite struct {
	baseTestSuite
}

func TestOrderBookService(t *testing.T) {
	suite.Run(t, new(orderBookServiceTestSuite))
}

func (s *orderBookServiceTestSuite) TestOrderBook() {
	orderbooks, err := s.client.NewListOrderBooksService().Markets("KRW-BTC").Do(newContext())
	if err != nil {
		fmt.Printf("%+v \n", err)
	}
	fmt.Printf("%+v \n", err)
	for _, orderbook := range orderbooks {
		fmt.Printf("%+v \n", *orderbook)
	}
}
