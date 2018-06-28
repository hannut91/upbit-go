package upbit_go

import (
	"strconv"

	"upbit-go/types"
	"upbit-go/util"
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

func NewClient() *Client {
	return new(Client)
}

func isValidMinuteCandleUnit(unit int) bool {
	return unit == 1 || unit == 3 || unit == 5 || unit == 10 || unit == 15 ||
		unit == 30 || unit == 60 || unit == 240
}
