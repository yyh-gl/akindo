package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/yyh-gl/fx-auto-trader/akindo"
	"github.com/yyh-gl/fx-auto-trader/oanda"
)

func main() {
	l := newLogger()

	logFileName := fmt.Sprintf("./logs/%s_trade_log.csv", time.Now().Format("060102150405"))
	f, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		l.Fatal(err)
	}
	defer func() { _ = f.Close() }()

	at := os.Getenv("API_ACCESS_TOKEN")
	id := os.Getenv("ACCOUNT_ID")
	oc := oanda.NewClient(at, id)
	a := akindo.New(oc, f, "USD_JPY", 10)

	// contextによって実行時間を指定
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	if err := a.GoToTrade(ctx); err != nil {
		l.Fatal(err)
	}
	defer cancel()

	os.Exit(0)
}
