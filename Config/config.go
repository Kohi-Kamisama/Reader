package config

import "github.com/spf13/viper"

type configEnv struct {
	User string `mapstructure:"User"`
}

func LoadConfigEnv(vip *viper.Viper) (*configEnv, error) {
	con := configEnv{}

	vip.BindEnv("user")
	vip.AutomaticEnv()

	err := vip.Unmarshal(&con)
	if err != nil {
		return nil, err
	}

	return &con, nil
}