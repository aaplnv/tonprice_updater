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

	// Record quotes for AED (United Arab Emirates Dirham)
	client.AEDQuote.Create().SetPrice(result.CurrentPrice["aed"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for ARS (Argentine Peso)
	client.ARSQuote.Create().SetPrice(result.CurrentPrice["ars"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for AUD (Australian Dollar)
	client.AUDQuote.Create().SetPrice(result.CurrentPrice["aud"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for BHD (Bahraini Dinar)
	client.BHDQuote.Create().SetPrice(result.CurrentPrice["bhd"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for BRL (Brazilian Real)
	client.BRLQuote.Create().SetPrice(result.CurrentPrice["brl"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for BTC (Bitcoin)
	client.BTCQuote.Create().SetPrice(result.CurrentPrice["btc"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for CAD (Canadian Dollar)
	client.CADQuote.Create().SetPrice(result.CurrentPrice["cad"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for CHF (Swiss Franc)
	client.CHFQuote.Create().SetPrice(result.CurrentPrice["chf"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for CLP (Chilean Peso)
	client.CLPQuote.Create().SetPrice(result.CurrentPrice["clp"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for CNY (Chinese Yuan)
	client.CNYQuote.Create().SetPrice(result.CurrentPrice["cny"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for CZK (Czech Koruna)
	client.CZKQuote.Create().SetPrice(result.CurrentPrice["czk"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for EUR (Euro)
	client.EUROQuote.Create().SetPrice(result.CurrentPrice["eur"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for GBP (British Pound)
	client.GBPQuote.Create().SetPrice(result.CurrentPrice["gbp"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for HKD (Hong Kong Dollar)
	client.HKDQuote.Create().SetPrice(result.CurrentPrice["hkd"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for HUF (Hungarian Forint)
	client.HUFQuote.Create().SetPrice(result.CurrentPrice["huf"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for IDR (Indonesian Rupiah)
	client.IDRQuote.Create().SetPrice(result.CurrentPrice["idr"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for ILS (Israel New Shekel)
	client.ILSQuote.Create().SetPrice(result.CurrentPrice["ils"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for INR (Indian Rupee)
	client.INRQuote.Create().SetPrice(result.CurrentPrice["inr"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for JPY (Japanese Yen)
	client.JPYQuote.Create().SetPrice(result.CurrentPrice["jpy"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for MXN (Mexican Peso)
	client.MXNQuote.Create().SetPrice(result.CurrentPrice["mxn"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for NOK (Norwegian Krone)
	client.NOKQuote.Create().SetPrice(result.CurrentPrice["nok"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for NZD (New Zealand Dollar)
	client.NZDQuote.Create().SetPrice(result.CurrentPrice["nzd"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for PKR (Pakistani Rupee)
	client.PKRQuote.Create().SetPrice(result.CurrentPrice["pkr"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for PLN (Polish Zloty)
	client.PLNQuote.Create().SetPrice(result.CurrentPrice["pln"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for RUB (Russian Ruble)
	client.RUBQuote.Create().SetPrice(result.CurrentPrice["rub"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for SAR (Saudi Riyal)
	client.SARQuote.Create().SetPrice(result.CurrentPrice["sar"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for SEK (Swedish Krona)
	client.SEKQuote.Create().SetPrice(result.CurrentPrice["sek"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for TRY (Turkish Lira)
	client.TRYQuote.Create().SetPrice(result.CurrentPrice["try"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for TWD (New Taiwan Dollar)
	client.TWDQuote.Create().SetPrice(result.CurrentPrice["twd"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for UAH (Ukrainian Hryvnia)
	client.UAHQuote.Create().SetPrice(result.CurrentPrice["uah"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for USD (United States Dollar)
	client.USDQuote.Create().SetPrice(result.CurrentPrice["usd"]).SetTimestamp(time.Now()).Save(context.Background())

	// Record quotes for ZAR (South African Rand)
	client.ZARQuote.Create().SetPrice(result.CurrentPrice["zar"]).SetTimestamp(time.Now()).Save(context.Background())

	log.Info("Rates successfully updated")
}
