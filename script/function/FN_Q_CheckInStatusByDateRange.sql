drop function if exists [dbo].[FN_Q_CheckInStatusByDateRange]
go

create function [dbo].[FN_Q_CheckInStatusByDateRange] (@emp_id bigint, @start_date date, @end_date date, @abnormal bit)
    returns @result table (
        id  bigint primary key,
        created_at datetimeoffset,
        updated_at datetimeoffset,
        deleted_at datetimeoffset,
        employee_id bigint,
        work_shift_id bigint,
        date date,
        work_check_in datetime,
        work_attend_status nvarchar(max),
        work_attend_proc_status nvarchar(max),
        off_work_check_in datetime,
        off_work_attend_status nvarchar(max),
        off_work_attend_proc_status nvarchar(max),
        absence_hours float,
        leave_hours float,
        sign_leave_hours float,
        overtime_hours float,
        sign_overtime_hours float
    )
as
begin

    ;with
         step1 as (
             select *
             from check_in_status
             where employee_id = @emp_id and
                   date between @start_date and @end_date and
                   deleted_at is null and
                   (@abnormal = 0 or
                    (@abnormal = 1 and (work_attend_proc_status = 'not processed' or off_work_attend_proc_status = 'not processed')))
         ),
         step2 as (
             select s1.*,
                    -- 實際需上班時間
                    dbo.FN_Convert_Datetime(date, work_start) as curr_work_start,
                    -- 實際需下班時間
                    dbo.FN_Convert_Datetime(iif(work_start > work_end, dateadd(day, 1, date), date), work_end) as curr_off_work_start
             from step1 s1
             join work_shift w on (s1.work_shift_id = w.id)
         ),
         step3 as (
             select *,
                    -- 上班補打卡時間
                    (select top 1 dbo.FN_Convert_Datetime(date, time)
                     from check_in_request_form f
                     join check_in_request_form_detail d on (f.id = d.check_in_request_form_id)
                     where f.employee_id = step2.employee_id and
                         f.deleted_at is null and
                         d.deleted_at is null and
                         f.sign_status = 1 and
                         check_in_type = 'work' and
                         dbo.FN_Convert_Datetime(d.date, d.time) between curr_work_start and curr_off_work_start
                     order by f.created_at desc, d.created_at desc) as 'new_check_in_work',
                    -- 下班打卡時間
                    (select top 1 dbo.FN_Convert_Datetime(date, time)
                     from check_in_request_form f
                     join check_in_request_form_detail d on (f.id = d.check_in_request_form_id)
                     where f.employee_id = step2.employee_id and
                         f.deleted_at is null and
                         d.deleted_at is null and
                         f.sign_status = 1 and
                         check_in_type = 'off work' and
                         dbo.FN_Convert_Datetime(d.date, d.time) between curr_work_start and curr_off_work_start
                     order by f.created_at desc, d.created_at desc) as 'new_check_in_off_work'
             from step2
         ),
         finally as (
             select id,
                    created_at,
                    updated_at,
                    deleted_at,
                    employee_id,
                    work_shift_id,
                    date,
                    iif(new_check_in_work is not null, new_check_in_work, work_check_in) as work_check_in,
                    work_attend_status,
                    work_attend_proc_status,
                    iif(new_check_in_off_work is not null, new_check_in_off_work, off_work_check_in) as off_work_check_in,
                    off_work_attend_status,
                    off_work_attend_proc_status,
                    absence_hours,
                    leave_hours,
                    sign_leave_hours,
                    overtime_hours,
                    sign_overtime_hours
             from step3
             union all
             select *
             from check_in_status
             where employee_id = @emp_id and
                   date between @start_date and @end_date and
                   deleted_at is null and
                   work_shift_id = 0 and
                   (@abnormal = 0 or
                    (@abnormal = 1 and (work_attend_proc_status = 'not processed' or off_work_attend_proc_status = 'not processed')))
         )
     insert into @result (id, created_at, updated_at, deleted_at, employee_id, work_shift_id, date, work_check_in, work_attend_status, work_attend_proc_status, off_work_check_in, off_work_attend_status, off_work_attend_proc_status, absence_hours, leave_hours, sign_leave_hours, overtime_hours, sign_overtime_hours)
     select * from finally;

    return;
end