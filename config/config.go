package config

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type Config struct {
	Laps        uint32     `json:"laps"`
	LapLen      float32    `json:"lapLen"`
	PenaltyLen  float32    `json:"penaltyLen"`
	FiringLines uint32     `json:"firingLines"`
	Start       time.Time  `json:"start"`       
	StartDelta  time.Time  `json:"startDelta"`
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
