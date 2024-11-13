package main

import (
	"fmt"

	"github.com/rx3lixir/urlshortener/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
