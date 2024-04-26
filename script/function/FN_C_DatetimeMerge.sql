drop function if exists [dbo].[FN_C_DatetimeMerge]
go

create function [dbo].[FN_C_DatetimeMerge] (@datetimes varchar(max))
    returns @result table(
        start_datetime datetime,
        end_datetime datetime
    )
as
begin
    -- test:2024-01-01 08:00:00~2024-01-01 17:00:00,2024-01-02 08:00:00~2024-01-02 17:00:00
    if (@datetimes = '')
    begin
        return
    end

    -- 暫存參數切割好的資料
    declare @temp_result table (
        start_datetime datetime,
        end_datetime datetime
    )

    -- 開啟cursor遍歷切割完的字串
    declare myCursor cursor for
    select value from string_split(@datetimes, ',');

    -- cursor暫存變數
    declare @temp_range varchar(max);
    declare @temp_start varchar(80);
    declare @temp_end varchar(80);

    open myCursor;
    fetch next from myCursor into @temp_range;
    -- 開始遍歷
    while (@@fetch_status <> -1)
    begin
        -- 切割2024-01-01 08:00:00~2024-01-01 17:00:00
        set @temp_start = (select value from string_split(@temp_range, '~', 1) where ordinal = 1)
        set @temp_end = (select value from string_split(@temp_range, '~', 1) where ordinal = 2)
        insert into @temp_result (start_datetime, end_datetime)
        values (@temp_start, @temp_end);

        fetch next from myCursor into @temp_range;
    end;

    -- 關閉&釋放cursor
    close myCursor;
    deallocate myCursor;

    -- 記錄總共有多少筆資料
    declare @len_temp_result int = (select count(*) from @temp_result);
    -- 開始合併
    declare @right_pointer int = 1;
    -- 記錄left指針
    declare @current_start datetime,
            @current_end datetime;
    -- 先取第一筆資料
    select
        @current_start = start_datetime, -- A1
        @current_end = end_datetime      -- A2
    from @temp_result order by start_datetime offset 0 row fetch next 1 rows only;

    -- 開始遍歷(演算法開始)
    while (@right_pointer < @len_temp_result)
    begin
        -- 紀錄right指針
        declare @temp_d_start datetime,
                @temp_d_end datetime;
        select
            @temp_d_start = start_datetime, -- B1
            @temp_d_end = end_datetime      -- B2
        from @temp_result order by start_datetime offset @right_pointer row fetch next 1 rows only;

        if (@current_start <= @temp_d_end and @temp_d_start <= @current_end)
        begin
            -- 如果有重疊
            set @current_start = iif(@current_start > @temp_d_start, @temp_d_start, @current_start);
            set @current_end = iif(@current_end < @temp_d_end, @temp_d_end, @current_end);
        end
        else
        begin
            -- 如果沒疊
            -- 將當前資料新增至@reulst
            insert into @result (start_datetime, end_datetime)
            values (@current_start, @current_end);

            -- 更新當前資料
            set @current_start = @temp_d_start;
            set @current_end = @temp_d_end;
        end

        -- select @right_pointer;
        -- step
        set @right_pointer = @right_pointer + 1;
    end

    -- 將最後結果新增至@result
    insert into @result (start_datetime, end_datetime)
    values (@current_start, @current_end);

    return
end
go