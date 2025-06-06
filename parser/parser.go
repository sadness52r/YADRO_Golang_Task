package parser

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"yadro_golang_task/api/handlers"
)

func ParseLine(line string) (*biathlon.Request, error) {
	parts := strings.Fields(line)
	if len(parts) < 3 {
		return nil, fmt.Errorf("parsing err : not enough parameters : %s", line)
	}

	timestamp, err := time.Parse("15:04:05.000", parts[0][1: len(parts[0])-1])
	if err != nil {
		return nil, fmt.Errorf("parsing err : incorrect params : invalid time : %v", err)
	}

	eventID, err := strconv.ParseUint(parts[1], 10, 8)
	if err != nil {
		return nil, fmt.Errorf("parsing err : incorrect eventId : %v", err)
	}

	competitorID, err := strconv.ParseUint(parts[2], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("parsing err : incorrect competitorId : %v", err)
	}

	var extra []any
	for _, p := range parts[3:] {
		extra = append(extra, p)
	}

	return &biathlon.Request{
		Time:         timestamp,
		EventID:      biathlon.EventID(eventID),
		CompetitorID: uint32(competitorID),
		ExtraParams:  extra,
	}, nil
}