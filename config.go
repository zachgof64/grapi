package main

import (
	"gopkg.in/yaml.v2"
	"github.com/zeuce/golang-api/utils"
)

//Config - Struct for config.yml 
type Config struct {
	Logging bool `yaml:"logging"`
	LogDir string `yaml:"logDir"`
	LogFile string `yaml:"logFile"`
	LogPrefix string `yaml:"logPrefix"`
	Port int `yaml:"port"`
	SSL bool `yaml:"ssl"`
	KeyFile string `yaml:"keyFile"`
	CertFile string `yaml:"certFile"`
}


//GetConfig - Returns Config
func GetConfig() Config {
	config := Config{}
	b := utils.GetByte("config.yml")
	 yaml.Unmarshal(b, &config)
	return config
}