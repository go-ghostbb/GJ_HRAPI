drop procedure if exists [dbo].[P_C_LeaveCorrectUpdateCount]
go

create procedure [dbo].[P_C_LeaveCorrectUpdateCount] (@leave_correct_id bigint)
as
begin
    -- 紀錄@leave_correct_id的生效時間和過期時間
    declare @temp_effective date,
            @temp_expired date;

    -- 紀錄對應的員工
    declare @temp_emp_id bigint;

    -- 紀錄對應的假別
    declare @temp_leave_id bigint;

    -- 紀錄可用天數
    declare @temp_available tinyint;

    select @temp_emp_id = employee_id,
           @temp_leave_id = leave_id,
           @temp_available = available,
           @temp_effective = effective,
           @temp_expired = expired
    from leave_correct
    where id = @leave_correct_id and deleted_at is null;

    -- 查詢一天算幾小時
    declare @hoursADay int = (select convert(int, value) from config_map where [key] = 'HoursADay');

    -- 紀錄@leave_correct_id區間內所有請假單
    declare @leave_history table (
        sign_status tinyint,
        leave_days_count float
    )

    insert into @leave_history (sign_status, leave_days_count)
    select sign_status, sum(leave_hours_count) / @hoursADay
    from FN_Q_LeaveRequestFormDetail(@temp_emp_id, @temp_effective, @temp_expired)
    where sign_status in (0, 1, 3) and leave_id = @temp_leave_id
    group by sign_status, leave_id;

    -- 紀錄簽核中和核准的總天數
    declare @total_signing float = (select leave_days_count from @leave_history where sign_status in (0, 3)),
            @total_used float = (select leave_days_count from @leave_history where sign_status = 1);

    -- 將null改為0
    set @total_signing = iif(@total_signing is null, 0, @total_signing);
    set @total_used = iif(@total_used is null, 0, @total_used);

    if (@temp_available >= @total_signing + @total_used)
    begin
        -- 如果沒超過可用天數，進行更新
        update leave_correct
        set signing = @total_signing,
            used = @total_used
        where id = @leave_correct_id;
    end
    else
    begin
        -- 否則拋出錯誤
        raiserror('the number of leave hours exceeds the available hours', 11, 1);
    end
end