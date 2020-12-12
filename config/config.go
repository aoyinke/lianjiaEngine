package config

const BaseUrl = "http://localhost:8080/api/houseInfo/"
type ServerConfig struct {
	Elastic Elastic `json:"elastic" yaml:"elastic"`
}

type Elastic struct {
	Host string `json:"host" yaml:"host"`
	Port int `json:"port" yaml:"port"`
}