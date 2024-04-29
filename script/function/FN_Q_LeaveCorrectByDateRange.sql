drop function if exists [dbo].[FN_Q_LeaveCorrectByDateRange]
go

create function [dbo].[FN_Q_LeaveCorrectByDateRange] (@emp_id bigint, @leave_id bigint, @start date, @end date)
    returns @result table (
        date date,
        leave_correct_id bigint
    )
begin
    declare @temp_start date = @start;

    while (@temp_start <= @end)
    begin
        insert into @result (date, leave_correct_id)
        values (
            @temp_start,
            (select id
             from leave_correct
             where employee_id = @emp_id and
                   leave_id = @leave_id and
                   @temp_start between effective and expired and
                   deleted_at is null)
        );

        -- step (1 day)
        set @temp_start = dateadd(day, 1, @temp_start);
    end;

    return;
end