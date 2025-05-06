package config

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type Config struct {
	Laps        uint32     `json:"laps"`
	LapLen      uint64     `json:"lapLen"`
	PenaltyLen  uint64     `json:"penaltyLen"`
	FiringLines uint32     `json:"firingLines"`
	Start       MyTime     `json:"start"`       
	StartDelta  MyTime     `json:"startDelta"`
}

var AppConfig *Config

type MyTime struct {
    time.Time
}

func (mt *MyTime) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }

    var layouts = []string{
        "15:04:05.000",
        "15:04:05",
    }

    var parsed time.Time
    var err error
    for _, layout := range layouts {
        parsed, err = time.Parse(layout, s)
        if err == nil {
            mt.Time = parsed
            return nil
        }
    }

    return nil
}

func LoadConfig(path string) *Config {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("config err : Failed to open config : %v", err)
	}

	AppConfig = &Config{}
	if err := json.Unmarshal(file, &AppConfig); err != nil {
		log.Fatalf("config err : config.json parsing error : %v", err)
	}

	if AppConfig.Laps == 0 {
		log.Fatalf("config err : laps count is 0")
	}
	if AppConfig.LapLen == 0 {
		log.Fatalf("config err : lap length is 0")
	}
	if AppConfig.PenaltyLen == 0 {
		log.Fatalf("config err : penalty length is 0")
	}
	if AppConfig.FiringLines == 0 {
		log.Fatalf("config err : firing lines count is 0")
	}

	return AppConfig
}
