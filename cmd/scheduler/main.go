package main

import (
	"fmt"
	"os"

	"github.com/dianrahmaji/star-vault/config"
	"github.com/dianrahmaji/star-vault/internal/delivery/scheduler"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		fmt.Println("%s", err)
		os.Exit(1)
	}

	scheduler.Start(cfg.IDNApiUrl)
}
