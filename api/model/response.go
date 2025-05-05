package model

import (
	"time"
)

type MainLapInfo struct {
	Time 	   time.Time
	AvgSpeed   float32
}
type PenaltyLapInfo MainLapInfo

type LogResponse struct {
	Time         time.Time
	Info		 string
}

type TableResponse struct {
	TotalTime       string
	CompetitorID    uint32
	MainLapsInfo    []MainLapInfo
	PenaltyLapsInfo PenaltyLapInfo
	Hits  	        uint32
	Shots		    uint32
}