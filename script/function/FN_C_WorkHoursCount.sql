drop function if exists [dbo].[FN_C_WorkHoursCount]
go

create function [dbo].[FN_C_WorkHoursCount] (@emp_id bigint, @start datetime, @end datetime)
    returns float
as
begin
    declare @result float = 0;

    -- 紀錄需要上班的時段
    declare @curr_work_start datetime,
            @curr_work_end datetime,
            @rest_start datetime,
            @rest_end datetime;

    select @curr_work_start = dbo.FN_Convert_Datetime(schedule_date, work_start),
           @curr_work_end = dbo.FN_Convert_Datetime(iif(work_start > work_end, dateadd(day, 1, schedule_date), schedule_date), work_end),
           @rest_start = dbo.FN_Convert_Datetime(iif(work_start > rest_start, dateadd(day, 1, schedule_date), schedule_date), rest_start),
           @rest_end = dbo.FN_Convert_Datetime(iif(work_start > rest_end, dateadd(day, 1, schedule_date), schedule_date), rest_end)
    from work_schedule schedule
    join work_shift shift on (schedule.work_shift_id = shift.id)
    where employee_id = @emp_id and
          @start <= dbo.FN_Convert_Datetime(iif(work_start > work_end, dateadd(day, 1, schedule_date), schedule_date), work_end) and
          dbo.FN_Convert_Datetime(schedule_date, work_start) <= @end;


    declare @overlap_start datetime,
            @overlap_end datetime;

    -- @curr_work_start -> @rest_start
    if (@curr_work_start <= @end and @start <= @rest_start)
    begin
        set @overlap_start = iif(@curr_work_start > @start, @curr_work_start, @start);
        set @overlap_end = iif(@rest_start < @end, @rest_start, @end);

        set @result = datediff(second, @overlap_start, @overlap_end);
    end

    -- @rest_end -> @curr_work_end
    if (@rest_end <= @end and @start <= @curr_work_end)
    begin
        set @overlap_start = iif(@rest_end > @start, @rest_end, @start);
        set @overlap_end = iif(@curr_work_end < @end, @curr_work_end, @end);

        set @result = @result + datediff(second, @overlap_start, @overlap_end);
    end

    return @result / 60 / 60;
end