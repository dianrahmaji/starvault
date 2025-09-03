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
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	err = scheduler.Start(cfg)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}
