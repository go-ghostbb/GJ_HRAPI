package utils

import "time"

// CalcSeniority 計算年資到@currentDate
func CalcSeniority(hireDate time.Time, currentDate time.Time) (years int, months int, days int) {
	years = currentDate.Year() - hireDate.Year()
	months = int(currentDate.Month()) - int(hireDate.Month())
	days = currentDate.Day() - hireDate.Day()

	// 處理月份和天數的負數情況
	if months < 0 || (months == 0 && days < 0) {
		years--
		if months < 0 {
			months += 12
		}
		if days < 0 {
			// 獲取上個月的天數
			lastMonthDays := time.Date(currentDate.Year(), currentDate.Month()-1, 0, 0, 0, 0, 0, time.UTC).Day()
			days += lastMonthDays + 1
			months--
		}
	}

	return
}

func GetDay(year int, month int) float64 {
	list := []float64{0.0, 31.0, 28.0, 31.0, 30.0, 31.0, 30.0, 31.0, 31.0, 30.0, 31.0, 30.0, 31.0}
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		list[2] = 29
	}
	return list[month]
}
