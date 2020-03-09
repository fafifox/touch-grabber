package main

import (
	"fmt"
	"github.com/farooch/touch-grabber/configs"
	"github.com/farooch/touch-grabber/pkg/data"
	"github.com/farooch/touch-grabber/pkg/versions"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

func main() {
	var cfg configs.Config

	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "CONFIG")
		return
	}
	//Get config.yml if set
	arg := os.Args[1]
	if arg != "" {
		err := cleanenv.ReadConfig(arg, &cfg)
		if err != nil {
			log.Println(err)
			os.Exit(2)
		}
	} else {
		err := cleanenv.ReadConfig("configs/config.yml", &cfg)
		if err != nil {
			log.Println(err)
			os.Exit(2)
		}
	}
	versions.GetVersions(cfg.SavePath)
	data.FetchAllDataMap(cfg.DataUrl, cfg.TargetBuildVersion, cfg.SavePath)
	data.FetchAllDataDictionary(cfg.DataUrl, cfg.TargetBuildVersion, cfg.SavePath)
}
