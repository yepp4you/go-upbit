package upbit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type tickerServiceTestSuite struct {
	baseTestSuite
}

func TestTickerService(t *testing.T) {
	suite.Run(t, new(tickerServiceTestSuite))
}

func (s *tickerServiceTestSuite) TestTickerService() {
	prices, err := s.client.NewListPricesService().Markets("KRW-BTC").Do(newContext())

	if err != nil {
		fmt.Printf("err => %+v \n", err)
	}

	for _, price := range prices {
		fmt.Printf("%+v \n", *price)
		fmt.Printf("%.4f \n", price.TradePrice)
		fmt.Printf("%.4f \n", price.HighPrice)
	}
}
