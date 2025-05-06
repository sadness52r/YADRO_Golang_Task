package biathlon

import (
	"log"
	"testing"
	"time"
	"yadro_golang_task/api/handlers"
	"yadro_golang_task/api/model"
	"yadro_golang_task/config"
	managers "yadro_golang_task/managers"
)

const (
	LAYOUT_TIME string = "15:04:05"
)

var manager *managers.BiathlonManager

func parseTime(t any) time.Time {
	tStr, ok := t.(string)
	if !ok {
		return time.Time{}
	}

	timestamp, err := time.Parse(biathlon.LAYOUT_TIME, tStr)
	if err != nil {
		return time.Time{}
	}
	
	return timestamp
}

func TestMain(m *testing.M) {
	defer TearDown()

	startTime, _ := time.Parse(LAYOUT_TIME, "09:30:00")
	timeDelta, _ := time.Parse(LAYOUT_TIME, "00:00:30")
	cfg := &config.Config {
		Laps: 2,
		LapLen: 3651,
		PenaltyLen: 50,
		FiringLines: 1,
		Start: config.MyTime{
			Time: startTime,
		},
		StartDelta: config.MyTime{
			Time: timeDelta,
		},
	}
	competitors := make(map[uint32]*model.Competitor)
	manager = managers.NewBiathlonManager(competitors, cfg)
	m.Run()
}

func TestRegister(t *testing.T) {
	testParams := []struct {
		CompetitorID uint32
		ExpectError bool
	}{
		{
			CompetitorID: 1,
			ExpectError: false,
		},
		{
			CompetitorID: 1,
			ExpectError: true,
		},
	}

	for _, testParam := range testParams {
		err := manager.RegisterCompetitor(testParam.CompetitorID)
		if testParam.ExpectError && err == nil {
			t.Error(err)
		}
		if !testParam.ExpectError && err != nil {
			t.Error(err)
		}
	}
}

func TestChooseStartTime(t *testing.T) {
	testParams := []struct {
		Request biathlon.Request
		CurTime time.Time
		ExpectError bool
	}{
		{
			Request: biathlon.Request{
				CompetitorID: 1,
				ExtraParams: []any{
					"09:30:00.000",
				},
			},
			CurTime: func() time.Time {
				t, _ := time.Parse(biathlon.LAYOUT_TIME, "09:25:00.000")
				return t
			}(),
			ExpectError: false,
		},
		{
			Request: biathlon.Request{
				CompetitorID: 2,
				ExtraParams: []any{
					"09:30:00.000",
				},
			},
			CurTime: func() time.Time {
				t, _ := time.Parse(biathlon.LAYOUT_TIME, "09:25:00.000")
				return t
			}(),
			ExpectError: true,
		},
		{
			Request: biathlon.Request{
				CompetitorID: 1,
				ExtraParams: []any{
					"09:30:00.000",
				},
			},
			CurTime: func() time.Time {
				t, _ := time.Parse(biathlon.LAYOUT_TIME, "09:35:00.000")
				return t
			}(),
			ExpectError: true,
		},
	}
	
	for _, testParam := range testParams {
		err := manager.ChooseStartTime(testParam.Request.CompetitorID, parseTime(testParam.Request.ExtraParams[0]), testParam.CurTime)
		if testParam.ExpectError && err == nil {
			t.Error(err)
		}
		if !testParam.ExpectError && err != nil {
			t.Error(err)
		}
	}
}

func TestReady(t *testing.T) {
	testParams := []struct {
		CompetitorID uint32
		ExpectError bool
	}{
		{
			CompetitorID: 1,
			ExpectError: false,
		},
		{
			CompetitorID: 2,
			ExpectError: true,
		},
	}

	for _, testParam := range testParams {
		err := manager.Ready(testParam.CompetitorID)
		if testParam.ExpectError && err == nil {
			t.Error(err)
		}
		if !testParam.ExpectError && err != nil {
			t.Error(err)
		}
	}
}

func TestStartMain(t *testing.T) {
	testParams := []struct {
		CompetitorID uint32
		ExpectError bool
	}{
		{
			CompetitorID: 1,
			ExpectError: false,
		},
		{
			CompetitorID: 2,
			ExpectError: true,
		},
	}

	for _, testParam := range testParams {
		startTime, _ := time.Parse(biathlon.LAYOUT_TIME, "09:30:02.000")
		err := manager.StartMain(testParam.CompetitorID, startTime)
		if testParam.ExpectError && err == nil {
			t.Error(err)
		}
		if !testParam.ExpectError && err != nil {
			t.Error(err)
		}
	}
}

