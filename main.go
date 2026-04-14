package main

import (
	"fmt"

	"github.com/DegsRed72/gator/internal/config"
)

func main() {
	cfg := config.Read()
	cfg.SetUser("Evan")
	cfg = config.Read()

	fmt.Println(cfg)
}
