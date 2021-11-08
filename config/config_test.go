package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/pelletier/go-toml"
)

func TestConfig(test *testing.T) {

	test.Run("Load Config File", func(t *testing.T) {

		data, err := os.ReadFile("../test.toml");
		if err != nil {
			t.Fatal(err);
		}

		fmt.Println(string(data))

	
		config := Config{}
		toml.Unmarshal(data, &config);
		if err != nil {
			t.Fatal(err);	
		}
		if config.DirName == "" {
			t.Fatal("failed to marshall config, properties are empty");
		}


		t.Cleanup(func() {
			
		})
	})
}