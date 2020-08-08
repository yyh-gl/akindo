package main

import (
	"context"
	"os"
	"time"

	"github.com/yyh-gl/fx-auto-trader/akindo"
	"github.com/yyh-gl/fx-auto-trader/oanda"
)

func main() {
	at := os.Getenv("API_ACCESS_TOKEN")
	id := os.Getenv("ACCOUNT_ID")
	oc := oanda.NewClient(at, id)

	a := akindo.New(oc, "USD_JPY")

	// contextによって実行時間を指定
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	if err := a.GoToTrade(ctx); err != nil {
		os.Exit(1)
	}
}
