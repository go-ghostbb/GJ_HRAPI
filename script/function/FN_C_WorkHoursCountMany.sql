drop function if exists [dbo].[FN_C_WorkHoursCountMany]
go

create function [dbo].[FN_C_WorkHoursCountMany] (@emp_id bigint, @start_date date, @end_date date, @start_time time(0), @end_time time(0))
    returns float
as
begin
    declare @temp_start_date date = @start_date;
    declare @result float = 0;

    while (@temp_start_date <= @end_date)
    begin
        set @result =
            @result
            +
            dbo.FN_C_WorkHoursCount(
                @emp_id,
                dbo.FN_Convert_Datetime(@temp_start_date, @start_time),
                dbo.FN_Convert_Datetime(iif(@start_time > @end_time, dateadd(day, 1, @temp_start_date), @temp_start_date), @end_time)
            )

        -- step(1 day)
        set @temp_start_date = dateadd(day, 1, @temp_start_date);
    end;

    return @result;
end