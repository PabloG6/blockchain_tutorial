package config

import (
	"fmt"
	"log"
	"os"

	"github.com/pelletier/go-toml"
)
var GlobalConfig Config;
var Environment string;
func init() {
	fmt.Printf("environment %s", Environment);
	
	LoadConfig("../config.toml");
}
type Config struct {
	DirName string
	TokenSupply int64;
}

func LoadConfig(configFile string) error {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return err;
	}


	 return  toml.Unmarshal(data, &GlobalConfig)

	
}
	

func CreateConfig(configFile string) (config Config) {

	data, err := os.ReadFile(configFile);
	if err != nil {
		log.Fatal(err);
	}

	toml.Unmarshal(data, config);
	return config;
}

	
