-- 在更新請假單時，出勤時數簽核時數扣掉原本請假單的時數
-- 只有更新時，才會觸發此procedure

DROP PROCEDURE IF EXISTS [dbo].[P_C_CheckInHourSubByLeave]
GO

create PROCEDURE [dbo].[P_C_CheckInHourSubByLeave] (@employee_id bigint, @leave_request_from_id bigint)
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

            -- 能更新的時候只有簽核中和送簽中
            -- 所以扣除簽核中時數
            update check_in_status
            set sign_leave_hours = sign_leave_hours - @temp_leave_hour_count
            from check_in_status c
            join work_shift w on (c.work_shift_id = w.id)
            where c.employee_id = @employee_id and
                @datetime_start < convert(datetime, convert(varchar, off_work_check_in_date) + ' ' + convert(varchar, w.work_end)) and
                @datetime_end > convert(datetime, convert(varchar, work_check_in_date) + ' ' + convert(varchar, w.work_start));

            -- 更新迴圈判定的變數
            set @temp_start_date = dateadd(day, 1, @temp_start_date)
        end;
end
go