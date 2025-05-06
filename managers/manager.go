package biathlon

import (
	"time"

	"github.com/pkg/errors"
	"yadro_golang_task/api/model"
	"yadro_golang_task/config"
)

const (
	TARGETS uint32 = 5
)

type BiathlonManager struct {
	Competitors map[uint32]*model.Competitor
	Cfg         *config.Config
}

func NewBiathlonManager(c map[uint32]*model.Competitor, config *config.Config) *BiathlonManager {
	return &BiathlonManager{
		Competitors: c,
		Cfg:         config,
	}
}

func (bm *BiathlonManager) CompetitorIsExist(competitorID uint32) bool {
	_, exists := bm.Competitors[competitorID]
	return exists
}

func (bm *BiathlonManager) RegisterCompetitor(competitorID uint32) error {
	if bm.CompetitorIsExist(competitorID) {
		return errors.Errorf("Competitor %d is already registered", competitorID)
	}

	comp := &model.Competitor{
		ID:             competitorID,
		State:          model.REGISTERED,
		MainLapsResult: make([]model.LapResult, 0, bm.Cfg.LapLen),
	}
	bm.Competitors[competitorID] = comp
	
	return nil
}

func (bm *BiathlonManager) ChooseStartTime(competitorID uint32, startTime time.Time, currentTime time.Time) error {
	if !bm.CompetitorIsExist(competitorID) {
		return errors.Errorf("Competitor %d is not exist", competitorID)
	}

	if bm.Competitors[competitorID].State > model.WAIT_START {
		return errors.Errorf("Competitor %d is no longer waiting", competitorID)
	}

	if currentTime.After(startTime) {
		return errors.Errorf("Competitor %d start time is in the past", competitorID)
	}

	bm.Competitors[competitorID].PlanStart = startTime
	bm.Competitors[competitorID].State = model.WAIT_START
	bm.Competitors[competitorID].MainLapStartTime = startTime
	
	return nil
}

func (bm *BiathlonManager) Ready(competitorID uint32) error {
	if !bm.CompetitorIsExist(competitorID) {
		return errors.Errorf("Competitor %d is not exist", competitorID)
	}

	if bm.Competitors[competitorID].State > model.ON_START_LINE {
		return errors.Errorf("Competitor %d is already started", competitorID)
	}

	if bm.Competitors[competitorID].State < model.WAIT_START { 
		return errors.Errorf("Competitor %d is not waiting for the start", competitorID)
	}

	bm.Competitors[competitorID].State = model.ON_START_LINE
	
	return nil
}

func (bm *BiathlonManager) StartMain(competitorID uint32, realStartTime time.Time) error {
	if !bm.CompetitorIsExist(competitorID) {
		return errors.Errorf("Competitor %d is not exist", competitorID)
	}

	if bm.Competitors[competitorID].State >= model.STARTED {
		return errors.Errorf("Competitor %d is already started", competitorID)
	}

	if bm.Competitors[competitorID].State < model.ON_START_LINE { 
		return errors.Errorf("Competitor %d is not on starting line", competitorID)
	}

	bm.Competitors[competitorID].State = model.STARTED
	bm.Competitors[competitorID].RealStart = realStartTime
	bm.Competitors[competitorID].CurrentLap = 1

	deltaDuration := time.Duration(bm.Cfg.StartDelta.Hour())*time.Hour + time.Duration(bm.Cfg.StartDelta.Minute())*time.Minute + 
		time.Duration(bm.Cfg.StartDelta.Second())*time.Second

	if (realStartTime.After(bm.Competitors[competitorID].PlanStart.Add(deltaDuration))) {
		bm.Competitors[competitorID].State = model.NOT_STARTED
		bm.Competitors[competitorID].Shots = bm.Cfg.FiringLines * TARGETS
	}
	
	return nil
}

func (bm *BiathlonManager) ArriveFiringRange(competitorID uint32) error {
	if !bm.CompetitorIsExist(competitorID) {
		return errors.Errorf("Competitor %d is not exist", competitorID)
	}

	if bm.Competitors[competitorID].State >= model.ON_FIRING_RANGE && bm.Competitors[competitorID].State < model.LEFT_FIRING_RANGE {
		return errors.Errorf("Competitor %d is already on firing range", competitorID)
	}

	if bm.Competitors[competitorID].State < model.STARTED { 
		return errors.Errorf("Competitor %d had not start yet", competitorID)
	}

	if bm.Competitors[competitorID].State >= model.FINISHED { 
		return errors.Errorf("Competitor %d is already finished or disqualified", competitorID)
	}

	bm.Competitors[competitorID].State = model.ON_FIRING_RANGE
	bm.Competitors[competitorID].VisitedFiringLines++
	
	return nil
}

func (bm *BiathlonManager) Hit(competitorID uint32) error {
	if !bm.CompetitorIsExist(competitorID) {
		return errors.Errorf("Competitor %d is not exist", competitorID)
	}

	if bm.Competitors[competitorID].State != model.ON_FIRING_RANGE {
		return errors.Errorf("Competitor %d is not on firing range", competitorID)
	}

	bm.Competitors[competitorID].Hits++
	
	return nil
}

