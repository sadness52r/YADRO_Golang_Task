package biathlon

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"yadro_golang_task/api"
	"yadro_golang_task/api/model"
	managers "yadro_golang_task/managers"
	"yadro_golang_task/table"
)

const (
	LAYOUT_TIME string = "15:04:05.000"
)

var hitTargets = make(map[uint32]bool)

type BiathlonHandler struct {
	manager *managers.BiathlonManager
	tableResponses []model.TableResponse
}

func NewBiathlonHandler(m *managers.BiathlonManager) *BiathlonHandler {
	return &BiathlonHandler{
		manager: m,
	}
}

func (h *BiathlonHandler) HandleRegister(req *Request) customresponse.CustomResponse {
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

	errDev := h.manager.RegisterCompetitor(req.CompetitorID)
	if errDev != nil {
		errorMsgUsr := "Error : register competitor : <Failed to register competitor>"
		errorMsgDev := fmt.Sprintf("Error : register competitor : <%s>", errDev)
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              errDev,
		}

		return response
	}

	infoMsg := fmt.Sprintf("[%s] The competitor(%d) registered", req.Time.Format(LAYOUT_TIME), req.CompetitorID)

	response = customresponse.CustomResponse{
		UserMessage:      infoMsg,
		DeveloperMessage: infoMsg,
		Err:              nil,
	}

	return response
}

func (h *BiathlonHandler) HandleChooseStartTime(req *Request) customresponse.CustomResponse {
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

	errDev := h.manager.ChooseStartTime(req.CompetitorID, timestamp, req.Time)
	if errDev != nil {
		errorMsgUsr := "Error : choose start time : <Failed to choose start time>"
		errorMsgDev := fmt.Sprintf("Error : choose start time : <%s>", errDev)
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              errDev,
		}

		return response
	}

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

func (h *BiathlonHandler) HandleReady(req *Request) customresponse.CustomResponse {
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

	errDev := h.manager.Ready(req.CompetitorID)
	if errDev != nil {
		errorMsgUsr := "Error : ready to start : <Failed to ready to start>"
		errorMsgDev := fmt.Sprintf("Error : ready to start : <%s>", errDev)
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              errDev,
		}

		return response
	}

	infoMsg := fmt.Sprintf("[%s] The competitor(%d) is on the start line", req.Time.Format(LAYOUT_TIME), req.CompetitorID)

	response = customresponse.CustomResponse{
		UserMessage:      infoMsg,
		DeveloperMessage: infoMsg,
		Err:              nil,
	}

	return response
}

func (h *BiathlonHandler) HandleStartMain(req *Request) customresponse.CustomResponse {
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

	errDev := h.manager.StartMain(req.CompetitorID, req.Time)
	if errDev != nil {
		errorMsgUsr := "Error : start main lap : <Failed to start main lap>"
		errorMsgDev := fmt.Sprintf("Error : start main lap : <%s>", errDev)
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              errDev,
		}

		return response
	}

	infoMsg := fmt.Sprintf("[%s] The competitor(%d) has started", req.Time.Format(LAYOUT_TIME), req.CompetitorID)

	response = customresponse.CustomResponse{
		UserMessage:      infoMsg,
		DeveloperMessage: infoMsg,
		Err:              nil,
	}

	return response
}

func (h *BiathlonHandler) HandleArriveFiringRange(req *Request) customresponse.CustomResponse {
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

	if h.manager.Competitors[req.CompetitorID].VisitedFiringLines == h.manager.Cfg.FiringLines {
		errorMsgUsr := "Error : arrive firing range : <Failed to decode request>"
		errorMsgDev := "Error : arrive firing range : <Firing lines limit exceeded>"
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              fmt.Errorf("firing lines expected %d", h.manager.Cfg.FiringLines),
		}

		return response
		
	}

	errDev := h.manager.ArriveFiringRange(req.CompetitorID)
	if errDev != nil {
		errorMsgUsr := "Error : arrive firing range : <Failed to arrive firing range>"
		errorMsgDev := fmt.Sprintf("Error : arrive firing range : <%s>", errDev)
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              errDev,
		}

		return response
	}

	infoMsg := fmt.Sprintf("[%s] The competitor(%d) is on the firing range(%d)", req.Time.Format(LAYOUT_TIME), req.CompetitorID, firingRange)

	response = customresponse.CustomResponse{
		UserMessage:      infoMsg,
		DeveloperMessage: infoMsg,
		Err:              nil,
	}

	return response
}

func (h *BiathlonHandler) HandleHit(req *Request) customresponse.CustomResponse {
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

	if _, ok := hitTargets[target]; ok {
		errorMsgUsr := "Error : hit target : <Failed to decode request>"
		errorMsgDev := "Error : hit target : <Target is already hit>"
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              fmt.Errorf("target %d is already hit", target),
		}

		return response
	}

	errDev := h.manager.Hit(req.CompetitorID)
	if errDev != nil {
		errorMsgUsr := "Error : hit target : <Failed to hit target>"
		errorMsgDev := fmt.Sprintf("Error : hit target : <%s>", errDev)
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              errDev,
		}

		return response
	}

	hitTargets[target] = true

	infoMsg := fmt.Sprintf("[%s] The target(%d) has been hit by competitor(%d)", req.Time.Format(LAYOUT_TIME), target, req.CompetitorID)

	response = customresponse.CustomResponse{
		UserMessage:      infoMsg,
		DeveloperMessage: infoMsg,
		Err:              nil,
	}

	return response
}

