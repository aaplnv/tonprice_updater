package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"main/ent"
	"main/markets"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	err := godotenv.Load("settings.env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	result, err := markets.GetAllWithCoinGecko()

	if err != nil {
		fmt.Println("Can't make request: ", err)
	}
	fmt.Println("Current ton price: ", result)

	client, err := ent.Open("mysql", fmt.Sprint(os.Getenv("MDB_LOGIN"), ":", os.Getenv("MDB_PASSWORD"), "@tcp(", os.Getenv("MDB_HOST"), ":", os.Getenv("MDB_PORT"), ")/charts?parseTime=True"))
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	ci, err := client.ChartItem.Create().SetPrice(result).Save(context.Background())
	if err != nil {
		fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", ci)
}
