package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	AppMode       string
	HttpPort      string
	TransferTable bool
	UseMultipoint bool
	LogPath       string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	RedisHost     string
	RedisPort     string
	RedisPassword string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Printf("配置文件路径有误", err)
	}
	LoadServer(file)
	LoadData(file)
	LoadRedis(file)

}

func LoadServer(file *ini.File) {
	// Section指定分区名，Key方法指定字段名，MustString允许给个默认值，String无默认值，不存在则为空
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8020")
	UseMultipoint = file.Section("server").Key("UseMultipoint").MustBool(false)
	TransferTable = file.Section("server").Key("TransferTable").MustBool(false)
	LogPath = file.Section("server").Key("LogPath").MustString("log/")
	//判断字符串是否以指定后缀结尾
	if !strings.HasSuffix(LogPath, "/") {
		LogPath += "/"
	}
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("debug")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("")
	DbName = file.Section("database").Key("DbName").MustString("goweb")
}

func LoadRedis(file *ini.File) {
	RedisHost = file.Section("redis").Key("RedisHost").MustString("localhost")
	RedisPort = file.Section("redis").Key("RedisPort").MustString("6379")
	RedisPassword = file.Section("redis").Key("RedisPassword").MustString("")
}
