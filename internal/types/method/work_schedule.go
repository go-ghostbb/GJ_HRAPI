package method

import "gorm.io/gen"

type WorkSchedule interface {
	// select * from @@table where schedule_date = @dateOnly and deleted_at is null
	QueryByDate(dateOnly string) ([]*gen.T, error)
	// select * from @@table where schedule_date between @dateOnly1 and @dateOnly2 and deleted_at is null
	QueryByDateRange(dateOnly1, dateOnly2 string) ([]*gen.T, error)
	// delete @@table where employee_id = @empID and schedule_date like @likeDate
	DeleteByDateRange(empID uint, likeDate string) (gen.RowsAffected, error)
	// select * from @@table where schedule_date between @dateOnly1 and @dateOnly2 and employee_id = @empID and deleted_at is null
	QueryByDateRangeAndEmpID(empID uint, dateOnly1, dateOnly2 string) ([]*gen.T, error)
	// select dbo.FN_C_WorkHoursCountMany(@empID, @startDate, @endDate, @startTime, @endTime);
	WorkHoursCountMany(empID uint, startDate, endDate, startTime, endTime string) (float32, error)
}
