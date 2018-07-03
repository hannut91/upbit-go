# Upbit-Go

A Golang SDK for Upbit API

Upbit full document is [here](https://docs.upbit.com/).

### Installation
```bash
go get github.com/hannut91/upbit-go
```

## Getting started

```go
package main

import (
	"fmt"

	"github.com/hannut91/upbit-go"
)

func main() {
	client := upbit.NewClient("YourAccessKey", "YourSecretKey")

	markets, err := client.Markets()
	if err != nil {
	    return
	}

	fmt.Println(markets[0].Market) // KRW-BTC
}
```
