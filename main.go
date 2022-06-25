package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"log"
	"main/ent"
	"main/markets"
	"os"
)

func init() {
	err := godotenv.Load("settings.env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	gocron.Every(1).Minutes().At("now").Do(updater)
	<-gocron.Start()
}

func updater() {
	fmt.Println("RUN")
	result, err := markets.GetAllWithCoinGecko()
	if err != nil {
		fmt.Println("Can't make request: ", err)
	}

	client, err := ent.Open("mysql", fmt.Sprint(os.Getenv("MDB_LOGIN"), ":", os.Getenv("MDB_PASSWORD"), "@tcp(", os.Getenv("MDB_HOST"), ":", os.Getenv("MDB_PORT"), ")/charts?parseTime=True"))
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	_, err = client.USDChart.Create().SetPrice(result.CurrentPrice["usd"]).Save(context.Background())
	if err != nil {
		fmt.Errorf("failed creating user: %w", err)
	}

	_, err = client.RUBChart.Create().SetPrice(result.CurrentPrice["rub"]).Save(context.Background())
	if err != nil {
		fmt.Errorf("failed creating user: %w", err)
	}
	fmt.Println("DONE")
}
