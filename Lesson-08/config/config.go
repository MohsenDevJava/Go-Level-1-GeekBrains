package config

import (
	"fmt"
	"os"
	"strconv"
)

type Configuration struct {
	DbUrl      string
	DbPort     int64
	DbUser     string
	DbPassword string
}

func ReadConfig() *Configuration {
	var envUrl string
	var envPort int64
	var envUser string
	var envPassword string

	envUrl = os.Getenv("DB_URL")
	envPort, _ = strconv.ParseInt(os.Getenv("DB_PORT"), 0, 64)
	if 1023 <= envPort && envPort <= 65536 {
		fmt.Println("Port entered correctly!")
	} else {
		fmt.Println("Port entered incorrectly!")
	}
	envUser = os.Getenv("DB_USER")
	if envUser == "" {
		fmt.Println("Enter username!")
	}
	fmt.Println("Username read successfully!")

	envPassword = os.Getenv("DB_PASSWORD")
	if envPassword == "" {
		fmt.Println("Enter password!")
	}
	fmt.Println("Password read successfully!")

	return &Configuration{
		DbUrl:      envUrl,
		DbPort:     envPort,
		DbUser:     envUser,
		DbPassword: envPassword,
	}
}
