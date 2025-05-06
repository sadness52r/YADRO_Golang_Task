package model

type TableResponse struct {
	TotalTime          string
	CompetitorID       uint32
	MainLapsInfo       []LapResult
	PenaltyLapsInfo    LapResult
	Hits  	           uint32
	Shots		       uint32
	VisitedPenaltyLaps uint32
}