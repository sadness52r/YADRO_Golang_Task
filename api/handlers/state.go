package biathlon

import "time"

type EventID uint8

const (
	EVENT_REGISTER              EventID = iota + 1
	EVENT_CHOOSE_START_TIME    
	EVENT_READY                
	EVENT_START_MAIN           
	EVENT_ARRIVE_FIRING_RANGE  
	EVENT_HIT                  
	EVENT_LEFT_FIRING_RANGE    
	EVENT_START_PENALTY        
	EVENT_END_PENALTY          
	EVENT_END_MAIN              
	EVENT_DISQUALIFIED       
)

type Request struct {
	Time         time.Time
	EventID      EventID
	CompetitorID uint32
	ExtraParams  []any
}