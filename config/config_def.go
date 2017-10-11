package config

import (
	"encoding/xml"
)

//配置文件相关结构定义
type Config struct {
	XMLName    xml.Name `xml:"Config"`
	ServerPort string   `xml:"ServerPort"`
	Corpid     string   `xml:"Corpid"`
	Corpsecret string   `xml:"Corpsecret"`
	Agentid    int64    `xml:"Agentid"`
	Touser     string   `xml:"Touser"`
	Toparty    string   `xml:"Toparty"`
	Totag      string   `xml:"Totag"`
	Safe       int64    `xml:"Safe"`
	Msgtype    string   `xml:"Msgtype"`
}
