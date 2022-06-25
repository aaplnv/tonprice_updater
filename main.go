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
	log.Info("Rates successfully updated")
}
