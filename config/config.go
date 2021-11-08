package config

import (
	"log"
	"os"

	"github.com/pelletier/go-toml"
)



type Config struct {
	FileName string
}

func LoadConfig(configFile string) (config Config) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}


	 toml.Unmarshal(data, config)

	 return config;

	


	

}