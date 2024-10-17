package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

var GlobalConfiguration *Configuration
var config *Configuration
var RemoteInUse bool = false

type Configuration struct {
	Server struct {
		Host    string `json:"Host"`
		Port    string `json:"Port"`
		Mode    string `json:"Mode"`
		Timeout struct {
			Read  int `json:"Read"`
			Write int `json:"Write"`
		} `json:"Timeout"`
		Logger struct {
			Directory string `json:"Directory"`
			Panic     string `json:"Panic"`
			Server    string `json:"Server"`
		} `json:"Logger"`
		Origin struct {
			Cors []string `json:"Cors"`
		} `json:"Origin"`
	}
	PostgresQL struct {
		Host     string `json:"Host"`
		Port     int64  `json:"Port"`
		User     string `json:"User"`
		Password string `json:"Password"`
		DBName   string `json:"DBName"`
		SSLMode  string `json:"SSLMode"`
	} `json:"PostgresQL"`

	Redis struct {
		Address  string `json:"Address"`
		Password string `json:"Password"`
		Redis_DB string `json:"Redis_DB"`
	}

	//EmailSender struct {
	//	Username string `json:"username"`
	//	From     string `json:"from"`
	//	Password string `json:"password"`
	//	Server   string `json:"server"`
	//	Port     int    `json:"port"`
	//}

	Cache map[string]struct {
		Address  string `json:"Address"`
		Password string `json:"Password"`
		Database int    `json:"Database"`
		Channel  string `json:"Channel"`
	} `json:"Cache"`
}

func Read(f string) *Configuration {
	file, err := os.Open(f)
	if err != nil {
		panic(fmt.Sprintf("Configuration file not exist !!! >>> %v", err.Error()))
	}
	defer func() {
		_ = file.Close()
	}()

	// Load configuration file
	config = &Configuration{}
	// Define JSON file decoder
	d := json.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		panic(fmt.Sprintf("Cant decode configuration file !!! >>> %v", err.Error()))
	}

	// Set default config values
	if config.Server.Timeout.Write < 1 {
		config.Server.Timeout.Write = 15
	}
	if config.Server.Timeout.Read < 1 {
		config.Server.Timeout.Read = 15
	}
	if len(config.Server.Logger.Panic) == 0 {
		config.Server.Logger.Panic = "/var/log/panic.log"
	}

	// Print configuration
	fmt.Printf("Configuration: Server: %+v\n", config.Server)

	return config
}

// GetConfig - get configuration file
func GetConfig() *Configuration {

	config = Read(os.Args[1])
	GlobalConfiguration = config

	return config
}
