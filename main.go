package main

import (
	"fmt"
	"os"


	"github.com/Kohi-Kamisama/Reader/config"
	"github.com/spf13/viper"
)
func main() {

	vip := viper.NewWithOptions()
	con, err := config.LoadConfigEnv(vip)
	if err != nil {
		fmt.Println(err)
	}
	if con != nil {
		fmt.Println(con.User)
		fmt.Println()
		fmt.Println(con.File)
		fmt.Println()

		rfile, err := os.ReadFile(con.File)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(rfile))
	}
	
}