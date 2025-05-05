package biathlon

import (
	//"errors"
	"fmt"
	"time"
	"strings"
	"log"
	"strconv"
	//biathlon "yadro_golang_task/api"
	"yadro_golang_task/api"
	//"yadro_golang_task/state"
)

const (
	LAYOUT_TIME string = "15:04:05.000"
)

type BiathlonHandler struct {
	//manager *vrrp.VRRPManager
}

func NewBiathlonHandler(/*manager *vrrp.VRRPManager*/) *BiathlonHandler {
	return &BiathlonHandler{
		//manager: manager,
	}
}


func (h *BiathlonHandler) HandleRegister(req *Request /*manager *state.StateManager*/) customresponse.CustomResponse {
	var response customresponse.CustomResponse

	if len(req.ExtraParams) > 0 {
		errorMsgUsr := "Error : register competitor : <Failed to decode request>"
		errorMsgDev := "Error : register competitor : <Too many extra parameters>"
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              fmt.Errorf("invalid number of parameters for event %d", req.EventID),
		}

		return response
	}

	//TODO: Call manager

	infoMsg := fmt.Sprintf("[%s] The competitor(%d) registered", req.Time.Format(LAYOUT_TIME), req.CompetitorID)

	response = customresponse.CustomResponse{
		UserMessage:      infoMsg,
		DeveloperMessage: infoMsg,
		Err:              nil,
	}

	return response
}

func (h *BiathlonHandler) HandleChooseStartTime(req *Request /*manager *state.StateManager*/) customresponse.CustomResponse {
	var response customresponse.CustomResponse

	if len(req.ExtraParams) != 1 {
		errorMsgUsr := "Error : choose start time : <Failed to decode request>"
		errorMsgDev := "Error : choose start time : <Invalid number of extra params>"
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              fmt.Errorf("invalid number of parameters for event %d", req.EventID),
		}

		return response
	}

	startTimePlan, ok := req.ExtraParams[0].(string)
	if !ok {
		errorMsgUsr := "Error : choose start time : <Failed to decode request>"
		errorMsgDev := "Error : choose start time : <Invalid type of extra params>"
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              fmt.Errorf("invalid time type: %v", req.ExtraParams[0]),
		}

		return response
	}

	timestamp, err := time.Parse(LAYOUT_TIME, startTimePlan)
	if err != nil {
		errorMsgUsr := "Error : choose start time : <Failed to decode request>"
		errorMsgDev := "Error : choose start time : <Unable to parse extra params into time>"
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              err,
		}

		return response
	}

	//TODO: Call manager

	infoMsg := fmt.Sprintf(
		"[%s] The start time for the competitor(%d) was set by a draw to %s", req.Time.Format(LAYOUT_TIME), req.CompetitorID, timestamp.Format(LAYOUT_TIME),
	)

	response = customresponse.CustomResponse{
		UserMessage:      infoMsg,
		DeveloperMessage: infoMsg,
		Err:              nil,
	}

	return response
}

func (h *BiathlonHandler) HandleReady(req *Request /*manager *state.StateManager*/) customresponse.CustomResponse {
	var response customresponse.CustomResponse

	if len(req.ExtraParams) > 0 {
		errorMsgUsr := "Error : ready to start : <Failed to decode request>"
		errorMsgDev := "Error : ready to start : <Too many extra parameters>"
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              fmt.Errorf("invalid number of parameters for event %d", req.EventID),
		}

		return response 
	}

	//TODO: Call manager

	infoMsg := fmt.Sprintf("[%s] The competitor(%d) is on the start line", req.Time.Format(LAYOUT_TIME), req.CompetitorID)

	response = customresponse.CustomResponse{
		UserMessage:      infoMsg,
		DeveloperMessage: infoMsg,
		Err:              nil,
	}

	return response
}

