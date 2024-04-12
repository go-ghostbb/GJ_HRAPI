DROP PROCEDURE IF EXISTS [dbo].[P_C_CheckInHourUpdateByLeave]
GO

create PROCEDURE [dbo].[P_C_CheckInHourUpdateByLeave] (@employee_id bigint, @leave_request_from_id bigint)
as
begin

    -- 儲存form一些重要資訊
    declare @temp_deleted_at datetimeoffset;
    declare @temp_start_date date;
    declare @temp_end_date date;
    declare @temp_start_time time(0);
    declare @temp_end_time time(0);
    declare @temp_sign_status tinyint;
    declare @temp_leave_hour_count float;

    -- 查詢請假單內容
    select @temp_deleted_at = deleted_at,
           @temp_start_date = start_date,
           @temp_end_date = end_date,
           @temp_start_time = start_time,
           @temp_end_time = end_time,
           @temp_leave_hour_count = leave_hour_count,
           @temp_sign_status = sign_status
    from leave_request_form
    where id = @leave_request_from_id;

    -- 遍歷時間
    while @temp_start_date <= @temp_end_date
    begin
        declare @datetime_start datetime;
        declare @datetime_end datetime;

        if @temp_end_time < @temp_start_time
        begin
            -- 如果請假時間到隔天
            set @datetime_start = convert(datetime, convert(varchar, @temp_start_date) + ' ' + convert(varchar, @temp_start_time));
            set @datetime_end = convert(datetime, convert(varchar, dateadd(day, 1, @temp_start_date)) + ' ' + convert(varchar, @temp_end_time));
        end
        else
        begin
            -- 不隔夜，請假當天
            set @datetime_start = convert(datetime, convert(varchar, @temp_start_date) + ' ' + convert(varchar, @temp_start_time));
            set @datetime_end = convert(datetime, convert(varchar, @temp_start_date) + ' ' + convert(varchar, @temp_end_time));
        end;

        -- 如果為送簽中或是簽核中
        if @temp_sign_status in (0, 3) and @temp_deleted_at is null
        begin
            update check_in_status
            set sign_leave_hours = sign_leave_hours + @temp_leave_hour_count
            from check_in_status c
            join work_shift w on (c.work_shift_id = w.id)
            where c.employee_id = @employee_id and
                @datetime_start < convert(datetime, convert(varchar, off_work_check_in_date) + ' ' + convert(varchar, w.work_end)) and
                @datetime_end > convert(datetime, convert(varchar, work_check_in_date) + ' ' + convert(varchar, w.work_start));
        end;

        -- 如果為核准
        if @temp_sign_status = 1 and @temp_deleted_at is null
        begin
            update check_in_status
                -- 扣掉簽核中天數
            set sign_leave_hours = sign_leave_hours - @temp_leave_hour_count,
                -- 加上請假天數
                leave_hours = leave_hours + @temp_leave_hour_count,
                -- 扣掉缺勤時數
                absence_hours = absence_hours - @temp_leave_hour_count
            from check_in_status c
            join work_shift w on (c.work_shift_id = w.id)
            where c.employee_id = @employee_id and
                @datetime_start < convert(datetime, convert(varchar, off_work_check_in_date) + ' ' + convert(varchar, w.work_end)) and
                @datetime_end > convert(datetime, convert(varchar, work_check_in_date) + ' ' + convert(varchar, w.work_start));
        end;

        -- 如果為駁回或是單子被刪除
        if @temp_sign_status = 2 or (@temp_deleted_at is not null and @temp_sign_status != 1)
        begin
            update check_in_status
                -- 扣掉簽核中天數
            set sign_leave_hours = sign_leave_hours - @temp_leave_hour_count
            from check_in_status c
            join work_shift w on (c.work_shift_id = w.id)
            where c.employee_id = @employee_id and
                @datetime_start < convert(datetime, convert(varchar, off_work_check_in_date) + ' ' + convert(varchar, w.work_end)) and
                @datetime_end > convert(datetime, convert(varchar, work_check_in_date) + ' ' + convert(varchar, w.work_start));
        end

        -- 如果簽核完畢被刪除
        if @temp_sign_status = 1 and @temp_deleted_at is not null
        begin
            update check_in_status
                -- 扣掉請假時數
            set leave_hours = leave_hours - @temp_leave_hour_count,
                -- 加上缺勤時數
                absence_hours = absence_hours + @temp_leave_hour_count
            from check_in_status c
            join work_shift w on (c.work_shift_id = w.id)
            where c.employee_id = @employee_id and
                @datetime_start < convert(datetime, convert(varchar, off_work_check_in_date) + ' ' + convert(varchar, w.work_end)) and
                @datetime_end > convert(datetime, convert(varchar, work_check_in_date) + ' ' + convert(varchar, w.work_start));
        end;

        -- 更新迴圈判定的變數
        set @temp_start_date = dateadd(day, 1, @temp_start_date)
    end;
end
go