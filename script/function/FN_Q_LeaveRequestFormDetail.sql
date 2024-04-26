drop function if exists [dbo].[FN_Q_LeaveRequestFormDetail]
go

create function [dbo].[FN_Q_LeaveRequestFormDetail] (@emp_id bigint, @start datetime, @end datetime)
    returns @result table (
        id bigint,
        sign_status tinyint,
        employee_id bigint,
        start datetime,
        [end] datetime,
        leave_hours_count float
    )
as
begin
    -- 根據員工id, 時間區間, 判斷有哪些時間是請假
    -- @emp_id 員工id
    -- @start 時間區間(開始)
    -- @end 時間區間(結束)

    ;with
         step1 as (
             select id,
                    start_date,
                    dbo.FN_Convert_Datetime(start_date, start_time) as start,
                    dbo.FN_Convert_Datetime(iif(start_time > end_time, dateadd(day, 1, start_date), start_date), end_time) as [end]
             from leave_request_form where employee_id = @emp_id and deleted_at is null
             union all
             select b.id,
                    dateadd(day, 1, a.start_date),
                    dbo.FN_Convert_Datetime(dateadd(day, 1, a.start_date), b.start_time) as start,
                    dbo.FN_Convert_Datetime(iif(b.start_time > b.end_time, dateadd(day, 2, a.start_date), dateadd(day, 1, a.start_date)), b.end_time) as [end]
             from step1 as a
             inner join leave_request_form as b on (a.id = b.id and dateadd(day, 1, a.start_date) <= b.end_date)
         ),
         step2 as (
             select a.id,
                    b.sign_status,
                    b.employee_id,
                    a.start,
                    a.[end],
                    dbo.FN_C_WorkHoursCount(b.employee_id, a.start, a.[end]) as hours_count
             from step1 as a
             join leave_request_form as b on (a.id = b.id)
         )
    insert into @result (id, sign_status, employee_id, start, [end], leave_hours_count)
    select * from step2 where start <= @end and @start <= [end];

    return;
end