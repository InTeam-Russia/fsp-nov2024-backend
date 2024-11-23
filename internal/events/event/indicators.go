package event

import "time"

func calculateIndicators(events []EventOut) *Indicators {
	minStartDate := events[0].StartDate
	maxEndDate := events[0].EndDate

	for _, event := range events {
		if minStartDate.After(event.StartDate) {
			minStartDate = event.StartDate
		}

		if maxEndDate.Before(event.EndDate) {
			maxEndDate = event.EndDate
		}
	}

	dates := make([]time.Time, 0)
	for i := minStartDate.Truncate(24 * time.Hour); i.Before(maxEndDate); i = i.Add(24 * time.Hour) {
		for _, event := range events {
			if i.After(event.StartDate) && i.Before(event.EndDate) || i.Equal(event.StartDate) || i.Equal(event.EndDate) {
				dates = append(dates, i)
				break
			}
		}
	}

	return &Indicators{dates}
}
