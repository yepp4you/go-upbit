package upbit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type klineServiceTestSuite struct {
	baseTestSuite
}

func TestKlineService(t *testing.T) {
	suite.Run(t, new(klineServiceTestSuite))
}

func (s *klineServiceTestSuite) TestMinutesKlines() {
	interval := "minutes"
	unit := "60"
	count := 20
	klines, err := s.client.NewKlinesService().Market("KRW-BTC").
		Interval(interval).Unit(unit).Count(count).Do(newContext())
	for _, kline := range klines {
		fmt.Printf("%+v \n", *kline)
	}
	fmt.Printf("%+v \n", klines)
	fmt.Printf("%+v \n", err)
}

func (s *klineServiceTestSuite) TestDaysKlines() {
	interval := "days"
	count := 2
	klines, err := s.client.NewKlinesService().Market("KRW-BTC").
		Interval(interval).Count(count).Do(newContext())

	if err != nil {
		fmt.Printf("err => %+v", err)
	}

	for _, kline := range klines {
		fmt.Printf("%+v \n", *kline)
	}
}

func (s *klineServiceTestSuite) TestMonthsKlines() {
	interval := "months"
	count := 1
	klines, err := s.client.NewKlinesService().Market("KRW-BTC").
		Interval(interval).Count(count).Do(newContext())

	fmt.Printf("%+v \n", klines)
	fmt.Printf("%+v \n", err)
}

func (s *klineServiceTestSuite) assertKlineEqual(e, a *Kline) {
	//r := s.r()
}
