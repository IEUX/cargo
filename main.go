package main

import "C"
import (
	logger "cargo/internal"
	"cargo/internal/views"
)

func init() {
	logger.Info.Println("Application launched")
}

func main() {
	//internal.Ps()
	views.Root()
}
