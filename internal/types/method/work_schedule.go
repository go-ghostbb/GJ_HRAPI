package method

type WorkSchedule interface {
	// select dbo.FN_C_WorkHoursCountMany(@empID, @startDate, @endDate, @startTime, @endTime)
	WorkHoursCountMany(empID uint, startDate, endDate, startTime, endTime string) (float32, error)
	// select dbo.FN_C_WorkHoursCount(@empID, @start, @end)
	WorkHourCount(empID uint, start, end string) (float32, error)
}
