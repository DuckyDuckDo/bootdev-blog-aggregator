package main

import (
	"fmt"
	"log"

	"github.com/DuckyDuckDo/bootdev-blog-aggregator/internal/config"
)

const configFileName = ".gatorconfig.json"

func main() {
	cfg, err := config.Read(configFileName)
	if err != nil {
		log.Fatalf("%v", err)
	}

	cfg.SetUser("DuckyDo", configFileName)
	cfg, err = config.Read(configFileName)

	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("%s\n%s\n\n", cfg.DbURL, cfg.CurrentUserName)
}
