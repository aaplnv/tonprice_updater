package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"main/ent"
	"main/markets"
	"os"
	"time"
)

func init() {
	// Set up logging level here
	log.SetLevel(log.InfoLevel)

	err := godotenv.Load("settings.env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	gocron.Every(1).Minutes().Do(updater)
	<-gocron.Start()
}

func updater() {
	log.Info("Updater working...")
	result, err := markets.GetAllWithCoinGecko()
	if err != nil {
		log.Error("Can't access CoinGecko", err)
		return
	}

	client, err := ent.Open("mysql", fmt.Sprint(os.Getenv("MDB_LOGIN"), ":", os.Getenv("MDB_PASSWORD"), "@tcp(", os.Getenv("MDB_HOST"), ":", os.Getenv("MDB_PORT"), ")/charts?parseTime=True"))
	if err != nil {
		log.Error("Can't connect to MySQL", err)
		return
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Error("failed creating schema resources: %v", err)
		return
	}

	_, err = client.USDQuote.Create().SetPrice(result.CurrentPrice["usd"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save USD rates: %w", err)
	}

	_, err = client.RUBQuote.Create().SetPrice(result.CurrentPrice["rub"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save RUB rates: %w", err)
	}

	_, err = client.EUROQuote.Create().SetPrice(result.CurrentPrice["eur"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save EUR rates: %w", err)
	}

	_, err = client.UAHQuote.Create().SetPrice(result.CurrentPrice["uah"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save UAH rates: %w", err)
	}

	_, err = client.AUDQuote.Create().SetPrice(result.CurrentPrice["aud"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save AUD rates: %w", err)
	}

	_, err = client.GBPQuote.Create().SetPrice(result.CurrentPrice["gbp"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save GBP rates: %w", err)
	}

	_, err = client.INRQuote.Create().SetPrice(result.CurrentPrice["inr"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save INR rates: %w", err)
	}

	_, err = client.NZDQuote.Create().SetPrice(result.CurrentPrice["nzd"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save NZD rates: %w", err)
	}

	_, err = client.ZARQuote.Create().SetPrice(result.CurrentPrice["zar"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save ZAR rates: %w", err)
	}

	_, err = client.PKRQuote.Create().SetPrice(result.CurrentPrice["pkr"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save PKR rates: %w", err)
	}

	_, err = client.CNYQuote.Create().SetPrice(result.CurrentPrice["cny"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save CNY rates: %w", err)
	}

	_, err = client.HKDQuote.Create().SetPrice(result.CurrentPrice["hkd"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save HKD rates: %w", err)
	}

	_, err = client.TWDQuote.Create().SetPrice(result.CurrentPrice["twd"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save TWD rates: %w", err)
	}

	_, err = client.CHFQuote.Create().SetPrice(result.CurrentPrice["chf"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save CHF rates: %w", err)
	}

	_, err = client.CZKQuote.Create().SetPrice(result.CurrentPrice["czk"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save CZK rates: %w", err)
	}

	_, err = client.HUFQuote.Create().SetPrice(result.CurrentPrice["huf"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save HUF rates: %w", err)
	}

	_, err = client.NOKQuote.Create().SetPrice(result.CurrentPrice["nok"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save NOK rates: %w", err)
	}

	_, err = client.PLNQuote.Create().SetPrice(result.CurrentPrice["pln"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save PLN rates: %w", err)
	}

	_, err = client.SEKQuote.Create().SetPrice(result.CurrentPrice["sek"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save SEK rates: %w", err)
	}

	_, err = client.ARSQuote.Create().SetPrice(result.CurrentPrice["ars"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save ARS rates: %w", err)
	}

	_, err = client.BRLQuote.Create().SetPrice(result.CurrentPrice["brl"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save BRL rates: %w", err)
	}

	_, err = client.BTCQuote.Create().SetPrice(result.CurrentPrice["btc"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save BTC rates: %w", err)
	}

	_, err = client.CADQuote.Create().SetPrice(result.CurrentPrice["cad"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save CAD rates: %w", err)
	}

	_, err = client.CLPQuote.Create().SetPrice(result.CurrentPrice["clp"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save CLP rates: %w", err)
	}

	_, err = client.MXNQuote.Create().SetPrice(result.CurrentPrice["mxn"]).SetTimestamp(time.Now()).Save(context.Background())
	if err != nil {
		fmt.Errorf("Can't save MXN rates: %w", err)
	}
	log.Info("Rates successfully updated")
}
