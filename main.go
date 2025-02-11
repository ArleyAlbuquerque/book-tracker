package main

import (
	"github.com/ArleyAlbuquerque/book-tracker/router"
)

func main() {
	// Inicializando as configurações do Gin atraves do subpackage router criado anteriormente
	router.Initialize()

}
