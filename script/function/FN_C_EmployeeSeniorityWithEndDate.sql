DROP FUNCTION IF EXISTS [dbo].[FN_C_EmployeeSeniorityWithEndDate]
    GO

CREATE function [dbo].[FN_C_EmployeeSeniorityWithEndDate] (@endDate smalldatetime)
    returns @retDateDuration table
                             (
                                 employee_id     int  null ,
                                 hire_date     smalldatetime null ,
                                 year        int  null ,
                                 month       int  null
                             )
as
begin
    -- 定義Cursor並打開
    declare idCursor cursor for select id, hire_date from employee
    open idCursor

    declare @id bigint
    declare @arrived smalldatetime

    -- 迴圈跑Cursor
    fetch next from idCursor into @id, @arrived
    while (@@fetch_status <> -1)
        begin
            declare @yearDiff int
            declare @monthDiff int

            -- 計算年資
            select @yearDiff = DATEDIFF(year, @arrived, @endDate)
            select @monthDiff = DATEDIFF(month, @arrived, @endDate)

            -- 未滿年資修正
            if ( DATEDIFF(day , DATEADD(year , @yearDiff, @arrived), @endDate) < 0 )
                set @yearDiff = @yearDiff - 1

            if ( DATEDIFF(day , DATEADD(month , @monthDiff, @arrived), @endDate) < 0 )
                set @monthDiff = @monthDiff -1

            -- 扣除年資後的月數
            set @monthDiff = @monthDiff - 12 * @yearDiff

            -- 回傳年資(id,年,月)
            insert @retDateDuration( employee_id , hire_date , year, month )
            select @id , @arrived , IIF( @yearDiff < 0, 0, @yearDiff ), IIF( @yearDiff < 0 OR @monthDiff < 0, 0, @monthDiff )

            -- 抓取下一個
            fetch next from idCursor into @id, @arrived
        end

    close idCursor
    deallocate idCursor

    return
end
go

