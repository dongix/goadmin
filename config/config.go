package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// RedisConfig redis配置
type RedisConfig struct {
	Host     string `json:"redis_host"`
	Port     string `json:"redis_port"`
	Password string `json:"redis_password"`
	DB       int    `json:"redis_db"`
}

// MysqlConfig mysql数据库配置
type MysqlConfig struct {
	Host     string `json:"mysql_host"`
	Port     int    `json:"mysql_port"`
	User     string `json:"mysql_user"`
	Password string `json:"mysql_password"`
	DB       string `json:"mysql_db"`
}

// GetConfig 读取配置文件并转成结构体
func GetConfig(v interface{}) error {
	data, error := ioutil.ReadFile("config.json")
	if error != nil {
		log.Println("load config file fail!")
		return error
	}
	error = json.Unmarshal(data, v)
	if error != nil {
		log.Fatal("parse json data fail!")
		return error
	}
	return nil
}
