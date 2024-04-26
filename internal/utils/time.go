package utils

import "time"

// TimeOverlapCount
// time overlap count
func TimeOverlapCount(timeRange []time.Time, timePeriod ...[]time.Time) (sec int64) {
	if timeRange[0].After(timeRange[1]) {
		// 如果開始時間>結束時間，代表過夜
		timeRange[1] = timeRange[1].AddDate(0, 0, 1)
	}

	for _, t := range timePeriod {
		if t[0].After(t[1]) {
			// 如果開始時間>結束時間，代表過夜
			t[1] = t[1].AddDate(0, 0, 1)
		}

		if (timeRange[0].Before(t[1]) || timeRange[0].Equal(t[1])) && (t[0].Before(timeRange[1]) || t[1].Equal(timeRange[0])) {
			overlapStart := max(timeRange[0], t[0])
			overlapEnd := min(timeRange[1], t[1])
			sec += int64(overlapEnd.Sub(overlapStart).Seconds())
		}
	}

	return sec
}

// max 函數返回兩個時間中的較大值
func max(a, b time.Time) time.Time {
	if a.After(b) {
		return a
	}
	return b
}

// min 函數返回兩個時間中的較小值
func min(a, b time.Time) time.Time {
	if a.Before(b) {
		return a
	}
	return b
}
