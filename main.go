package main

import (
	"github.com/ArleyAlbuquerque/book-tracker/config"
	"github.com/ArleyAlbuquerque/book-tracker/router"
)

var (
	logger *config.Logger
)

func main() {
	// Inicializando as configurações do Gin atraves do subpackage router criado anteriormente
	logger = config.GetLogger("main")
	router.Initialize()
	err := config.Init()
	if err != nil {
		logger.ErrF("CONFIG INITIALIZATION ERROR: %v", err)
		return
	}
}
