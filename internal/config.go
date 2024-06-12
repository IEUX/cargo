package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Config struct {
	ConfigName string `json:"ConfigName"`
	App        struct {
		Version string `json:"Version"`
	}
	Logger struct {
		Dir  string `json:"Dir"`
		File string `json:"File"`
	}
}

var C Config
var confFile = "configs/config.json"

func init() {
	f, err := os.Open(confFile)
	if err != nil {
		log.Fatalf(ErrorColor, fmt.Sprintf("[CONFIG INIT FAILED]: %s", err))
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf(ErrorColor, fmt.Sprintf("[CONFIG INIT FAILED]: %s", err))
		}
	}(f)
	byteValue, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf(ErrorColor, fmt.Sprintf("[CONFIG INIT FAILED]: %s", err))
	}
	err = json.Unmarshal(byteValue, &C)
	if err != nil {
		log.Fatalf(ErrorColor, fmt.Sprintf("[CONFIG INIT FAILED]: %s", err))
	}
}
