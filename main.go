package main

import (
	"github.com/scripted-wings/todo-server/config"
	"github.com/scripted-wings/todo-server/server"
)

func main() {
	config.LoadConfig()
	config.InitLogger()
	server.StartServer()
}
