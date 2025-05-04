package model

import "time"

type EventID uint8

const (
	EVENT_REGISTER              EventID = 1
	EVENT_CHOOSE_START_TIME     EventID = 2
	EVENT_READY                 EventID = 3
	EVENT_START_MAIN            EventID = 4
	EVENT_ARRIVE_FIRING_RANGE   EventID = 5
	EVENT_HIT                   EventID = 6
	EVENT_LEFT_FIRING_RANGE     EventID = 7
	EVENT_START_PENALTY         EventID = 8
	EVENT_END_PENALTY           EventID = 9
	EVENT_END_MAIN              EventID = 10
	EVENT_DISQUALIFIED          EventID = 11
)

type Request struct {
	Time         time.Time
	EventID      EventID
	CompetitorID uint32
	ExtraParams  []any
}
