package main

import (
	"fmt"

	"github.com/Kohi-Kamisama/Reader/Config"
	"github.com/spf13/viper"
)
func main() {

	vip := viper.NewWithOptions()
	con, err := config.LoadConfigEnv(vip)
	if err != nil {
		fmt.Println(err)
	}
	if con != nil {
		fmt.Println(con)
	}
	
}