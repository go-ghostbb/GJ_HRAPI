drop procedure if exists [dbo].[P_C_CheckInStatusUpdateStatus]
go

create procedure [dbo].[P_C_CheckInStatusUpdateStatus] (@start_date date, @end_date date, @employee_id varchar(max))
as
begin
    -- 按照請假單, 加班單, 補打卡單更新出勤狀態表
    -- @start_date 時間區間(開始)
    -- @end_date 時間區間(結束)
    -- @employee_id 員工id, 多筆用","隔開 (ex:"1,2,4,5")

    -- 暫存表
    declare @temp_table table (
        id bigint,
        employee_id bigint,
        -- 實際需上班時間
        curr_work_start datetime,
        -- 實際需下班時間
        curr_work_end datetime,
        -- 休息開始時間
        rest_start datetime,
        -- 休息結束時間
        rest_end datetime,
        -- 上班打卡時間
        check_in_work datetime,
        -- 下班打卡時間
        check_in_off_work datetime,
        -- 缺勤時數
        absence_hours float,
        -- 請假時數
        leave_hours float,
        -- 簽核中請假時數
        sign_leave_hours float,
        -- 上班狀態
        work_attend_status nvarchar(max),
        -- 上班處理狀態
        work_attend_proc_status nvarchar(max),
        -- 下班狀態
        off_work_attend_status nvarchar(max),
        -- 下班處理狀態
        off_work_attend_proc_status nvarchar(max)
    )

    -- 查詢基本訊息(id, employee_id, 打卡時間等...)
    ;with
         -- step.1 查詢基本內容，並且過濾不必要資料
         step1 as (
             select c.id, c.employee_id, c.date, c.work_check_in, c.off_work_check_in, w.work_start, w.work_end, w.rest_start, w.rest_end
             from check_in_status c
             -- 這邊會自動過濾work_shift_id = 0的問題(休息日)
             join work_shift w on (c.work_shift_id = w.id)
             where c.date between @start_date and @end_date and c.deleted_at is null and w.deleted_at is null and
                   employee_id in (select convert(bigint, value) from string_split(@employee_id, ','))
         ),
         -- step.2 計算實際需上下班時間
         step2 as (
             select id, employee_id,
                    -- 只計算到分鐘，秒自動捨去
                    dateadd(minute, datediff(minute, 0, work_check_in), 0) as work_check_in,
                    dateadd(minute, datediff(minute, 0, off_work_check_in), 0) as off_work_check_in,
                    -- 實際需上班時間
                    dbo.FN_Convert_Datetime(date, work_start) as curr_work_start,
                    -- 實際需下班時間
                    dbo.FN_Convert_Datetime(iif(work_start > work_end, dateadd(day, 1, date), date), work_end) as curr_off_work_start,
                    -- 休息開始時間
                    dbo.FN_Convert_Datetime(iif(work_start > rest_start, dateadd(day, 1, date), date), rest_start) as rest_start,
                    -- 休息結束時間
                    dbo.FN_Convert_Datetime(iif(work_start > rest_end, dateadd(day, 1, date), date), rest_end) as rest_end
             from step1
         ),
         -- step.3 查詢是否有補打卡
         step3 as (
             select *,
                    -- 上班補打卡時間
                    (select top 1 dateadd(minute, datediff(minute, 0, dbo.FN_Convert_Datetime(date, time)), 0)
                     from check_in_request_form f
                     join check_in_request_form_detail d on (f.id = d.check_in_request_form_id)
                     where f.employee_id = step2.employee_id and
                         f.deleted_at is null and
                         d.deleted_at is null and
                         f.sign_status = 1 and
                         check_in_type = 'work' and
                         dbo.FN_Convert_Datetime(d.date, d.time) between curr_work_start and curr_off_work_start
                     order by f.created_at desc, d.created_at desc) as 'new_check_in_work',
                    -- 下班打卡時間
                    (select top 1 dateadd(minute, datediff(minute, 0, dbo.FN_Convert_Datetime(date, time)), 0)
                     from check_in_request_form f
                     join check_in_request_form_detail d on (f.id = d.check_in_request_form_id)
                     where f.employee_id = step2.employee_id and
                         f.deleted_at is null and
                         d.deleted_at is null and
                         f.sign_status = 1 and
                         check_in_type = 'off work' and
                         dbo.FN_Convert_Datetime(d.date, d.time) between curr_work_start and curr_off_work_start
                     order by f.created_at desc, d.created_at desc) as 'new_check_in_off_work'
             from step2
         ),
         -- finally 回傳最終結果
         finally as (
             select id, employee_id, curr_work_start, curr_off_work_start, rest_start, rest_end,
                    -- 上班時間，有補打卡優先使用補打卡資料
                    iif(new_check_in_work is not null, new_check_in_work, work_check_in) as work_check_in,
                    -- 下班時間，有補打卡優先使用補打卡資料
                    iif(new_check_in_off_work is not null, new_check_in_off_work, off_work_check_in) as off_work_check_in
             from step3
         )
     insert into @temp_table (id, employee_id, curr_work_start, curr_work_end, rest_start, rest_end, check_in_work, check_in_off_work)
     select * from finally;

    -- 更新狀態
    update @temp_table
    set
        work_attend_status = iif(
            check_in_work is null,
            -- 如果上班沒打卡
            'not swiped',
            -- 如果有打卡
            iif(check_in_work > curr_work_start, 'late', 'normal')
        ),
        off_work_attend_status = iif(
            check_in_off_work is null,
            -- 如果上班沒打卡
            'not swiped',
            -- 如果有打卡
            iif(check_in_off_work < curr_work_end, 'early', 'normal')
        ),
        work_attend_proc_status = iif(
            check_in_work is null or check_in_work > curr_work_start,
            'not processed',
            'normal'
        ),
        off_work_attend_proc_status = iif(
            check_in_off_work is null or check_in_off_work < curr_work_end,
            'not processed',
            'normal'
        )
    where id != 0;

    -- 更新請假單
    update @temp_table
    set
        leave_hours = (
            select iif(sum(leave_hours_count) is not null, sum(leave_hours_count), 0)
            from FN_Q_LeaveRequestFormDetail(employee_id, curr_work_start, curr_work_end)
            where sign_status = 1
        ),
        sign_leave_hours = (
            select iif(sum(leave_hours_count) is not null, sum(leave_hours_count), 0)
            from FN_Q_LeaveRequestFormDetail(employee_id, curr_work_start, curr_work_end)
            where sign_status in (0, 3)
        )
    where id != 0;

    -- 更新缺勤時數
    update @temp_table
    set absence_hours = iif(
        check_in_work is null or check_in_off_work is null,
        -- 如果上下班其中一個沒打卡，缺勤時數=整天
        -- 下班-上班 - (休息結束-休息開始)
        dbo.FN_C_WorkHoursCount(employee_id, curr_work_start, curr_work_end),
        -- 如果都有打卡
        -- result:遲到時間+早退時間
        iif(work_attend_status = 'late', dbo.FN_C_WorkHoursCount(employee_id, curr_work_start, check_in_work), 0)
        +
        iif(off_work_attend_status = 'early', dbo.FN_C_WorkHoursCount(employee_id, check_in_off_work, curr_work_end), 0)
    ) - leave_hours
    where id != 0;

    -- 更新處理狀態
    update @temp_table
    set
        work_attend_proc_status = iif(work_attend_status != 'normal' and absence_hours <= 0, 'processed', work_attend_proc_status),
        off_work_attend_proc_status = iif(off_work_attend_status != 'normal' and absence_hours <= 0, 'processed', off_work_attend_proc_status)
    where id != 0;

    print @@rowcount;

    -- 合併
    merge check_in_status as T using @temp_table S on (T.id = S.id)
    when matched then update set
        T.updated_at = getdate(),
        T.work_attend_status = S.work_attend_status,
        T.work_attend_proc_status = S.work_attend_proc_status,
        T.off_work_attend_status = S.off_work_attend_status,
        T.off_work_attend_proc_status = S.off_work_attend_proc_status,
        T.absence_hours = S.absence_hours,
        T.leave_hours = S.leave_hours,
        T.sign_leave_hours = S.sign_leave_hours;
end
go