func (h *BiathlonHandler) HandleLeftFiringRange(req *Request) customresponse.CustomResponse {
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

	for k := range hitTargets {
		delete(hitTargets, k)
	}

	errDev := h.manager.LeftFiringRange(req.CompetitorID)
	if errDev != nil {
		errorMsgUsr := "Error : left firing range : <Failed to left firing range>"
		errorMsgDev := fmt.Sprintf("Error : left firing range : <%s>", errDev)
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              errDev,
		}

		return response
	}

	infoMsg := fmt.Sprintf("[%s] The competitor(%d) left the firing range", req.Time.Format(LAYOUT_TIME), req.CompetitorID)

	response = customresponse.CustomResponse{
		UserMessage:      infoMsg,
		DeveloperMessage: infoMsg,
		Err:              nil,
	}

	return response
}

func (h *BiathlonHandler) HandleStartPenalty(req *Request) customresponse.CustomResponse {
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

	errDev := h.manager.StartPenalty(req.CompetitorID, req.Time)
	if errDev != nil {
		errorMsgUsr := "Error : start penalty lap : <Failed to start penalty lap>"
		errorMsgDev := fmt.Sprintf("Error : start penalty lap : <%s>", errDev)
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              errDev,
		}

		return response
	}

	infoMsg := fmt.Sprintf("[%s] The competitor(%d) entered the penalty laps", req.Time.Format(LAYOUT_TIME), req.CompetitorID)

	response = customresponse.CustomResponse{
		UserMessage:      infoMsg,
		DeveloperMessage: infoMsg,
		Err:              nil,
	}

	return response
}

func (h *BiathlonHandler) HandleEndPenalty(req *Request) customresponse.CustomResponse {
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

	errDev := h.manager.EndPenalty(req.CompetitorID, req.Time)
	if errDev != nil {
		errorMsgUsr := "Error : end penalty lap : <Failed to end penalty lap>"
		errorMsgDev := fmt.Sprintf("Error : end penalty lap : <%s>", errDev)
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              errDev,
		}

		return response
	}

	infoMsg := fmt.Sprintf("[%s] The competitor(%d) left the penalty laps", req.Time.Format(LAYOUT_TIME), req.CompetitorID)

	response = customresponse.CustomResponse{
		UserMessage:      infoMsg,
		DeveloperMessage: infoMsg,
		Err:              nil,
	}

	return response
}

func (h *BiathlonHandler) HandleEndMain(req *Request) customresponse.CustomResponse {
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

	if h.manager.Competitors[req.CompetitorID].CurrentLap > h.manager.Cfg.Laps {
		errorMsgUsr := "Error : end main lap : <Failed to decode request>"
		errorMsgDev := "Error : end main lap : <Main laps limit exceeded>"
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              fmt.Errorf("main laps expected %d", h.manager.Cfg.Laps),
		}

		return response
		
	}

	errDev := h.manager.EndMain(req.CompetitorID, req.Time)
	if errDev != nil {
		errorMsgUsr := "Error : end main lap : <Failed to end main lap>"
		errorMsgDev := fmt.Sprintf("Error : end main lap : <%s>", errDev)
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              errDev,
		}

		return response
	}

	infoMsg := fmt.Sprintf("[%s] The competitor(%d) ended the main lap", req.Time.Format(LAYOUT_TIME), req.CompetitorID)

	response = customresponse.CustomResponse{
		UserMessage:      infoMsg,
		DeveloperMessage: infoMsg,
		Err:              nil,
	}

	return response
}

func (h *BiathlonHandler) HandleDisqualified(req *Request) customresponse.CustomResponse {
	var response customresponse.CustomResponse

	if len(req.ExtraParams) == 0 {
		errorMsgUsr := "Error : disqualified : <Failed to decode request>"
		errorMsgDev := "Error : disqualified : <Not enough extra parameters>"
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

	errDev := h.manager.Disqualified(req.CompetitorID)
	if errDev != nil {
		errorMsgUsr := "Error : disqualified : <Failed to disqualified>"
		errorMsgDev := fmt.Sprintf("Error : disqualified : <%s>", errDev)
		log.Println(errorMsgUsr)

		response = customresponse.CustomResponse{
			UserMessage:      errorMsgUsr,
			DeveloperMessage: errorMsgDev,
			Err:              errDev,
		}

		return response
	}

	infoMsg := fmt.Sprintf("[%s] The competitor(%d) can`t continue: %s", req.Time.Format(LAYOUT_TIME), req.CompetitorID, builder.String())

	response = customresponse.CustomResponse{
		UserMessage:      infoMsg,
		DeveloperMessage: infoMsg,
		Err:              nil,
	}

	return response
}

func (h *BiathlonHandler) HandleGetStatus() customresponse.CustomResponse {
	var response customresponse.CustomResponse

	for key, value := range h.manager.Competitors {
		tableResponse := model.TableResponse{}
		if value.State == model.NOT_STARTED || value.State == model.NOT_FINISHED {
			tableResponse.TotalTime = value.State.String()
		} else {
			tableResponse.TotalTime = model.FormatDuration(value.TotalTime)
		}
		tableResponse.CompetitorID = key
		tableResponse.MainLapsInfo = value.MainLapsResult
		tableResponse.PenaltyLapsInfo = value.PenaltyLapsResult
		tableResponse.Hits = value.Hits
		tableResponse.Shots = value.Shots

		h.tableResponses = append(h.tableResponses, tableResponse)
	}

	infoMsg := "Status was saved to result.txt"

	response = customresponse.CustomResponse{
		UserMessage:      infoMsg,
		DeveloperMessage: infoMsg,
		Data:             h.tableResponses,
		Err:              nil,
	}

	table.ProcessAndSave("result.txt", h.tableResponses, h.manager.Cfg)

	return response
}