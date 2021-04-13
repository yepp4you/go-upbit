package upbit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type accountServiceTestSuite struct {
	baseTestSuite
}

func TestAccountService(t *testing.T) {
	suite.Run(t, new(accountServiceTestSuite))
}

func (s *accountServiceTestSuite) TestGetAccount() {
	accounts, err := s.client.NewGetAccountService().Do(newContext())
	if err != nil {
		fmt.Printf("err => %+v", err)
		return
	}
	for _, account := range accounts {
		fmt.Printf("%+v \n", account)
	}
}
