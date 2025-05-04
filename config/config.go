package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Laps        int    `json:"laps"`
	LapLen      int    `json:"lapLen"`
	PenaltyLen  int    `json:"penaltyLen"`
	FiringLines int    `json:"firingLines"`
	Start       string `json:"start"`       
	StartDelta  string `json:"startDelta"`
}

var AppConfig *Config

func LoadConfig(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("config err : Failed to open config : %v", err)
	}
	defer file.Close()

	AppConfig = &Config{}
	if err := json.NewDecoder(file).Decode(AppConfig); err != nil {
		log.Fatalf("config err : config.json parsing error : %v", err)
	}
}
