package config

import (

	"os"
	"testing"

)

func TestConfig(test *testing.T) {

	test.Run("Load Config File", func(t *testing.T) {

		err := LoadConfig("../test.toml");
		if err != nil {
			t.Fatal(err);	
		}
		if GlobalConfig.DirName == "" {
			t.Fatal("failed to marshall config, properties are empty");
		}


		t.Cleanup(func() {
			
			os.RemoveAll(GlobalConfig.DirName)
		})
	})
	
}

