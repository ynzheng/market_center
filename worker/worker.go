package worker

import (
	"context"
	"github.com/goex-top/market_center/data"
	goex "github.com/nntaoli-project/GoEx"
	"log"
	"time"
)

var (
	defaultDepthSize = 20
)

func SetDefaultDepthSize(size int) {
	defaultDepthSize = size
}

func GetDefaultDepthSize() int {
	return defaultDepthSize
}

func NewDepthWorker(ctx context.Context, depthData *data.Data, api goex.API, pair goex.CurrencyPair, ticker *time.Ticker) {
	log.Printf("new depth worker for [%s] %s ", api.GetExchangeName(), pair.String())

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			dep, err := api.GetDepth(defaultDepthSize, pair)
			if err != nil {
				log.Printf("[%s] refresh depth error:%s", api.GetExchangeName(), err.Error())
			}
			//log.Println("DEPTH:", dep)
			depthData.UpdateDepth(api.GetExchangeName(), pair.String(), dep)
		}
	}
}

func NewTickerWorker(ctx context.Context, tickerData *data.Data, api goex.API, pair goex.CurrencyPair, ticker *time.Ticker) {
	log.Printf("new ticker worker for [%s] %s ", api.GetExchangeName(), pair.String())
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			tick, err := api.GetTicker(pair)
			if err != nil {
				log.Printf("[%s] refresh ticker error:%s", api.GetExchangeName(), err.Error())
			}
			//log.Println("TICKER:", tick)
			tickerData.UpdateTicker(api.GetExchangeName(), pair.String(), tick)
		}
	}
}
