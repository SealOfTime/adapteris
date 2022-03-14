package app

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

const (
	DefaultHostURL    = "http://localhost:8080"
	DefaultHttpPort   = 8080
	DefaultConfigPath = "./cfg.json"
)

type Config struct {
	ConfigPath string         `json:"-"`
	HostURL    string         `json:"hostUrl"`
	Port       uint           `json:"port"`
	Vk         OAuth2Provider `json:"vk"`
	DbUri      string         `json:"db"`
	JWTKey     string         `json:"jwtKey"`
}

type OAuth2Provider struct {
	ClientId string `json:"clientId"`
	Secret   string `json:"secret"`
}

func NewConfig() Config {
	return Config{
		ConfigPath: DefaultConfigPath,
		HostURL:    DefaultHostURL,
		Port:       DefaultHttpPort,
	}
}

//Setup fills a Config with data from application flags and JSON in this order
func (cfg *Config) Setup() {
	cfg.parseFlags()
	cfg.parseJson(loadJsonConfig(cfg.ConfigPath))
}

func loadJsonConfig(path string) []byte {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("error loading config file '%s': %+v", path, err)
	}

	return c
}

func (cfg *Config) parseFlags() {
	flag.Parse()
	flag.StringVar(&cfg.ConfigPath, "cfg", cfg.ConfigPath, "-cfg for setting up config path")

	flag.StringVar(&cfg.HostURL, "host", cfg.HostURL, "-host to specify url that this app is hosted on(for redirect purposes)")
	flag.UintVar(&cfg.Port, "port", cfg.Port, "-port to specify port to start the app on")

	flag.StringVar(&cfg.Vk.ClientId, "vkClientId", cfg.Vk.ClientId, "-vkClientId for OAuth2")
	flag.StringVar(&cfg.Vk.Secret, "vkSecret", cfg.Vk.Secret, "-vkSecret for OAuth2")

	flag.StringVar(&cfg.DbUri, "db", cfg.DbUri, "-db to provide connection url")

	flag.StringVar(&cfg.JWTKey, "jwtKey", cfg.JWTKey, "jwtKey for JWT signing")
}

func (cfg *Config) parseJson(rawCfg []byte) {
	err := json.Unmarshal(rawCfg, cfg)
	if err != nil {
		log.Fatalf("error parsing config: %+v", err)
	}
}
