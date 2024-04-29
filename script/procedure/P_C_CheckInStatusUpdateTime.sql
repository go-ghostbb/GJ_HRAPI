
drop procedure if exists [dbo].[P_C_CheckInStatusUpdateTime]
go

create procedure [dbo].[P_C_CheckInStatusUpdateTime] (@datetime datetime, @work_shift_code nvarchar(100), @card_number nvarchar(191), @isWork bit)
as
begin
    -- 匯入打卡資料
    -- @datetime 日期時間
    -- @work_shift_code 班別代碼
    -- @card_number 卡號
    -- @isWork 上班或下班

    if (@isWork = 1)
    begin
        -- 上班
        update check_in_status
        set work_check_in = @datetime
        from check_in_status c
        join work_shift w on (c.work_shift_id = w.id)
        where work_shift_id = (select id from work_shift where code = @work_shift_code and deleted_at is null) and
              employee_id = (select id from employee where card_number = @card_number and deleted_at is null) and
              -- 只更新時間有在實際上下班區間內
              @datetime between dbo.FN_Convert_Datetime(date, w.work_start) and
                                dbo.FN_Convert_Datetime(iif(w.work_start > w.work_end, dateadd(day, 1, date), date), w.work_end);

        if (@@rowcount = 0)
        begin
            -- 如果沒更新到
            -- 抓取最接近的上班的時間
            update check_in_status
            set work_check_in = @datetime
            where -- 抓取最接近的上班的時間
                  id = (select top 1 i_c.id
                        from check_in_status i_c
                        join work_shift i_w on (i_c.work_shift_id = i_w.id)
                        where work_shift_id = (select id from work_shift where code = @work_shift_code and deleted_at is null) and
                              employee_id = (select id from employee where card_number = @card_number and deleted_at is null) and
                              i_c.deleted_at is null
                        order by abs(datediff(second, @datetime, dbo.FN_Convert_Datetime(i_c.date, i_w.work_start))))
        end
    end
    else
    begin
        -- 下班
        update check_in_status
        set off_work_check_in = @datetime
        from check_in_status c
        join work_shift w on (c.work_shift_id = w.id)
        where work_shift_id = (select id from work_shift where code = @work_shift_code and deleted_at is null) and
              employee_id = (select id from employee where card_number = @card_number and deleted_at is null) and
              -- 只更新時間有在實際上下班區間內
              @datetime between dbo.FN_Convert_Datetime(date, w.work_start) and
                                dbo.FN_Convert_Datetime(iif(w.work_start > w.work_end, dateadd(day, 1, date), date), w.work_end);

        if (@@rowcount = 0)
            begin
                -- 如果沒更新到
                -- 抓取最接近的上班的時間
                update check_in_status
                set off_work_check_in = @datetime
                where -- 抓取最接近的上班的時間
                      id = (select top 1 i_c.id
                            from check_in_status i_c
                            join work_shift i_w on (i_c.work_shift_id = i_w.id)
                            where work_shift_id = (select id from work_shift where code = @work_shift_code and deleted_at is null) and
                                  employee_id = (select id from employee where card_number = @card_number and deleted_at is null) and
                                  i_c.deleted_at is null
                            order by abs(datediff(second, @datetime, dbo.FN_Convert_Datetime(iif(i_w.work_start > i_w.work_end, dateadd(day, 1, i_c.date), i_c.date), i_w.work_end))))
            end
    end
    
    declare @date date = convert(date, @datetime);
    declare @start_date date = dateadd(day, -1, @date);
    declare @end_date date = dateadd(day, 1, @date);
    declare @emp_id bigint = (select id from employee where card_number = @card_number and deleted_at is null);

    exec P_C_CheckInStatusUpdateStatus
        @start_date , @end_date, @emp_id;
end