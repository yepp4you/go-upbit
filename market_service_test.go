package upbit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type marketServiceTestSuite struct {
	baseTestSuite
}

func TestMarketService(t *testing.T) {
	suite.Run(t, new(marketServiceTestSuite))
}

func (s *marketServiceTestSuite) TestMarket() {
	markets, err := s.client.NewMarketService().IsDetails(true).Do(newContext())
	if err != nil {
		fmt.Printf("%+v \n", err)
	}
	for _, market := range markets {
		fmt.Printf("%+v \n", *market)
	}
}
