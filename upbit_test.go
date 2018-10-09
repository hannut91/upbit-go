package upbit

import (
	"fmt"
	"log"
)

var client *Client

const (
	accessKey = ""
	secretKey = ""
)

func setUp() {
	client = NewClient(accessKey, secretKey)
}

func ExampleGetMarkets() {
	setUp()

	markets, err := client.Markets()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(markets[0].Market)

	// Output:
	// KRW-BTC
}

func ExampleGetMinuteCandles() {
	setUp()

	candles, err := client.MinuteCandles(1, "KRW-BTC", map[string]string{
		"count": "1",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(candles))
	fmt.Println(candles[0].Market)

	// Output:
	// 1
	// KRW-BTC
}

func ExampleWrongUnitGetMinuteCandles() {
	setUp()

	_, err := client.MinuteCandles(2, "KRW-BTC")

	fmt.Println(err)

	// Output:
	// Invalid unit
}

func ExampleDayCandles() {
	setUp()

	candles, err := client.DayCandles("KRW-BTC", map[string]string{
		"count": "5",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(candles))
	fmt.Println(candles[0].Market)

	// Output:
	// 5
	// KRW-BTC
}

func ExampleWeekCandles() {
	setUp()

	candles, err := client.WeekCandles("KRW-BTC", map[string]string{
		"count": "5",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(candles))
	fmt.Println(candles[0].Market)

	// Output:
	// 5
	// KRW-BTC
}

func ExampleMonthCandles() {
	setUp()

	candles, err := client.MonthCandles("KRW-BTC", map[string]string{
		"count": "5",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(candles))
	fmt.Println(candles[0].Market)

	// Output:
	// 5
	// KRW-BTC
}

func ExampleTradeTicks() {
	setUp()

	tradeTicks, err := client.TradeTicks("KRW-BTC", map[string]string{
		"count": "5",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(tradeTicks))
	fmt.Println(tradeTicks[0].Market)

	// Output:
	// 5
	// KRW-BTC
}

func ExampleGetTickers() {
	setUp()

	ticks, err := client.Ticker("KRW-BTC, KRW-TRX")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(ticks))
	fmt.Println(ticks[0].Market, ticks[1].Market)

	// Output:
	// 2
	// KRW-BTC KRW-TRX
}

func ExampleOrderbooks() {
	setUp()

	ticks, err := client.Orderbooks("KRW-BTC, KRW-TRX")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(ticks))
	fmt.Println(ticks[0].Market, ticks[1].Market)
	fmt.Println(len(ticks[0].OrderbookUnits))

	// Output:
	// 2
	// KRW-BTC KRW-TRX
	// 10
}

func ExampleAccounts() {
	setUp()

	accounts, err := client.Accounts()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(accounts[0].Currency)

	// Output:
	// KRW
}

func ExampleOrderChange() {
	setUp()

	orderChance, err := client.OrderChance("KRW-BTC")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(orderChance.Market.Id)

	// Output:
	// KRW-BTC
}

// func ExampleSell() {
// 	setUp()
//
// 	order, err := client.Order(
// 		strconv.Itoa(int(util.TimeStamp())),
// 		"bid",
// 		"BTC-TRX",
// 		"0.000003",
// 		"1000",
// 		"limit",
// 	)
// 	if err != nil {
// 		log.Panicln(err)
// 	}
//
// 	fmt.Println(order)
//
// 	// Output:
// 	// BTC-TRX
// }
