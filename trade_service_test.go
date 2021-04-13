package upbit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type tradesServiceTestSuite struct {
	baseTestSuite
}

func TestTradesService(t *testing.T) {
	suite.Run(t, new(tradesServiceTestSuite))
}

func (s *tradesServiceTestSuite) TestTradesService() {
	count := 20
	trades, err := s.client.NewListTradesService().Market("KRW-BTC").Count(count).Do(newContext())

	if err != nil {
		fmt.Printf("err => %+v \n", err)
	}

	for _, trade := range trades {
		fmt.Printf("%+v \n", *trade)
	}
}
