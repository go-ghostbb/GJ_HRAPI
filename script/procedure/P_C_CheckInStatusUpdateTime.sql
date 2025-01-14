
drop procedure if exists [dbo].[P_C_CheckInStatusUpdateTime]
go

create procedure [dbo].[P_C_CheckInStatusUpdateTime] (@datetime datetime, @card_number nvarchar(191), @type nvarchar(20))
as
begin
    -- 匯入打卡資料
    -- @datetime 日期時間
    -- @card_number 卡號
    -- @type 類型: 上班或下班

    -- 查詢一天算幾小時
    -- 預設不插入大於一天的資料
    declare @hoursADay float = convert(float, (select value from config_map where [key] = 'HoursADay'));

    if (@type = 'work')
    begin
        -- 上班
        update check_in_status
        set work_check_in = @datetime
        where -- 抓取最接近的上班的時間
            id = (select top 1 i_c.id
                  from check_in_status i_c
                  join work_shift i_w on (i_c.work_shift_id = i_w.id)
                  where employee_id = (select id from employee where card_number = @card_number and deleted_at is null) and
                        abs(datediff(second, @datetime, dbo.FN_Convert_Datetime(i_c.date, i_w.work_start))) / 60 / 60 <= @hoursADay and
                        i_c.deleted_at is null
                  order by abs(datediff(second, @datetime, dbo.FN_Convert_Datetime(i_c.date, i_w.work_start))));
    end
    else if (@type = 'off work')
    begin
        -- 下班
        update check_in_status
        set off_work_check_in = @datetime
        where -- 抓取最接近的下班的時間
            id = (select top 1 i_c.id
                  from check_in_status i_c
                  join work_shift i_w on (i_c.work_shift_id = i_w.id)
                  where employee_id = (select id from employee where card_number = @card_number and deleted_at is null) and
                        abs(datediff(second, @datetime, dbo.FN_Convert_Datetime(iif(i_w.work_start > i_w.work_end, dateadd(day, 1, i_c.date), i_c.date), i_w.work_end))) / 60 / 60 <= @hoursADay and
                        i_c.deleted_at is null
                  order by abs(datediff(second, @datetime, dbo.FN_Convert_Datetime(iif(i_w.work_start > i_w.work_end, dateadd(day, 1, i_c.date), i_c.date), i_w.work_end))));
    end;

    -- 新增至打卡資料表
    insert into check_in_data (created_at, updated_at, deleted_at, employee_id, datetime, type, card_number)
    values (sysdatetimeoffset(), sysdatetimeoffset(), null, (select id from employee where card_number = @card_number and deleted_at is null), @datetime, @type, @card_number);
    
    declare @date date = convert(date, @datetime);
    declare @start_date date = dateadd(day, -1, @date);
    declare @end_date date = dateadd(day, 1, @date);
    declare @emp_id bigint = (select id from employee where card_number = @card_number and deleted_at is null);

    exec P_C_CheckInStatusUpdateStatus
        @start_date , @end_date, @emp_id;
end