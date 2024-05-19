package main

import (
	"gin/cmd/server"
)

func main() {
	type todo struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Status bool   `json:"status"`
	}
	server.Init()
	server.Run()
}
