package config

import (
	"fmt"
	"os"
	"rustdesk-api-server/constant"
	"rustdesk-api-server/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Load the configuration item
func init() {
	fmt.Println("Load the configuration item")
	var config string
	if configEnv := os.Getenv(constant.ConfigEnv); configEnv == "" {
		config = constant.ConfigFile
	} else {
		config = configEnv
	}

	// Determine whether there is a configuration file
	_, err := os.Stat(config)
	if err != nil && os.IsNotExist(err) {
		// The configuration file does not exist
		err := os.WriteFile(config, []byte(`dbtype: 'sqlite3'
mysql:
  host: '127.0.0.1'
  port: 3306
  database: 'rustdesk'
  username: 'root'
  password: ''
app:
  authkey: 123456
  cryptkey: NanEVhjEwuPSemoJkwcYEcjDJRVWcZfb9bIIZcBeswhPP`), 0777)
		if err != nil {
			panic(err)
		}
	}

	v := viper.New()

	// Set up a profile
	v.SetConfigFile(config)

	// Read the configuration
	err = v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to load configuration file: %s", err))
	}

	// Monitor configuration updates
	v.WatchConfig()

	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("The configuration file is modified:", in.Name)
		if err := v.Unmarshal(&global.ConfigVar); err != nil {
			panic(err)
		}
	})

	if err := v.Unmarshal(&global.ConfigVar); err != nil {
		panic(err)
	}
	fmt.Println("Loading the configuration item is complete")

}
