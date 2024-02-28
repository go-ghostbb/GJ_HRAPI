package utils

import "time"

// TimeOverlapCount
// time overlap count
func TimeOverlapCount(timeRange []time.Time, timePeriod ...[]time.Time) (sec int64) {
	start, end := timeRange[0].Unix(), timeRange[1].Unix()

	for _, t := range timePeriod {
		checkStart, checkEnd := t[0].Unix(), t[1].Unix()
		if start <= checkEnd && end >= checkStart {
			// if overlap
			// Divided into 4 situations
			if checkStart < start && start < checkEnd {
				sec += checkEnd - start
				continue
			}
			if checkStart < end && end < checkEnd {
				sec += end - checkStart
				continue
			}
			if start <= checkStart && end >= checkEnd {
				sec += checkEnd - checkStart
				continue
			}
			if start >= checkStart && end <= checkEnd {
				sec += end - start
				continue
			}
		}
	}

	return sec
}
