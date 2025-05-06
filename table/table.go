package table

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"yadro_golang_task/api/model"
	"yadro_golang_task/config"
)

func parseTime(t string) time.Duration {
	if t == "NotFinished" || t == "NotStarted" || t == "" {
		return time.Hour * 999
	}
	d, err := time.ParseDuration(fmt.Sprintf("%s%s", t[:2]+"h", t[3:5]+"m")+t[6:]+"s")
	if err != nil {
		return time.Hour * 999
	}
	return d
}

func formatCompetitor(c model.TableResponse, cfg *config.Config) string {
	var b strings.Builder

	fmt.Fprintf(&b, "[%s] %d [", c.TotalTime, c.CompetitorID)

	for i := 0; i < int(cfg.Laps); i++ {
		if i > 0 {
			b.WriteString("}, {")
		} else {
			b.WriteString("{")
		}

		if i < len(c.MainLapsInfo) {
			lapResult := c.MainLapsInfo[i]
			b.WriteString(fmt.Sprintf("%s, %.3f", model.FormatDuration(lapResult.Time), lapResult.AvgSpeed))
		} else {
			b.WriteString(",")
		}
	}
	b.WriteString("}]")

	if c.Shots - c.Hits > 0 {
		b.WriteString(fmt.Sprintf(" {%s, %.3f}", model.FormatDuration(c.PenaltyLapsInfo.Time), c.PenaltyLapsInfo.AvgSpeed))
	} else {
		b.WriteString(" {,}")
	}

	b.WriteString(fmt.Sprintf(" %d/%d", c.Hits, c.Shots))

	return b.String()
}

func ProcessAndSave(filename string, tableResponses []model.TableResponse, cfg *config.Config) error {
	sort.Slice(tableResponses, func(i, j int) bool {
		return parseTime(tableResponses[i].TotalTime) < parseTime(tableResponses[j].TotalTime)
	})

	var result string
	for _, c := range tableResponses {
		result += formatCompetitor(c, cfg) + "\n"
	}

	return os.WriteFile(filename, []byte(result), 0644)
}