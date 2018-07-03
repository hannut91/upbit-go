package upbit

import (
	"net/url"
	"strconv"

	"github.com/dgrijalva/jwt-go"

	"github.com/hannut91/upbit-go/types"
	"github.com/hannut91/upbit-go/util"
)

const (
	baseUrl = "https://api.upbit.com/v1"
)

type InvalidParams struct {
	message string
	Err     error
}

func (e *InvalidParams) Error() string {
	return e.message
}

type Client struct {
	accessKey string
	secretKey string
}

func (client *Client) Token(query map[string]string) (tokenStr string, err error) {
	claim := jwt.MapClaims{
		"access_key": client.accessKey,
		"nonce":      util.TimeStamp(),
	}

	if query != nil {
		url := new(url.URL)

		q := url.Query()

		for i, value := range query {
			q.Add(i, value)
		}

		rawQuery := q.Encode()

		claim["query"] = rawQuery
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenStr, err = token.SignedString([]byte(client.secretKey[:]))
	if err != nil {
		return
	}

	return
}

func (client *Client) Markets() (markets []*types.Market, err error) {
	options := &util.RequestOptions{Url: baseUrl + "/market/all"}
	err = util.Request(options, &markets)
	return
}

func (client *Client) MinuteCandles(
	unit int,
	market string,
	params ...map[string]string,
) (candles []*types.MinuteCandle, err error) {
	if !isValidMinuteCandleUnit(unit) {
		err = &InvalidParams{
			message: "Invalid unit",
		}
		return
	}

	query := map[string]string{
		"market": market,
	}

	if len(params) > 0 {
		for _, param := range params {
			for index, value := range param {
				query[index] = value
			}
		}
	}

	options := &util.RequestOptions{
		Url:   baseUrl + "/candles/minutes/" + strconv.Itoa(unit),
		Query: query,
	}
	err = util.Request(options, &candles)
	return
}

func (client *Client) DayCandles(
	market string,
	params ...map[string]string,
) (candles []*types.DayCandle, err error) {
	query := map[string]string{
		"market": market,
	}

	if len(params) > 0 {
		for _, param := range params {
			for index, value := range param {
				query[index] = value
			}
		}
	}

	options := &util.RequestOptions{
		Url:   baseUrl + "/candles/days",
		Query: query,
	}
	err = util.Request(options, &candles)
	return
}

func (client *Client) WeekCandles(
	market string,
	params ...map[string]string,
) (candles []*types.WeekCandle, err error) {
	query := map[string]string{
		"market": market,
	}

	if len(params) > 0 {
		for _, param := range params {
			for index, value := range param {
				query[index] = value
			}
		}
	}

	options := &util.RequestOptions{
		Url:   baseUrl + "/candles/weeks",
		Query: query,
	}
	err = util.Request(options, &candles)
	return
}

func (client *Client) MonthCandles(
	market string,
	params ...map[string]string,
) (candles []*types.MonthCandle, err error) {
	query := map[string]string{
		"market": market,
	}

	if len(params) > 0 {
		for _, param := range params {
			for index, value := range param {
				query[index] = value
			}
		}
	}

	options := &util.RequestOptions{
		Url:   baseUrl + "/candles/months",
		Query: query,
	}
	err = util.Request(options, &candles)
	return
}

func (client *Client) TradeTicks(
	market string,
	params ...map[string]string,
) (tradeTicks []*types.TradeTicks, err error) {
	query := map[string]string{
		"market": market,
	}

	if len(params) > 0 {
		for _, param := range params {
			for index, value := range param {
				query[index] = value
			}
		}
	}

	options := &util.RequestOptions{
		Url:   baseUrl + "/trades/ticks",
		Query: query,
	}
	err = util.Request(options, &tradeTicks)
	return
}

func (client *Client) Ticker(
	markets string,
) (tickers []*types.Ticker, err error) {
	query := map[string]string{
		"markets": markets,
	}

	options := &util.RequestOptions{
		Url:   baseUrl + "/ticker",
		Query: query,
	}
	err = util.Request(options, &tickers)
	return
}

func (client *Client) Orderbooks(
	markets string,
) (orderbooks []*types.Orderbook, err error) {
	query := map[string]string{
		"markets": markets,
	}

	options := &util.RequestOptions{
		Url:   baseUrl + "/orderbook",
		Query: query,
	}
	err = util.Request(options, &orderbooks)
	return
}

func (client *Client) Accounts() (balances []*types.Balance, err error) {
	token, err := client.Token(nil)
	if err != nil {
		return
	}

	options := &util.RequestOptions{
		Url: baseUrl + "/accounts",
		Headers: map[string]string{
			"Authorization": "Bearer " + token,
		},
	}
	err = util.Request(options, &balances)
	return
}

func (client *Client) OrderChance(
	marketId string,
) (orderChance types.OrderChance, err error) {
	query := map[string]string{
		"market": marketId,
	}

	token, err := client.Token(query)
	if err != nil {
		return
	}

	options := &util.RequestOptions{
		Url: baseUrl + "/orders/chance",
		Headers: map[string]string{
			"Authorization": "Bearer " + token,
		},
		Query: query,
	}
	err = util.Request(options, &orderChance)
	return
}

func NewClient(accessKey, secretKey string) *Client {
	return &Client{accessKey, secretKey}
}

func isValidMinuteCandleUnit(unit int) bool {
	return unit == 1 || unit == 3 || unit == 5 || unit == 10 || unit == 15 ||
		unit == 30 || unit == 60 || unit == 240
}