func (h *BiathlonHandler) HandleStartMain(req *Request /*manager *state.StateManager*/) customresponse.CustomResponse {
	var response customresponse.CustomResponse

	if len(req.ExtraParams) > 0 {
		errorMsgUsr := "Error : start main lap : <Failed to decode request>"
		errorMsgDev := "Error : start main lap : <Too many extra parameters>"
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              fmt.Errorf("invalid number of parameters for event %d", req.EventID),
		}

		return response 
	}

	//TODO: Call manager

	infoMsg := fmt.Sprintf("[%s] The competitor(%d) has started", req.Time.Format(LAYOUT_TIME), req.CompetitorID)

	response = customresponse.CustomResponse{
		UserMessage:      infoMsg,
		DeveloperMessage: infoMsg,
		Err:              nil,
	}

	return response
}

func (h *BiathlonHandler) HandleArriveFiringRange(req *Request /*manager *state.StateManager*/) customresponse.CustomResponse {
	var response customresponse.CustomResponse

	if len(req.ExtraParams) != 1 {
		errorMsgUsr := "Error : arrive firing range : <Failed to decode request>"
		errorMsgDev := "Error : arrive firing range : <Invalid number of extra params>"
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              fmt.Errorf("invalid number of parameters for event %d", req.EventID),
		}

		return response
	}

	strFiringRange := req.ExtraParams[0].(string)
	firingRange64, err := strconv.ParseUint(strFiringRange, 10, 32)
	if err != nil {
		errorMsgUsr := "Error : arrive firing range : <Failed to decode request>"
		errorMsgDev := "Error : arrive firing range : <Unable to parse extra params into uint32>"
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              err,
		}

		return response
	}
	firingRange := uint32(firingRange64)

	//TODO: Call manager

	infoMsg := fmt.Sprintf("[%s] The competitor(%d) is on the firing range(%d)", req.Time.Format(LAYOUT_TIME), req.CompetitorID, firingRange)

	response = customresponse.CustomResponse{
		UserMessage:      infoMsg,
		DeveloperMessage: infoMsg,
		Err:              nil,
	}

	return response
}

func (h *BiathlonHandler) HandleHit(req *Request /*manager *state.StateManager*/) customresponse.CustomResponse {
	var response customresponse.CustomResponse

	if len(req.ExtraParams) != 1 {
		errorMsgUsr := "Error : hit target : <Failed to decode request>"
		errorMsgDev := "Error : hit target : <Invalid number of extra params>"
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              fmt.Errorf("invalid number of parameters for event %d", req.EventID),
		}

		return response
	}

	strTarget := req.ExtraParams[0].(string)
	target64, err := strconv.ParseUint(strTarget, 10, 32)
	if err != nil {
		errorMsgUsr := "Error : hit target : <Failed to decode request>"
		errorMsgDev := "Error : hit target : <Unable to parse extra params into uint32>"
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              err,
		}

		return response
	}
	target := uint32(target64)

	//TODO: Call manager

	infoMsg := fmt.Sprintf("[%s] The target(%d) has been hit by competitor(%d)", req.Time.Format(LAYOUT_TIME), target, req.CompetitorID)

	response = customresponse.CustomResponse{
		UserMessage:      infoMsg,
		DeveloperMessage: infoMsg,
		Err:              nil,
	}

	return response
}

func (h *BiathlonHandler) HandleLeftFiringRange(req *Request /*manager *state.StateManager*/) customresponse.CustomResponse {
	var response customresponse.CustomResponse

	if len(req.ExtraParams) > 0 {
		errorMsgUsr := "Error : left firing range : <Failed to decode request>"
		errorMsgDev := "Error : left firing range : <Too many extra parameters>"
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              fmt.Errorf("invalid number of parameters for event %d", req.EventID),
		}

		return response 
	}

	//TODO: Call manager

	infoMsg := fmt.Sprintf("[%s] The competitor(%d) left the firing range", req.Time.Format(LAYOUT_TIME), req.CompetitorID)

	response = customresponse.CustomResponse{
		UserMessage:      infoMsg,
		DeveloperMessage: infoMsg,
		Err:              nil,
	}

	return response
}

