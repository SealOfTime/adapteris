package setup

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

type config struct {
	ConfigPath string `json:"-"`
	HostURL    string `json:"hostUrl"`
	Port       int    `json:"port"`
	Vk         struct {
		ClientId string `json:"clientId"`
		Secret   string `json:"secret"`
	}
	JWTKey string `json:"jwtKey"`
}

func (cfg *config) parseFlags() {
	flag.Parse()
	flag.StringVar(&cfg.ConfigPath, "cfg", "./cfg.json", "-cfg for setting up config path")
	flag.StringVar(&cfg.HostURL, "host", "", "-host to specify url that this app is hosted on(for redirect purposes)")
	flag.IntVar(&cfg.Port, "port", -1, "-port to specify port to start the app on")
	flag.StringVar(&cfg.Vk.ClientId, "vkClientId", "", "-vkClientId for OAuth2")
	flag.StringVar(&cfg.Vk.Secret, "vkSecret", "", "-vkSecret for OAuth2")
	flag.StringVar(&cfg.JWTKey, "jwtKey", "", "jwtKey for JWT signing")
}

func loadJsonConfig(path string) []byte {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("error loading config file '%s': %+v", path, err)
	}

	return c
}

func (cfg *config) parseJson(rawCfg []byte) {
	err := json.Unmarshal(rawCfg, cfg)
	if err != nil {
		log.Fatalf("error parsing config: %+v", err)
	}
}