func (bm *BiathlonManager) LeftFiringRange(competitorID uint32) error {
	if !bm.CompetitorIsExist(competitorID) {
		return errors.Errorf("Competitor %d is not exist", competitorID)
	}

	if bm.Competitors[competitorID].State != model.ON_FIRING_RANGE {
		return errors.Errorf("Competitor %d is not on firing range", competitorID)
	}

	bm.Competitors[competitorID].State = model.LEFT_FIRING_RANGE
	bm.Competitors[competitorID].Shots += TARGETS 
	
	return nil
}

func (bm *BiathlonManager) StartPenalty(competitorID uint32, curTime time.Time) error {
	if !bm.CompetitorIsExist(competitorID) {
		return errors.Errorf("Competitor %d is not exist", competitorID)
	}

	if bm.Competitors[competitorID].State != model.LEFT_FIRING_RANGE && bm.Competitors[competitorID].State != model.LEFT_PENALTY_LAP {
		return errors.Errorf("Competitor %d can start penalty laps only after firing line or other penalty lap", competitorID)
	}

	bm.Competitors[competitorID].State = model.ON_PENALTY_LAP
	bm.Competitors[competitorID].PenaltyLapStartTime = curTime
	
	return nil
}

func (bm *BiathlonManager) EndPenalty(competitorID uint32, curTime time.Time) error {
	if !bm.CompetitorIsExist(competitorID) {
		return errors.Errorf("Competitor %d is not exist", competitorID)
	}

	if bm.Competitors[competitorID].State != model.ON_PENALTY_LAP {
		return errors.Errorf("Competitor %d cannot finish penalty lap before start it", competitorID)
	}

	bm.Competitors[competitorID].State = model.LEFT_PENALTY_LAP
	bm.Competitors[competitorID].TotalPenaltyTime += curTime.Sub(bm.Competitors[competitorID].PenaltyLapStartTime)
	
	return nil
}

func (bm *BiathlonManager) EndMain(competitorID uint32, curTime time.Time) error {
	if !bm.CompetitorIsExist(competitorID) {
		return errors.Errorf("competitor %d is not exist", competitorID)
	}

	if bm.Competitors[competitorID].State != model.LEFT_FIRING_RANGE && bm.Competitors[competitorID].State != model.LEFT_PENALTY_LAP {
		return errors.Errorf("competitor %d can finish main lap only after fire line or penalty laps", competitorID)
	}

	bm.Competitors[competitorID].State = model.END_MAIN_LAP
	timeDuration := curTime.Sub(bm.Competitors[competitorID].MainLapStartTime)
	bm.Competitors[competitorID].MainLapsResult = append(bm.Competitors[competitorID].MainLapsResult, model.LapResult{
		Time: timeDuration,
		AvgSpeed: float32(bm.Cfg.LapLen) / float32(timeDuration.Seconds()),
	})
	bm.Competitors[competitorID].MainLapStartTime = curTime

	if bm.Cfg.Laps > bm.Competitors[competitorID].CurrentLap {
		bm.Competitors[competitorID].CurrentLap++
		bm.Competitors[competitorID].MainLapStartTime = curTime
	} else {
		bm.Competitors[competitorID].State = model.FINISHED
		bm.Competitors[competitorID].TotalTime = curTime.Sub(bm.Competitors[competitorID].PlanStart)
		bm.Competitors[competitorID].PenaltyLapsResult = model.LapResult{
			Time: bm.Competitors[competitorID].TotalPenaltyTime,
			AvgSpeed: float32(bm.Cfg.PenaltyLen) * float32(bm.Competitors[competitorID].Shots - bm.Competitors[competitorID].Hits) /
			 float32(bm.Competitors[competitorID].TotalPenaltyTime.Seconds()),
		}
	}

	return nil
}

func (bm *BiathlonManager) Disqualified(competitorID uint32) error {
	if !bm.CompetitorIsExist(competitorID) {
		return errors.Errorf("Competitor %d is not exist", competitorID)
	}

	if bm.Competitors[competitorID].State < model.STARTED {
		return errors.Errorf("Competitor %d cannot be disqualified before start competitions", competitorID)
	}

	if bm.Competitors[competitorID].State == model.NOT_STARTED || bm.Competitors[competitorID].State == model.NOT_FINISHED {
		return errors.Errorf("Competitor %d is already disqualified", competitorID)
	}

	bm.Competitors[competitorID].State = model.NOT_FINISHED
	bm.Competitors[competitorID].PenaltyLapsResult = model.LapResult{
		Time: bm.Competitors[competitorID].TotalPenaltyTime,
		AvgSpeed: float32(bm.Cfg.PenaltyLen) * float32(bm.Competitors[competitorID].Shots - bm.Competitors[competitorID].Hits) /
		 float32(bm.Competitors[competitorID].TotalPenaltyTime.Seconds()),
	}

	return nil
}