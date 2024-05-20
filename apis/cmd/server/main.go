package main

import "github.com/becardine/learning-go/apis/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBHost)
}
