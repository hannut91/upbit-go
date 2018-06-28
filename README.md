#Upbit-Go

A Golang SDK for Upbit API

Upbit full document is [here](https://docs.upbit.com/).

###Installation
```
go get github.com/hannut91/upbit-go
```

## Getting started

```
package main

import (
    "fmt"
    
    "github.com/hannut91/upbit-go"
)

func main() {
    client := client.NewClient()
    
    markets, err := client.Markets()
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(markets[0].Market) // KRW-BTC
}

```
