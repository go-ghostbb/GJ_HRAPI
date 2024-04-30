drop function if exists [dbo].[FN_Q_OvertimeCheckIn]
go

create function [dbo].[FN_Q_OvertimeCheckIn] (@emp_id bigint, @start_date date, @end_date date)
    returns @result table (
        employee_id bigint,
        start datetime,
        [end] datetime
    )
begin

    -- 暫存cursor資料
    declare @temp_datetime datetime,
            @temp_check_in_type nvarchar(max)

    -- 定義cursor
    declare myCursor cursor for
    select dbo.FN_Convert_Datetime(date, time), check_in_type
    from overtime_check_in
    where employee_id = @emp_id and date between @start_date and @end_date
    order by dbo.FN_Convert_Datetime(date, time);

    -- 使用雙指標充當stack
    declare @left datetime;

    open myCursor;
    fetch next from myCursor into @temp_datetime, @temp_check_in_type;
    while (@@fetch_status <> -1)
    begin
        if (@temp_check_in_type = 'work')
        begin
            if (@left is not null)
            begin
                -- 如果不為null, 直接插入結果
                insert into @result (employee_id, start, [end])
                values (@emp_id, @left, null);
            end
            set @left = @temp_datetime;
        end
        else if (@temp_check_in_type = 'off work')
        begin
            if (@left is not null)
            begin
                insert into @result (employee_id, start, [end])
                values (@emp_id, @left, @temp_datetime);

                set @left = null;
            end
            else
            begin
                insert into @result (employee_id, start, [end])
                values (@emp_id, null, @temp_datetime);
            end
        end

        -- step
        fetch next from myCursor into @temp_datetime, @temp_check_in_type;
    end

    close myCursor;
    deallocate myCursor;
    return;
end