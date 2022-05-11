package main

import (
	"log"

	"github.com/huangchao308/go-tools/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal("execute cmd err: ", err.Error())
	}
}
