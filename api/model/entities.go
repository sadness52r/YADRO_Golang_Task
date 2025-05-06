package model

import (
	"time"
)

type CompetitorState uint8

const (
	REGISTERED         CompetitorState = iota
	WAIT_START
	ON_START_LINE
	STARTED
	ON_FIRING_RANGE
	LEFT_FIRING_RANGE
	ON_PENALTY_LAP
	LEFT_PENALTY_LAP
	END_MAIN_LAP
	FINISHED
	NOT_FINISHED
	NOT_STARTED
)

func (cs CompetitorState) String() string {
	switch cs {
	case REGISTERED:
		return "Registered"
	case WAIT_START:
		return "WaitStart"
	case ON_START_LINE:
		return "OnStartLine"
	case STARTED:
		return "Started"
	case ON_FIRING_RANGE:
		return "OnFiringRange"
	case LEFT_FIRING_RANGE:
		return "LeftFiringRange"
	case ON_PENALTY_LAP:
		return "OnPenaltyLap"
	case LEFT_PENALTY_LAP:
		return "LeftPenaltyLap"
	case END_MAIN_LAP:
		return "EndMainLap"
	case FINISHED:
		return "Finished"
	case NOT_FINISHED:
		return "NotFinished"
	case NOT_STARTED:
		return "NotStarted"
	default:
		return "Unknown"
	}
}

type LapResult struct {
	Time 	   time.Duration
	AvgSpeed   float32
}

type Competitor struct {
	ID                   uint32
	State                CompetitorState
	PlanStart            time.Time
	RealStart            time.Time
	CurrentLap           uint32
	MainLapsResult       []LapResult
	PenaltyLapsResult    LapResult
	MainLapStartTime     time.Time
	PenaltyLapStartTime  time.Time
	TotalPenaltyTime     time.Duration
	VisitedFiringLines   uint32
	Hits                 uint32
	Shots                uint32
	TotalTime            time.Duration
}