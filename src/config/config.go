package config

import (
	"fmt"
	"log"

	gcfg "gopkg.in/gcfg.v1"
)

type (
	Config struct {
		Server struct {
			Port string
		}
		Database struct {
			DBMaster string
			DBSlave  string
		}
	}
)

// FilePathList list of possible config file relative path to binary location
func FilePathList() []string {
	return []string{
		"files/etc/testingwebapp/",
		"../etc/testingwebapp/",
		"../../etc/testingwebapp/",
	}
}

// ReadConfig read `*.ini` configuration file and save it to variable of `*Config` type
func ReadConfig() *Config {
	var (
		cfg     Config
		err     error
		environ string
	)

	path := FilePathList()

	for _, val := range path {
		file := fmt.Sprintf("%sconfig.ini", val)
		log.Printf("%s\n", file)
		err := gcfg.ReadFileInto(&cfg, file)
		if err == nil {
			break
		}
	}

	if err != nil {
		log.Fatalf("[testingwebapp] Cannot load config env:%s :%+v\n ", environ, err)
	}

	log.Printf("[testingwebapp] Config load success, using \"%s\".\n", environ)

	return &cfg
}
