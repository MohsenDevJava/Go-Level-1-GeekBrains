package main

import (
	"flag"
	"fmt"
	"msa/config"
	"os"
)

func main() {
	os.Setenv("SERVER_HOST", "http://google.com")
	os.Setenv("SERVER_PORT", "80")
	os.Setenv("SERVER_PATH", "/index.html")
	os.Setenv("SERVER_TIMEOUT", "0")
	os.Setenv("SERVER_CHECK", "True")

	configPath := flag.String("config", "", "Config file path")
	//host := flag.String("host", "http://msa.com", "Set host")

	flag.Parse()
	cfg := config.GetConfig(*configPath)
	fmt.Println(*cfg)
}