func TestArriveFiringRange(t *testing.T) {
	testParams := []struct {
		CompetitorID uint32
		ExpectError bool
	}{
		{
			CompetitorID: 1,
			ExpectError: false,
		},
		{
			CompetitorID: 2,
			ExpectError: true,
		},
	}

	for _, testParam := range testParams {
		err := manager.ArriveFiringRange(testParam.CompetitorID)
		if testParam.ExpectError && err == nil {
			t.Error(err)
		}
		if !testParam.ExpectError && err != nil {
			t.Error(err)
		}
	}
}

func TestHit(t *testing.T) {
	testParams := []struct {
		CompetitorID uint32
		ExpectError bool
	}{
		{
			CompetitorID: 1,
			ExpectError: false,
		},
		{
			CompetitorID: 2,
			ExpectError: true,
		},
	}

	for _, testParam := range testParams {
		err := manager.Hit(testParam.CompetitorID)
		if testParam.ExpectError && err == nil {
			t.Error(err)
		}
		if !testParam.ExpectError && err != nil {
			t.Error(err)
		}
	}
}

func TestLeftFiringRange(t *testing.T) {
	testParams := []struct {
		CompetitorID uint32
		ExpectError bool
	}{
		{
			CompetitorID: 1,
			ExpectError: false,
		},
		{
			CompetitorID: 2,
			ExpectError: true,
		},
	}

	for _, testParam := range testParams {
		err := manager.LeftFiringRange(testParam.CompetitorID)
		if testParam.ExpectError && err == nil {
			t.Error(err)
		}
		if !testParam.ExpectError && err != nil {
			t.Error(err)
		}
	}
}

func TestStartPenalty(t *testing.T) {
	testParams := []struct {
		CompetitorID uint32
		ExpectError bool
	}{
		{
			CompetitorID: 1,
			ExpectError: false,
		},
		{
			CompetitorID: 2,
			ExpectError: true,
		},
	}

	for _, testParam := range testParams {
		startTime, _ := time.Parse(biathlon.LAYOUT_TIME, "09:44:02.000")
		err := manager.StartPenalty(testParam.CompetitorID, startTime)
		if testParam.ExpectError && err == nil {
			t.Error(err)
		}
		if !testParam.ExpectError && err != nil {
			t.Error(err)
		}
	}
}

func TestEndPenalty(t *testing.T) {
	testParams := []struct {
		CompetitorID uint32
		ExpectError bool
	}{
		{
			CompetitorID: 1,
			ExpectError: false,
		},
		{
			CompetitorID: 2,
			ExpectError: true,
		},
	}

	for _, testParam := range testParams {
		curTime, _ := time.Parse(biathlon.LAYOUT_TIME, "09:46:02.000")
		err := manager.EndPenalty(testParam.CompetitorID, curTime)
		if testParam.ExpectError && err == nil {
			t.Error(err)
		}
		if !testParam.ExpectError && err != nil {
			t.Error(err)
		}
	}
}

func TestEndMain(t *testing.T) {
	testParams := []struct {
		CompetitorID uint32
		ExpectError bool
	}{
		{
			CompetitorID: 1,
			ExpectError: false,
		},
		{
			CompetitorID: 2,
			ExpectError: true,
		},
	}

	for _, testParam := range testParams {
		curTime, _ := time.Parse(biathlon.LAYOUT_TIME, "09:50:02.000")
		err := manager.EndMain(testParam.CompetitorID, curTime)
		if testParam.ExpectError && err == nil {
			t.Error(err)
		}
		if !testParam.ExpectError && err != nil {
			t.Error(err)
		}
	}
}

func TestDisqualified(t *testing.T) {
	testParams := []struct {
		CompetitorID uint32
		ExpectError bool
	}{
		{
			CompetitorID: 1,
			ExpectError: false,
		},
		{
			CompetitorID: 2,
			ExpectError: true,
		},
	}

	for _, testParam := range testParams {
		err := manager.Disqualified(testParam.CompetitorID)
		if testParam.ExpectError && err == nil {
			t.Error(err)
		}
		if !testParam.ExpectError && err != nil {
			t.Error(err)
		}
	}
}

func TearDown() {
	for _, comp := range manager.Competitors {
		log.Println(comp)
	}
}
