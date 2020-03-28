package main

import (
	"fmt"
	"s2-geojson/internal/app/server"
)

const (
	rootPath = "./website"
)

func main() {
	if err := server.Init(rootPath); err != nil {
		fmt.Printf("failed to init: %v", err)
	}
}
