package main

import (
	"fmt"
	"github.com/joho/godotenv" //using the "godotenv" library
	"md/config"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Configuration file not found!")
	}
	fmt.Println("Configuration file uploaded successfully!")
}

// In the main, we call the imported ReadConfig function from the config package
func main() {
	fmt.Println("Reading data from virtual environment: ")
	configurations := config.ReadConfig()
	fmt.Printf("DB_URL: %s\nDB_PORT: %d\nDB_USER: %s\nDB_PASSWORD: %s\n",
		configurations.DbUrl,
		configurations.DbPort,
		configurations.DbUser,
		configurations.DbPassword)

}
