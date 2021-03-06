package main

import (
  // "fmt"
  "context"
  luno "github.com/luno/luno-go"
	"github.com/luno/luno-go/decimal"
)

func getTickerRequest() (*luno.Client, *luno.GetTickerRequest){
  lunoClient := luno.NewClient()
  lunoClient.SetAuth("k527sj3ed4sd8", "bCf0zAnNL0Su1SZIKGmo43LAPZUfem0SvzwZZAdv7Kk")

  return lunoClient, &luno.GetTickerRequest{Pair: pair}
}

func getCurrBid() decimal.Decimal{
  res, err := client.GetTicker(context.Background(), reqPointer)
  if err != nil{
    panic(err)
  }
  return res.Bid
}

func getCurrAsk() decimal.Decimal{
  res, err := client.GetTicker(context.Background(), reqPointer)
  if err != nil{
    panic(err)
  }
  return res.Ask
}

func getTicker() (decimal.Decimal, decimal.Decimal, luno.Time) {
  res, err := client.GetTicker(context.Background(), reqPointer)
  if err != nil{
    panic(err)
  }
  return res.Bid, res.Ask, res.Timestamp
}

func getAsset (currency string) decimal.Decimal{
	balancesReq := luno.GetBalancesRequest{}
	balances, err := client.GetBalances(context.Background(), &balancesReq)
	if err != nil {panic(err)}

  for _, accBalance := range balances.Balance {
    if accBalance.Asset == currency {return accBalance.Balance}
  }

  panic("Cannot retrieve account balance")
}
