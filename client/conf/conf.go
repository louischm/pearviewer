package conf

import (
	"encoding/json"
	"github.com/louischm/logger"
	"os"
	"strconv"
)

var log = logger.NewLog()

type Conf struct {
	ServerPort    int32
	ServerAddress string
}

func GetConf() *Conf {
	log.Info("Reading conf")
	conf, err := os.ReadFile("./conf/conf.json")

	if err != nil {
		log.Fatal("Error on Reading conf file: " + err.Error())
	}

	var confData Conf
	log.Info("Parsing json conf")
	err = json.Unmarshal(conf, &confData)
	if err != nil {
		log.Fatal("Error on Unmarshal conf: " + err.Error())
	}

	return &confData
}

func (c *Conf) GetServerAddress() string {
	return c.ServerAddress + ":" + strconv.Itoa(int(c.ServerPort))
}
