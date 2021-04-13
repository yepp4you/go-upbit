package upbit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type baseOrderTestSuite struct {
	baseTestSuite
}

type orderServiceTestSuite struct {
	baseOrderTestSuite
}

func TestOrderService(t *testing.T) {
	suite.Run(t, new(orderServiceTestSuite))
}

func (s *orderServiceTestSuite) TestBuyOrderMarket() {
	market := "KRW-EOS"
	side := SideTypeBuy
	orderType := OrderTypePrice
	price := "10000"
	res, err := s.client.NewCreateOrderService().Market(market).Side(side).
		Type(orderType).Price(price).Do(newContext())

	if err != nil {
		fmt.Printf("%+v \n", err)
	}
	fmt.Printf("%+v \n", res)
}

func (s *orderServiceTestSuite) TestSellOrderMarket() {
	market := "KRW-EOS"
	side := SideTypeSell
	orderType := OrderTypeMarket
	volume := "2.5"
	res, err := s.client.NewCreateOrderService().Market(market).Side(side).
		Type(orderType).Volume(volume).Do(newContext())

	if err != nil {
		fmt.Printf("%+v \n", err)
	}
	fmt.Printf("%+v \n", res)
}

func (s *orderServiceTestSuite) TestBuyOrderLimit() {
	market := "KRW-EOS"
	side := SideTypeBuy
	orderType := OrderTypeLimit
	price := "8000"
	volume := "2"
	res, err := s.client.NewCreateOrderService().Market(market).Side(side).
		Type(orderType).Price(price).Volume(volume).Do(newContext())

	if err != nil {
		fmt.Printf("%+v \n", err)
	}
	fmt.Printf("%+v \n", res)
}

func (s *orderServiceTestSuite) TestSellOrderLimit() {
	market := "KRW-EOS"
	side := SideTypeSell
	orderType := OrderTypeLimit
	price := "8400"
	volume := "1.1"
	res, err := s.client.NewCreateOrderService().Market(market).Side(side).
		Type(orderType).Price(price).Volume(volume).Do(newContext())

	if err != nil {
		fmt.Printf("%+v \n", err)
	}
	fmt.Printf("%+v \n", res)
}

func (s *orderServiceTestSuite) TestGetOrder() {
	market := "KRW-EOS"
	uuid := "6c29dbf5-949a-4f3a-89f6-126c7621e6b9"
	// id := "myOrder1"
	order, err := s.client.NewGetOrderService().Market(market).
		UUID(uuid).Do(newContext())

	if err != nil {
		fmt.Printf("%+v \n", err)
	}
	fmt.Printf("%+v \n", order)
}

func (s *orderServiceTestSuite) TestListOrders() {
	market := "KRW-EOS"
	state := OrderStateDone
	page := 2
	limit := 10
	orderBy := "desc"
	orders, err := s.client.NewListOrdersService().Market(market).
		State(state).Page(page).Limit(limit).OrderBy(orderBy).Do(newContext())

	if err != nil {
		fmt.Printf("%+v \n", err)
	}
	fmt.Printf("%+v \n", orders)

}
