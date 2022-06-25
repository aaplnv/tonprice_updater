package markets

import (
	coingecko "github.com/superoo7/go-gecko/v3"
	"github.com/superoo7/go-gecko/v3/types"
	"net/http"
	"time"
)

func GetAllWithCoinGecko() (marketdata *types.MarketDataItem, err error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	CG := coingecko.NewClient(httpClient)

	responce, err := CG.CoinsID("the-open-network", false, false, true, false, false, false)
	if err != nil {
		return nil, err
	}

	return responce.MarketData, nil
}
