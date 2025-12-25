package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type configEnv struct {
	User string `mapstructure:"DEV_USER"`
}

func LoadConfigEnv(vip *viper.Viper) (*configEnv, error) {
	con := configEnv{}

	vip.BindEnv("dev_user")
	vip.AutomaticEnv()
	
	err := vip.Unmarshal(&con)
	user := con.User

	if err != nil {
		return nil, err
	}

	if len(user) == 0 {
		return nil, fmt.Errorf("User Not Found")
	}

	return &con, nil
}