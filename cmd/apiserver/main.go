package main

import (
	"log"

	"github.com/zhashkevych/todo-app"
)


func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error while running http server ")
	}

}