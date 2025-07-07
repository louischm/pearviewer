package conf

import (
	"encoding/json"
	"github.com/louischm/pkg/logger"
	"os"
	"strconv"
)

var confSingleton *Conf = nil
var log = logger.NewLog()

type Conf struct {
	ServerPort    int32
	ServerAddress string
	ServerCert    string
	ServerKey     string
	CaCertPath    string
	LogPath       string
	DataPath      string
}

func NewConf() *Conf {
	if confSingleton != nil {
		return confSingleton
	}
	confSingleton = getConf()
	return confSingleton
}

func getConf() *Conf {
	conf, err := os.ReadFile("./conf/conf.json")

	if err != nil {
		panic(err)
	}

	var confData Conf
	err = json.Unmarshal(conf, &confData)
	if err != nil {
		panic(err)
	}

	return &confData
}

func (c *Conf) GetServerAddress() string {
	return c.ServerAddress + ":" + strconv.Itoa(int(c.ServerPort))
}

func (c *Conf) GetServerCert() string { return c.ServerCert }

func (c *Conf) GetServerKey() string { return c.ServerKey }
