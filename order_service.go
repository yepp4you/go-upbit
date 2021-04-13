package upbit

import (
	"context"
	"encoding/json"
)

// CreateOrderService create order
type CreateOrderService struct {
	c          *Client
	market     string
	side       SideType
	volume     *string
	price      *string
	ordType    OrderType
	identifier string
}

func (s *CreateOrderService) Market(market string) *CreateOrderService {
	s.market = market
	return s
}

func (s *CreateOrderService) Price(price string) *CreateOrderService {
	s.price = &price
	return s
}

func (s *CreateOrderService) Volume(volume string) *CreateOrderService {
	s.volume = &volume
	return s
}

func (s *CreateOrderService) Side(side SideType) *CreateOrderService {
	s.side = side
	return s
}

func (s *CreateOrderService) Type(ordType OrderType) *CreateOrderService {
	s.ordType = ordType
	return s
}

func (s *CreateOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "POST",
		endpoint: endpoint,
	}
	m := params{
		"market":   s.market,
		"side":     s.side,
		"ord_type": s.ordType,
	}
	if s.volume != nil {
		m["volume"] = *s.volume
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *CreateOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateOrderResponse, err error) {
	data, err := s.createOrder(ctx, "/v1/orders", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateOrderResponse define create order response
type CreateOrderResponse struct {
	UUID            string `json:"uuid"`             // "cdd92199-2897-4e14-9448-f923320408ad",
	Side            string `json:"side"`             //:"bid",
	OrdType         string `json:"ord_type"`         //:"limit",
	Price           string `json:"price"`            //:"100.0",
	AvgPrice        string `json:"avg_price"`        //:"0.0",
	State           string `json:"state"`            //:"wait",
	Market          string `json:"market"`           //:"KRW-BTC",
	CreatedAt       string `json:"created_at"`       //:"2018-04-10T15:42:23+09:00",
	Volume          string `json:"volume"`           //:"0.01",
	RemainingVolume string `json:"remaining_volume"` //:"0.01",
	ReservedFee     string `json:"reserved_fee"`     //:"0.0015",
	RemainingFee    string `json:"remaining_fee"`    //:"0.0015",
	PaidFee         string `json:"paid_fee"`         //:"0.0",
	Locked          string `json:"locked"`           //:"1.0015",
	ExecutedVolume  string `json:"executed_volume"`  //:"0.0",
	TradersCount    int64  `json:"trades_count"`     //:0
}

// ListOpenOrdersService list opened orders
// GetOrderService get an order
type GetOrderService struct {
	c          *Client
	market     string
	uuid       *string
	identifier *string
}

// Symbol set symbol
func (s *GetOrderService) Market(market string) *GetOrderService {
	s.market = market
	return s
}

func (s *GetOrderService) UUID(uuid string) *GetOrderService {
	s.uuid = &uuid
	return s
}

func (s *GetOrderService) Identifier(identifier string) *GetOrderService {
	s.identifier = &identifier
	return s
}

// Do send request
func (s *GetOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Order, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/v1/order",
	}
	r.setParam("market", s.market)
	if s.uuid != nil {
		r.setParam("uuid", *s.uuid)
	}
	if s.identifier != nil {
		r.setParam("identifier", *s.identifier)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(Order)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Order define order info
type Order struct {
	UUID            string `json:"uuid"`             // "cdd92199-2897-4e14-9448-f923320408ad",
	Side            string `json:"side"`             //:"bid",
	OrdType         string `json:"ord_type"`         //:"limit",
	Price           string `json:"price"`            //:"100.0",
	AvgPrice        string `json:"avg_price"`        //:"0.0",
	State           string `json:"state"`            //:"wait",
	Market          string `json:"market"`           //:"KRW-BTC",
	CreatedAt       string `json:"created_at"`       //:"2018-04-10T15:42:23+09:00",
	Volume          string `json:"volume"`           //:"0.01",
	RemainingVolume string `json:"remaining_volume"` //:"0.01",
	ReservedFee     string `json:"reserved_fee"`     //:"0.0015",
	RemainingFee    string `json:"remaining_fee"`    //:"0.0015",
	PaidFee         string `json:"paid_fee"`         //:"0.0",
	Locked          string `json:"locked"`           //:"1.0015",
	ExecutedVolume  string `json:"executed_volume"`  //:"0.0",
	TradersCount    string `json:"trades_count"`     //:0
	Trades          []Fill `json:"trades,omitempty"`
}

type Fill struct {
	Market string `json:"market"` //: "KRW-BTC",
	UUID   string `json:"uuid"`   //: "9e8f8eba-7050-4837-8969-cfc272cbe083",
	Price  string `json:"price"`  //: "4280000.0",
	Volume string `json:"volume"` // : "1.0",
	Funds  string `json:"funds"`  //: "4280000.0",
	Side   string `json:"side"`   //: "ask"
}

// ListOrdersService all account orders; active, canceled, or filled
type ListOrdersService struct {
	c       *Client
	market  string
	state   *OrderState
	page    *int
	limit   *int
	orderBy *string
}

func (s *ListOrdersService) Market(market string) *ListOrdersService {
	s.market = market
	return s
}

func (s *ListOrdersService) State(state OrderState) *ListOrdersService {
	s.state = &state
	return s
}

func (s *ListOrdersService) Page(page int) *ListOrdersService {
	s.page = &page
	return s
}

func (s *ListOrdersService) Limit(limit int) *ListOrdersService {
	s.limit = &limit
	return s
}

func (s *ListOrdersService) OrderBy(orderBy string) *ListOrdersService {
	s.orderBy = &orderBy
	return s
}

// Do send request
func (s *ListOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*Order, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/v1/orders",
	}
	r.setParam("market", s.market)
	if s.state != nil {
		r.setParam("state", *s.state)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.orderBy != nil {
		r.setParam("order_by", *s.orderBy)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Order{}, err
	}
	res = make([]*Order, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Order{}, err
	}
	return res, nil
}

// CancelOrderService cancel an order
type CancelOrderService struct {
	c          *Client
	market     string
	uuid       *string
	identifier *string
}

// Symbol set symbol
func (s *CancelOrderService) Market(market string) *CancelOrderService {
	s.market = market
	return s
}

func (s *CancelOrderService) UUID(uuid string) *CancelOrderService {
	s.uuid = &uuid
	return s
}

func (s *CancelOrderService) Identifier(identifier string) *CancelOrderService {
	s.identifier = &identifier
	return s
}

// Do send request
func (s *CancelOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Order, err error) {
	r := &request{
		method:   "DELETE",
		endpoint: "/v1/order",
	}
	r.setParam("market", s.market)
	if s.uuid != nil {
		r.setParam("uuid", *s.uuid)
	}
	if s.identifier != nil {
		r.setParam("identifier", *s.identifier)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(Order)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
