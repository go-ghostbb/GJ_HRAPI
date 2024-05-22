drop procedure if exists [dbo].[P_C_LeaveCorrectUpdateByFormSigning]
go

create procedure [dbo].[P_C_LeaveCorrectUpdateByFormSigning] (@leave_request_form_id bigint)
as
begin
    -- 查詢表單基本訊息
    declare @temp_start_date date,
            @temp_end_date date,
            @temp_start_time time(0),
            @temp_end_time time(0),
            @temp_leave_id bigint,
            @temp_emp_id bigint,
            @temp_sign_status tinyint;

    select @temp_start_date = start_date,
           @temp_end_date = end_date,
           @temp_start_time = start_time,
           @temp_end_time = end_time,
           @temp_leave_id = leave_id,
           @temp_emp_id = employee_id,
           @temp_sign_status = sign_status
    from leave_request_form
    where id = @leave_request_form_id;

    -- 查詢一天算幾小時
    declare @hoursADay int = (select convert(int, value) from config_map where [key] = 'HoursADay');

    -- 暫存最後要更新的表
    declare @final_map table (
        leave_request_form_id bigint,
        leave_correct_id bigint,
        day float
    );

    declare @final_leave_correct table (
        id bigint,
        available float,
        signing float,
        used float,
        minimum bigint
    );

    -- 遍歷時間
    while (@temp_start_date <= @temp_end_date)
    begin
        declare @temp_leave_correct table (
            id bigint,
            available float,
            signing float,
            used float,
            minimum bigint
        );

        -- 找出日期所套用的leave_correct
        insert into @temp_leave_correct (id, available, signing, used, minimum)
        select a.id, available, signing, used, b.minimum
        from leave_correct as a
        join leave as b on (a.leave_id = b.id)
        where a.deleted_at is null and b.deleted_at is null and
              @temp_start_date between effective and expired and
              employee_id = @temp_emp_id and
              leave_id = @temp_leave_id;

        -- 如果沒找到可用請假，回傳錯誤
        if ((select count(*) from @temp_leave_correct) = 0)
        begin
            raiserror('no leave available', 11, 1);
            return;
        end;

        declare @minimum bigint = (select top 1 minimum from @temp_leave_correct);

        -- 計算請假時間
        declare @start datetime = dbo.FN_Convert_Datetime(@temp_start_date, @temp_start_time),
                @end datetime = dbo.FN_Convert_Datetime(iif(@temp_start_time > @temp_end_time, dateadd(day, 1, @temp_start_date), @temp_start_date), @temp_end_time);

        declare @hour float = dbo.FN_C_WorkHoursCount(@temp_emp_id, @start, @end);
        declare @day float = iif(@minimum = 0, @hour, ceiling(@hour * 60 / @minimum) * @minimum) / 60 / @hoursADay;

        -- 遍歷@temp_leave_correct
        declare myCursor cursor for
        select id, available, signing, used from @temp_leave_correct;

        -- 暫存cursor
        declare @cur_id bigint,
                @cur_available float,
                @cur_signing float,
                @cur_used float;

        open myCursor;
        fetch next from myCursor into @cur_id, @cur_available, @cur_signing, @cur_used;
        while (@@fetch_status <> -1)
        begin
            if (@day <= 0)
            begin
                break;
            end;

            -- signing
            if (@temp_sign_status = 0 or @temp_sign_status = 3)
            begin
                if (@cur_signing + @cur_used + @day <= @cur_available)
                begin
                    -- 如果夠用
                    insert into @final_map (leave_request_form_id, leave_correct_id, day)
                    values (@leave_request_form_id, @cur_id, @day);

                    set @day = 0;
                end
                else
                begin
                    -- 如果不夠，放入一部分
                    declare @left float = @cur_available - (@cur_signing + @cur_used);

                    insert into @final_map (leave_request_form_id, leave_correct_id, day)
                    values (@leave_request_form_id, @cur_id, @left);

                    set @day = @day - @left;
                end;
            end;

            -- step
            fetch next from myCursor into @cur_id, @cur_available, @cur_signing, @cur_used;
        end;
        close myCursor;
        deallocate myCursor;

        -- 如果可用假別都用完
        if (@day > 0)
        begin
            -- 回報錯誤
            raiserror('the number of leave days exceeds the available days', 11, 1);
            return;
        end

        -- 如果能全部放完，丟進最後要更新的暫存表
        insert into @final_leave_correct
        select * from @temp_leave_correct;

        -- step
        set @temp_start_date = dateadd(day, 1, @temp_start_date);
    end;

    -- 使用try catch確認是否更新過
    -- 如果更新過不進行merge操作
    begin try
        insert into leave_request_form_correct_map (leave_request_form_id, leave_correct_id, day)
        select leave_request_form_id, leave_correct_id, sum(day) from @final_map group by leave_request_form_id, leave_correct_id;
    end try
    begin catch
        throw;
    end catch;

    declare updateCursor cursor for
    select leave_correct_id, sum(day) from @final_map group by leave_request_form_id, leave_correct_id;

    -- 暫存updateCursor
    declare @cur_leave_correct_id bigint,
            @cur_day float;

    open updateCursor;
    fetch next from updateCursor into @cur_leave_correct_id, @cur_day;
    while (@@fetch_status <> -1)
    begin
        update leave_correct
        set signing = signing + @cur_day
        where id = @cur_leave_correct_id;

        -- step
        fetch next from updateCursor into @cur_leave_correct_id, @cur_day;
    end;
    close updateCursor;
    deallocate updateCursor;
end