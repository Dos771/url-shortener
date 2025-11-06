package main

import (
	"url-shortener/internal/config"
	"fmt"
)
func main() {
	// TODO: init config: cleanenv
	cfg := config.MustLoad()

	fmt.Println(cfg)
	// TODO: init logger: slog

	// TODO: init storage: sqlite

	// TODO: init router: chi, "chi render" 

	// TODO: init server
}