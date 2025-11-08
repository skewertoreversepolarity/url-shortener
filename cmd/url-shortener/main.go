package main

import "github.com/skewertoreversepolarity/url-shortener.git/cmd/internal/config"

func main() {
	cfg := config.MustLoad()

	//TODO: init config: cleanenv

	//TODO: init logger: slog

	//TODO: init storage: SQL lite

	//TODO: init router: chi, "chi render"

	//TODO: run server
}
