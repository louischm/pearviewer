package conf

import (
	"encoding/json"
	"os"
	"strconv"
)

var confSingleton *Conf = nil

type Conf struct {
	ServerPort    int32
	ServerAddress string
}

func NewConf() *Conf {
	if confSingleton != nil {
		return confSingleton
	}
	confSingleton = getConf()
	return confSingleton
}

func getConf() *Conf {
	println("Reading conf")
	conf, err := os.ReadFile("./conf/conf.json")

	if err != nil {
		println("Error on Reading conf file: " + err.Error())
	}

	var confData Conf
	println("Parsing json conf")
	err = json.Unmarshal(conf, &confData)
	if err != nil {
		println("Error on Unmarshal conf: " + err.Error())
	}
	println("Conf loaded")
	return &confData
}

func (c *Conf) GetServerAddress() string {
	return c.ServerAddress + ":" + strconv.Itoa(int(c.ServerPort))
}
