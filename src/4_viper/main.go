package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config") // call multiple times to add many search paths
	viper.AddConfigPath(".")        // optionally look for config in the working directory
	viper.AddConfigPath("C:/Project/Golang/EDD/src/4_viper")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	} else {
		fmt.Println("Name of Candidate: ", viper.GetString("name"))
	}
}