func (h *BiathlonHandler) HandleStartPenalty(req *Request /*manager *state.StateManager*/) customresponse.CustomResponse {
	var response customresponse.CustomResponse

	if len(req.ExtraParams) > 0 {
		errorMsgUsr := "Error : start penalty lap : <Failed to decode request>"
		errorMsgDev := "Error : start penalty lap : <Too many extra parameters>"
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              fmt.Errorf("invalid number of parameters for event %d", req.EventID),
		}

		return response 
	}

	//TODO: Call manager

	infoMsg := fmt.Sprintf("[%s] The competitor(%d) entered the penalty laps", req.Time.Format(LAYOUT_TIME), req.CompetitorID)

	response = customresponse.CustomResponse{
		UserMessage:      infoMsg,
		DeveloperMessage: infoMsg,
		Err:              nil,
	}

	return response
}

func (h *BiathlonHandler) HandleEndPenalty(req *Request /*manager *state.StateManager*/) customresponse.CustomResponse {
	var response customresponse.CustomResponse

	if len(req.ExtraParams) > 0 {
		errorMsgUsr := "Error : end penalty lap : <Failed to decode request>"
		errorMsgDev := "Error : end penalty lap : <Too many extra parameters>"
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              fmt.Errorf("invalid number of parameters for event %d", req.EventID),
		}

		return response 
	}

	//TODO: Call manager

	infoMsg := fmt.Sprintf("[%s] The competitor(%d) left the penalty laps", req.Time.Format(LAYOUT_TIME), req.CompetitorID)

	response = customresponse.CustomResponse{
		UserMessage:      infoMsg,
		DeveloperMessage: infoMsg,
		Err:              nil,
	}

	return response
}

func (h *BiathlonHandler) HandleEndMain(req *Request /*manager *state.StateManager*/) customresponse.CustomResponse {
	var response customresponse.CustomResponse

	if len(req.ExtraParams) > 0 {
		errorMsgUsr := "Error : end main lap : <Failed to decode request>"
		errorMsgDev := "Error : end main lap : <Too many extra parameters>"
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              fmt.Errorf("invalid number of parameters for event %d", req.EventID),
		}

		return response
	}

	//TODO: Call manager

	infoMsg := fmt.Sprintf("[%s] The competitor(%d) ended the main lap", req.Time.Format(LAYOUT_TIME), req.CompetitorID)

	response = customresponse.CustomResponse{
		UserMessage:      infoMsg,
		DeveloperMessage: infoMsg,
		Err:              nil,
	}

	return response
}

func (h *BiathlonHandler) HandleDisqualified(req *Request /*manager *state.StateManager*/) customresponse.CustomResponse {
	var response customresponse.CustomResponse

	if len(req.ExtraParams) > 0 {
		errorMsgUsr := "Error : disqualified : <Failed to decode request>"
		errorMsgDev := "Error : disqualified : <Too many extra parameters>"
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              fmt.Errorf("invalid number of parameters for event %d", req.EventID),
		}

		return response
	}

	var builder strings.Builder
	for _, param := range req.ExtraParams {
		strParam, ok := param.(string)
		if !ok {
			errorMsgUsr := "Error : disqualified : <Failed to decode request>"
			errorMsgDev := "Error : disqualified : <Invalid type of extra params>"
			log.Println(errorMsgUsr)

			response = customresponse.CustomResponse{
				UserMessage:      errorMsgUsr,
				DeveloperMessage: errorMsgDev,
				Err:              fmt.Errorf("invalid disqualified description : %v", param),
			}

			return response
		}
		builder.WriteString(strParam)
		builder.WriteString(" ")
	}

	//TODO: Call manager

	infoMsg := fmt.Sprintf("[%s] The target(%d) can`t continue: %s", req.Time.Format(LAYOUT_TIME), req.CompetitorID, builder.String())

	response = customresponse.CustomResponse{
		UserMessage:      infoMsg,
		DeveloperMessage: infoMsg,
		Err:              nil,
	}

	return response
}