package markets

import (
	coingecko "github.com/superoo7/go-gecko/v3"
	"net/http"
	"time"
)

func GetAllWithCoinGecko() (price float64, err error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	CG := coingecko.NewClient(httpClient)

	responce, err := CG.CoinsID("the-open-network", false, false, true, false, false, false)
	if err != nil {
		return 0, err
	}

	return responce.MarketData.CurrentPrice["usd"], nil
}
