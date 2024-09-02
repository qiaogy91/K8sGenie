package alert

import "time"

func (a *Alert) GetStartTime() string {
	utcTime, err := time.Parse(time.RFC3339, a.StartsAt)
	if err != nil {
		return "time parse failed"
	}
	return utcTime.Local().Format("2006-01-02 15:04:05")
}

func (a *Alert) GetEndTime() string {
	utcTime, err := time.Parse(time.RFC3339, a.EndsAt)
	if err != nil {
		return "time parse failed"
	}
	return utcTime.Local().Format("2006-01-02 15:04:05")
}
