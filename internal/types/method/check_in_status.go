package method

import "gorm.io/gen"

type CheckInStatus interface {
	// delete @@table where work_check_in_date = @dateOnly
	DeleteByDate(dateOnly string) (gen.RowsAffected, error)
	// delete @@table where work_check_in_date between @dateOnly1 and @dateOnly2
	DeleteByDateRange(dateOnly1, dateOnly2 string) (gen.RowsAffected, error)
	// update @@table
	// set
	//     work_attend_status = (
	//         select
	//             case
	//                 when check_in_status.work_attend_time = '00:00:00' then 'not swiped'
	//                 when work_start >= convert(varchar(5), check_in_status.work_attend_time, 108) then 'normal'
	//                 when work_start < convert(varchar(5), check_in_status.work_attend_time, 108) then 'late'
	//                 end
	//         from work_shift
	//         where id = work_shift_id
	//     ),
	//     work_attend_proc_status = (
	//         select
	//             case
	//                 when check_in_status.work_attend_time = '00:00:00' then 'not processed'
	//                 when work_start >= convert(varchar(5), check_in_status.work_attend_time, 108) then 'normal'
	//                 when work_start < convert(varchar(5), check_in_status.work_attend_time, 108) then 'not processed'
	//                 end
	//         from work_shift
	//         where id = work_shift_id
	//     ),
	//     off_work_attend_status = (
	//         select
	//             case
	//                 when check_in_status.off_work_attend_time = '00:00:00' then 'not swiped'
	//                 when work_end <= check_in_status.off_work_attend_time then 'normal'
	//                 when work_end > check_in_status.off_work_attend_time then 'early'
	//                 end
	//         from work_shift
	//         where id = work_shift_id
	//     ),
	//     off_work_attend_proc_status = (
	//         select
	//             case
	//                 when check_in_status.off_work_attend_time = '00:00:00' then 'not processed'
	//                 when work_end <= check_in_status.off_work_attend_time then 'normal'
	//                 when work_end > check_in_status.off_work_attend_time then 'not processed'
	//                 end
	//         from work_shift
	//         where id = work_shift_id
	//     ),
	//     absence_hours = (
	//         select
	//             iif(check_in_status.work_attend_time != '00:00:00' and check_in_status.off_work_attend_time != '00:00:00',
	//                 iif(work_start < convert(varchar(5), check_in_status.work_attend_time, 108),
	//                     convert(float, datediff(second, work_start, work_attend_time)) / 60 / 60,
	//                     0)
	//                     +
	//                 iif(work_end > check_in_status.off_work_attend_time,
	//                     convert(float, datediff(second , off_work_attend_time, '17:00:00')) / 60 / 60,
	//                     0)
	//                 , absence_hours)
	//         from work_shift
	//         where id = work_shift_id
	//     )
	// where work_check_in_date in (@dateOnly)
	UpdateStatus(dateOnly []string) (gen.RowsAffected, error)
	// update @@table
	// {{if isWork}}
	//     set work_attend_time = @timeOnly, updated_at = getdate()
	// {{else}}
	//     set off_work_attend_time = @timeOnly, updated_at = getdate()
	// {{end}}
	// where
	//     work_shift_id = (select id from work_shift where code = @workShiftCode) and
	//     employee_id = (select id from employee where card_number = @cardNum) and
	//     {{if isWork}}
	//         work_check_in_date = @dateOnly
	//     {{else}}
	//         off_work_check_in_date = @dateOnly
	//     {{end}}
	UpdateTime(timeOnly, dateOnly, workShiftCode, cardNum string, isWork bool) (gen.RowsAffected, error)
}
