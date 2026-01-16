package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type configEnv struct {
	User string `mapstructure:"DEV_USER"`
	File string `mapstructure:"FILE"`
}

func LoadConfigEnv(vip *viper.Viper) (*configEnv, error) {
	con := configEnv{}

	vip.BindEnv("dev_user")
	vip.BindEnv("file")
	vip.AutomaticEnv()
	
	err := vip.Unmarshal(&con)
	user := con.User
	file := con.File

	if err != nil {
		return nil, err
	}

	if len(user) == 0 {
		return nil, fmt.Errorf("User Not Found")
	}
	if len(file) == 0 {
		return nil, fmt.Errorf("File Not Found")
	}

	return &con, nil
}