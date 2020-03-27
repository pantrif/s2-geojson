package main

import (
	"fmt"
	"s2-geojson/internal/app/server"
)

func main() {
	if err := server.Init(); err != nil {
		fmt.Printf("failed to init: %v", err)
	}
}
