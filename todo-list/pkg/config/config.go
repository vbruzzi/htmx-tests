package config

import "github.com/joeshaw/envdecode"

type DbConf struct {
	Name     string `env:"POSTGRES_DB,required"`
	Host     string `env:"POSTGRES_HOST,required"`
	Port     string `env:"POSTGRES_PORT,required"`
	Pw       string `env:"POSTGRES_PW,required"`
	Username string `env:"POSTGRES_USER,required"`
}

type OidcConfig struct {
	Authority    string `env:"OIDC_AUTHORITY,required"`
	ClientId     string `env:"OIDC_CLIENT_ID,required"`
	ClientSecret string `env:"OIDC_CLIENT_SECRET,required"`
	RedirectUrl  string `env:"OIDC_REDIRECT_URL,required"`
}

type Conf struct {
	Db   DbConf
	Oidc OidcConfig
}

func NewConfig() (*Conf, error) {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
