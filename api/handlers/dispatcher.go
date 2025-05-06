package biathlon

import (
	"yadro_golang_task/api"
	"log"
)

func Dispatch(req *Request, h *BiathlonHandler) {
	var eventHandlers = map[EventID]func(*Request) customresponse.CustomResponse {
		EVENT_REGISTER:            h.HandleRegister,
		EVENT_CHOOSE_START_TIME:   h.HandleChooseStartTime,
		EVENT_READY:               h.HandleReady,
		EVENT_START_MAIN:          h.HandleStartMain,
		EVENT_ARRIVE_FIRING_RANGE: h.HandleArriveFiringRange,
		EVENT_HIT:                 h.HandleHit,
		EVENT_LEFT_FIRING_RANGE:   h.HandleLeftFiringRange,
		EVENT_START_PENALTY:       h.HandleStartPenalty,
		EVENT_END_PENALTY:         h.HandleEndPenalty,
		EVENT_END_MAIN:            h.HandleEndMain,
		EVENT_DISQUALIFIED:        h.HandleDisqualified,
	}

    if handler, ok := eventHandlers[req.EventID]; ok {
        response := handler(req)
		if response.Err != nil {
			log.Fatal(response.DeveloperMessage)
		} else {
			log.Println(response.UserMessage)
		}
    } else {
        log.Fatal("unknown event ID : ", req.EventID)
    }
}