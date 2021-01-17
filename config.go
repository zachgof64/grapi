package main

import (
	"gopkg.in/yaml.v2"
	"fmt"
	"os"
	"io/ioutil"
)

//Config ... 
type Config struct {
	Logging bool `yaml:"logging"`
	LogFile string `yaml:"logFile"`
	LogPrefix string `yaml:"logPrefix"`
	Port int `yaml:"port"`
	SSL bool `yaml:"ssl"`
	KeyFile string `yaml:"keyFile"`
	CertFile string `yaml:"certFile"`
}


//GetConfig ... Returns Config Struct
func GetConfig() Config {
	config := Config{}
	b := OpenConfigFile("config.yml")
	 yaml.Unmarshal(b, &config)
	return config
}

//OpenConfigFile ... Returns []byte of config
func OpenConfigFile(configPath string) []byte {
	file, fileErr := os.Open(configPath)
	if (fileErr != nil){
		fmt.Printf("ERROR:\nCODE: FILEERR\nMESSAGE: File was not found\n StackTrace: %s", fileErr)
	}
	b, bErr := ioutil.ReadAll(file)
	if(bErr != nil){
		fmt.Printf("ERROR:\nCODE: BYTEERR\nMESSAGE: Couldn't Read File\n StackTrace: %s", bErr)
	}
	file.Close()
	return b
}