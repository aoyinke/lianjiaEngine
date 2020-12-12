package main

import (
	"fmt"
	"gengycSrc/gycSearchEngine/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)


func InitConfig() config.ServerConfig {
	file,err := ioutil.ReadFile("./config.yaml")
	if err != nil{
		fmt.Println("read yaml file failed")
	}

	conf := config.ServerConfig{}
	err = yaml.UnmarshalStrict(file,&conf)
	if err != nil{
		fmt.Println("yaml unmarshal failed")
	}

	return conf
}
