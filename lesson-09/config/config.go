package config

import (
	"encoding/json"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path"
	"strings"
)

type Config struct {
	Server struct {
		Host    string  `envconfig:"SERVER_HOST" yaml:"host" json:"host"`
		Port    int     `envconfig:"SERVER_PORT" yaml:"port" json:"port"`
		Path    string  `envconfig:"SERVER_PATH" yaml:"path" json:"path"`
		Timeout float32 `envconfig:"SERVER_TIMEOUT" yaml:"timeout" json:"timeout"`
		Check   bool    `envconfig:"SERVER_CHECK" yaml:"check" json:"check"`
	} `yaml:"server" json:"server"`
}

func processError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func readFile(Msa *Config, configPath string) {
	if configPath == "" {
		return
	}
	b, err := ioutil.ReadFile(configPath)
	processError(err)
	configType := strings.ToLower(path.Ext(configPath))[1:]
	switch configType {
	case "yaml", "yml":
		err = yaml.Unmarshal(b, &Msa)
	case "json", "jsn":
		err = json.Unmarshal(b, &Msa)
	default:
		log.Printf("Unexpected config file type %s", configPath)
	}
	processError(err)
}

func readEnv(ms *Config) {
	err := envconfig.Process("", ms)
	processError(err)
}

func GetConfig(configPath string) *Config {
	var ms Config // set defaults
	readFile(&ms, configPath)
	readEnv(&ms)
	//readCli(&ms)
	return &ms
}
