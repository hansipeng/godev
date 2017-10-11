package config

import (
	"encoding/xml"

	"io/ioutil"

	l4g "github.com/alecthomas/log4go"
	goconfig "github.com/zsounder/zgo/config"
)

var CfgServer *Config = &Config{}

func Init() {
	goconfig.AddLoader("config.xml", loadServerConfig)
}

func loadServerConfig(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		l4g.Error("loadServerConfig read file:%s fail err:%v\n", path, err)
		return err
	}
	CfgServer = &Config{}

	if err := xml.Unmarshal(data, CfgServer); err != nil {
		l4g.Error("loadServerConfig %s Unmarshal fail err:%v\n", path, err)
		return err
	}
	l4g.Info("Server config:%v", *CfgServer)
	return nil
}

func GetServerConfig() *Config {
	goconfig.ReadBegin()
	defer goconfig.ReadEnd()
	return CfgServer
}

// ------------------------------------------
func GetServerPort() string {
	return GetServerConfig().ServerPort
}

func GetCorpid() string {
	return GetServerConfig().Corpid
}

func GetCorpsecret() string {
	return GetServerConfig().Corpsecret
}

func GetAgentid() int64 {
	return GetServerConfig().Agentid
}

func GetSafe() int64 {
	return GetServerConfig().Safe
}

func GetToparty() string {
	return GetServerConfig().Toparty
}

func GetTotag() string {
	return GetServerConfig().Totag
}

func GetTouser() string {
	return GetServerConfig().Touser
}

func GetMsgtype() string {
	return GetServerConfig().Msgtype
}

// ------------------------------------------
