package g

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/toolkits/file"
)

type PostgresConfig struct {
	Addr    string `json:"addr"`
	MaxIdle int    `json:"maxIdle"`
}

type GlobalConfig struct {
	Debug      bool            `json:"debug"`
	Hosts      string          `json:"hosts"`
	Postgresql *PostgresConfig `json:"postgresql"`
	Listen     string          `json:"listen"`
	Trustable  []string        `json:"trustable"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	configLock = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}

	configLock.Lock()
	defer configLock.Unlock()

	config = &c

	log.Println("read config file:", cfg, "successfully")
}
