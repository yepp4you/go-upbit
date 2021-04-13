package upbit

import (
	"bytes"
	"context"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/yepp4you/go-upbit/common"
)

// SideType define side type of order
type SideType string

// OrderType define order type
type OrderType string

// OrderState define order state
type OrderState string

// Endpoints
const (
	baseAPIMainURL = "https://api.upbit.com"
)

// UseTestnet switch all the API endpoints from production to the testnet
var UseTestnet = false

// Global enums
const (
	SideTypeBuy  SideType = "bid"
	SideTypeSell SideType = "ask"

	OrderTypeLimit  OrderType = "limit"  // 지정가 주문
	OrderTypePrice  OrderType = "price"  // 시장가 주문 (매수)
	OrderTypeMarket OrderType = "market" // 시장가 주문 (매도)

	OrderStateWait   OrderState = "wait"   // 체결 대기 (default)
	OrderStateWatch  OrderState = "watch"  // 예약주문 대기
	OrderStateDone   OrderState = "done"   // 전체 체결 완료
	OrderStateCancel OrderState = "cancel" // 주문 취소

	timestampKey  = "timestamp"
	signatureKey  = "signature"
	recvWindowKey = "recvWindow"
)

func currentTimestamp() int64 {
	return FormatTimestamp(time.Now())
}

// FormatTimestamp formats a time into Unix timestamp in milliseconds, as requested by Binance.
func FormatTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func newJSON(data []byte) (j *simplejson.Json, err error) {
	j, err = simplejson.NewJson(data)
	if err != nil {
		return nil, err
	}
	return j, nil
}

// getAPIEndpoint return the base endpoint of the Rest API according the UseTestnet flag
func getAPIEndpoint() string {
	return baseAPIMainURL
}

// NewClient initialize an API client instance with API key and secret key.
// You should always call this function before using this SDK.
// Services will be created by the form client.NewXXXService().
func NewClient(accessKey, secretKey string) *Client {
	return &Client{
		AccessKey:  accessKey,
		SecretKey:  secretKey,
		BaseURL:    getAPIEndpoint(),
		UserAgent:  "Upbit/golang",
		QueryHash:  sha512.New(),
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "Upbit-golang ", log.LstdFlags),
	}
}

type doFunc func(req *http.Request) (*http.Response, error)

// Client define API client
type Client struct {
	AccessKey  string
	SecretKey  string
	BaseURL    string
	UserAgent  string
	QueryHash  hash.Hash
	HTTPClient *http.Client
	Debug      bool
	Logger     *log.Logger
	TimeOffset int64
	do         doFunc
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

func (c *Client) parseRequest(r *request, opts ...RequestOption) (err error) {
	// set request options from user
	for _, opt := range opts {
		opt(r)
	}
	err = r.validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)
	claim := jwt.MapClaims{
		"access_key": c.AccessKey,
		"nonce":      uuid.New().String(),
	}
	queryString := r.query.Encode()
	body := &bytes.Buffer{}
	bodyString := r.form.Encode()
	header := http.Header{}
	if bodyString != "" {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		body = bytes.NewBufferString(bodyString)
	}

	if queryString != "" {
		// claim["query"] = queryString
		c.QueryHash.Reset()
		c.QueryHash.Write([]byte(queryString))
		claim["query_hash"] = hex.EncodeToString(c.QueryHash.Sum(nil))
		claim["query_hash_alg"] = "SHA512"
	}

	fmt.Println(queryString)

	if bodyString != "" {
		// claim["query"] = bodyString
		c.QueryHash.Reset()
		c.QueryHash.Write([]byte(bodyString))
		claim["query_hash"] = hex.EncodeToString(c.QueryHash.Sum(nil))
		claim["query_hash_alg"] = "SHA512"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, e := token.SignedString([]byte(c.SecretKey[:]))
	if e != nil {
		return err
	}
	header.Add("Authorization", "Bearer "+signedToken)

	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	c.debug("full url: %s, body: %s", fullURL, bodyString)

	r.fullURL = fullURL
	r.header = header
	r.body = body
	return nil
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	c.debug("request: %#v", req)
	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}
	res, err := f(req)
	if err != nil {
		return []byte{}, err
	}
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		// Only overwrite the retured error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && cerr != nil {
			err = cerr
		}
	}()
	c.debug("response: %#v", res)
	c.debug("response body: %s", string(data))
	c.debug("response status code: %d", res.StatusCode)

	// fmt.Printf("%+v", string(data))

	if res.StatusCode >= 400 {
		apiErr := new(common.APIError)
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s", e)
		}
		return nil, apiErr
	}
	return data, nil
}

// NewKlinesService init klines service
func (c *Client) NewKlinesService() *KlinesService {
	return &KlinesService{c: c}
}

// NewCreateOrderService init creating order service
func (c *Client) NewCreateOrderService() *CreateOrderService {
	return &CreateOrderService{c: c}
}

// NewGetOrderService init get order service
func (c *Client) NewGetOrderService() *GetOrderService {
	return &GetOrderService{c: c}
}

// NewCancelOrderService init cancel order service
func (c *Client) NewCancelOrderService() *CancelOrderService {
	return &CancelOrderService{c: c}
}

// NewListOrdersService init listing orders service
func (c *Client) NewListOrdersService() *ListOrdersService {
	return &ListOrdersService{c: c}
}

// NewGetAccountService init getting account service
func (c *Client) NewGetAccountService() *GetAccountService {
	return &GetAccountService{c: c}
}

// NewListTradesService init listing trades service
func (c *Client) NewListPricesService() *ListPricesService {
	return &ListPricesService{c: c}
}

// NewListTradesService init listing trades service
func (c *Client) NewListTradesService() *ListTradesService {
	return &ListTradesService{c: c}
}

func (c *Client) NewListOrderbooksService() *ListOrderBooksService {
	return &ListOrderBooksService{c: c}
}
