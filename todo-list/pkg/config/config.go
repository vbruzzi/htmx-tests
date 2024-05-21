package config

import "github.com/joeshaw/envdecode"

type DbConf struct {
	Name     string `env:"POSTGRES_DB,required"`
	Host     string `env:"POSTGRES_HOST,required"`
	Port     string `env:"POSTGRES_PORT,required"`
	Pw       string `env:"POSTGRES_PW,required"`
	Username string `env:"POSTGRES_USER,required"`
}

type Conf struct {
	Db DbConf
}

func NewConfig() (*Conf, error) {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
