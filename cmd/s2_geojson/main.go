package main

import (
	"fmt"
	"github.com/pantrif/s2_geojson/internal/app/server"
)

const (
	rootPath = "./website"
)

func main() {
	if err := server.Init(rootPath); err != nil {
		fmt.Printf("failed to init: %v", err)
	}
}
