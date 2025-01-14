USE [hr]
GO
ALTER TABLE [dbo].[work_shift] DROP CONSTRAINT [DF__work_shif__last___56E8E7AB]
GO
ALTER TABLE [dbo].[work_shift] DROP CONSTRAINT [DF__work_shif__statu__55F4C372]
GO
ALTER TABLE [dbo].[work_schedule] DROP CONSTRAINT [DF__work_sche__work___1F63A897]
GO
ALTER TABLE [dbo].[work_schedule] DROP CONSTRAINT [DF__work_sche__user___1E6F845E]
GO
ALTER TABLE [dbo].[work_schedule] DROP CONSTRAINT [DF__work_schedu__day__1D7B6025]
GO
ALTER TABLE [dbo].[work_schedule] DROP CONSTRAINT [DF__work_sche__month__1C873BEC]
GO
ALTER TABLE [dbo].[work_schedule] DROP CONSTRAINT [DF__work_sched__year__1B9317B3]
GO
ALTER TABLE [dbo].[vacations] DROP CONSTRAINT [DF__vacations__last___339FAB6E]
GO
ALTER TABLE [dbo].[vacations] DROP CONSTRAINT [DF__vacations__day__32AB8735]
GO
ALTER TABLE [dbo].[vacations] DROP CONSTRAINT [DF__vacations__month__31B762FC]
GO
ALTER TABLE [dbo].[vacations] DROP CONSTRAINT [DF__vacations__year__30C33EC3]
GO
ALTER TABLE [dbo].[vacations] DROP CONSTRAINT [DF__vacations__type___2FCF1A8A]
GO
ALTER TABLE [dbo].[vacation_types] DROP CONSTRAINT [DF__vacation___last___02FC7413]
GO
ALTER TABLE [dbo].[vacation_types] DROP CONSTRAINT [DF__vacation___statu__02084FDA]
GO
ALTER TABLE [dbo].[users] DROP CONSTRAINT [DF__users__backend__7849DB76]
GO
ALTER TABLE [dbo].[users] DROP CONSTRAINT [DF__users__last_move__68487DD7]
GO
ALTER TABLE [dbo].[users] DROP CONSTRAINT [DF__users__grade_id__5DCAEF64]
GO
ALTER TABLE [dbo].[users] DROP CONSTRAINT [DF__users__rank_id__5CD6CB2B]
GO
ALTER TABLE [dbo].[users] DROP CONSTRAINT [DF__users__status__4316F928]
GO
ALTER TABLE [dbo].[roles] DROP CONSTRAINT [DF__roles__last_move__693CA210]
GO
ALTER TABLE [dbo].[roles] DROP CONSTRAINT [DF__roles__status__46E78A0C]
GO
ALTER TABLE [dbo].[ranks] DROP CONSTRAINT [DF__ranks__last_move__66603565]
GO
ALTER TABLE [dbo].[ranks] DROP CONSTRAINT [DF__ranks__status__5812160E]
GO
ALTER TABLE [dbo].[permissions] DROP CONSTRAINT [DF__permissio__last___74AE54BC]
GO
ALTER TABLE [dbo].[permissions] DROP CONSTRAINT [DF__permissio__statu__73BA3083]
GO
ALTER TABLE [dbo].[overtime_request_form] DROP CONSTRAINT [DF__overtime___depar__66EA454A]
GO
ALTER TABLE [dbo].[overtime_request_form] DROP CONSTRAINT [DF__overtime___user___65F62111]
GO
ALTER TABLE [dbo].[overtime_request_form] DROP CONSTRAINT [DF__overtime___sign___6501FCD8]
GO
ALTER TABLE [dbo].[overtime_rate] DROP CONSTRAINT [DF__overtime___last___6166761E]
GO
ALTER TABLE [dbo].[menus] DROP CONSTRAINT [DF__menus__show__793DFFAF]
GO
ALTER TABLE [dbo].[menus] DROP CONSTRAINT [DF__menus__last_move__6A30C649]
GO
ALTER TABLE [dbo].[menus] DROP CONSTRAINT [DF__menus__status__5441852A]
GO
ALTER TABLE [dbo].[menus] DROP CONSTRAINT [DF__menus__sort__534D60F1]
GO
ALTER TABLE [dbo].[menus] DROP CONSTRAINT [DF__menus__parent_id__52593CB8]
GO
ALTER TABLE [dbo].[leave_request_form] DROP CONSTRAINT [DF__leave_req__proxy__5006DFF2]
GO
ALTER TABLE [dbo].[leave_request_form] DROP CONSTRAINT [DF__leave_req__proxy__4F12BBB9]
GO
ALTER TABLE [dbo].[leave_request_form] DROP CONSTRAINT [DF__leave_req__depar__4E1E9780]
GO
ALTER TABLE [dbo].[leave_request_form] DROP CONSTRAINT [DF__leave_req__user___4D2A7347]
GO
ALTER TABLE [dbo].[leave_request_form] DROP CONSTRAINT [DF__leave_req__leave__4C364F0E]
GO
ALTER TABLE [dbo].[leave_request_form] DROP CONSTRAINT [DF__leave_req__sign___4B422AD5]
GO
ALTER TABLE [dbo].[leave_defer] DROP CONSTRAINT [DF__leave_def__next___29AC2CE0]
GO
ALTER TABLE [dbo].[leave_defer] DROP CONSTRAINT [DF__leave_def__previ__28B808A7]
GO
ALTER TABLE [dbo].[leave_defer] DROP CONSTRAINT [DF__leave_def__leave__27C3E46E]
GO
ALTER TABLE [dbo].[leave_defer] DROP CONSTRAINT [DF__leave_def__user___26CFC035]
GO
ALTER TABLE [dbo].[leave_correct] DROP CONSTRAINT [DF__leave_cor__leave__23F3538A]
GO
ALTER TABLE [dbo].[leave_correct] DROP CONSTRAINT [DF__leave_cor__user___22FF2F51]
GO
ALTER TABLE [dbo].[leave_core] DROP CONSTRAINT [DF__leave_cor__is_ma__0C1BC9F9]
GO
ALTER TABLE [dbo].[leave_core] DROP CONSTRAINT [DF__leave_cor__matur__047AA831]
GO
ALTER TABLE [dbo].[leave_core] DROP CONSTRAINT [DF__leave_cor__cycle__038683F8]
GO
ALTER TABLE [dbo].[leave_core] DROP CONSTRAINT [DF__leave_cor__excha__00AA174D]
GO
ALTER TABLE [dbo].[leave_core] DROP CONSTRAINT [DF__leave_cor__defer__7FB5F314]
GO
ALTER TABLE [dbo].[leave_core] DROP CONSTRAINT [DF__leave_cor__last___5E8A0973]
GO
ALTER TABLE [dbo].[leave_core] DROP CONSTRAINT [DF__leave_cor__statu__5D95E53A]
GO
ALTER TABLE [dbo].[grades] DROP CONSTRAINT [DF__grades__last_mov__6754599E]
GO
ALTER TABLE [dbo].[grades] DROP CONSTRAINT [DF__grades__status__5BE2A6F2]
GO
ALTER TABLE [dbo].[factory_type] DROP CONSTRAINT [DF__factory_t__last___656C112C]
GO
ALTER TABLE [dbo].[factory_type] DROP CONSTRAINT [DF__factory_t__statu__403A8C7D]
GO
ALTER TABLE [dbo].[factory] DROP CONSTRAINT [DF__factory__last_mo__6477ECF3]
GO
ALTER TABLE [dbo].[factory] DROP CONSTRAINT [DF__factory__status__3C69FB99]
GO
ALTER TABLE [dbo].[department] DROP CONSTRAINT [DF__departmen__manag__2CBDA3B5]
GO
ALTER TABLE [dbo].[department] DROP CONSTRAINT [DF__departmen__last___6383C8BA]
GO
ALTER TABLE [dbo].[department] DROP CONSTRAINT [DF__departmen__paren__5EBF139D]
GO
ALTER TABLE [dbo].[department] DROP CONSTRAINT [DF__departmen__statu__38996AB5]
GO
ALTER TABLE [dbo].[check-in_data] DROP CONSTRAINT [DF__check-in___work___59C55456]
GO
/****** Object:  Index [UQ__work_shi__357D4CF91A3ADDC1]    Script Date: 2024/2/26 上午 10:00:27 ******/
ALTER TABLE [dbo].[work_shift] DROP CONSTRAINT [UQ__work_shi__357D4CF91A3ADDC1]
GO
/****** Object:  Index [UQ__vacation__357D4CF9544A43A4]    Script Date: 2024/2/26 上午 10:00:27 ******/
ALTER TABLE [dbo].[vacation_types] DROP CONSTRAINT [UQ__vacation__357D4CF9544A43A4]
GO
/****** Object:  Index [UQ__roles__357D4CF9E8DA137E]    Script Date: 2024/2/26 上午 10:00:27 ******/
ALTER TABLE [dbo].[roles] DROP CONSTRAINT [UQ__roles__357D4CF9E8DA137E]
GO
/****** Object:  Index [UQ__ranks__357D4CF9F05C8630]    Script Date: 2024/2/26 上午 10:00:27 ******/
ALTER TABLE [dbo].[ranks] DROP CONSTRAINT [UQ__ranks__357D4CF9F05C8630]
GO
/****** Object:  Index [UQ__permissi__357D4CF9BBC75B51]    Script Date: 2024/2/26 上午 10:00:27 ******/
ALTER TABLE [dbo].[permissions] DROP CONSTRAINT [UQ__permissi__357D4CF9BBC75B51]
GO
/****** Object:  Index [UQ__leave_co__357D4CF96B61C62E]    Script Date: 2024/2/26 上午 10:00:27 ******/
ALTER TABLE [dbo].[leave_core] DROP CONSTRAINT [UQ__leave_co__357D4CF96B61C62E]
GO
/****** Object:  Index [UQ__grades__357D4CF908CAA2F0]    Script Date: 2024/2/26 上午 10:00:27 ******/
ALTER TABLE [dbo].[grades] DROP CONSTRAINT [UQ__grades__357D4CF908CAA2F0]
GO
/****** Object:  Index [UQ__factory___357D4CF9F4573E10]    Script Date: 2024/2/26 上午 10:00:27 ******/
ALTER TABLE [dbo].[factory_type] DROP CONSTRAINT [UQ__factory___357D4CF9F4573E10]
GO
/****** Object:  Index [UQ__factory__357D4CF997DBC82F]    Script Date: 2024/2/26 上午 10:00:27 ******/
ALTER TABLE [dbo].[factory] DROP CONSTRAINT [UQ__factory__357D4CF997DBC82F]
GO
/****** Object:  Index [UQ__departme__357D4CF9B7B5F09A]    Script Date: 2024/2/26 上午 10:00:27 ******/
ALTER TABLE [dbo].[department] DROP CONSTRAINT [UQ__departme__357D4CF9B7B5F09A]
GO
/****** Object:  Table [dbo].[work_shift]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[work_shift]') AND type in (N'U'))
DROP TABLE [dbo].[work_shift]
GO
/****** Object:  Table [dbo].[work_schedule]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[work_schedule]') AND type in (N'U'))
DROP TABLE [dbo].[work_schedule]
GO
/****** Object:  Table [dbo].[vacations]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[vacations]') AND type in (N'U'))
DROP TABLE [dbo].[vacations]
GO
/****** Object:  Table [dbo].[vacation_types]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[vacation_types]') AND type in (N'U'))
DROP TABLE [dbo].[vacation_types]
GO
/****** Object:  Table [dbo].[users]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[users]') AND type in (N'U'))
DROP TABLE [dbo].[users]
GO
/****** Object:  Table [dbo].[user_role]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[user_role]') AND type in (N'U'))
DROP TABLE [dbo].[user_role]
GO
/****** Object:  Table [dbo].[user_department]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[user_department]') AND type in (N'U'))
DROP TABLE [dbo].[user_department]
GO
/****** Object:  Table [dbo].[system_logs]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[system_logs]') AND type in (N'U'))
DROP TABLE [dbo].[system_logs]
GO
/****** Object:  Table [dbo].[sign_off_flow]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[sign_off_flow]') AND type in (N'U'))
DROP TABLE [dbo].[sign_off_flow]
GO
/****** Object:  Table [dbo].[roles]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[roles]') AND type in (N'U'))
DROP TABLE [dbo].[roles]
GO
/****** Object:  Table [dbo].[ranks]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[ranks]') AND type in (N'U'))
DROP TABLE [dbo].[ranks]
GO
/****** Object:  Table [dbo].[permissions]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[permissions]') AND type in (N'U'))
DROP TABLE [dbo].[permissions]
GO
/****** Object:  Table [dbo].[permission_role]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[permission_role]') AND type in (N'U'))
DROP TABLE [dbo].[permission_role]
GO
/****** Object:  Table [dbo].[overtime_sign_off_flow]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[overtime_sign_off_flow]') AND type in (N'U'))
DROP TABLE [dbo].[overtime_sign_off_flow]
GO
/****** Object:  Table [dbo].[overtime_request_form]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[overtime_request_form]') AND type in (N'U'))
DROP TABLE [dbo].[overtime_request_form]
GO
/****** Object:  Table [dbo].[overtime_rate]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[overtime_rate]') AND type in (N'U'))
DROP TABLE [dbo].[overtime_rate]
GO
/****** Object:  Table [dbo].[menus]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[menus]') AND type in (N'U'))
DROP TABLE [dbo].[menus]
GO
/****** Object:  Table [dbo].[menu_role]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[menu_role]') AND type in (N'U'))
DROP TABLE [dbo].[menu_role]
GO
/****** Object:  Table [dbo].[leave_sign_off_flow]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[leave_sign_off_flow]') AND type in (N'U'))
DROP TABLE [dbo].[leave_sign_off_flow]
GO
/****** Object:  Table [dbo].[leave_request_form]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[leave_request_form]') AND type in (N'U'))
DROP TABLE [dbo].[leave_request_form]
GO
/****** Object:  Table [dbo].[leave_defer]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[leave_defer]') AND type in (N'U'))
DROP TABLE [dbo].[leave_defer]
GO
/****** Object:  Table [dbo].[leave_correct]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[leave_correct]') AND type in (N'U'))
DROP TABLE [dbo].[leave_correct]
GO
/****** Object:  Table [dbo].[leave_core_group_user]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[leave_core_group_user]') AND type in (N'U'))
DROP TABLE [dbo].[leave_core_group_user]
GO
/****** Object:  Table [dbo].[leave_core_group_condition]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[leave_core_group_condition]') AND type in (N'U'))
DROP TABLE [dbo].[leave_core_group_condition]
GO
/****** Object:  Table [dbo].[leave_core_group]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[leave_core_group]') AND type in (N'U'))
DROP TABLE [dbo].[leave_core_group]
GO
/****** Object:  Table [dbo].[leave_core]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[leave_core]') AND type in (N'U'))
DROP TABLE [dbo].[leave_core]
GO
/****** Object:  Table [dbo].[grades]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[grades]') AND type in (N'U'))
DROP TABLE [dbo].[grades]
GO
/****** Object:  Table [dbo].[factory_type]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[factory_type]') AND type in (N'U'))
DROP TABLE [dbo].[factory_type]
GO
/****** Object:  Table [dbo].[factory]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[factory]') AND type in (N'U'))
DROP TABLE [dbo].[factory]
GO
/****** Object:  Table [dbo].[department]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[department]') AND type in (N'U'))
DROP TABLE [dbo].[department]
GO
/****** Object:  Table [dbo].[check-in_data]    Script Date: 2024/2/26 上午 10:00:27 ******/
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[check-in_data]') AND type in (N'U'))
DROP TABLE [dbo].[check-in_data]
GO
/****** Object:  UserDefinedFunction [dbo].[FN_CalculateUserSeniorityWithEndDate]    Script Date: 2024/2/26 上午 10:00:27 ******/
DROP FUNCTION [dbo].[FN_CalculateUserSeniorityWithEndDate]
GO
/****** Object:  UserDefinedFunction [dbo].[FN_CalculateUserSeniority]    Script Date: 2024/2/26 上午 10:00:27 ******/
DROP FUNCTION [dbo].[FN_CalculateUserSeniority]
GO
/****** Object:  UserDefinedFunction [dbo].[FN_CalculateSeniorityMonth]    Script Date: 2024/2/26 上午 10:00:27 ******/
DROP FUNCTION [dbo].[FN_CalculateSeniorityMonth]
GO
/****** Object:  UserDefinedFunction [dbo].[FN_CalculateSeniority]    Script Date: 2024/2/26 上午 10:00:27 ******/
DROP FUNCTION [dbo].[FN_CalculateSeniority]
GO
/****** Object:  UserDefinedFunction [dbo].[FN_CalculateSeniority]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE FUNCTION [dbo].[FN_CalculateSeniority] (@startDate SMALLDATETIME, @endDate SMALLDATETIME)
    RETURNS @retDateDuration TABLE
                             (
                                 DurationYear  INT  NULL,
                                 DurationMonth INT  NULL
                             )
AS
BEGIN

    DECLARE @yearDiff INT
    DECLARE @monthDiff INT

    -- 計算年資
    SELECT @yearDiff = DATEDIFF(YEAR, @startDate, @endDate)
    SELECT @monthDiff = DATEDIFF(MONTH, @startDate, @endDate)

    -- 未滿年資修正
    IF( DATEDIFF(DAY, DATEADD(YEAR, @yearDiff, @startDate), @endDate) < 0 )
        SET @yearDiff = @yearDiff - 1

    IF( DATEDIFF(DAY, DATEADD(MONTH, @monthDiff, @startDate), @endDate) < 0 )
        SET @monthDiff = @monthDiff -1

    -- 扣除年資後的月數
    SET @monthDiff = @monthDiff - 12 * @yearDiff

    -- 回傳年資(年,月)
    INSERT @retDateDuration( DurationYear, DurationMonth )
    SELECT IIF( @yearDiff < 0, 0, @yearDiff ), IIF( @yearDiff < 0 OR @monthDiff < 0, 0, @monthDiff )

    RETURN
END
GO
/****** Object:  UserDefinedFunction [dbo].[FN_CalculateSeniorityMonth]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE FUNCTION [dbo].[FN_CalculateSeniorityMonth] (@startDate SMALLDATETIME, @endDate SMALLDATETIME)
    RETURNS int
AS
BEGIN
    DECLARE @monthDiff INT

    -- 計算年資
    SELECT @monthDiff = DATEDIFF(MONTH, @startDate, @endDate)

    -- 未滿年資修正
    IF( DATEDIFF(DAY, DATEADD(MONTH, @monthDiff, @startDate), @endDate) < 0 )
        SET @monthDiff = @monthDiff -1

    RETURN @monthDiff
END
GO
/****** Object:  UserDefinedFunction [dbo].[FN_CalculateUserSeniority]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
create function [dbo].[FN_CalculateUserSeniority] ()
    returns @retDateDuration table
                             (
                                 user_id     int  null ,
                                 arrived     smalldatetime null ,
                                 year        int  null ,
                                 month       int  null
                             )
as
begin
    -- 定義Cursor並打開
    declare idCursor cursor for select id, arrived from users where leaved is not null or leaved != '0001-01-01 00:00:00.0000000 +00:00'
    open idCursor

    -- 定義存放id,arrived變數
    declare @id bigint
    declare @arrived smalldatetime

    -- 今天日期
    declare @today smalldatetime
    set @today = getdate()

    -- 迴圈跑Cursor
    fetch next from idCursor into @id, @arrived
    while (@@fetch_status <> -1)
    begin
        declare @yearDiff int
        declare @monthDiff int

        -- 計算年資
        select @yearDiff = DATEDIFF(year, @arrived, @today)
        select @monthDiff = DATEDIFF(month, @arrived, @today)

        -- 未滿年資修正
        if ( DATEDIFF(day , DATEADD(year , @yearDiff, @arrived), @today) < 0 )
            set @yearDiff = @yearDiff - 1

        if ( DATEDIFF(day , DATEADD(month , @monthDiff, @arrived), @today) < 0 )
            set @monthDiff = @monthDiff -1

        -- 扣除年資後的月數
        set @monthDiff = @monthDiff - 12 * @yearDiff

        -- 回傳年資(id,年,月)
        insert @retDateDuration( user_id , arrived , year, month )
        select @id, @arrived , IIF( @yearDiff < 0, 0, @yearDiff ), IIF( @yearDiff < 0 OR @monthDiff < 0, 0, @monthDiff )

        -- 抓取下一個
        fetch next from idCursor into @id, @arrived
    end

    close idCursor
    deallocate idCursor

    return
end
GO
/****** Object:  UserDefinedFunction [dbo].[FN_CalculateUserSeniorityWithEndDate]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE function [dbo].[FN_CalculateUserSeniorityWithEndDate] (@endDate smalldatetime)
    returns @retDateDuration table
                             (
                                 user_id     int  null ,
                                 arrived     smalldatetime null ,
                                 year        int  null ,
                                 month       int  null
                             )
as
begin
    -- 定義Cursor並打開
    declare idCursor cursor for select id, arrived from users
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
            insert @retDateDuration( user_id , arrived , year, month )
            select @id , @arrived , IIF( @yearDiff < 0, 0, @yearDiff ), IIF( @yearDiff < 0 OR @monthDiff < 0, 0, @monthDiff )

            -- 抓取下一個
            fetch next from idCursor into @id, @arrived
        end

    close idCursor
    deallocate idCursor

    return
end
GO
/****** Object:  Table [dbo].[check-in_data]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[check-in_data](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[date] [nvarchar](15) NOT NULL,
	[time] [nvarchar](15) NOT NULL,
	[check_in_code] [nvarchar](50) NULL,
	[action] [tinyint] NOT NULL,
	[employee_name] [nvarchar](10) NOT NULL,
	[work_shift_id] [bigint] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[department]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[department](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[code] [nvarchar](100) NOT NULL,
	[name] [nvarchar](100) NOT NULL,
	[status] [bit] NOT NULL,
	[remark] [nvarchar](max) NULL,
	[parent_id] [bigint] NOT NULL,
	[last_mover] [bigint] NOT NULL,
	[manager_id] [bigint] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[factory]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[factory](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[code] [nvarchar](100) NOT NULL,
	[name] [nvarchar](100) NOT NULL,
	[status] [bit] NOT NULL,
	[remark] [nvarchar](max) NULL,
	[last_mover] [bigint] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[factory_type]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[factory_type](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[code] [nvarchar](100) NOT NULL,
	[name] [nvarchar](100) NOT NULL,
	[status] [bit] NOT NULL,
	[remark] [nvarchar](max) NULL,
	[factory_id] [bigint] NOT NULL,
	[last_mover] [bigint] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[grades]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[grades](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[code] [nvarchar](100) NOT NULL,
	[name] [nvarchar](100) NOT NULL,
	[status] [bit] NOT NULL,
	[remark] [nvarchar](max) NULL,
	[rank_id] [bigint] NULL,
	[last_mover] [bigint] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[leave_core]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[leave_core](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[code] [nvarchar](100) NOT NULL,
	[name] [nvarchar](100) NOT NULL,
	[status] [bit] NOT NULL,
	[remark] [nvarchar](max) NULL,
	[last_mover] [bigint] NOT NULL,
	[is_pay] [float] NULL,
	[minimum] [tinyint] NULL,
	[default_limit] [tinyint] NULL,
	[deferred] [bit] NULL,
	[defer_day] [bigint] NULL,
	[exchange] [bit] NULL,
	[cycle] [tinyint] NULL,
	[maturity] [bigint] NULL,
	[is_maturity] [bit] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[leave_core_group]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[leave_core_group](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[leave_core_id] [bigint] NOT NULL,
	[name] [nvarchar](max) NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[leave_core_group_condition]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[leave_core_group_condition](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[leave_core_group_id] [bigint] NOT NULL,
	[condition] [tinyint] NULL,
	[cond1] [bigint] NULL,
	[day] [bigint] NULL,
	[level] [tinyint] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[leave_core_group_user]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[leave_core_group_user](
	[leave_core_group_id] [bigint] NOT NULL,
	[user_id] [bigint] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[leave_core_group_id] ASC,
	[user_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[leave_correct]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[leave_correct](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[user_id] [bigint] NOT NULL,
	[leave_core_id] [bigint] NOT NULL,
	[available] [float] NULL,
	[used] [float] NULL,
	[signing] [float] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[leave_defer]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[leave_defer](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[user_id] [bigint] NOT NULL,
	[leave_core_id] [bigint] NOT NULL,
	[effective] [nvarchar](max) NULL,
	[expired] [nvarchar](max) NULL,
	[available] [float] NULL,
	[this_time] [float] NULL,
	[extra] [float] NULL,
	[previous_id] [bigint] NOT NULL,
	[next_id] [bigint] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[leave_request_form]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[leave_request_form](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[order] [nvarchar](50) NULL,
	[start_date] [nvarchar](50) NOT NULL,
	[start_time] [nvarchar](50) NOT NULL,
	[end_date] [nvarchar](50) NOT NULL,
	[end_time] [nvarchar](50) NOT NULL,
	[remark] [nvarchar](max) NULL,
	[sign_status] [tinyint] NOT NULL,
	[leave_minute_count] [float] NULL,
	[leave_hour_count] [float] NULL,
	[leave_day_count] [float] NULL,
	[attach] [nvarchar](max) NULL,
	[leave_core_id] [bigint] NOT NULL,
	[user_id] [bigint] NOT NULL,
	[department_id] [bigint] NOT NULL,
	[proxy_user_id] [bigint] NOT NULL,
	[proxy_department_id] [bigint] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[leave_sign_off_flow]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[leave_sign_off_flow](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[leave_request_form_id] [bigint] NOT NULL,
	[sign_off_user_id] [bigint] NOT NULL,
	[level] [bigint] NULL,
	[sign_type] [bigint] NULL,
	[notify] [bigint] NULL,
	[remark] [nvarchar](max) NULL,
	[comment] [nvarchar](max) NULL,
	[status] [bigint] NULL,
	[sign_date] [datetimeoffset](7) NULL,
	[uuid] [nvarchar](max) NULL,
	[locale] [nvarchar](max) NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[menu_role]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[menu_role](
	[menu_id] [bigint] NOT NULL,
	[role_id] [bigint] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[menu_id] ASC,
	[role_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[menus]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[menus](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[type] [nvarchar](10) NULL,
	[parent_id] [bigint] NOT NULL,
	[path] [nvarchar](max) NOT NULL,
	[name] [nvarchar](max) NOT NULL,
	[component] [nvarchar](max) NOT NULL,
	[redirect] [nvarchar](max) NULL,
	[sort] [bigint] NOT NULL,
	[status] [bit] NULL,
	[title] [nvarchar](max) NULL,
	[dynamic_level] [bigint] NULL,
	[real_path] [nvarchar](max) NULL,
	[ignore_keep_alive] [bit] NULL,
	[affix] [bit] NULL,
	[icon] [nvarchar](max) NULL,
	[frame_src] [nvarchar](max) NULL,
	[transition_name] [nvarchar](max) NULL,
	[hide_breadcrumb] [bit] NULL,
	[carry_param] [bit] NULL,
	[hide_children_in_menu] [bit] NULL,
	[current_active_menu] [nvarchar](max) NULL,
	[hide_tab] [bit] NULL,
	[hide_menu] [bit] NULL,
	[ignore_route] [bit] NULL,
	[hide_path_for_children] [bit] NULL,
	[last_mover] [bigint] NOT NULL,
	[show] [bigint] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[overtime_rate]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[overtime_rate](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[start] [tinyint] NULL,
	[end] [tinyint] NULL,
	[multiply] [float] NULL,
	[vacation_type_id] [bigint] NOT NULL,
	[last_mover] [bigint] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[overtime_request_form]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[overtime_request_form](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[order] [nvarchar](50) NULL,
	[start] [datetimeoffset](7) NOT NULL,
	[end] [datetimeoffset](7) NOT NULL,
	[remark] [nvarchar](max) NULL,
	[sign_status] [tinyint] NOT NULL,
	[user_id] [bigint] NOT NULL,
	[department_id] [bigint] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[overtime_sign_off_flow]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[overtime_sign_off_flow](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[overtime_request_form_id] [bigint] NOT NULL,
	[sign_off_user_id] [bigint] NOT NULL,
	[level] [bigint] NULL,
	[sign_type] [bigint] NULL,
	[notify] [bigint] NULL,
	[remark] [nvarchar](max) NULL,
	[comment] [nvarchar](max) NULL,
	[status] [bigint] NULL,
	[sign_date] [datetimeoffset](7) NULL,
	[uuid] [nvarchar](max) NULL,
	[locale] [nvarchar](max) NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[permission_role]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[permission_role](
	[permission_id] [bigint] NOT NULL,
	[role_id] [bigint] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[permission_id] ASC,
	[role_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[permissions]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[permissions](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[code] [nvarchar](100) NOT NULL,
	[name] [nvarchar](100) NOT NULL,
	[status] [bit] NULL,
	[remark] [nvarchar](max) NULL,
	[last_mover] [bigint] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[ranks]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[ranks](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[code] [nvarchar](100) NOT NULL,
	[name] [nvarchar](100) NOT NULL,
	[status] [bit] NOT NULL,
	[remark] [nvarchar](max) NULL,
	[last_mover] [bigint] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[roles]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[roles](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[code] [nvarchar](100) NOT NULL,
	[name] [nvarchar](100) NOT NULL,
	[status] [bit] NOT NULL,
	[remark] [nvarchar](max) NULL,
	[last_mover] [bigint] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[sign_off_flow]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[sign_off_flow](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[department_id] [bigint] NOT NULL,
	[leave_core_id] [bigint] NOT NULL,
	[level] [bigint] NULL,
	[exceed_day] [bigint] NULL,
	[user_id] [bigint] NULL,
	[sign_type] [bigint] NULL,
	[notify] [bigint] NULL,
	[remark] [nvarchar](max) NULL,
	[rule_type] [bigint] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[system_logs]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[system_logs](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[user_id] [bigint] NOT NULL,
	[table] [nvarchar](100) NOT NULL,
	[crud] [nvarchar](2) NOT NULL,
	[message] [nvarchar](1000) NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[user_department]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[user_department](
	[department_id] [bigint] NOT NULL,
	[user_id] [bigint] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[department_id] ASC,
	[user_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[user_role]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[user_role](
	[role_id] [bigint] NOT NULL,
	[user_id] [bigint] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[role_id] ASC,
	[user_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[users]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[users](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[username] [nvarchar](100) NOT NULL,
	[password] [nvarchar](100) NOT NULL,
	[real_name] [nvarchar](100) NOT NULL,
	[card] [nvarchar](11) NOT NULL,
	[birth] [datetimeoffset](7) NOT NULL,
	[arrived] [datetimeoffset](7) NOT NULL,
	[leaved] [datetimeoffset](7) NULL,
	[status] [bit] NOT NULL,
	[remark] [nvarchar](max) NULL,
	[email] [nvarchar](100) NULL,
	[mobile] [nvarchar](11) NULL,
	[department_id] [bigint] NOT NULL,
	[factory_id] [bigint] NOT NULL,
	[factory_type_id] [bigint] NOT NULL,
	[rank_id] [bigint] NOT NULL,
	[grade_id] [bigint] NOT NULL,
	[desc] [nvarchar](max) NULL,
	[avatar] [nvarchar](max) NULL,
	[last_mover] [bigint] NOT NULL,
	[check_in_code] [nvarchar](50) NULL,
	[backend] [bit] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[vacation_types]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[vacation_types](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[code] [nvarchar](100) NOT NULL,
	[name] [nvarchar](100) NOT NULL,
	[status] [bit] NOT NULL,
	[remark] [nvarchar](max) NULL,
	[last_mover] [bigint] NOT NULL,
	[color] [nvarchar](10) NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[vacations]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[vacations](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[general_key] [nvarchar](129) NOT NULL,
	[internal_key] [nvarchar](129) NULL,
	[type_id] [bigint] NOT NULL,
	[year] [smallint] NOT NULL,
	[month] [tinyint] NOT NULL,
	[day] [tinyint] NOT NULL,
	[remark] [nvarchar](max) NULL,
	[last_mover] [bigint] NOT NULL,
	[start_date] [nvarchar](max) NULL,
	[end_date] [nvarchar](max) NULL,
	[repeat] [nvarchar](max) NULL,
	[end_repeat] [nvarchar](max) NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[work_schedule]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[work_schedule](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[year] [smallint] NOT NULL,
	[month] [tinyint] NOT NULL,
	[day] [tinyint] NOT NULL,
	[user_id] [bigint] NOT NULL,
	[work_shift_id] [bigint] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[work_shift]    Script Date: 2024/2/26 上午 10:00:27 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[work_shift](
	[id] [bigint] IDENTITY(1,1) NOT NULL,
	[created_at] [datetimeoffset](7) NULL,
	[updated_at] [datetimeoffset](7) NULL,
	[deleted_at] [datetimeoffset](7) NULL,
	[code] [nvarchar](100) NOT NULL,
	[name] [nvarchar](100) NOT NULL,
	[status] [bit] NOT NULL,
	[remark] [nvarchar](max) NULL,
	[start_time] [nvarchar](10) NOT NULL,
	[end_time] [nvarchar](10) NOT NULL,
	[rest_start_time] [nvarchar](10) NULL,
	[rest_end_time] [nvarchar](10) NOT NULL,
	[last_mover] [bigint] NOT NULL,
	[color] [nvarchar](10) NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
SET IDENTITY_INSERT [dbo].[check-in_data] ON 

INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (1, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:54:00', N'0000010004', 0, N'曾明慧', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (2, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:18:32', N'0000010004', 1, N'曾明慧', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (3, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:54:25', N'0000010005', 0, N'廖家稘', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (4, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:10:19', N'0000010005', 1, N'廖家稘', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (5, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'08:00:18', N'0000010006', 0, N'張孟莉', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (6, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:20:04', N'0000010006', 1, N'張孟莉', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (7, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:55:40', N'0000010016', 0, N'洪藝芩', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (8, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:01:05', N'0000010016', 1, N'洪藝芩', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (9, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:54:36', N'0000010018', 0, N'顏婉玲', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (10, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:24:01', N'0000010018', 1, N'顏婉玲', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (11, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:57:20', N'0000010021', 0, N'洪孟芳', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (12, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'18:01:48', N'0000010021', 1, N'洪孟芳', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (13, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:56:35', N'0000010022', 0, N'林芊逸', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (14, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:14:21', N'0000010022', 1, N'林芊逸', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (15, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:59:38', N'0000010024', 0, N'劉青佩', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (16, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:19:45', N'0000010024', 1, N'劉青佩', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (17, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'08:04:05', N'0000010025', 0, N'林敬芳', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (18, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:25:31', N'0000010025', 1, N'林敬芳', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (19, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:58:09', N'0000010026', 0, N'廖麗君', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (20, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:12:47', N'0000010026', 1, N'廖麗君', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (21, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:59:07', N'0000010030', 0, N'黃玲文', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (22, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:02:44', N'0000010030', 1, N'黃玲文', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (23, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:55:30', N'0000010031', 0, N'鄭欣怡', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (24, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:05:49', N'0000010031', 1, N'鄭欣怡', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (25, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:58:45', N'0000010033', 0, N'陳立宜', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (26, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:17:37', N'0000010033', 1, N'陳立宜', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (27, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:45:51', N'0000010035', 0, N'孫維偵', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (28, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:59:45', N'0000010035', 1, N'孫維偵', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (29, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:59:11', N'0000010039', 0, N'林家瑜', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (30, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:10:30', N'0000010039', 1, N'林家瑜', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (31, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:03:14', N'0000010040', 0, N'蕭雪梅', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (32, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:00:58', N'0000010040', 1, N'蕭雪梅', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (33, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'08:01:05', N'0000010047', 0, N'王淑媛', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (34, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:27:14', N'0000010047', 1, N'王淑媛', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (35, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:45:26', N'0000010055', 0, N'李進成', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (36, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:02:22', N'0000010055', 1, N'李進成', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (37, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:52:39', N'0000010056', 0, N'朱來好', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (38, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'18:05:03', N'0000010056', 1, N'朱來好', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (39, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:59:04', N'0000010058', 0, N'何虹億', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (40, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'15:32:25', N'0000010058', 1, N'何虹億', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (41, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'08:23:13', N'0000010059', 0, N'林智鴻', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (42, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:04:48', N'0000010059', 1, N'林智鴻', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (43, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:57:14', N'0000010062', 0, N'安多', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (44, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:18:26', N'0000010062', 1, N'安多', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (45, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:47:33', N'0000010073', 0, N'林慶敦', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (46, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:03:41', N'0000010073', 1, N'林慶敦', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (47, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'08:00:51', N'0000010081', 0, N'陳莉雯', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (48, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:39:12', N'0000010081', 1, N'陳莉雯', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (49, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:54:22', N'0000010105', 0, N'周伶穗', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (50, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:04:19', N'0000010105', 1, N'周伶穗', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (51, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:50:11', N'0000010111', 0, N'李伊萍', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (52, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:04:58', N'0000010111', 1, N'李伊萍', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (53, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:57:06', N'0000010114', 0, N'分迪', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (54, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:18:22', N'0000010114', 1, N'分迪', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (55, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:52:32', N'0000010118', 0, N'葉淳昱', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (56, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:40:45', N'0000010118', 1, N'葉淳昱', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (57, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:56:06', N'0000010121', 0, N'陳映汶', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (58, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:20:40', N'0000010121', 1, N'陳映汶', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (59, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:59:52', N'0000010122', 0, N'曾鈺婷', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (60, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:03:22', N'0000010122', 1, N'曾鈺婷', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (61, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:53:06', N'0000010124', 0, N'高嘉莉', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (62, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:11:54', N'0000010124', 1, N'高嘉莉', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (63, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:58:22', N'0000010126', 0, N'詹婉青', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (64, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:59:17', N'0000010126', 1, N'詹婉青', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (65, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:56:21', N'0000010136', 0, N'林劻郁', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (66, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'15:32:12', N'0000010136', 1, N'林劻郁', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (67, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:59:46', N'0000010140', 0, N'黃秋玲', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (68, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:06:37', N'0000010140', 1, N'黃秋玲', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (69, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:31:19', N'0000010143', 0, N'葉沛彤', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (70, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:01:12', N'0000010143', 1, N'葉沛彤', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (71, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:49:03', N'0000010144', 0, N'洪青妙', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (72, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:40:07', N'0000010144', 1, N'洪青妙', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (73, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:55:14', N'0000010147', 0, N'呂旻樺', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (74, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:34:11', N'0000010147', 1, N'呂旻樺', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (75, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:58:50', N'0000010152', 0, N'林彥宏', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (76, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:09:55', N'0000010152', 1, N'林彥宏', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (77, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:44:47', N'0000010155', 0, N'戴榕昀', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (78, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:55:33', N'0000010156', 0, N'劉怡伶', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (79, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:01:41', N'0000010156', 1, N'劉怡伶', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (80, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:52:58', N'0000010157', 0, N'張冠惠', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (81, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:11:00', N'0000010157', 1, N'張冠惠', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (82, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:51:21', N'0000010161', 0, N'賴宛萱', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (83, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'16:04:56', N'0000010161', 1, N'賴宛萱', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (84, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'07:58:06', N'0000010163', 0, N'張育維', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (85, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/20', N'17:01:18', N'0000010163', 1, N'張育維', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (86, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/21', N'08:16:51', N'0000010056', 0, N'朱來好', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (87, CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:51:10.6937779+08:00' AS DateTimeOffset), NULL, N'2023/10/21', N'16:42:15', N'0000010056', 1, N'朱來好', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (88, CAST(N'2023-10-24T15:52:38.8100000+00:00' AS DateTimeOffset), CAST(N'2023-10-24T15:52:35.2110000+00:00' AS DateTimeOffset), NULL, N'2023/10/21', N'08:58:06', N'0000010163', 0, N'張育維', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (89, CAST(N'2023-10-24T15:52:40.5560000+00:00' AS DateTimeOffset), CAST(N'2023-10-24T15:52:36.6400000+00:00' AS DateTimeOffset), NULL, N'2023/10/21', N'17:00:05', N'0000010163', 1, N'張育維', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (90, CAST(N'2023-10-24T15:52:38.8100000+00:00' AS DateTimeOffset), CAST(N'2023-10-24T15:52:35.2110000+00:00' AS DateTimeOffset), NULL, N'2023/10/21', N'09:50:06', N'0000010163', 0, N'張育維', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (91, CAST(N'2023-10-24T15:52:40.5560000+00:00' AS DateTimeOffset), CAST(N'2023-10-24T15:52:36.6400000+00:00' AS DateTimeOffset), NULL, N'2023/10/21', N'16:50:05', N'0000010163', 1, N'張育維', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (92, CAST(N'2023-10-24T15:52:38.8100000+00:00' AS DateTimeOffset), CAST(N'2023-10-24T15:52:35.2110000+00:00' AS DateTimeOffset), NULL, N'2023/10/22', N'09:50:06', N'0000010163', 0, N'張育維', 1)
INSERT [dbo].[check-in_data] ([id], [created_at], [updated_at], [deleted_at], [date], [time], [check_in_code], [action], [employee_name], [work_shift_id]) VALUES (93, CAST(N'2023-10-24T15:52:40.5560000+00:00' AS DateTimeOffset), CAST(N'2023-10-24T15:52:36.6400000+00:00' AS DateTimeOffset), NULL, N'2023/10/22', N'', N'0000010163', 1, N'張育維', 1)
SET IDENTITY_INSERT [dbo].[check-in_data] OFF
GO
SET IDENTITY_INSERT [dbo].[department] ON 

INSERT [dbo].[department] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [parent_id], [last_mover], [manager_id]) VALUES (1, CAST(N'2023-09-27T10:47:24.7210000+00:00' AS DateTimeOffset), CAST(N'2023-12-11T15:28:24.0932059+08:00' AS DateTimeOffset), NULL, N'test1', N'測試部門1', 1, N'', 0, 1, 3)
INSERT [dbo].[department] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [parent_id], [last_mover], [manager_id]) VALUES (2, CAST(N'2023-09-27T10:48:01.7120000+00:00' AS DateTimeOffset), CAST(N'2023-09-27T10:48:02.8180000+00:00' AS DateTimeOffset), NULL, N'test1-1', N'測試部門1-1', 1, NULL, 1, 0, 0)
INSERT [dbo].[department] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [parent_id], [last_mover], [manager_id]) VALUES (3, CAST(N'2023-09-27T10:48:20.0020000+00:00' AS DateTimeOffset), CAST(N'2023-09-27T10:48:20.8580000+00:00' AS DateTimeOffset), NULL, N'test1-2', N'測試部門1-2', 1, NULL, 1, 0, 0)
INSERT [dbo].[department] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [parent_id], [last_mover], [manager_id]) VALUES (4, CAST(N'2023-09-27T10:48:36.9790000+00:00' AS DateTimeOffset), CAST(N'2023-09-27T10:48:37.7760000+00:00' AS DateTimeOffset), NULL, N'test1-1-1', N'測試部門1-1-1', 1, NULL, 2, 0, 0)
INSERT [dbo].[department] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [parent_id], [last_mover], [manager_id]) VALUES (5, CAST(N'2023-09-27T10:49:01.2910000+00:00' AS DateTimeOffset), CAST(N'2023-09-27T13:52:26.0569126+08:00' AS DateTimeOffset), NULL, N'test2', N'測試部門2', 0, N'', 0, 0, 0)
INSERT [dbo].[department] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [parent_id], [last_mover], [manager_id]) VALUES (6, CAST(N'2023-09-27T10:49:26.4940000+00:00' AS DateTimeOffset), CAST(N'2023-09-27T10:49:27.3340000+00:00' AS DateTimeOffset), NULL, N'test1-1-2', N'測試部門1-1-2', 1, NULL, 2, 0, 0)
INSERT [dbo].[department] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [parent_id], [last_mover], [manager_id]) VALUES (7, CAST(N'2023-09-27T13:32:06.4517161+08:00' AS DateTimeOffset), CAST(N'2023-09-27T13:32:06.4517161+08:00' AS DateTimeOffset), NULL, N'test2-1', N'測試部門2-1', 1, N'321', 5, 0, 0)
INSERT [dbo].[department] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [parent_id], [last_mover], [manager_id]) VALUES (9, CAST(N'2023-09-27T13:35:26.1752556+08:00' AS DateTimeOffset), CAST(N'2023-09-27T13:45:32.1358982+08:00' AS DateTimeOffset), NULL, N'test2-1-1', N'測試部門2-1-1', 1, N'123123', 7, 0, 0)
INSERT [dbo].[department] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [parent_id], [last_mover], [manager_id]) VALUES (10, CAST(N'2023-11-22T10:05:10.0519382+08:00' AS DateTimeOffset), CAST(N'2023-11-22T10:07:05.0548660+08:00' AS DateTimeOffset), NULL, N'test1-1-1-1', N'測試1-1-1-1', 1, N'', 4, 1, 1)
SET IDENTITY_INSERT [dbo].[department] OFF
GO
SET IDENTITY_INSERT [dbo].[factory] ON 

INSERT [dbo].[factory] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (1, CAST(N'2023-10-02T09:24:16.5170000+00:00' AS DateTimeOffset), CAST(N'2023-10-06T09:08:15.6994438+08:00' AS DateTimeOffset), NULL, N'f_test_1', N'工廠1', 1, N'測試用工廠', 1)
INSERT [dbo].[factory] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (2, CAST(N'2023-10-03T15:46:57.4317113+08:00' AS DateTimeOffset), CAST(N'2023-10-04T10:07:11.7240764+08:00' AS DateTimeOffset), NULL, N'f_test_2', N'工廠2', 1, N'123', 1)
SET IDENTITY_INSERT [dbo].[factory] OFF
GO
SET IDENTITY_INSERT [dbo].[factory_type] ON 

INSERT [dbo].[factory_type] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [factory_id], [last_mover]) VALUES (1, CAST(N'2023-10-02T09:24:47.7730000+00:00' AS DateTimeOffset), CAST(N'2023-12-25T10:00:01.1922194+08:00' AS DateTimeOffset), NULL, N'ft_test_1', N'廠別1', 1, N'測試用廠別1', 1, 1)
INSERT [dbo].[factory_type] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [factory_id], [last_mover]) VALUES (4, CAST(N'2023-10-03T15:32:47.7711927+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:32:49.7396059+08:00' AS DateTimeOffset), CAST(N'2023-10-04T10:59:21.3783044+08:00' AS DateTimeOffset), N'ft_test_2', N'廠別2', 1, N'12345', 1, 1)
INSERT [dbo].[factory_type] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [factory_id], [last_mover]) VALUES (6, CAST(N'2023-10-03T15:48:55.4566153+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:48:55.4566153+08:00' AS DateTimeOffset), NULL, N'ft_test_2-1', N'廠別2-1', 1, N'', 2, 1)
INSERT [dbo].[factory_type] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [factory_id], [last_mover]) VALUES (7, CAST(N'2023-10-03T15:49:11.2732615+08:00' AS DateTimeOffset), CAST(N'2023-10-06T09:10:47.0712162+08:00' AS DateTimeOffset), NULL, N'ft_test_2-2', N'廠別2-2', 1, N'', 2, 1)
SET IDENTITY_INSERT [dbo].[factory_type] OFF
GO
SET IDENTITY_INSERT [dbo].[grades] ON 

INSERT [dbo].[grades] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [rank_id], [last_mover]) VALUES (1, CAST(N'2023-10-02T09:26:26.1760000+00:00' AS DateTimeOffset), CAST(N'2023-10-02T09:26:26.8920000+00:00' AS DateTimeOffset), NULL, N'g_test_1', N'職等1', 1, N'測試用', 1, 0)
INSERT [dbo].[grades] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [rank_id], [last_mover]) VALUES (2, CAST(N'2023-10-04T11:52:17.6453934+08:00' AS DateTimeOffset), CAST(N'2023-10-04T11:52:22.1912657+08:00' AS DateTimeOffset), CAST(N'2023-10-04T11:52:26.1474506+08:00' AS DateTimeOffset), N'g_tset_2', N'職等2', 0, N'asafsdasdfa', 1, 1)
INSERT [dbo].[grades] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [rank_id], [last_mover]) VALUES (3, CAST(N'2023-10-04T11:52:41.1698124+08:00' AS DateTimeOffset), CAST(N'2023-10-06T09:11:11.8388372+08:00' AS DateTimeOffset), NULL, N'asdfasf', N'廠別1', 1, N'', 2, 1)
SET IDENTITY_INSERT [dbo].[grades] OFF
GO
SET IDENTITY_INSERT [dbo].[leave_core] ON 

INSERT [dbo].[leave_core] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover], [is_pay], [minimum], [default_limit], [deferred], [defer_day], [exchange], [cycle], [maturity], [is_maturity]) VALUES (2, CAST(N'2023-11-01T09:20:16.5089542+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:49:19.9259351+08:00' AS DateTimeOffset), NULL, N'0001', N'事假', 1, N'勞工請假規則：第 7 條 勞工因有事故必須親自處理者，得請事假，一年內合計不得超過十四日。 事假期間不給工資。', 1, 0, 15, 14, 0, 0, 0, 1, NULL, NULL)
INSERT [dbo].[leave_core] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover], [is_pay], [minimum], [default_limit], [deferred], [defer_day], [exchange], [cycle], [maturity], [is_maturity]) VALUES (4, CAST(N'2023-11-01T09:21:20.7634313+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:14:06.6877264+08:00' AS DateTimeOffset), NULL, N'0002', N'超額-事假', 1, N'勞工請假規則：第 7 條 勞工因有事故必須親自處理者，得請事假，一年內合計不得超過十四日。 事假期間不給工資。 超過法定(累計天數)：改請<超額-事假>使用的假別', 1, 0, 15, 255, 0, 0, 0, 1, NULL, NULL)
INSERT [dbo].[leave_core] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover], [is_pay], [minimum], [default_limit], [deferred], [defer_day], [exchange], [cycle], [maturity], [is_maturity]) VALUES (5, CAST(N'2023-11-01T09:22:46.0119694+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:13:46.1582088+08:00' AS DateTimeOffset), NULL, N'0003', N'病假', 1, N'勞工請假規則：第 4 條 勞工因普通傷害、疾病或生理原因必須治療或休養者，得在下列規定範圍 內請普通傷病假： 一、未住院者，一年內合計不得超過三十日。 二、住院者，二年內合計不得超過一年。 三、未住院傷病假與住院傷病假二年內合計不得超過一年。 經醫師診斷，罹患癌症（含原位癌）採門診方式治療或懷孕期間需安胎休 養者，其治療或休養期間，併入住院傷病假計算。 普通傷病假一年內未超過三十日部分，工資折半發給，其領有勞工保險普 通傷病給付未達工資半數者，由雇主補足之。', 1, 0.5, 0, 30, 0, 0, 0, 1, NULL, NULL)
INSERT [dbo].[leave_core] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover], [is_pay], [minimum], [default_limit], [deferred], [defer_day], [exchange], [cycle], [maturity], [is_maturity]) VALUES (6, CAST(N'2023-11-01T09:23:40.5696033+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:03:57.5792895+08:00' AS DateTimeOffset), NULL, N'1000', N'特休', 1, N'', 1, 1, 0, 2, 1, 9, 0, 3, 5, 1)
INSERT [dbo].[leave_core] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover], [is_pay], [minimum], [default_limit], [deferred], [defer_day], [exchange], [cycle], [maturity], [is_maturity]) VALUES (7, CAST(N'2023-11-01T09:24:31.5396405+08:00' AS DateTimeOffset), CAST(N'2024-01-22T10:07:02.9992322+08:00' AS DateTimeOffset), NULL, N'0004', N'婚假', 1, N'勞工請假規則第 2 條 勞工結婚者給予婚假八日，工資照給。', 1, 1, 0, 8, 0, 0, 0, 1, 0, 0)
INSERT [dbo].[leave_core] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover], [is_pay], [minimum], [default_limit], [deferred], [defer_day], [exchange], [cycle], [maturity], [is_maturity]) VALUES (8, CAST(N'2023-11-01T09:25:07.3610568+08:00' AS DateTimeOffset), CAST(N'2024-01-22T10:07:11.4199905+08:00' AS DateTimeOffset), NULL, N'0005', N'產檢假', 1, N'性別工作平等法：第 15 條 受僱者妊娠期間，雇主應給予產檢假五日。 產檢假期間，薪資照給。 必要時，雇主得要求受僱者提出證明文件。 2021/7/1修正，將產檢假日數由5日增加為7日', 1, 1, 0, 5, 0, 0, 0, 1, 0, 0)
INSERT [dbo].[leave_core] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover], [is_pay], [minimum], [default_limit], [deferred], [defer_day], [exchange], [cycle], [maturity], [is_maturity]) VALUES (9, CAST(N'2023-11-01T09:26:01.6725941+08:00' AS DateTimeOffset), CAST(N'2024-01-22T10:07:21.2830263+08:00' AS DateTimeOffset), NULL, N'0006', N'陪產假', 1, N'性別工作平等法：第 15 條 受僱者於其配偶分娩時，雇主應給予陪產假五日。 陪產假期間，薪資照給。', 1, 1, 0, 5, 0, 0, 0, 1, 0, 0)
INSERT [dbo].[leave_core] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover], [is_pay], [minimum], [default_limit], [deferred], [defer_day], [exchange], [cycle], [maturity], [is_maturity]) VALUES (10, CAST(N'2023-11-01T09:29:27.8447766+08:00' AS DateTimeOffset), CAST(N'2024-01-22T10:07:29.0773117+08:00' AS DateTimeOffset), NULL, N'0007', N'喪假8天', 1, N'勞工請假規則：第 3 條 父母、養父母、繼父母、配偶喪亡者，給予喪假八日，工資照給。', 1, 1, 0, 8, 0, 0, 0, 1, 0, 0)
INSERT [dbo].[leave_core] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover], [is_pay], [minimum], [default_limit], [deferred], [defer_day], [exchange], [cycle], [maturity], [is_maturity]) VALUES (11, CAST(N'2023-11-01T09:29:47.5188540+08:00' AS DateTimeOffset), CAST(N'2024-01-22T10:07:52.3765938+08:00' AS DateTimeOffset), NULL, N'0008', N'喪假6天', 1, N'勞工請假規則：第 3 條 祖父母、子女、配偶之父母、配偶之養父母或繼父母喪亡者，給予喪假六日，工資照給。', 1, 1, 0, 6, 0, 0, 0, 1, 0, 0)
INSERT [dbo].[leave_core] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover], [is_pay], [minimum], [default_limit], [deferred], [defer_day], [exchange], [cycle], [maturity], [is_maturity]) VALUES (12, CAST(N'2023-11-01T09:30:28.7005056+08:00' AS DateTimeOffset), CAST(N'2024-01-22T10:07:58.8442782+08:00' AS DateTimeOffset), NULL, N'0009', N'喪假3天', 1, N'勞工請假規則：第 3 條 曾祖父母、兄弟姊妹、配偶之祖父母喪亡者，給予喪假三日，工資照給。', 1, 1, 0, 3, 0, 0, 0, 1, 0, 0)
INSERT [dbo].[leave_core] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover], [is_pay], [minimum], [default_limit], [deferred], [defer_day], [exchange], [cycle], [maturity], [is_maturity]) VALUES (13, CAST(N'2023-11-01T09:33:30.6886740+08:00' AS DateTimeOffset), CAST(N'2024-01-22T10:08:49.1293220+08:00' AS DateTimeOffset), NULL, N'0010', N'遲到', 1, N'緩衝次數5次，遲到5分鐘內可請"遲到假"， 超過6分鐘即須請事假或做其他假別申請，5次內不影響薪資不扣薪。', 1, 1, 0, 5, 0, 0, 0, 1, 0, 0)
SET IDENTITY_INSERT [dbo].[leave_core] OFF
GO
SET IDENTITY_INSERT [dbo].[leave_core_group] ON 

INSERT [dbo].[leave_core_group] ([id], [created_at], [updated_at], [deleted_at], [leave_core_id], [name]) VALUES (12, CAST(N'2024-01-17T14:59:47.5285662+08:00' AS DateTimeOffset), CAST(N'2024-01-17T14:59:47.5285662+08:00' AS DateTimeOffset), NULL, 6, N'群組1')
INSERT [dbo].[leave_core_group] ([id], [created_at], [updated_at], [deleted_at], [leave_core_id], [name]) VALUES (13, CAST(N'2024-01-17T15:12:11.8860126+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:12:11.8860126+08:00' AS DateTimeOffset), NULL, 6, N'群組2')
SET IDENTITY_INSERT [dbo].[leave_core_group] OFF
GO
SET IDENTITY_INSERT [dbo].[leave_core_group_condition] ON 

INSERT [dbo].[leave_core_group_condition] ([id], [created_at], [updated_at], [deleted_at], [leave_core_group_id], [condition], [cond1], [day], [level]) VALUES (13, CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), NULL, 12, 1, 6, 3, 1)
INSERT [dbo].[leave_core_group_condition] ([id], [created_at], [updated_at], [deleted_at], [leave_core_group_id], [condition], [cond1], [day], [level]) VALUES (14, CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), NULL, 12, 2, 1, 7, 2)
INSERT [dbo].[leave_core_group_condition] ([id], [created_at], [updated_at], [deleted_at], [leave_core_group_id], [condition], [cond1], [day], [level]) VALUES (15, CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), NULL, 12, 2, 2, 10, 3)
INSERT [dbo].[leave_core_group_condition] ([id], [created_at], [updated_at], [deleted_at], [leave_core_group_id], [condition], [cond1], [day], [level]) VALUES (16, CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), NULL, 12, 2, 3, 14, 4)
INSERT [dbo].[leave_core_group_condition] ([id], [created_at], [updated_at], [deleted_at], [leave_core_group_id], [condition], [cond1], [day], [level]) VALUES (17, CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), NULL, 12, 2, 5, 15, 5)
INSERT [dbo].[leave_core_group_condition] ([id], [created_at], [updated_at], [deleted_at], [leave_core_group_id], [condition], [cond1], [day], [level]) VALUES (18, CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), NULL, 12, 2, 10, 16, 6)
INSERT [dbo].[leave_core_group_condition] ([id], [created_at], [updated_at], [deleted_at], [leave_core_group_id], [condition], [cond1], [day], [level]) VALUES (19, CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), NULL, 12, 2, 11, 17, 7)
INSERT [dbo].[leave_core_group_condition] ([id], [created_at], [updated_at], [deleted_at], [leave_core_group_id], [condition], [cond1], [day], [level]) VALUES (20, CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), NULL, 12, 2, 12, 18, 8)
INSERT [dbo].[leave_core_group_condition] ([id], [created_at], [updated_at], [deleted_at], [leave_core_group_id], [condition], [cond1], [day], [level]) VALUES (21, CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), NULL, 12, 2, 13, 19, 9)
INSERT [dbo].[leave_core_group_condition] ([id], [created_at], [updated_at], [deleted_at], [leave_core_group_id], [condition], [cond1], [day], [level]) VALUES (22, CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), NULL, 12, 2, 14, 20, 10)
INSERT [dbo].[leave_core_group_condition] ([id], [created_at], [updated_at], [deleted_at], [leave_core_group_id], [condition], [cond1], [day], [level]) VALUES (23, CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), NULL, 12, 2, 15, 21, 11)
INSERT [dbo].[leave_core_group_condition] ([id], [created_at], [updated_at], [deleted_at], [leave_core_group_id], [condition], [cond1], [day], [level]) VALUES (24, CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), NULL, 12, 2, 16, 22, 12)
INSERT [dbo].[leave_core_group_condition] ([id], [created_at], [updated_at], [deleted_at], [leave_core_group_id], [condition], [cond1], [day], [level]) VALUES (25, CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), NULL, 12, 2, 17, 23, 13)
INSERT [dbo].[leave_core_group_condition] ([id], [created_at], [updated_at], [deleted_at], [leave_core_group_id], [condition], [cond1], [day], [level]) VALUES (26, CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), NULL, 12, 2, 18, 24, 14)
INSERT [dbo].[leave_core_group_condition] ([id], [created_at], [updated_at], [deleted_at], [leave_core_group_id], [condition], [cond1], [day], [level]) VALUES (27, CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), NULL, 12, 2, 19, 25, 15)
INSERT [dbo].[leave_core_group_condition] ([id], [created_at], [updated_at], [deleted_at], [leave_core_group_id], [condition], [cond1], [day], [level]) VALUES (28, CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), NULL, 12, 2, 20, 26, 16)
INSERT [dbo].[leave_core_group_condition] ([id], [created_at], [updated_at], [deleted_at], [leave_core_group_id], [condition], [cond1], [day], [level]) VALUES (29, CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), NULL, 12, 2, 21, 27, 17)
INSERT [dbo].[leave_core_group_condition] ([id], [created_at], [updated_at], [deleted_at], [leave_core_group_id], [condition], [cond1], [day], [level]) VALUES (30, CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), NULL, 12, 2, 22, 28, 18)
INSERT [dbo].[leave_core_group_condition] ([id], [created_at], [updated_at], [deleted_at], [leave_core_group_id], [condition], [cond1], [day], [level]) VALUES (31, CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), NULL, 12, 2, 23, 29, 19)
INSERT [dbo].[leave_core_group_condition] ([id], [created_at], [updated_at], [deleted_at], [leave_core_group_id], [condition], [cond1], [day], [level]) VALUES (32, CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:06:01.8009098+08:00' AS DateTimeOffset), NULL, 12, 2, 24, 30, 20)
INSERT [dbo].[leave_core_group_condition] ([id], [created_at], [updated_at], [deleted_at], [leave_core_group_id], [condition], [cond1], [day], [level]) VALUES (33, CAST(N'2024-01-17T15:12:11.8969013+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:12:11.8969013+08:00' AS DateTimeOffset), NULL, 13, 1, 0, 0, 1)
SET IDENTITY_INSERT [dbo].[leave_core_group_condition] OFF
GO
INSERT [dbo].[leave_core_group_user] ([leave_core_group_id], [user_id]) VALUES (12, 1)
INSERT [dbo].[leave_core_group_user] ([leave_core_group_id], [user_id]) VALUES (12, 3)
GO
SET IDENTITY_INSERT [dbo].[leave_correct] ON 

INSERT [dbo].[leave_correct] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [available], [used], [signing]) VALUES (287, CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), NULL, 1, 2, 12.25, 0, 0)
INSERT [dbo].[leave_correct] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [available], [used], [signing]) VALUES (288, CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), NULL, 1, 5, 30, 0, 0)
INSERT [dbo].[leave_correct] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [available], [used], [signing]) VALUES (289, CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), NULL, 1, 8, 5, 0, 0)
INSERT [dbo].[leave_correct] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [available], [used], [signing]) VALUES (290, CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), NULL, 1, 10, 8, 0, 0)
INSERT [dbo].[leave_correct] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [available], [used], [signing]) VALUES (291, CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), NULL, 1, 11, 6, 0, 0)
INSERT [dbo].[leave_correct] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [available], [used], [signing]) VALUES (292, CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), NULL, 1, 12, 3, 0, 0)
INSERT [dbo].[leave_correct] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [available], [used], [signing]) VALUES (293, CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), NULL, 1, 13, 5, 0, 0)
INSERT [dbo].[leave_correct] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [available], [used], [signing]) VALUES (294, CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), NULL, 1, 4, 255, 0, 0)
INSERT [dbo].[leave_correct] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [available], [used], [signing]) VALUES (295, CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), NULL, 1, 6, 4.8666666666666663, 0, 0)
INSERT [dbo].[leave_correct] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [available], [used], [signing]) VALUES (296, CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), NULL, 1, 7, 8, 0, 0)
INSERT [dbo].[leave_correct] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [available], [used], [signing]) VALUES (297, CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), NULL, 1, 9, 5, 0, 0)
INSERT [dbo].[leave_correct] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [available], [used], [signing]) VALUES (298, CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), NULL, 3, 6, 8.5416666666666679, 0, 0)
INSERT [dbo].[leave_correct] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [available], [used], [signing]) VALUES (299, CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), NULL, 3, 8, 5, 0, 0)
INSERT [dbo].[leave_correct] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [available], [used], [signing]) VALUES (300, CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), NULL, 3, 9, 5, 0, 0)
INSERT [dbo].[leave_correct] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [available], [used], [signing]) VALUES (301, CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), NULL, 3, 10, 8, 0, 0)
INSERT [dbo].[leave_correct] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [available], [used], [signing]) VALUES (302, CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), NULL, 3, 11, 6, 0, 0)
INSERT [dbo].[leave_correct] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [available], [used], [signing]) VALUES (303, CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), NULL, 3, 12, 3, 0, 0)
INSERT [dbo].[leave_correct] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [available], [used], [signing]) VALUES (304, CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), NULL, 3, 2, 14, 0, 0)
INSERT [dbo].[leave_correct] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [available], [used], [signing]) VALUES (305, CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), NULL, 3, 4, 255, 0, 0)
INSERT [dbo].[leave_correct] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [available], [used], [signing]) VALUES (306, CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), NULL, 3, 5, 30, 0, 0)
INSERT [dbo].[leave_correct] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [available], [used], [signing]) VALUES (307, CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), NULL, 3, 7, 8, 0, 0)
INSERT [dbo].[leave_correct] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [available], [used], [signing]) VALUES (308, CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7512748+08:00' AS DateTimeOffset), NULL, 3, 13, 5, 0, 0)
SET IDENTITY_INSERT [dbo].[leave_correct] OFF
GO
SET IDENTITY_INSERT [dbo].[leave_defer] ON 

INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (1, CAST(N'2024-01-25T11:07:34.7695336+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:17:30.5774947+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:17:30.5798363+08:00' AS DateTimeOffset), 1, 6, N'2024-01-25', N'2024-01-25', 0, 0, 0, 0, 10)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (2, CAST(N'2024-01-25T11:07:34.8144416+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:17:22.2664872+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:17:22.2706956+08:00' AS DateTimeOffset), 3, 6, N'2024-01-25', N'2024-01-25', 0, 0, 0, 0, 9)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (9, CAST(N'2024-01-25T11:17:22.2616921+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:20:36.8490321+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:20:36.8510631+08:00' AS DateTimeOffset), 3, 6, N'2024-01-25', N'2024-01-25', 9, 15, 6, 2, 13)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (10, CAST(N'2024-01-25T11:17:30.5750040+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:20:36.7898318+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:20:36.7940262+08:00' AS DateTimeOffset), 1, 6, N'2024-01-25', N'2024-01-25', 2, 2, 0, 1, 12)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (12, CAST(N'2024-01-25T11:20:36.7863580+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:23:16.7165223+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:23:16.7208497+08:00' AS DateTimeOffset), 1, 6, N'2024-01-25', N'2024-01-25', 4, 2, 0, 10, 16)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (13, CAST(N'2024-01-25T11:20:36.8467590+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:23:16.7749960+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:23:16.7775364+08:00' AS DateTimeOffset), 3, 6, N'2024-01-25', N'2024-01-25', 9, 15, 15, 9, 17)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (16, CAST(N'2024-01-25T11:23:16.7123258+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:24:13.0952733+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:24:13.0973924+08:00' AS DateTimeOffset), 1, 6, N'2024-01-25', N'2024-01-25', 6, 2, 0, 12, 19)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (17, CAST(N'2024-01-25T11:23:16.7722262+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:24:10.2260374+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:24:10.2295934+08:00' AS DateTimeOffset), 3, 6, N'2024-01-25', N'2024-01-25', 9, 10, 10, 13, 18)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (18, CAST(N'2024-01-25T11:24:10.2223257+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:24:20.3214182+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:24:20.3237226+08:00' AS DateTimeOffset), 3, 6, N'2024-01-25', N'2024-01-25', 9, 13.731182795698924, 13.731182795698924, 17, 21)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (19, CAST(N'2024-01-25T11:24:13.0930798+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:24:20.2644420+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:24:20.2672069+08:00' AS DateTimeOffset), 1, 6, N'2024-01-25', N'2024-01-25', 8, 2, 0, 16, 20)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (20, CAST(N'2024-01-25T11:24:20.2621088+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:24:53.7025582+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:24:53.7049223+08:00' AS DateTimeOffset), 1, 6, N'2024-01-25', N'2024-01-25', 9, 2, 1, 19, 22)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (21, CAST(N'2024-01-25T11:24:20.3186632+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:24:53.7419762+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:24:53.7445522+08:00' AS DateTimeOffset), 3, 6, N'2024-01-25', N'2024-01-25', 9, 14, 14, 18, 23)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (22, CAST(N'2024-01-25T11:24:53.7000153+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:26:41.6350959+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:26:41.6394904+08:00' AS DateTimeOffset), 1, 6, N'2024-01-25', N'2024-01-25', 9, 2, 2, 20, 26)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (23, CAST(N'2024-01-25T11:24:53.7397432+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:26:41.6688499+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:26:41.6712371+08:00' AS DateTimeOffset), 3, 6, N'2024-01-25', N'2024-01-25', 9, 14, 14, 21, 27)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (26, CAST(N'2024-01-25T11:26:41.6305571+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:26:52.0986355+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:26:52.1019351+08:00' AS DateTimeOffset), 1, 6, N'2024-01-25', N'2024-01-25', 9, 2, 2, 22, 28)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (27, CAST(N'2024-01-25T11:26:41.6662255+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:26:52.1439464+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:26:52.1464844+08:00' AS DateTimeOffset), 3, 6, N'2024-01-25', N'2024-01-25', 9, 9.7983870967741939, 9.7983870967741922, 23, 29)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (28, CAST(N'2024-01-25T11:26:52.0957996+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:28:55.5156139+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:28:55.5214233+08:00' AS DateTimeOffset), 1, 6, N'2024-01-25', N'2024-01-25', 9, 2, 2, 26, 30)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (29, CAST(N'2024-01-25T11:26:52.1416432+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:33:59.4976051+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:33:59.5022915+08:00' AS DateTimeOffset), 3, 6, N'2024-01-25', N'2024-01-25', 9, 2, 2, 27, 31)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (30, CAST(N'2024-01-25T11:28:55.5102950+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:55:01.8579564+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:55:01.8626012+08:00' AS DateTimeOffset), 1, 6, N'2024-01-25', N'2024-01-25', 9, 2, 2, 28, 38)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (31, CAST(N'2024-01-25T11:33:59.4937569+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:55:03.4094187+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:55:03.4117549+08:00' AS DateTimeOffset), 3, 6, N'2024-01-25', N'2024-01-25', 9, 2, 2, 29, 39)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (38, CAST(N'2024-01-25T11:55:01.8540755+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:58:37.6830334+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:58:37.6872684+08:00' AS DateTimeOffset), 1, 6, N'2024-01-25', N'2024-01-25', 9, 2, 2, 30, 40)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (39, CAST(N'2024-01-25T11:55:03.4064417+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:58:37.7601204+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:58:37.7637635+08:00' AS DateTimeOffset), 3, 6, N'2024-01-25', N'2024-01-25', 9, 2, 2, 31, 41)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (40, CAST(N'2024-01-25T11:58:37.6777600+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:33.4217899+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:33.4259342+08:00' AS DateTimeOffset), 1, 6, N'2024-01-25', N'2024-01-25', 9, 4.8666666666666663, 4.8666666666666671, 38, 42)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (41, CAST(N'2024-01-25T11:58:37.7575356+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:33.4758798+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:33.4786734+08:00' AS DateTimeOffset), 3, 6, N'2024-01-25', N'2024-01-25', 9, 5.446236559139785, 5.4462365591397841, 39, 43)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (42, CAST(N'2024-01-25T12:00:33.4164444+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.6948344+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.6975648+08:00' AS DateTimeOffset), 1, 6, N'2024-01-25', N'2024-01-25', 9, 4.8666666666666663, 4.8666666666666671, 40, 44)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (43, CAST(N'2024-01-25T12:00:33.4732325+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7182725+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7206551+08:00' AS DateTimeOffset), 3, 6, N'2024-01-25', N'2024-01-25', 9, 6.029569892473118, 6.029569892473118, 41, 45)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (44, CAST(N'2024-01-25T12:00:53.6919368+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.6919368+08:00' AS DateTimeOffset), NULL, 1, 6, N'2024-01-25', N'2024-06-25', 9, 4.8666666666666663, 4.8666666666666671, 42, 0)
INSERT [dbo].[leave_defer] ([id], [created_at], [updated_at], [deleted_at], [user_id], [leave_core_id], [effective], [expired], [available], [this_time], [extra], [previous_id], [next_id]) VALUES (45, CAST(N'2024-01-25T12:00:53.7155879+08:00' AS DateTimeOffset), CAST(N'2024-01-25T12:00:53.7155879+08:00' AS DateTimeOffset), NULL, 3, 6, N'2024-01-25', N'2024-06-25', 9, 6.5972222222222223, 6.5972222222222214, 43, 0)
SET IDENTITY_INSERT [dbo].[leave_defer] OFF
GO
SET IDENTITY_INSERT [dbo].[leave_request_form] ON 

INSERT [dbo].[leave_request_form] ([id], [created_at], [updated_at], [deleted_at], [order], [start_date], [start_time], [end_date], [end_time], [remark], [sign_status], [leave_minute_count], [leave_hour_count], [leave_day_count], [attach], [leave_core_id], [user_id], [department_id], [proxy_user_id], [proxy_department_id]) VALUES (3, CAST(N'2023-12-13T16:46:56.8381482+08:00' AS DateTimeOffset), CAST(N'2024-01-02T15:22:12.6368792+08:00' AS DateTimeOffset), NULL, N'A20231213000003', N'2023-12-01', N'08:00', N'2023-12-01', N'17:00', N'', 2, 480, 8, 1, N'', 2, 1, 1, 3, 1)
INSERT [dbo].[leave_request_form] ([id], [created_at], [updated_at], [deleted_at], [order], [start_date], [start_time], [end_date], [end_time], [remark], [sign_status], [leave_minute_count], [leave_hour_count], [leave_day_count], [attach], [leave_core_id], [user_id], [department_id], [proxy_user_id], [proxy_department_id]) VALUES (19, CAST(N'2023-12-19T16:24:42.9777518+08:00' AS DateTimeOffset), CAST(N'2024-01-02T13:40:08.0221404+08:00' AS DateTimeOffset), NULL, N'A20231219000019', N'2023-12-21', N'08:00', N'2023-12-21', N'23:00', N'', 1, 960, 16, 2, N'', 2, 1, 1, 3, 1)
INSERT [dbo].[leave_request_form] ([id], [created_at], [updated_at], [deleted_at], [order], [start_date], [start_time], [end_date], [end_time], [remark], [sign_status], [leave_minute_count], [leave_hour_count], [leave_day_count], [attach], [leave_core_id], [user_id], [department_id], [proxy_user_id], [proxy_department_id]) VALUES (34, CAST(N'2023-12-27T10:44:07.5585247+08:00' AS DateTimeOffset), CAST(N'2023-12-27T10:44:46.6862458+08:00' AS DateTimeOffset), NULL, N'A20231227000034', N'2023-12-22', N'08:00', N'2023-12-22', N'17:00', N'', 3, 480, 8, 1, N'', 2, 1, 1, 3, 1)
INSERT [dbo].[leave_request_form] ([id], [created_at], [updated_at], [deleted_at], [order], [start_date], [start_time], [end_date], [end_time], [remark], [sign_status], [leave_minute_count], [leave_hour_count], [leave_day_count], [attach], [leave_core_id], [user_id], [department_id], [proxy_user_id], [proxy_department_id]) VALUES (38, CAST(N'2023-12-27T16:45:03.0859908+08:00' AS DateTimeOffset), CAST(N'2023-12-27T16:47:25.2711496+08:00' AS DateTimeOffset), NULL, N'A20231227000038', N'2023-12-27', N'08:00', N'2023-12-27', N'17:00', N'', 2, 0, 0, 0, N'', 2, 1, 1, 3, 1)
INSERT [dbo].[leave_request_form] ([id], [created_at], [updated_at], [deleted_at], [order], [start_date], [start_time], [end_date], [end_time], [remark], [sign_status], [leave_minute_count], [leave_hour_count], [leave_day_count], [attach], [leave_core_id], [user_id], [department_id], [proxy_user_id], [proxy_department_id]) VALUES (40, CAST(N'2024-01-10T14:04:25.3144785+08:00' AS DateTimeOffset), CAST(N'2024-01-10T14:32:14.7308800+08:00' AS DateTimeOffset), NULL, N'A20240110000040', N'2024-01-10', N'08:00', N'2024-01-12', N'14:00', N'', 2, 900, 15, 1.8799999952316284, N'', 2, 1, 1, 3, 1)
INSERT [dbo].[leave_request_form] ([id], [created_at], [updated_at], [deleted_at], [order], [start_date], [start_time], [end_date], [end_time], [remark], [sign_status], [leave_minute_count], [leave_hour_count], [leave_day_count], [attach], [leave_core_id], [user_id], [department_id], [proxy_user_id], [proxy_department_id]) VALUES (41, CAST(N'2024-01-10T15:18:34.5435972+08:00' AS DateTimeOffset), CAST(N'2024-01-10T15:45:30.7442200+08:00' AS DateTimeOffset), NULL, N'A20240110000041', N'2024-01-09', N'08:00', N'2024-01-09', N'12:00', N'', 3, 240, 4, 0.5, N'', 2, 1, 1, 3, 1)
INSERT [dbo].[leave_request_form] ([id], [created_at], [updated_at], [deleted_at], [order], [start_date], [start_time], [end_date], [end_time], [remark], [sign_status], [leave_minute_count], [leave_hour_count], [leave_day_count], [attach], [leave_core_id], [user_id], [department_id], [proxy_user_id], [proxy_department_id]) VALUES (42, CAST(N'2024-01-10T16:01:34.9429434+08:00' AS DateTimeOffset), CAST(N'2024-01-10T16:01:53.6340463+08:00' AS DateTimeOffset), NULL, N'A20240110000042', N'2024-01-08', N'08:00', N'2024-01-08', N'14:00', N'', 2, 300, 5, 0.62999999523162842, N'', 2, 1, 1, 3, 1)
INSERT [dbo].[leave_request_form] ([id], [created_at], [updated_at], [deleted_at], [order], [start_date], [start_time], [end_date], [end_time], [remark], [sign_status], [leave_minute_count], [leave_hour_count], [leave_day_count], [attach], [leave_core_id], [user_id], [department_id], [proxy_user_id], [proxy_department_id]) VALUES (51, CAST(N'2024-01-12T15:30:26.5872818+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:30:56.7604417+08:00' AS DateTimeOffset), NULL, N'A20240112000051', N'2024-01-15', N'08:00', N'2024-01-15', N'14:00', N'', 2, 300, 5, 0.62999999523162842, N'', 2, 1, 1, 3, 1)
INSERT [dbo].[leave_request_form] ([id], [created_at], [updated_at], [deleted_at], [order], [start_date], [start_time], [end_date], [end_time], [remark], [sign_status], [leave_minute_count], [leave_hour_count], [leave_day_count], [attach], [leave_core_id], [user_id], [department_id], [proxy_user_id], [proxy_department_id]) VALUES (55, CAST(N'2024-02-16T09:29:58.5020843+08:00' AS DateTimeOffset), CAST(N'2024-02-16T14:15:07.5520726+08:00' AS DateTimeOffset), NULL, N'A20240216000055', N'2024-02-16', N'08:00', N'2024-02-16', N'17:00', N'', 2, 480, 8, 1, N'', 2, 1, 1, 3, 1)
INSERT [dbo].[leave_request_form] ([id], [created_at], [updated_at], [deleted_at], [order], [start_date], [start_time], [end_date], [end_time], [remark], [sign_status], [leave_minute_count], [leave_hour_count], [leave_day_count], [attach], [leave_core_id], [user_id], [department_id], [proxy_user_id], [proxy_department_id]) VALUES (61, CAST(N'2024-02-16T14:35:14.1136574+08:00' AS DateTimeOffset), CAST(N'2024-02-16T14:35:51.6801295+08:00' AS DateTimeOffset), NULL, N'A20240216000061', N'2024-02-17', N'08:00', N'2024-02-17', N'17:00', N'', 2, 480, 8, 1, N'assets/A20240216000061/ghostbb-og.png', 2, 1, 1, 3, 1)
INSERT [dbo].[leave_request_form] ([id], [created_at], [updated_at], [deleted_at], [order], [start_date], [start_time], [end_date], [end_time], [remark], [sign_status], [leave_minute_count], [leave_hour_count], [leave_day_count], [attach], [leave_core_id], [user_id], [department_id], [proxy_user_id], [proxy_department_id]) VALUES (65, CAST(N'2024-02-19T10:21:31.9328234+08:00' AS DateTimeOffset), CAST(N'2024-02-19T10:22:35.8284051+08:00' AS DateTimeOffset), NULL, N'A20240219000065', N'2024-02-19', N'08:00', N'2024-02-19', N'17:00', N'', 2, 480, 8, 1, N'', 2, 1, 1, 3, 1)
INSERT [dbo].[leave_request_form] ([id], [created_at], [updated_at], [deleted_at], [order], [start_date], [start_time], [end_date], [end_time], [remark], [sign_status], [leave_minute_count], [leave_hour_count], [leave_day_count], [attach], [leave_core_id], [user_id], [department_id], [proxy_user_id], [proxy_department_id]) VALUES (66, CAST(N'2024-02-19T10:25:17.5921363+08:00' AS DateTimeOffset), CAST(N'2024-02-19T11:02:46.5377798+08:00' AS DateTimeOffset), NULL, N'A20240219000066', N'2024-02-20', N'08:00', N'2024-02-20', N'15:00', N'', 2, 360, 6, 0.75, N'', 2, 1, 1, 3, 1)
INSERT [dbo].[leave_request_form] ([id], [created_at], [updated_at], [deleted_at], [order], [start_date], [start_time], [end_date], [end_time], [remark], [sign_status], [leave_minute_count], [leave_hour_count], [leave_day_count], [attach], [leave_core_id], [user_id], [department_id], [proxy_user_id], [proxy_department_id]) VALUES (67, CAST(N'2024-02-19T11:29:16.0537918+08:00' AS DateTimeOffset), CAST(N'2024-02-19T11:29:27.0202544+08:00' AS DateTimeOffset), NULL, N'A20240219000067', N'2024-02-21', N'08:00', N'2024-02-21', N'15:00', N'', 3, 360, 6, 0.75, N'', 2, 1, 1, 3, 1)
INSERT [dbo].[leave_request_form] ([id], [created_at], [updated_at], [deleted_at], [order], [start_date], [start_time], [end_date], [end_time], [remark], [sign_status], [leave_minute_count], [leave_hour_count], [leave_day_count], [attach], [leave_core_id], [user_id], [department_id], [proxy_user_id], [proxy_department_id]) VALUES (69, CAST(N'2024-02-19T11:35:27.1032107+08:00' AS DateTimeOffset), CAST(N'2024-02-19T11:35:27.1302306+08:00' AS DateTimeOffset), NULL, N'A20240219000069', N'2024-02-21', N'08:00', N'2024-02-21', N'14:00', N'', 0, 300, 5, 0.625, N'assets/A20240219000069/ghostbb-og.png', 2, 1, 1, 3, 1)
SET IDENTITY_INSERT [dbo].[leave_request_form] OFF
GO
SET IDENTITY_INSERT [dbo].[leave_sign_off_flow] ON 

INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (11, CAST(N'2023-12-13T16:47:00.7231740+08:00' AS DateTimeOffset), CAST(N'2023-12-13T16:48:48.6704704+08:00' AS DateTimeOffset), NULL, 3, 3, 1, 3, 1, N'', N'核准', 1, CAST(N'2023-12-13T16:48:48.6698163+08:00' AS DateTimeOffset), N'940c8ce0-bb0c-49d3-9568-03ca66b54a5c', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (12, CAST(N'2023-12-13T16:47:00.7231740+08:00' AS DateTimeOffset), CAST(N'2024-01-02T15:22:12.5736486+08:00' AS DateTimeOffset), NULL, 3, 3, 2, 1, 1, N'test', N'test', 1, CAST(N'2024-01-02T15:22:12.5726802+08:00' AS DateTimeOffset), N'0736a827-de62-436f-b6ca-6fd2e57770e4', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (13, CAST(N'2023-12-13T16:47:00.7231740+08:00' AS DateTimeOffset), CAST(N'2024-01-02T15:22:12.6120337+08:00' AS DateTimeOffset), NULL, 3, 3, 3, 4, 2, N'', N'', 3, CAST(N'2024-01-02T15:22:12.6103253+08:00' AS DateTimeOffset), N'207e3b23-c039-4b4d-a6b4-7598f64a869a', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (14, CAST(N'2023-12-19T16:24:43.0678891+08:00' AS DateTimeOffset), CAST(N'2024-01-02T13:40:08.0160615+08:00' AS DateTimeOffset), NULL, 19, 3, 1, 3, 1, N'', N'test', 1, CAST(N'2024-01-02T13:40:08.0160615+08:00' AS DateTimeOffset), N'96707828-b404-43c6-a5ad-8cd71d6ad5d8', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (15, CAST(N'2023-12-19T16:24:43.0678891+08:00' AS DateTimeOffset), CAST(N'2024-01-02T13:40:08.0594453+08:00' AS DateTimeOffset), NULL, 19, 3, 2, 1, 1, N'test', N'', 4, CAST(N'0001-01-01T00:00:00.0000000+00:00' AS DateTimeOffset), N'28bf1c7d-de5d-44f4-8054-696479339424', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (16, CAST(N'2023-12-19T16:24:43.0678891+08:00' AS DateTimeOffset), CAST(N'2023-12-19T16:24:43.0678891+08:00' AS DateTimeOffset), NULL, 19, 3, 3, 4, 2, N'', N'', 0, CAST(N'0001-01-01T00:00:00.0000000+00:00' AS DateTimeOffset), N'778ea36f-f2bc-4cff-812d-fe9cc346809a', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (61, CAST(N'2023-12-27T10:44:07.6033044+08:00' AS DateTimeOffset), CAST(N'2023-12-27T10:44:46.7196172+08:00' AS DateTimeOffset), NULL, 34, 3, 1, 3, 1, N'', N'', 1, CAST(N'2023-12-27T10:44:25.4387693+08:00' AS DateTimeOffset), N'b3ecb3f2-2298-4c81-bb67-7701b8bd2c8e', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (62, CAST(N'2023-12-27T10:44:07.6033044+08:00' AS DateTimeOffset), CAST(N'2023-12-27T10:44:46.7196172+08:00' AS DateTimeOffset), NULL, 34, 3, 2, 1, 1, N'test', N'', 2, CAST(N'2023-12-27T10:44:46.6862458+08:00' AS DateTimeOffset), N'7733ac66-9980-48a7-8bcd-97779a776544', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (63, CAST(N'2023-12-27T10:44:07.6033044+08:00' AS DateTimeOffset), CAST(N'2023-12-27T10:44:46.7196172+08:00' AS DateTimeOffset), NULL, 34, 1, 3, 4, 2, N'', N'', 0, CAST(N'0001-01-01T00:00:00.0000000+00:00' AS DateTimeOffset), N'9eceb60a-047e-4dd7-b59f-7d2d1a0c6d81', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (70, CAST(N'2023-12-27T16:45:03.1392764+08:00' AS DateTimeOffset), CAST(N'2023-12-27T16:47:13.0404552+08:00' AS DateTimeOffset), NULL, 38, 3, 1, 3, 1, N'', N'', 1, CAST(N'2023-12-27T16:47:13.0398381+08:00' AS DateTimeOffset), N'55a90982-5a6d-4f52-96b9-b310c9919d81', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (71, CAST(N'2023-12-27T16:45:03.1392764+08:00' AS DateTimeOffset), CAST(N'2023-12-27T16:47:25.2474402+08:00' AS DateTimeOffset), NULL, 38, 3, 2, 1, 1, N'test', N'', 1, CAST(N'2023-12-27T16:47:25.2468975+08:00' AS DateTimeOffset), N'3c93e788-6d07-4f7f-87e7-179040b71e3e', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (72, CAST(N'2023-12-27T16:45:03.1392764+08:00' AS DateTimeOffset), CAST(N'2023-12-27T16:47:25.2618150+08:00' AS DateTimeOffset), NULL, 38, 1, 3, 4, 2, N'', N'', 3, CAST(N'2023-12-27T16:47:25.2612153+08:00' AS DateTimeOffset), N'303dcc4b-c320-47a3-bb76-012404a921b2', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (76, CAST(N'2024-01-10T14:04:25.4460799+08:00' AS DateTimeOffset), CAST(N'2024-01-10T14:30:35.2227632+08:00' AS DateTimeOffset), NULL, 40, 3, 1, 3, 1, N'', N'', 1, CAST(N'2024-01-10T14:30:35.2227632+08:00' AS DateTimeOffset), N'3e4e5d35-6d3c-4a18-b6fb-93a83fe5459f', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (77, CAST(N'2024-01-10T14:04:25.4460799+08:00' AS DateTimeOffset), CAST(N'2024-01-10T14:32:07.5747443+08:00' AS DateTimeOffset), NULL, 40, 3, 2, 1, 1, N'test', N'', 1, CAST(N'2024-01-10T14:32:07.5747443+08:00' AS DateTimeOffset), N'2f2044ce-0aa0-4b72-8d44-df57a3bbdc19', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (78, CAST(N'2024-01-10T14:04:25.4460799+08:00' AS DateTimeOffset), CAST(N'2024-01-10T14:30:53.2522233+08:00' AS DateTimeOffset), NULL, 40, 1, 3, 4, 2, N'', N'', 3, CAST(N'2024-01-10T14:30:53.2516085+08:00' AS DateTimeOffset), N'019c32dd-0ed9-48c8-8554-0ed524381d5e', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (79, CAST(N'2024-01-10T15:18:34.6432372+08:00' AS DateTimeOffset), CAST(N'2024-01-10T15:45:30.7542215+08:00' AS DateTimeOffset), NULL, 41, 3, 1, 3, 1, N'', N'test', 2, CAST(N'2024-01-10T15:45:30.7542215+08:00' AS DateTimeOffset), N'a2c48dd8-4752-46d1-ad07-2f06a2394d0b', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (80, CAST(N'2024-01-10T15:18:34.6432372+08:00' AS DateTimeOffset), CAST(N'2024-01-10T15:18:34.6432372+08:00' AS DateTimeOffset), NULL, 41, 3, 2, 1, 1, N'test', N'', 0, CAST(N'0001-01-01T00:00:00.0000000+00:00' AS DateTimeOffset), N'5e582e91-747e-4400-b920-dffa57bdd39a', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (81, CAST(N'2024-01-10T15:18:34.6432372+08:00' AS DateTimeOffset), CAST(N'2024-01-10T15:18:34.6432372+08:00' AS DateTimeOffset), NULL, 41, 1, 3, 4, 2, N'', N'', 0, CAST(N'0001-01-01T00:00:00.0000000+00:00' AS DateTimeOffset), N'7beeec30-d022-453b-8665-ff9bdbd905aa', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (82, CAST(N'2024-01-10T16:01:35.0389765+08:00' AS DateTimeOffset), CAST(N'2024-01-10T16:01:44.8123952+08:00' AS DateTimeOffset), NULL, 42, 3, 1, 3, 1, N'', N'', 1, CAST(N'2024-01-10T16:01:44.8123952+08:00' AS DateTimeOffset), N'c6e34d03-607f-4be2-95f5-281f06c85c8c', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (83, CAST(N'2024-01-10T16:01:35.0389765+08:00' AS DateTimeOffset), CAST(N'2024-01-10T16:01:53.5961031+08:00' AS DateTimeOffset), NULL, 42, 3, 2, 1, 1, N'test', N'', 1, CAST(N'2024-01-10T16:01:53.5961031+08:00' AS DateTimeOffset), N'2bf26532-6b38-4f8e-b6a0-7f0be1fa3d96', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (84, CAST(N'2024-01-10T16:01:35.0389765+08:00' AS DateTimeOffset), CAST(N'2024-01-10T16:01:53.6269163+08:00' AS DateTimeOffset), NULL, 42, 1, 3, 4, 2, N'', N'', 3, CAST(N'2024-01-10T16:01:53.6269163+08:00' AS DateTimeOffset), N'7d6691f7-11a0-444c-8700-25bdd7cd6770', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (94, CAST(N'2024-01-12T15:30:26.6789699+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:30:45.1084998+08:00' AS DateTimeOffset), NULL, 51, 3, 1, 3, 1, N'', N'test', 1, CAST(N'2024-01-12T15:30:45.1084998+08:00' AS DateTimeOffset), N'377cc5b7-dd9b-4119-8f14-0990a9237fb0', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (95, CAST(N'2024-01-12T15:30:26.6789699+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:30:56.6656941+08:00' AS DateTimeOffset), NULL, 51, 3, 2, 1, 1, N'test', N't1', 1, CAST(N'2024-01-12T15:30:56.6656941+08:00' AS DateTimeOffset), N'ea1ef4d9-b81b-4e45-b257-501729d7cf8a', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (96, CAST(N'2024-01-12T15:30:26.6789699+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:30:56.7338056+08:00' AS DateTimeOffset), NULL, 51, 1, 3, 4, 2, N'', N'', 3, CAST(N'2024-01-12T15:30:56.7338056+08:00' AS DateTimeOffset), N'921583fe-2620-413f-8e3e-5f26baeb3302', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (103, CAST(N'2024-02-16T09:29:58.5266465+08:00' AS DateTimeOffset), CAST(N'2024-02-16T14:02:48.6971824+08:00' AS DateTimeOffset), NULL, 55, 3, 1, 3, 1, N'', N'test', 1, CAST(N'2024-02-16T14:02:48.6971824+08:00' AS DateTimeOffset), N'34c7f2e7-893e-4a87-a432-7ba7ff45c1dc', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (104, CAST(N'2024-02-16T09:29:58.5266465+08:00' AS DateTimeOffset), CAST(N'2024-02-16T14:15:07.4748181+08:00' AS DateTimeOffset), NULL, 55, 3, 2, 1, 1, N'test', N'test123', 1, CAST(N'2024-02-16T14:15:07.4748181+08:00' AS DateTimeOffset), N'6ee13d2d-e2b9-4008-be0f-a2c87a6b502c', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (105, CAST(N'2024-02-16T09:29:58.5266465+08:00' AS DateTimeOffset), CAST(N'2024-02-16T14:15:07.5400939+08:00' AS DateTimeOffset), NULL, 55, 1, 3, 4, 2, N'', N'', 3, CAST(N'2024-02-16T14:15:07.5400939+08:00' AS DateTimeOffset), N'4c36b1e5-3b4a-4af7-a736-c0636561b48a', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (112, CAST(N'2024-02-16T14:35:14.1571848+08:00' AS DateTimeOffset), CAST(N'2024-02-16T14:35:40.2356065+08:00' AS DateTimeOffset), NULL, 61, 3, 1, 3, 1, N'', N'', 1, CAST(N'2024-02-16T14:35:40.2356065+08:00' AS DateTimeOffset), N'bc592a37-eccb-4e2d-8f32-78008e6bd0da', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (113, CAST(N'2024-02-16T14:35:14.1571848+08:00' AS DateTimeOffset), CAST(N'2024-02-16T14:35:51.6242934+08:00' AS DateTimeOffset), NULL, 61, 3, 2, 1, 1, N'test', N'', 1, CAST(N'2024-02-16T14:35:51.6242934+08:00' AS DateTimeOffset), N'16e7a7e0-8ae9-4660-94f6-65e9d1fbb8bc', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (114, CAST(N'2024-02-16T14:35:14.1571848+08:00' AS DateTimeOffset), CAST(N'2024-02-16T14:35:51.6641122+08:00' AS DateTimeOffset), NULL, 61, 1, 3, 4, 2, N'', N'', 3, CAST(N'2024-02-16T14:35:51.6641122+08:00' AS DateTimeOffset), N'5066a35c-db06-4726-91e7-7e738440c095', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (124, CAST(N'2024-02-19T10:21:32.0036936+08:00' AS DateTimeOffset), CAST(N'2024-02-19T10:22:09.5687086+08:00' AS DateTimeOffset), NULL, 65, 3, 1, 3, 1, N'', N'', 1, CAST(N'2024-02-19T10:22:09.5687086+08:00' AS DateTimeOffset), N'078562ff-e07f-4e63-9233-5f67d0d6653b', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (125, CAST(N'2024-02-19T10:21:32.0036936+08:00' AS DateTimeOffset), CAST(N'2024-02-19T10:22:19.1533516+08:00' AS DateTimeOffset), NULL, 65, 3, 2, 1, 1, N'test', N'', 1, CAST(N'2024-02-19T10:22:19.1528455+08:00' AS DateTimeOffset), N'6cbdd085-2dd6-49f8-9c11-717742192f5c', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (126, CAST(N'2024-02-19T10:21:32.0036936+08:00' AS DateTimeOffset), CAST(N'2024-02-19T10:22:35.7648023+08:00' AS DateTimeOffset), NULL, 65, 1, 3, 4, 2, N'', N'', 3, CAST(N'2024-02-19T10:22:35.7648023+08:00' AS DateTimeOffset), N'52f81e04-d05f-4c4c-a934-40bb294e822f', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (127, CAST(N'2024-02-19T10:25:17.6291352+08:00' AS DateTimeOffset), CAST(N'2024-02-19T10:25:28.2934925+08:00' AS DateTimeOffset), NULL, 66, 3, 1, 3, 1, N'', N'', 1, CAST(N'2024-02-19T10:25:28.2934925+08:00' AS DateTimeOffset), N'91916a32-ed86-40ff-913e-c028d4c90dcd', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (128, CAST(N'2024-02-19T10:25:17.6291352+08:00' AS DateTimeOffset), CAST(N'2024-02-19T11:02:37.4731650+08:00' AS DateTimeOffset), NULL, 66, 3, 2, 1, 1, N'test', N'', 1, CAST(N'2024-02-19T11:02:37.4731650+08:00' AS DateTimeOffset), N'd8354f97-5213-4646-aa7a-e78d8027fbb4', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (129, CAST(N'2024-02-19T10:25:17.6291352+08:00' AS DateTimeOffset), CAST(N'2024-02-19T11:02:46.5169546+08:00' AS DateTimeOffset), NULL, 66, 1, 3, 4, 2, N'', N'', 3, CAST(N'2024-02-19T11:02:46.5169546+08:00' AS DateTimeOffset), N'856122ec-76ac-45a4-afcd-9e8af03ff161', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (130, CAST(N'2024-02-19T11:29:16.1329327+08:00' AS DateTimeOffset), CAST(N'2024-02-19T11:29:27.0346011+08:00' AS DateTimeOffset), NULL, 67, 3, 1, 3, 1, N'', N'', 2, CAST(N'2024-02-19T11:29:27.0340069+08:00' AS DateTimeOffset), N'fcf29223-26c2-4105-80a8-f048c8639dc3', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (131, CAST(N'2024-02-19T11:29:16.1329327+08:00' AS DateTimeOffset), CAST(N'2024-02-19T11:29:16.1329327+08:00' AS DateTimeOffset), NULL, 67, 3, 2, 1, 1, N'test', N'', 0, CAST(N'0001-01-01T00:00:00.0000000+00:00' AS DateTimeOffset), N'672c572e-2aba-402f-a05f-aa05e25b7181', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (132, CAST(N'2024-02-19T11:29:16.1329327+08:00' AS DateTimeOffset), CAST(N'2024-02-19T11:29:16.1329327+08:00' AS DateTimeOffset), NULL, 67, 1, 3, 4, 2, N'', N'', 0, CAST(N'0001-01-01T00:00:00.0000000+00:00' AS DateTimeOffset), N'd99519bc-ec88-48ae-8251-2276aae7f3dd', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (133, CAST(N'2024-02-19T11:35:27.1451117+08:00' AS DateTimeOffset), CAST(N'2024-02-19T11:35:27.1451117+08:00' AS DateTimeOffset), NULL, 69, 3, 1, 3, 1, N'', N'', 4, CAST(N'0001-01-01T00:00:00.0000000+00:00' AS DateTimeOffset), N'1e7907d9-1324-4bf8-b061-4a96c653757a', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (134, CAST(N'2024-02-19T11:35:27.1451117+08:00' AS DateTimeOffset), CAST(N'2024-02-19T11:35:27.1451117+08:00' AS DateTimeOffset), NULL, 69, 3, 2, 1, 1, N'test', N'', 0, CAST(N'0001-01-01T00:00:00.0000000+00:00' AS DateTimeOffset), N'79a46598-c930-467c-b07a-2b62ed0bb000', N'zh_TW')
INSERT [dbo].[leave_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [leave_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (135, CAST(N'2024-02-19T11:35:27.1451117+08:00' AS DateTimeOffset), CAST(N'2024-02-19T11:35:27.1451117+08:00' AS DateTimeOffset), NULL, 69, 1, 3, 4, 2, N'', N'', 0, CAST(N'0001-01-01T00:00:00.0000000+00:00' AS DateTimeOffset), N'166cdc89-8c44-414f-90de-16f7faabed3b', N'zh_TW')
SET IDENTITY_INSERT [dbo].[leave_sign_off_flow] OFF
GO
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (4, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (4, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (5, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (5, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (7, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (7, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (8, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (8, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (9, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (9, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (10, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (10, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (11, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (11, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (12, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (12, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (13, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (13, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (14, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (14, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (15, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (15, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (16, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (16, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (17, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (17, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (18, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (18, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (19, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (19, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (20, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (20, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (21, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (21, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (22, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (22, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (23, 1)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (23, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (23, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (24, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (24, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (25, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (25, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (26, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (26, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (27, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (27, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (28, 8)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (29, 8)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (30, 8)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (31, 8)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (32, 8)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (33, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (33, 9)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (34, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (35, 8)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (36, 2)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (38, 8)
INSERT [dbo].[menu_role] ([menu_id], [role_id]) VALUES (39, 8)
GO
SET IDENTITY_INSERT [dbo].[menus] ON 

INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (1, CAST(N'2023-09-25T16:05:48.4570000+00:00' AS DateTimeOffset), CAST(N'2023-09-25T16:05:49.3150000+00:00' AS DateTimeOffset), CAST(N'2023-10-27T13:19:54.7764760+08:00' AS DateTimeOffset), N'Dir', 0, N'/dashboard', N'Dashboard', N'LAYOUT', N'/dashboard/analysis', 1, 1, N'routes.dashboard.dashboard', NULL, NULL, NULL, NULL, N'bx:bx-home', NULL, NULL, NULL, NULL, 1, NULL, NULL, NULL, NULL, NULL, 0, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (2, CAST(N'2023-09-25T16:08:25.4090000+00:00' AS DateTimeOffset), CAST(N'2023-09-25T16:08:26.2150000+00:00' AS DateTimeOffset), CAST(N'2023-10-27T13:19:51.6518738+08:00' AS DateTimeOffset), N'Menu', 1, N'analysis', N'Analysis', N'/dashboard/analysis/index', NULL, 1, 1, N'routes.dashboard.analysis', NULL, NULL, NULL, NULL, N'bx:bx-home', NULL, NULL, 1, NULL, NULL, N'/dashboard', NULL, 1, NULL, NULL, 0, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (3, CAST(N'2023-09-25T16:09:47.2620000+00:00' AS DateTimeOffset), CAST(N'2023-09-25T16:09:48.0820000+00:00' AS DateTimeOffset), CAST(N'2023-10-27T13:19:44.4671117+08:00' AS DateTimeOffset), N'Menu', 1, N'workbench', N'Workbench', N'/dashboard/workbench/index', NULL, 2, 1, N'routes.dashboard.workbench', NULL, NULL, NULL, NULL, N'bx:bx-home', NULL, NULL, 1, NULL, NULL, N'/dashboard', NULL, 1, NULL, NULL, 0, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (4, CAST(N'2023-09-26T14:22:44.9520000+00:00' AS DateTimeOffset), CAST(N'2023-11-07T10:45:34.6499248+08:00' AS DateTimeOffset), NULL, N'Dir', 0, N'/system', N'System', N'LAYOUT', N'/system/menu', 3, 1, N'manager.title', 0, N'', 0, 0, N'ant-design:android-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (5, CAST(N'2023-09-26T14:22:40.8530000+00:00' AS DateTimeOffset), CAST(N'2023-11-10T13:16:07.4249858+08:00' AS DateTimeOffset), NULL, N'Menu', 4, N'menu', N'MenuManagement', N'/manager/menu/index', N'', 4, 1, N'manager.menu.title', 0, N'', 0, 0, N'ant-design:menu-unfold-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (7, CAST(N'2023-09-26T15:50:16.1096812+08:00' AS DateTimeOffset), CAST(N'2023-09-27T10:26:14.1633037+08:00' AS DateTimeOffset), NULL, N'Menu', 4, N'role', N'RoleManagement', N'/manager/role/index', N'', 2, 1, N'manager.role.title', 0, N'', 0, 0, N'ant-design:android-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 0, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (8, CAST(N'2023-09-27T10:27:41.0805844+08:00' AS DateTimeOffset), CAST(N'2023-09-27T10:27:41.0805844+08:00' AS DateTimeOffset), NULL, N'Menu', 4, N'department', N'DepartmentManagement', N'/manager/department/index', N'', 3, 1, N'manager.department.title', 0, N'', 0, 0, N'ant-design:codepen-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 0, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (9, CAST(N'2023-10-02T08:42:57.7473992+08:00' AS DateTimeOffset), CAST(N'2023-11-10T13:15:53.7298851+08:00' AS DateTimeOffset), NULL, N'Menu', 4, N'user', N'UserManagement', N'/manager/user/index', N'', 1, 1, N'manager.user.title', 0, N'', 0, 0, N'ant-design:user-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (10, CAST(N'2023-10-02T09:02:21.6539282+08:00' AS DateTimeOffset), CAST(N'2023-10-02T13:56:28.7200260+08:00' AS DateTimeOffset), NULL, N'Menu', 4, N'user/:id/detail', N'UserDetail', N'/manager/user/detail', N'', 0, 1, N'manager.user.detailTitle', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'/system/user', 0, 1, 0, 0, 0, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (11, CAST(N'2023-10-02T14:25:37.8100914+08:00' AS DateTimeOffset), CAST(N'2023-11-07T10:46:36.4946403+08:00' AS DateTimeOffset), NULL, N'Menu', 4, N'changePassword', N'ChangePassword', N'/sys/password/index', N'', 0, 1, N'sys.password.title', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 1, 0, 0, 1, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (12, CAST(N'2023-10-03T13:36:13.8745616+08:00' AS DateTimeOffset), CAST(N'2023-10-03T13:36:13.8745616+08:00' AS DateTimeOffset), NULL, N'Menu', 4, N'factory', N'FactoryManagement', N'/manager/factory/index', N'', 5, 1, N'manager.factory.title', 0, N'', 0, 0, N'ant-design:shop-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (13, CAST(N'2023-10-04T11:49:25.1907882+08:00' AS DateTimeOffset), CAST(N'2023-10-04T11:49:25.1907882+08:00' AS DateTimeOffset), NULL, N'Menu', 4, N'rank', N'RankManagement', N'/manager/rank/index', N'', 6, 1, N'manager.rank.title', 0, N'', 0, 0, N'ant-design:aliwangwang-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (14, CAST(N'2023-10-04T15:15:14.4822722+08:00' AS DateTimeOffset), CAST(N'2023-10-04T15:15:14.4822722+08:00' AS DateTimeOffset), NULL, N'Menu', 4, N'permission', N'PermissionManagement', N'/manager/permission/index', N'', 7, 1, N'manager.permission.title', 0, N'', 0, 0, N'ant-design:key-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (15, CAST(N'2023-10-11T11:28:51.1096319+08:00' AS DateTimeOffset), CAST(N'2023-10-30T14:10:54.3391556+08:00' AS DateTimeOffset), NULL, N'Dir', 0, N'/setting', N'Setting', N'LAYOUT', N'', 4, 1, N'setting.title', 0, N'', 0, 0, N'ant-design:setting-twotone', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (16, CAST(N'2023-10-11T11:53:50.9686172+08:00' AS DateTimeOffset), CAST(N'2023-10-11T11:53:50.9686172+08:00' AS DateTimeOffset), NULL, N'Menu', 15, N'vacation', N'VacationSetting', N'/setting/vacation/index', N'', 1, 1, N'setting.vacation.title', 0, N'', 0, 0, N'ant-design:calendar-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (17, CAST(N'2023-10-13T09:00:28.6620149+08:00' AS DateTimeOffset), CAST(N'2023-10-27T13:22:24.6248213+08:00' AS DateTimeOffset), NULL, N'Menu', 0, N'/formDesign', N'FormDesign', N'/form-design/index', N'', 10, 0, N'表單設計', 0, N'', 0, 0, N'ant-design:format-painter-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (18, CAST(N'2023-10-13T10:55:59.3744341+08:00' AS DateTimeOffset), CAST(N'2023-10-27T14:18:05.3740207+08:00' AS DateTimeOffset), NULL, N'Menu', 4, N'vacationType', N'VacationTypeManagement', N'/manager/vacation_type/index', N'', 8, 1, N'manager.vacationType.title', 0, N'', 0, 0, N'ant-design:aliwangwang-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (19, CAST(N'2023-10-20T11:46:01.2561761+08:00' AS DateTimeOffset), CAST(N'2023-11-01T09:51:09.4447333+08:00' AS DateTimeOffset), NULL, N'Menu', 15, N'vacation/preview/:year', N'VacationPreview', N'/setting/vacation/preview', N'', 1, 1, N'setting.vacation.preview', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'/setting/vacation', 0, 1, 0, 0, 1, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (20, CAST(N'2023-10-23T09:12:36.1359958+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:49:11.5298805+08:00' AS DateTimeOffset), NULL, N'Menu', 15, N'workShift', N'WorkShiftSetting', N'/setting/work_shift/index', N'', 2, 1, N'setting.workShift.title', 0, N'', 0, 0, N'ant-design:medium-workmark-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (21, CAST(N'2023-10-23T16:37:05.4839167+08:00' AS DateTimeOffset), CAST(N'2023-10-30T14:10:49.5230857+08:00' AS DateTimeOffset), NULL, N'Dir', 0, N'/upload', N'Upload', N'LAYOUT', N'', 5, 1, N'上傳資料', 0, N'', 0, 0, N'ant-design:cloud-upload-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (22, CAST(N'2023-10-23T16:39:13.5053494+08:00' AS DateTimeOffset), CAST(N'2023-10-23T16:39:13.5053494+08:00' AS DateTimeOffset), NULL, N'Menu', 21, N'check-in', N'CheckInUpload', N'/upload/check-in/index', N'', 1, 1, N'打卡資料', 0, N'', 0, 0, N'ant-design:database-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (23, CAST(N'2023-10-27T13:13:03.4991002+08:00' AS DateTimeOffset), CAST(N'2023-11-07T10:50:55.7935871+08:00' AS DateTimeOffset), NULL, N'Menu', 0, N'/home', N'HomePage', N'/sys/home/index', N'', 1, 1, N'sys.home.title', 0, N'', 0, 1, N'ant-design:home-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (24, CAST(N'2023-10-27T14:17:53.7207489+08:00' AS DateTimeOffset), CAST(N'2023-10-27T14:17:53.7207489+08:00' AS DateTimeOffset), NULL, N'Menu', 4, N'leave_core', N'LeaveCoreManagement', N'/manager/leave_core/index', N'', 9, 1, N'manager.leaveCore.title', 0, N'', 0, 0, N'ant-design:book-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (25, CAST(N'2023-10-30T14:13:26.7289793+08:00' AS DateTimeOffset), CAST(N'2023-11-06T13:28:35.8737634+08:00' AS DateTimeOffset), NULL, N'Dir', 0, N'/query', N'Query', N'LAYOUT', N'', 2, 1, N'query.title', 0, N'', 0, 0, N'ant-design:search-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (26, CAST(N'2023-10-30T14:14:36.1895323+08:00' AS DateTimeOffset), CAST(N'2023-11-06T13:28:42.2789782+08:00' AS DateTimeOffset), NULL, N'Menu', 25, N'attend', N'AttendQuery', N'/query/attend/index', N'', 1, 1, N'query.attend.title', 0, N'', 0, 0, N'ant-design:search-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (27, CAST(N'2023-11-01T09:49:08.3608836+08:00' AS DateTimeOffset), CAST(N'2023-11-01T09:49:08.3608836+08:00' AS DateTimeOffset), NULL, N'Menu', 25, N'vacation', N'VacationQuery', N'/query/vacation/index', N'', 2, 1, N'特/補休/預支假/彈休查詢', 0, N'', 0, 0, N'ant-design:search-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (28, CAST(N'2023-11-07T10:48:06.4108045+08:00' AS DateTimeOffset), CAST(N'2023-11-07T15:16:36.5139478+08:00' AS DateTimeOffset), NULL, N'Menu', 30, N'changePassword', N'ChangePassword', N'/sys/password/index', N'', 0, 1, N'修改密碼', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'/user', 0, 1, 0, 0, 1, 2)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (29, CAST(N'2023-10-27T13:13:03.4991000+08:00' AS DateTimeOffset), CAST(N'2023-11-07T10:50:55.7935870+08:00' AS DateTimeOffset), NULL, N'Menu', 0, N'/home', N'HomePage', N'/sys/home/index', N'', 1, 1, N'sys.home.title', 0, N'', 0, 1, N'ant-design:home-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 2)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (30, CAST(N'2023-11-07T15:13:05.3502557+08:00' AS DateTimeOffset), CAST(N'2023-11-07T15:15:40.4082117+08:00' AS DateTimeOffset), NULL, N'Dir', 0, N'/user', N'User', N'LAYOUT', N'/user/detail', 100, 1, N'User', 0, N'', 0, 0, N'ant-design:user-outlined', N'', N'', 0, 0, 1, N'', 0, 0, 0, 0, 1, 2)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (31, CAST(N'2023-11-08T09:17:32.8014668+08:00' AS DateTimeOffset), CAST(N'2023-11-08T09:17:32.8014668+08:00' AS DateTimeOffset), NULL, N'Dir', 0, N'/daliy', N'Daliy', N'LAYOUT', N'', 1, 1, N'日常作業', 0, N'', 0, 0, N'ant-design:calendar-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 2)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (32, CAST(N'2023-11-08T09:23:59.2508675+08:00' AS DateTimeOffset), CAST(N'2023-12-20T09:25:28.1580814+08:00' AS DateTimeOffset), NULL, N'Menu', 31, N'leave', N'LeaveDaily', N'/daily/leave/index', N'', 1, 1, N'請假單/申請', 0, N'', 0, 0, N'ant-design:form-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 2)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (33, CAST(N'2023-11-10T13:32:11.8495397+08:00' AS DateTimeOffset), CAST(N'2023-11-10T13:32:11.8495397+08:00' AS DateTimeOffset), NULL, N'Menu', 15, N'schedule', N'WorkScheduleSetting', N'/setting/work_schedule/index', N'', 3, 1, N'班表設定', 0, N'', 0, 0, N'ant-design:schedule-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (34, CAST(N'2023-11-24T09:51:43.4833697+08:00' AS DateTimeOffset), CAST(N'2023-12-20T10:52:23.8905679+08:00' AS DateTimeOffset), NULL, N'Menu', 36, N'leave', N'LeaveSignOffSetting', N'/setting/sign_off/index', N'', 1, 1, N'請假單', 0, N'', 0, 0, N'ant-design:cluster-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (35, CAST(N'2023-12-19T10:21:36.6907096+08:00' AS DateTimeOffset), CAST(N'2023-12-20T08:54:57.9404739+08:00' AS DateTimeOffset), NULL, N'Menu', 31, N'overtime', N'OvertimeDaily', N'/daily/overtime/index', N'', 2, 1, N'加班單/申請', 0, N'', 0, 0, N'ant-design:form-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 2)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (36, CAST(N'2023-12-20T10:47:30.3786798+08:00' AS DateTimeOffset), CAST(N'2023-12-20T10:47:30.3786798+08:00' AS DateTimeOffset), NULL, N'Dir', 15, N'signOff', N'SignOffSetting', N'', N'', 4, 1, N'簽核設定', 0, N'', 0, 0, N'ant-design:bank-filled', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 1)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (37, CAST(N'2023-12-25T10:04:28.4763096+08:00' AS DateTimeOffset), CAST(N'2023-12-25T10:04:28.4763096+08:00' AS DateTimeOffset), NULL, N'', 0, N'', N'', N'', N'', 0, 1, N'', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 0, 0)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (38, CAST(N'2024-01-02T09:40:10.4370725+08:00' AS DateTimeOffset), CAST(N'2024-01-02T16:45:48.7062326+08:00' AS DateTimeOffset), NULL, N'Menu', 0, N'/formDesign', N'FormDesign', N'/form-design/index', N'', 100, 1, N'表單設計', 0, N'', 0, 0, N'ant-design:format-painter-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 2)
INSERT [dbo].[menus] ([id], [created_at], [updated_at], [deleted_at], [type], [parent_id], [path], [name], [component], [redirect], [sort], [status], [title], [dynamic_level], [real_path], [ignore_keep_alive], [affix], [icon], [frame_src], [transition_name], [hide_breadcrumb], [carry_param], [hide_children_in_menu], [current_active_menu], [hide_tab], [hide_menu], [ignore_route], [hide_path_for_children], [last_mover], [show]) VALUES (39, CAST(N'2024-01-02T13:04:14.5774300+08:00' AS DateTimeOffset), CAST(N'2024-01-02T13:04:14.5774300+08:00' AS DateTimeOffset), NULL, N'Menu', 31, N'signOff', N'SignOffDaily', N'/daily/signOff/index.vue', N'', 3, 1, N'簽核作業', 0, N'', 0, 0, N'ant-design:snippets-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0, 1, 2)
SET IDENTITY_INSERT [dbo].[menus] OFF
GO
SET IDENTITY_INSERT [dbo].[overtime_rate] ON 

INSERT [dbo].[overtime_rate] ([id], [created_at], [updated_at], [deleted_at], [start], [end], [multiply], [vacation_type_id], [last_mover]) VALUES (1, CAST(N'2023-10-30T11:16:52.7610000+00:00' AS DateTimeOffset), CAST(N'2023-10-30T11:14:10.0207549+08:00' AS DateTimeOffset), NULL, 1, 2, 1.3400000333786011, 1, 1)
INSERT [dbo].[overtime_rate] ([id], [created_at], [updated_at], [deleted_at], [start], [end], [multiply], [vacation_type_id], [last_mover]) VALUES (2, CAST(N'2023-10-30T10:59:37.3590447+08:00' AS DateTimeOffset), CAST(N'2023-10-30T10:59:37.3590447+08:00' AS DateTimeOffset), NULL, 3, 8, 1.6699999570846558, 1, 1)
INSERT [dbo].[overtime_rate] ([id], [created_at], [updated_at], [deleted_at], [start], [end], [multiply], [vacation_type_id], [last_mover]) VALUES (17, CAST(N'2023-10-30T13:14:55.4233662+08:00' AS DateTimeOffset), CAST(N'2023-10-30T13:15:07.2742059+08:00' AS DateTimeOffset), NULL, 1, 8, 1, 2, 1)
INSERT [dbo].[overtime_rate] ([id], [created_at], [updated_at], [deleted_at], [start], [end], [multiply], [vacation_type_id], [last_mover]) VALUES (18, CAST(N'2023-10-30T13:15:03.9595069+08:00' AS DateTimeOffset), CAST(N'2023-10-30T13:15:03.9595069+08:00' AS DateTimeOffset), NULL, 9, 12, 2, 2, 1)
INSERT [dbo].[overtime_rate] ([id], [created_at], [updated_at], [deleted_at], [start], [end], [multiply], [vacation_type_id], [last_mover]) VALUES (19, CAST(N'2023-10-30T13:18:37.2165376+08:00' AS DateTimeOffset), CAST(N'2023-10-30T13:18:37.2165376+08:00' AS DateTimeOffset), NULL, 1, 8, 2, 3, 1)
INSERT [dbo].[overtime_rate] ([id], [created_at], [updated_at], [deleted_at], [start], [end], [multiply], [vacation_type_id], [last_mover]) VALUES (20, CAST(N'2023-10-30T13:18:49.8051618+08:00' AS DateTimeOffset), CAST(N'2023-10-30T13:18:49.8051618+08:00' AS DateTimeOffset), NULL, 9, 10, 1.3400000333786011, 3, 1)
INSERT [dbo].[overtime_rate] ([id], [created_at], [updated_at], [deleted_at], [start], [end], [multiply], [vacation_type_id], [last_mover]) VALUES (21, CAST(N'2023-10-30T13:20:14.3722784+08:00' AS DateTimeOffset), CAST(N'2023-10-30T13:20:14.3722784+08:00' AS DateTimeOffset), NULL, 11, 12, 1.6699999570846558, 3, 1)
INSERT [dbo].[overtime_rate] ([id], [created_at], [updated_at], [deleted_at], [start], [end], [multiply], [vacation_type_id], [last_mover]) VALUES (26, CAST(N'2023-10-31T14:09:00.3586768+08:00' AS DateTimeOffset), CAST(N'2023-10-31T14:09:00.3586768+08:00' AS DateTimeOffset), NULL, 9, 12, 2.6700000762939453, 1, 1)
INSERT [dbo].[overtime_rate] ([id], [created_at], [updated_at], [deleted_at], [start], [end], [multiply], [vacation_type_id], [last_mover]) VALUES (27, CAST(N'2023-12-22T15:21:55.4541712+08:00' AS DateTimeOffset), CAST(N'2023-12-22T15:21:55.4541712+08:00' AS DateTimeOffset), NULL, 1, 2, 1.3300000429153442, 6, 1)
INSERT [dbo].[overtime_rate] ([id], [created_at], [updated_at], [deleted_at], [start], [end], [multiply], [vacation_type_id], [last_mover]) VALUES (28, CAST(N'2023-12-22T15:23:06.4921574+08:00' AS DateTimeOffset), CAST(N'2023-12-22T15:23:06.4921574+08:00' AS DateTimeOffset), NULL, 3, 4, 1.6699999570846558, 6, 1)
INSERT [dbo].[overtime_rate] ([id], [created_at], [updated_at], [deleted_at], [start], [end], [multiply], [vacation_type_id], [last_mover]) VALUES (29, CAST(N'2023-12-22T15:25:09.8132436+08:00' AS DateTimeOffset), CAST(N'2023-12-22T15:25:09.8132436+08:00' AS DateTimeOffset), NULL, 5, 12, 2, 6, 1)
SET IDENTITY_INSERT [dbo].[overtime_rate] OFF
GO
SET IDENTITY_INSERT [dbo].[overtime_request_form] ON 

INSERT [dbo].[overtime_request_form] ([id], [created_at], [updated_at], [deleted_at], [order], [start], [end], [remark], [sign_status], [user_id], [department_id]) VALUES (4, CAST(N'2023-12-22T14:40:53.8296428+08:00' AS DateTimeOffset), CAST(N'2023-12-22T14:41:08.5431799+08:00' AS DateTimeOffset), NULL, N'O20231222000004', CAST(N'2023-12-22T17:00:00.0000000+00:00' AS DateTimeOffset), CAST(N'2023-12-22T22:00:00.0000000+00:00' AS DateTimeOffset), N'', 2, 1, 1)
INSERT [dbo].[overtime_request_form] ([id], [created_at], [updated_at], [deleted_at], [order], [start], [end], [remark], [sign_status], [user_id], [department_id]) VALUES (7, CAST(N'2024-01-02T13:30:25.6827664+08:00' AS DateTimeOffset), CAST(N'2024-01-12T16:38:22.8647654+08:00' AS DateTimeOffset), NULL, N'O20240102000007', CAST(N'2024-01-02T08:30:00.0000000+00:00' AS DateTimeOffset), CAST(N'2024-01-02T09:00:00.0000000+00:00' AS DateTimeOffset), N'', 2, 1, 1)
INSERT [dbo].[overtime_request_form] ([id], [created_at], [updated_at], [deleted_at], [order], [start], [end], [remark], [sign_status], [user_id], [department_id]) VALUES (8, CAST(N'2024-01-10T16:28:44.1128957+08:00' AS DateTimeOffset), CAST(N'2024-01-10T16:28:55.2489465+08:00' AS DateTimeOffset), NULL, N'O20240110000008', CAST(N'2024-01-10T16:28:00.0000000+00:00' AS DateTimeOffset), CAST(N'2024-01-10T21:28:00.0000000+00:00' AS DateTimeOffset), N'', 2, 1, 1)
INSERT [dbo].[overtime_request_form] ([id], [created_at], [updated_at], [deleted_at], [order], [start], [end], [remark], [sign_status], [user_id], [department_id]) VALUES (9, CAST(N'2024-01-10T16:29:31.1461170+08:00' AS DateTimeOffset), CAST(N'2024-01-10T16:29:41.3112847+08:00' AS DateTimeOffset), NULL, N'O20240110000009', CAST(N'2024-01-08T16:29:00.0000000+00:00' AS DateTimeOffset), CAST(N'2024-01-08T20:29:00.0000000+00:00' AS DateTimeOffset), N'', 3, 1, 1)
INSERT [dbo].[overtime_request_form] ([id], [created_at], [updated_at], [deleted_at], [order], [start], [end], [remark], [sign_status], [user_id], [department_id]) VALUES (10, CAST(N'2024-01-12T16:21:03.6141053+08:00' AS DateTimeOffset), CAST(N'2024-01-12T16:39:00.8073839+08:00' AS DateTimeOffset), NULL, N'O20240112000010', CAST(N'2024-01-15T16:20:00.0000000+00:00' AS DateTimeOffset), CAST(N'2024-01-15T18:00:00.0000000+00:00' AS DateTimeOffset), N'', 3, 1, 1)
INSERT [dbo].[overtime_request_form] ([id], [created_at], [updated_at], [deleted_at], [order], [start], [end], [remark], [sign_status], [user_id], [department_id]) VALUES (15, CAST(N'2024-02-19T11:24:10.0425348+08:00' AS DateTimeOffset), CAST(N'2024-02-19T11:24:26.7397595+08:00' AS DateTimeOffset), NULL, N'O20240219000015', CAST(N'2024-02-19T11:24:00.0000000+00:00' AS DateTimeOffset), CAST(N'2024-02-19T16:00:00.0000000+00:00' AS DateTimeOffset), N'', 3, 1, 1)
SET IDENTITY_INSERT [dbo].[overtime_request_form] OFF
GO
SET IDENTITY_INSERT [dbo].[overtime_sign_off_flow] ON 

INSERT [dbo].[overtime_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [overtime_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (7, CAST(N'2023-12-22T14:40:53.8686861+08:00' AS DateTimeOffset), CAST(N'2023-12-22T14:41:08.4441536+08:00' AS DateTimeOffset), NULL, 4, 3, 1, 1, 1, N'', N'test', 1, CAST(N'2023-12-22T14:41:08.4422180+08:00' AS DateTimeOffset), N'5ce4c012-7fd0-47e0-bf39-4a105053c5d9', NULL)
INSERT [dbo].[overtime_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [overtime_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (8, CAST(N'2023-12-22T14:40:53.8686861+08:00' AS DateTimeOffset), CAST(N'2023-12-22T14:41:08.5337187+08:00' AS DateTimeOffset), NULL, 4, 1, 2, 3, 2, N'', N'', 3, CAST(N'2023-12-22T14:41:08.5333618+08:00' AS DateTimeOffset), N'5f2314d6-f4b6-41d9-b64e-627c8896d695', NULL)
INSERT [dbo].[overtime_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [overtime_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (13, CAST(N'2024-01-02T13:30:25.7257148+08:00' AS DateTimeOffset), CAST(N'2024-01-12T16:38:22.7635010+08:00' AS DateTimeOffset), NULL, 7, 3, 1, 1, 1, N'', N'', 1, CAST(N'2024-01-12T16:38:22.7635010+08:00' AS DateTimeOffset), N'ce22fb94-88ed-4ae2-919b-3d567e8003b0', N'zh_TW')
INSERT [dbo].[overtime_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [overtime_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (14, CAST(N'2024-01-02T13:30:25.7257148+08:00' AS DateTimeOffset), CAST(N'2024-01-12T16:38:22.8555348+08:00' AS DateTimeOffset), NULL, 7, 1, 2, 3, 2, N'', N'', 3, CAST(N'2024-01-12T16:38:22.8555348+08:00' AS DateTimeOffset), N'521273d2-155d-41bd-8729-9f420735a838', N'zh_TW')
INSERT [dbo].[overtime_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [overtime_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (15, CAST(N'2024-01-10T16:28:44.1477702+08:00' AS DateTimeOffset), CAST(N'2024-01-10T16:28:55.2029719+08:00' AS DateTimeOffset), NULL, 8, 3, 1, 1, 1, N'', N'test', 1, CAST(N'2024-01-10T16:28:55.2029719+08:00' AS DateTimeOffset), N'0734dcc8-efd2-4f85-8caf-35681783cd6a', N'zh_TW')
INSERT [dbo].[overtime_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [overtime_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (16, CAST(N'2024-01-10T16:28:44.1477702+08:00' AS DateTimeOffset), CAST(N'2024-01-10T16:28:55.2384951+08:00' AS DateTimeOffset), NULL, 8, 1, 2, 3, 2, N'', N'', 3, CAST(N'2024-01-10T16:28:55.2384951+08:00' AS DateTimeOffset), N'621e2a91-d2cd-49ac-b79a-22ab6efe36c7', N'zh_TW')
INSERT [dbo].[overtime_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [overtime_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (17, CAST(N'2024-01-10T16:29:31.1706398+08:00' AS DateTimeOffset), CAST(N'2024-01-10T16:29:41.3194001+08:00' AS DateTimeOffset), NULL, 9, 3, 1, 1, 1, N'', N't', 2, CAST(N'2024-01-10T16:29:41.3194001+08:00' AS DateTimeOffset), N'a4bf7185-80ba-4a27-917a-9a7a5868d45e', N'zh_TW')
INSERT [dbo].[overtime_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [overtime_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (18, CAST(N'2024-01-10T16:29:31.1706398+08:00' AS DateTimeOffset), CAST(N'2024-01-10T16:29:31.1706398+08:00' AS DateTimeOffset), NULL, 9, 1, 2, 3, 2, N'', N'', 0, CAST(N'0001-01-01T00:00:00.0000000+00:00' AS DateTimeOffset), N'47d9448a-e00d-4ad0-bd23-786190a3c0c5', N'zh_TW')
INSERT [dbo].[overtime_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [overtime_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (19, CAST(N'2024-01-12T16:21:03.6565917+08:00' AS DateTimeOffset), CAST(N'2024-01-12T16:39:00.8217937+08:00' AS DateTimeOffset), NULL, 10, 3, 1, 1, 1, N'', N'test', 2, CAST(N'2024-01-12T16:39:00.8217937+08:00' AS DateTimeOffset), N'b9edf9cb-2300-4bec-a66e-dc722364777b', N'zh_TW')
INSERT [dbo].[overtime_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [overtime_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (20, CAST(N'2024-01-12T16:21:03.6565917+08:00' AS DateTimeOffset), CAST(N'2024-01-12T16:21:03.6565917+08:00' AS DateTimeOffset), NULL, 10, 1, 2, 3, 2, N'', N'', 0, CAST(N'0001-01-01T00:00:00.0000000+00:00' AS DateTimeOffset), N'abf1f491-31ac-497a-8461-3487c034eb88', N'zh_TW')
INSERT [dbo].[overtime_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [overtime_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (29, CAST(N'2024-02-19T11:24:10.0919572+08:00' AS DateTimeOffset), CAST(N'2024-02-19T11:24:26.7592916+08:00' AS DateTimeOffset), NULL, 15, 3, 1, 1, 1, N'', N'', 2, CAST(N'2024-02-19T11:24:26.7592916+08:00' AS DateTimeOffset), N'e37d9ef1-02f4-4829-9c1d-507c1230aeb8', N'zh_TW')
INSERT [dbo].[overtime_sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [overtime_request_form_id], [sign_off_user_id], [level], [sign_type], [notify], [remark], [comment], [status], [sign_date], [uuid], [locale]) VALUES (30, CAST(N'2024-02-19T11:24:10.0919572+08:00' AS DateTimeOffset), CAST(N'2024-02-19T11:24:10.0919572+08:00' AS DateTimeOffset), NULL, 15, 1, 2, 3, 2, N'', N'', 0, CAST(N'0001-01-01T00:00:00.0000000+00:00' AS DateTimeOffset), N'2265fcfc-9c09-4c01-b582-095ab911e6ff', N'zh_TW')
SET IDENTITY_INSERT [dbo].[overtime_sign_off_flow] OFF
GO
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (1, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (1, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (7, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (7, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (8, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (8, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (9, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (9, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (10, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (10, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (11, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (11, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (12, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (12, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (13, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (13, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (14, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (14, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (15, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (15, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (16, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (16, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (17, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (17, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (18, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (18, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (19, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (19, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (20, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (20, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (21, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (21, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (22, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (22, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (23, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (23, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (24, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (24, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (25, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (25, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (26, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (27, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (27, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (28, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (28, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (29, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (29, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (30, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (30, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (31, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (31, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (32, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (32, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (33, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (33, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (34, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (34, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (35, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (35, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (36, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (36, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (37, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (37, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (38, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (38, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (39, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (39, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (40, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (40, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (41, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (41, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (42, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (42, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (43, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (43, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (47, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (47, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (48, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (48, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (49, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (49, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (50, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (50, 9)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (51, 2)
INSERT [dbo].[permission_role] ([permission_id], [role_id]) VALUES (51, 9)
GO
SET IDENTITY_INSERT [dbo].[permissions] ON 

INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (1, CAST(N'2023-10-04T15:22:18.8950417+08:00' AS DateTimeOffset), CAST(N'2023-10-04T15:22:37.4082782+08:00' AS DateTimeOffset), NULL, N'user:status:edit', N'使用者狀態按鈕', 1, N'設定使用者狀態', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (7, CAST(N'2023-10-04T16:56:18.4197436+08:00' AS DateTimeOffset), CAST(N'2023-10-06T09:11:16.5707226+08:00' AS DateTimeOffset), NULL, N'user:password:reset', N'重製使用者密碼按鈕', 1, N'重製密碼', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (8, CAST(N'2023-10-06T08:20:54.7074305+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:20:54.7074305+08:00' AS DateTimeOffset), NULL, N'perm:add', N'新增權限按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (9, CAST(N'2023-10-06T08:21:13.9089799+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:21:13.9089799+08:00' AS DateTimeOffset), NULL, N'perm:edit', N'編輯權限按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (10, CAST(N'2023-10-06T08:21:31.2122433+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:21:31.2122433+08:00' AS DateTimeOffset), NULL, N'perm:delete', N'刪除權限按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (11, CAST(N'2023-10-06T08:22:55.1116353+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:22:55.1116353+08:00' AS DateTimeOffset), NULL, N'rank:add', N'新增職級按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (12, CAST(N'2023-10-06T08:23:06.4132463+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:23:06.4132463+08:00' AS DateTimeOffset), NULL, N'rank:edit', N'編輯職級按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (13, CAST(N'2023-10-06T08:23:22.0745177+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:23:22.0745177+08:00' AS DateTimeOffset), NULL, N'rank:delete', N'刪除職級按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (14, CAST(N'2023-10-06T08:23:39.4900154+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:23:39.4900154+08:00' AS DateTimeOffset), NULL, N'grade:add', N'新增職等按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (15, CAST(N'2023-10-06T08:23:51.7666681+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:23:51.7666681+08:00' AS DateTimeOffset), NULL, N'grade:edit', N'編輯職等按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (16, CAST(N'2023-10-06T08:24:03.9359466+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:24:03.9359466+08:00' AS DateTimeOffset), NULL, N'grade:delete', N'刪除職等按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (17, CAST(N'2023-10-06T08:24:40.9311987+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:25:56.8599112+08:00' AS DateTimeOffset), NULL, N'factory:add', N'新增工廠按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (18, CAST(N'2023-10-06T08:24:56.7445965+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:25:49.9233176+08:00' AS DateTimeOffset), NULL, N'factory:edit', N'編輯工廠按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (19, CAST(N'2023-10-06T08:25:11.5643241+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:25:44.9755634+08:00' AS DateTimeOffset), NULL, N'factory:delete', N'刪除工廠按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (20, CAST(N'2023-10-06T08:25:38.8668631+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:26:46.9554431+08:00' AS DateTimeOffset), NULL, N'factoryType:add', N'新增廠別按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (21, CAST(N'2023-10-06T08:26:14.3053666+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:26:40.1171879+08:00' AS DateTimeOffset), NULL, N'factoryType:edit', N'編輯廠別按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (22, CAST(N'2023-10-06T08:26:31.2291787+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:26:31.2291787+08:00' AS DateTimeOffset), NULL, N'factoryType:delete', N'刪除廠別按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (23, CAST(N'2023-10-06T08:27:18.1577965+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:27:18.1577965+08:00' AS DateTimeOffset), NULL, N'user:add', N'新增使用者按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (24, CAST(N'2023-10-06T08:27:35.6907916+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:27:35.6907916+08:00' AS DateTimeOffset), NULL, N'user:edit', N'編輯使用者按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (25, CAST(N'2023-10-06T08:27:49.1384170+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:27:49.1384170+08:00' AS DateTimeOffset), NULL, N'user:delete', N'刪除使用者按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (26, CAST(N'2023-10-06T08:28:10.1058664+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:28:10.1058664+08:00' AS DateTimeOffset), NULL, N'dept:add', N'新增部門按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (27, CAST(N'2023-10-06T08:28:45.0998314+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:28:45.0998314+08:00' AS DateTimeOffset), NULL, N'dept:edit', N'編輯部門按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (28, CAST(N'2023-10-06T08:28:59.1625191+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:28:59.1625191+08:00' AS DateTimeOffset), NULL, N'dept:delete', N'刪除部門按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (29, CAST(N'2023-10-06T08:29:16.2870024+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:29:16.2870024+08:00' AS DateTimeOffset), NULL, N'role:add', N'新增角色按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (30, CAST(N'2023-10-06T08:29:35.6196459+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:29:35.6196459+08:00' AS DateTimeOffset), NULL, N'role:edit', N'編輯角色按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (31, CAST(N'2023-10-06T08:29:47.1810584+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:29:47.1810584+08:00' AS DateTimeOffset), NULL, N'role:delete', N'刪除角色按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (32, CAST(N'2023-10-06T08:30:12.8536542+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:30:12.8536542+08:00' AS DateTimeOffset), NULL, N'menu:add', N'新增菜單按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (33, CAST(N'2023-10-06T08:30:23.9656039+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:30:23.9656039+08:00' AS DateTimeOffset), NULL, N'menu:edit', N'編輯菜單按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (34, CAST(N'2023-10-06T08:30:38.1058660+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:30:38.1058660+08:00' AS DateTimeOffset), NULL, N'menu:delete', N'刪除菜單按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (35, CAST(N'2023-10-06T08:54:12.6114212+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:54:12.6114212+08:00' AS DateTimeOffset), NULL, N'factory:status:edit', N'設定工廠狀態按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (36, CAST(N'2023-10-06T08:54:37.1818991+08:00' AS DateTimeOffset), CAST(N'2023-10-06T09:10:32.5108301+08:00' AS DateTimeOffset), NULL, N'factoryType:status:edit', N'設定廠別狀態按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (37, CAST(N'2023-10-06T08:56:09.8548396+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:56:09.8548396+08:00' AS DateTimeOffset), NULL, N'grade:status:edit', N'設定職等狀態按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (38, CAST(N'2023-10-06T08:56:26.9486176+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:56:26.9486176+08:00' AS DateTimeOffset), NULL, N'rank:status:edit', N'設定職級狀態按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (39, CAST(N'2023-10-06T08:56:53.6158910+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:56:53.6158910+08:00' AS DateTimeOffset), NULL, N'perm:status:edit', N'設定權限狀態按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (40, CAST(N'2023-10-06T08:57:18.6132810+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:57:18.6132810+08:00' AS DateTimeOffset), NULL, N'role:status:edit', N'設定角色狀態按鈕', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (41, CAST(N'2023-10-23T09:17:10.5087093+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:50:03.2155578+08:00' AS DateTimeOffset), NULL, N'workShift:add', N'新增班別', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (42, CAST(N'2023-10-23T09:17:23.0704086+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:50:08.8609664+08:00' AS DateTimeOffset), NULL, N'workShift:edit', N'編輯班別', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (43, CAST(N'2023-10-23T09:17:37.4902791+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:50:16.1450780+08:00' AS DateTimeOffset), NULL, N'workShift:delete', N'刪除班別', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (44, CAST(N'2023-10-23T10:29:04.0790335+08:00' AS DateTimeOffset), CAST(N'2023-10-23T10:29:04.0790335+08:00' AS DateTimeOffset), CAST(N'2023-12-20T10:39:22.8585155+08:00' AS DateTimeOffset), N'vacation:add', N'新增休假日', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (45, CAST(N'2023-10-23T10:29:17.7443941+08:00' AS DateTimeOffset), CAST(N'2023-10-23T10:29:17.7443941+08:00' AS DateTimeOffset), CAST(N'2023-12-20T10:39:24.1474618+08:00' AS DateTimeOffset), N'vacation:edit', N'編輯休假日', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (46, CAST(N'2023-10-23T10:29:49.2914173+08:00' AS DateTimeOffset), CAST(N'2023-10-23T10:29:49.2914173+08:00' AS DateTimeOffset), CAST(N'2023-12-20T10:39:25.5209975+08:00' AS DateTimeOffset), N'vacation:delete', N'刪除休假日', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (47, CAST(N'2023-10-24T14:38:04.3725420+08:00' AS DateTimeOffset), CAST(N'2023-10-24T15:38:49.5704456+08:00' AS DateTimeOffset), NULL, N'check-in:delete', N'刪除打卡資料', 0, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (48, CAST(N'2023-10-27T14:23:11.2823920+08:00' AS DateTimeOffset), CAST(N'2023-10-27T14:23:11.2823920+08:00' AS DateTimeOffset), NULL, N'leaveCore:add', N'新增假別', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (49, CAST(N'2023-10-27T14:23:23.2108728+08:00' AS DateTimeOffset), CAST(N'2023-10-27T14:23:23.2108728+08:00' AS DateTimeOffset), NULL, N'leaveCore:edit', N'編輯假別', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (50, CAST(N'2023-10-27T14:23:33.8272499+08:00' AS DateTimeOffset), CAST(N'2023-10-27T14:23:33.8272499+08:00' AS DateTimeOffset), NULL, N'leaveCore:delete', N'刪除假別', 1, N'', 1)
INSERT [dbo].[permissions] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (51, CAST(N'2023-10-27T16:16:06.9693407+08:00' AS DateTimeOffset), CAST(N'2023-10-27T16:16:06.9693407+08:00' AS DateTimeOffset), NULL, N'factoryType:overtime', N'編輯加班倍率', 1, N'', 1)
SET IDENTITY_INSERT [dbo].[permissions] OFF
GO
SET IDENTITY_INSERT [dbo].[ranks] ON 

INSERT [dbo].[ranks] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (1, CAST(N'2023-10-02T09:25:38.1750000+00:00' AS DateTimeOffset), CAST(N'2023-10-04T11:51:19.4375487+08:00' AS DateTimeOffset), NULL, N'r_test_1', N'職級1', 1, N'測試用123', 1)
INSERT [dbo].[ranks] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (2, CAST(N'2023-10-04T11:51:40.6400623+08:00' AS DateTimeOffset), CAST(N'2023-10-06T09:11:06.8494966+08:00' AS DateTimeOffset), NULL, N'r_test_2', N'職級2', 1, N'tafafsdf', 1)
INSERT [dbo].[ranks] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (3, CAST(N'2023-10-04T15:19:30.1348881+08:00' AS DateTimeOffset), CAST(N'2023-10-04T15:19:30.1348881+08:00' AS DateTimeOffset), NULL, N'r_test_3', N'職級3', 1, N'', 1)
SET IDENTITY_INSERT [dbo].[ranks] OFF
GO
SET IDENTITY_INSERT [dbo].[roles] ON 

INSERT [dbo].[roles] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (1, CAST(N'2023-09-25T16:03:59.5940000+00:00' AS DateTimeOffset), CAST(N'2023-11-01T11:25:02.1481816+08:00' AS DateTimeOffset), NULL, N'user', N'一般使用者', 1, N'', 1)
INSERT [dbo].[roles] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (2, CAST(N'2023-09-25T16:04:21.2920000+00:00' AS DateTimeOffset), CAST(N'2023-12-20T10:48:18.1219994+08:00' AS DateTimeOffset), NULL, N'manager', N'管理員', 1, N'', 1)
INSERT [dbo].[roles] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (4, CAST(N'2023-09-26T16:38:17.1060281+08:00' AS DateTimeOffset), CAST(N'2023-09-26T16:38:17.1060281+08:00' AS DateTimeOffset), CAST(N'2023-09-26T16:40:57.2024071+08:00' AS DateTimeOffset), N't', N'T', 1, N'', 0)
INSERT [dbo].[roles] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (5, CAST(N'2023-09-26T16:39:50.5320765+08:00' AS DateTimeOffset), CAST(N'2023-09-26T16:39:50.5320765+08:00' AS DateTimeOffset), CAST(N'2023-09-26T16:40:55.5464385+08:00' AS DateTimeOffset), N'321', N'123', 1, N'', 0)
INSERT [dbo].[roles] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (6, CAST(N'2023-09-26T16:40:50.8096780+08:00' AS DateTimeOffset), CAST(N'2023-09-26T16:40:50.8096780+08:00' AS DateTimeOffset), CAST(N'2023-09-26T16:40:53.3229932+08:00' AS DateTimeOffset), N'786', N'768', 1, N'', 0)
INSERT [dbo].[roles] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (7, CAST(N'2023-09-26T16:43:10.6554241+08:00' AS DateTimeOffset), CAST(N'2023-09-26T16:43:10.6588630+08:00' AS DateTimeOffset), CAST(N'2023-09-26T16:43:13.6029829+08:00' AS DateTimeOffset), N'dfg', N'fgs', 0, N'', 0)
INSERT [dbo].[roles] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (8, CAST(N'2023-11-07T10:48:59.6623527+08:00' AS DateTimeOffset), CAST(N'2024-01-02T13:04:52.0375958+08:00' AS DateTimeOffset), NULL, N'root', N'超級管理員(web)', 1, N'', 1)
INSERT [dbo].[roles] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover]) VALUES (9, CAST(N'2023-12-25T10:04:27.8625733+08:00' AS DateTimeOffset), CAST(N'2023-12-25T10:04:51.3174517+08:00' AS DateTimeOffset), CAST(N'2023-12-25T10:05:10.9745801+08:00' AS DateTimeOffset), N'test_1', N'test', 1, N'測試用', 1)
SET IDENTITY_INSERT [dbo].[roles] OFF
GO
SET IDENTITY_INSERT [dbo].[sign_off_flow] ON 

INSERT [dbo].[sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [department_id], [leave_core_id], [level], [exceed_day], [user_id], [sign_type], [notify], [remark], [rule_type]) VALUES (5, CAST(N'2023-11-27T15:30:11.1087932+08:00' AS DateTimeOffset), CAST(N'2023-11-29T08:22:42.4808748+08:00' AS DateTimeOffset), NULL, 1, 2, 1, 0, 0, 1, 1, N'test', 1)
INSERT [dbo].[sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [department_id], [leave_core_id], [level], [exceed_day], [user_id], [sign_type], [notify], [remark], [rule_type]) VALUES (6, CAST(N'2023-11-27T15:30:11.1087932+08:00' AS DateTimeOffset), CAST(N'2023-11-27T16:44:58.3588455+08:00' AS DateTimeOffset), NULL, 1, 2, 2, 2, 0, 1, 1, N'', 1)
INSERT [dbo].[sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [department_id], [leave_core_id], [level], [exceed_day], [user_id], [sign_type], [notify], [remark], [rule_type]) VALUES (7, CAST(N'2023-11-27T16:44:13.1879905+08:00' AS DateTimeOffset), CAST(N'2023-12-11T16:04:38.4709278+08:00' AS DateTimeOffset), NULL, 1, 2, 3, 3, 3, 2, 1, N'', 1)
INSERT [dbo].[sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [department_id], [leave_core_id], [level], [exceed_day], [user_id], [sign_type], [notify], [remark], [rule_type]) VALUES (9, CAST(N'2023-11-27T16:51:07.7807768+08:00' AS DateTimeOffset), CAST(N'2023-11-27T16:51:07.7807768+08:00' AS DateTimeOffset), NULL, 1, 4, 2, 2, 0, 1, 1, N'', 1)
INSERT [dbo].[sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [department_id], [leave_core_id], [level], [exceed_day], [user_id], [sign_type], [notify], [remark], [rule_type]) VALUES (10, CAST(N'2023-11-27T16:51:07.7807768+08:00' AS DateTimeOffset), CAST(N'2023-11-27T16:51:07.7807768+08:00' AS DateTimeOffset), NULL, 1, 4, 1, 0, 0, 1, 1, N'', 1)
INSERT [dbo].[sign_off_flow] ([id], [created_at], [updated_at], [deleted_at], [department_id], [leave_core_id], [level], [exceed_day], [user_id], [sign_type], [notify], [remark], [rule_type]) VALUES (11, CAST(N'2023-11-29T10:05:01.0225500+08:00' AS DateTimeOffset), CAST(N'2023-11-29T10:05:01.0225500+08:00' AS DateTimeOffset), CAST(N'2023-11-29T10:05:18.5840558+08:00' AS DateTimeOffset), 1, 2, 4, 0, 0, 2, 0, N'', 1)
SET IDENTITY_INSERT [dbo].[sign_off_flow] OFF
GO
SET IDENTITY_INSERT [dbo].[system_logs] ON 

INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (1, CAST(N'2023-10-03T11:31:21.9790909+08:00' AS DateTimeOffset), CAST(N'2023-10-03T11:31:21.9790909+08:00' AS DateTimeOffset), NULL, 1, N'user', N'U', N'user:1 use function "SetStatus()", set user status to false id:3')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (2, CAST(N'2023-10-03T11:31:22.9340865+08:00' AS DateTimeOffset), CAST(N'2023-10-03T11:31:22.9340865+08:00' AS DateTimeOffset), NULL, 1, N'user', N'U', N'user:1 use function "SetStatus()", set user status to true id:3')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (3, CAST(N'2023-10-03T11:35:03.2812065+08:00' AS DateTimeOffset), CAST(N'2023-10-03T11:35:03.2812065+08:00' AS DateTimeOffset), NULL, 1, N'user', N'U', N'user:1 use function "SetStatus()", set user status to false id:3')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (4, CAST(N'2023-10-03T11:35:03.9864624+08:00' AS DateTimeOffset), CAST(N'2023-10-03T11:35:03.9864624+08:00' AS DateTimeOffset), NULL, 1, N'user', N'U', N'user:1 use function "SetStatus()", set user status to true id:3')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (5, CAST(N'2023-10-03T12:01:51.9160252+08:00' AS DateTimeOffset), CAST(N'2023-10-03T12:01:51.9160252+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "SetRoleStatus()", set roles status to false id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (6, CAST(N'2023-10-03T12:01:52.8717019+08:00' AS DateTimeOffset), CAST(N'2023-10-03T12:01:52.8717019+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "SetRoleStatus()", set roles status to true id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (7, CAST(N'2023-10-03T13:11:18.3804024+08:00' AS DateTimeOffset), CAST(N'2023-10-03T13:11:18.3804024+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "SetRoleStatus()", set roles status to false id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (8, CAST(N'2023-10-03T13:11:19.1461456+08:00' AS DateTimeOffset), CAST(N'2023-10-03T13:11:19.1461456+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "SetRoleStatus()", set roles status to true id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (9, CAST(N'2023-10-03T13:36:13.8835170+08:00' AS DateTimeOffset), CAST(N'2023-10-03T13:36:13.8835170+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:12')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (10, CAST(N'2023-10-03T13:36:26.5572493+08:00' AS DateTimeOffset), CAST(N'2023-10-03T13:36:26.5572493+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (11, CAST(N'2023-10-03T13:41:00.1725261+08:00' AS DateTimeOffset), CAST(N'2023-10-03T13:41:00.1725261+08:00' AS DateTimeOffset), NULL, 1, N'factory', N'U', N'user:1 use function "SetFactoryStatus()", set factory status to false id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (12, CAST(N'2023-10-03T13:41:26.1736305+08:00' AS DateTimeOffset), CAST(N'2023-10-03T13:41:26.1736305+08:00' AS DateTimeOffset), NULL, 1, N'factory', N'U', N'user:1 use function "SetFactoryStatus()", set factory status to true id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (13, CAST(N'2023-10-03T15:01:57.5632185+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:01:57.5632185+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "DeleteFactoryType()", set factory type status to false id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (14, CAST(N'2023-10-03T15:01:58.5313536+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:01:58.5313536+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "DeleteFactoryType()", set factory type status to true id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (15, CAST(N'2023-10-03T15:02:18.2699800+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:02:18.2699800+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "DeleteFactoryType()", set factory type status to false id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (16, CAST(N'2023-10-03T15:02:25.0004239+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:02:25.0004239+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "DeleteFactoryType()", set factory type status to true id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (17, CAST(N'2023-10-03T15:26:43.5469803+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:26:43.5469803+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "EditFactoryType()", edit factory type id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (18, CAST(N'2023-10-03T15:26:49.0139869+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:26:49.0139869+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "EditFactoryType()", edit factory type id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (19, CAST(N'2023-10-03T15:26:53.0571145+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:26:53.0571145+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "EditFactoryType()", edit factory type id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (20, CAST(N'2023-10-03T15:28:41.7019583+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:28:41.7019583+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "EditFactoryType()", edit factory type id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (21, CAST(N'2023-10-03T15:28:45.6375609+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:28:45.6375609+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "EditFactoryType()", edit factory type id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (22, CAST(N'2023-10-03T15:28:49.3281711+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:28:49.3281711+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "EditFactoryType()", edit factory type id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (23, CAST(N'2023-10-03T15:28:52.6717698+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:28:52.6717698+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "EditFactoryType()", edit factory type id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (24, CAST(N'2023-10-03T15:29:00.4071086+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:29:00.4071086+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "EditFactoryType()", edit factory type id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (25, CAST(N'2023-10-03T15:29:03.6774294+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:29:03.6774294+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "EditFactoryType()", edit factory type id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (26, CAST(N'2023-10-03T15:29:29.5525278+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:29:29.5525278+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "EditFactoryType()", edit factory type id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (27, CAST(N'2023-10-03T15:29:33.3590120+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:29:33.3590120+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "EditFactoryType()", edit factory type id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (28, CAST(N'2023-10-03T15:29:42.6458575+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:29:42.6458575+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "EditFactoryType()", edit factory type id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (29, CAST(N'2023-10-03T15:29:45.6069889+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:29:45.6069889+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "EditFactoryType()", edit factory type id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (30, CAST(N'2023-10-03T15:29:52.5114141+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:29:52.5114141+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "DeleteFactoryType()", set factory type status to false id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (31, CAST(N'2023-10-03T15:29:52.9901044+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:29:52.9901044+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "DeleteFactoryType()", set factory type status to true id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (32, CAST(N'2023-10-03T15:32:47.7930906+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:32:47.7930906+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'C', N'user:1 use function "AddFactoryType()", create factory type id:4')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (33, CAST(N'2023-10-03T15:32:49.7410366+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:32:49.7410366+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "DeleteFactoryType()", set factory type status to true id:4')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (34, CAST(N'2023-10-03T15:46:29.2918396+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:46:29.2918396+08:00' AS DateTimeOffset), NULL, 1, N'factory', N'U', N'user:1 use function "EditFactory()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (35, CAST(N'2023-10-03T15:46:33.3151558+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:46:33.3151558+08:00' AS DateTimeOffset), NULL, 1, N'factory', N'U', N'user:1 use function "EditFactory()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (36, CAST(N'2023-10-03T15:46:37.9681886+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:46:37.9681886+08:00' AS DateTimeOffset), NULL, 1, N'factory', N'U', N'user:1 use function "SetFactoryStatus()", set factory status to false id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (37, CAST(N'2023-10-03T15:46:38.5115764+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:46:38.5115764+08:00' AS DateTimeOffset), NULL, 1, N'factory', N'U', N'user:1 use function "SetFactoryStatus()", set factory status to true id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (38, CAST(N'2023-10-03T15:46:57.4515407+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:46:57.4515407+08:00' AS DateTimeOffset), NULL, 1, N'factory', N'C', N'user:1 use function "AddFactory()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (39, CAST(N'2023-10-03T15:48:55.4627117+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:48:55.4627117+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'C', N'user:1 use function "AddFactoryType()", create factory type id:6')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (40, CAST(N'2023-10-03T15:49:11.2788258+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:49:11.2788258+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'C', N'user:1 use function "AddFactoryType()", create factory type id:7')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (41, CAST(N'2023-10-03T15:49:12.7269653+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:49:12.7269653+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "DeleteFactoryType()", set factory type status to true id:7')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (42, CAST(N'2023-10-03T15:49:13.3557441+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:49:13.3557441+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "DeleteFactoryType()", set factory type status to false id:7')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (43, CAST(N'2023-10-03T15:49:18.0839527+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:49:18.0839527+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "EditFactoryType()", edit factory type id:7')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (44, CAST(N'2023-10-03T15:49:21.1889281+08:00' AS DateTimeOffset), CAST(N'2023-10-03T15:49:21.1889281+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "EditFactoryType()", edit factory type id:7')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (45, CAST(N'2023-10-04T10:07:11.7433677+08:00' AS DateTimeOffset), CAST(N'2023-10-04T10:07:11.7433677+08:00' AS DateTimeOffset), NULL, 1, N'factory', N'U', N'user:1 use function "SetFactoryStatus()", set factory status to true id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (46, CAST(N'2023-10-04T10:59:15.1362164+08:00' AS DateTimeOffset), CAST(N'2023-10-04T10:59:15.1362164+08:00' AS DateTimeOffset), NULL, 1, N'factory', N'D', N'user:1 use function "DeleteFactory()", delete factory id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (47, CAST(N'2023-10-04T10:59:21.3825867+08:00' AS DateTimeOffset), CAST(N'2023-10-04T10:59:21.3825867+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'D', N'user:1 use function "DeleteFactoryType()", delete factory type id:4')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (48, CAST(N'2023-10-04T10:59:39.9532452+08:00' AS DateTimeOffset), CAST(N'2023-10-04T10:59:39.9532452+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "DeleteFactoryType()", set factory type status to true id:7')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (49, CAST(N'2023-10-04T11:49:25.1985468+08:00' AS DateTimeOffset), CAST(N'2023-10-04T11:49:25.1985468+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:13')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (50, CAST(N'2023-10-04T11:50:22.9543625+08:00' AS DateTimeOffset), CAST(N'2023-10-04T11:50:22.9543625+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (51, CAST(N'2023-10-04T11:51:17.0470976+08:00' AS DateTimeOffset), CAST(N'2023-10-04T11:51:17.0470976+08:00' AS DateTimeOffset), NULL, 1, N'rank', N'U', N'user:1 use function "EditRank()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (52, CAST(N'2023-10-04T11:51:18.7833334+08:00' AS DateTimeOffset), CAST(N'2023-10-04T11:51:18.7833334+08:00' AS DateTimeOffset), NULL, 1, N'rank', N'U', N'user:1 use function "SetRankStatus()", set rank status to false id 1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (53, CAST(N'2023-10-04T11:51:19.4390883+08:00' AS DateTimeOffset), CAST(N'2023-10-04T11:51:19.4390883+08:00' AS DateTimeOffset), NULL, 1, N'rank', N'U', N'user:1 use function "SetRankStatus()", set rank status to true id 1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (54, CAST(N'2023-10-04T11:51:40.6471573+08:00' AS DateTimeOffset), CAST(N'2023-10-04T11:51:40.6471573+08:00' AS DateTimeOffset), NULL, 1, N'rank', N'C', N'user:1 use function "AddRank()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (55, CAST(N'2023-10-04T11:51:42.8883698+08:00' AS DateTimeOffset), CAST(N'2023-10-04T11:51:42.8883698+08:00' AS DateTimeOffset), NULL, 1, N'rank', N'U', N'user:1 use function "SetRankStatus()", set rank status to true id 2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (56, CAST(N'2023-10-04T11:51:43.4473663+08:00' AS DateTimeOffset), CAST(N'2023-10-04T11:51:43.4473663+08:00' AS DateTimeOffset), NULL, 1, N'rank', N'U', N'user:1 use function "SetRankStatus()", set rank status to false id 2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (57, CAST(N'2023-10-04T11:51:50.9733512+08:00' AS DateTimeOffset), CAST(N'2023-10-04T11:51:50.9733512+08:00' AS DateTimeOffset), NULL, 1, N'rank', N'U', N'user:1 use function "EditRank()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (58, CAST(N'2023-10-04T11:51:52.1102019+08:00' AS DateTimeOffset), CAST(N'2023-10-04T11:51:52.1102019+08:00' AS DateTimeOffset), NULL, 1, N'rank', N'U', N'user:1 use function "SetRankStatus()", set rank status to false id 2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (59, CAST(N'2023-10-04T11:52:17.6865624+08:00' AS DateTimeOffset), CAST(N'2023-10-04T11:52:17.6865624+08:00' AS DateTimeOffset), NULL, 1, N'grade', N'C', N'user:1 use function "AddGrade()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (60, CAST(N'2023-10-04T11:52:19.0187968+08:00' AS DateTimeOffset), CAST(N'2023-10-04T11:52:19.0187968+08:00' AS DateTimeOffset), NULL, 1, N'grade', N'U', N'user:1 use function "SetGradeStatus()", set grade status to true id 2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (61, CAST(N'2023-10-04T11:52:19.5671981+08:00' AS DateTimeOffset), CAST(N'2023-10-04T11:52:19.5671981+08:00' AS DateTimeOffset), NULL, 1, N'grade', N'U', N'user:1 use function "SetGradeStatus()", set grade status to false id 2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (62, CAST(N'2023-10-04T11:52:22.2110632+08:00' AS DateTimeOffset), CAST(N'2023-10-04T11:52:22.2110632+08:00' AS DateTimeOffset), NULL, 1, N'grade', N'U', N'user:1 use function "EditGrade()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (63, CAST(N'2023-10-04T11:52:26.1606565+08:00' AS DateTimeOffset), CAST(N'2023-10-04T11:52:26.1606565+08:00' AS DateTimeOffset), NULL, 1, N'grade', N'D', N'user:1 use function "DeleteGrade()", delete rank id 2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (64, CAST(N'2023-10-04T11:52:41.1727300+08:00' AS DateTimeOffset), CAST(N'2023-10-04T11:52:41.1727300+08:00' AS DateTimeOffset), NULL, 1, N'grade', N'C', N'user:1 use function "AddGrade()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (65, CAST(N'2023-10-04T15:15:14.4956073+08:00' AS DateTimeOffset), CAST(N'2023-10-04T15:15:14.4956073+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:14')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (66, CAST(N'2023-10-04T15:15:28.7146043+08:00' AS DateTimeOffset), CAST(N'2023-10-04T15:15:28.7146043+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (67, CAST(N'2023-10-04T15:17:39.4246462+08:00' AS DateTimeOffset), CAST(N'2023-10-04T15:17:39.4246462+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (68, CAST(N'2023-10-04T15:17:43.0803156+08:00' AS DateTimeOffset), CAST(N'2023-10-04T15:17:43.0803156+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'U', N'user:1 use function "SetPermissionStatus()", set permission status to true id 1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (69, CAST(N'2023-10-04T15:19:30.1403897+08:00' AS DateTimeOffset), CAST(N'2023-10-04T15:19:30.1403897+08:00' AS DateTimeOffset), NULL, 1, N'rank', N'C', N'user:1 use function "AddRank()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (70, CAST(N'2023-10-04T15:22:18.9005276+08:00' AS DateTimeOffset), CAST(N'2023-10-04T15:22:18.9005276+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (71, CAST(N'2023-10-04T15:22:27.7945537+08:00' AS DateTimeOffset), CAST(N'2023-10-04T15:22:27.7945537+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'U', N'user:1 use function "EditPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (72, CAST(N'2023-10-04T15:22:28.9964735+08:00' AS DateTimeOffset), CAST(N'2023-10-04T15:22:28.9964735+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'U', N'user:1 use function "SetPermissionStatus()", set permission status to true id 1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (73, CAST(N'2023-10-04T15:22:33.5230974+08:00' AS DateTimeOffset), CAST(N'2023-10-04T15:22:33.5230974+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'U', N'user:1 use function "EditPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (74, CAST(N'2023-10-04T15:22:37.4128113+08:00' AS DateTimeOffset), CAST(N'2023-10-04T15:22:37.4128113+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'U', N'user:1 use function "EditPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (75, CAST(N'2023-10-04T15:22:38.8898101+08:00' AS DateTimeOffset), CAST(N'2023-10-04T15:22:38.8898101+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'D', N'user:1 use function "DeletePermission()", delete permission id 1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (76, CAST(N'2023-10-04T15:24:18.2591987+08:00' AS DateTimeOffset), CAST(N'2023-10-04T15:24:18.2591987+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (77, CAST(N'2023-10-04T16:24:34.4010887+08:00' AS DateTimeOffset), CAST(N'2023-10-04T16:24:34.4010887+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (78, CAST(N'2023-10-04T16:24:44.9529782+08:00' AS DateTimeOffset), CAST(N'2023-10-04T16:24:44.9529782+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (79, CAST(N'2023-10-04T16:24:54.8819281+08:00' AS DateTimeOffset), CAST(N'2023-10-04T16:24:54.8819281+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (80, CAST(N'2023-10-04T16:25:41.6711755+08:00' AS DateTimeOffset), CAST(N'2023-10-04T16:25:41.6711755+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (81, CAST(N'2023-10-04T16:25:51.9096623+08:00' AS DateTimeOffset), CAST(N'2023-10-04T16:25:51.9096623+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (82, CAST(N'2023-10-04T16:33:40.4780093+08:00' AS DateTimeOffset), CAST(N'2023-10-04T16:33:40.4780093+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (83, CAST(N'2023-10-04T16:35:20.8487213+08:00' AS DateTimeOffset), CAST(N'2023-10-04T16:35:20.8487213+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (84, CAST(N'2023-10-04T16:35:25.0545802+08:00' AS DateTimeOffset), CAST(N'2023-10-04T16:35:25.0545802+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (85, CAST(N'2023-10-04T16:56:18.4260726+08:00' AS DateTimeOffset), CAST(N'2023-10-04T16:56:18.4260726+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (86, CAST(N'2023-10-04T16:56:22.9300829+08:00' AS DateTimeOffset), CAST(N'2023-10-04T16:56:22.9300829+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'D', N'user:1 use function "DeletePermission()", delete permission id 2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (87, CAST(N'2023-10-04T16:56:24.7040379+08:00' AS DateTimeOffset), CAST(N'2023-10-04T16:56:24.7040379+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'D', N'user:1 use function "DeletePermission()", delete permission id 3')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (88, CAST(N'2023-10-04T16:56:25.7762664+08:00' AS DateTimeOffset), CAST(N'2023-10-04T16:56:25.7762664+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'D', N'user:1 use function "DeletePermission()", delete permission id 4')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (89, CAST(N'2023-10-04T16:56:27.5622391+08:00' AS DateTimeOffset), CAST(N'2023-10-04T16:56:27.5622391+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'D', N'user:1 use function "DeletePermission()", delete permission id 5')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (90, CAST(N'2023-10-04T16:56:28.7502238+08:00' AS DateTimeOffset), CAST(N'2023-10-04T16:56:28.7502238+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'D', N'user:1 use function "DeletePermission()", delete permission id 6')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (91, CAST(N'2023-10-04T16:56:44.8032305+08:00' AS DateTimeOffset), CAST(N'2023-10-04T16:56:44.8032305+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (92, CAST(N'2023-10-04T16:58:52.7836891+08:00' AS DateTimeOffset), CAST(N'2023-10-04T16:58:52.7836891+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (93, CAST(N'2023-10-04T16:59:56.5848249+08:00' AS DateTimeOffset), CAST(N'2023-10-04T16:59:56.5848249+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (94, CAST(N'2023-10-06T08:19:41.0686205+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:19:41.0686205+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'U', N'user:1 use function "EditPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (95, CAST(N'2023-10-06T08:20:54.7123773+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:20:54.7123773+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (96, CAST(N'2023-10-06T08:21:13.9117375+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:21:13.9117375+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (97, CAST(N'2023-10-06T08:21:31.2163920+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:21:31.2163920+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (98, CAST(N'2023-10-06T08:22:55.1159362+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:22:55.1159362+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (99, CAST(N'2023-10-06T08:23:06.4156319+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:23:06.4156319+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
GO
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (100, CAST(N'2023-10-06T08:23:22.0773035+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:23:22.0773035+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (101, CAST(N'2023-10-06T08:23:39.4930009+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:23:39.4930009+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (102, CAST(N'2023-10-06T08:23:51.7707814+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:23:51.7707814+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (103, CAST(N'2023-10-06T08:24:03.9384233+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:24:03.9384233+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (104, CAST(N'2023-10-06T08:24:40.9323197+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:24:40.9323197+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (105, CAST(N'2023-10-06T08:24:56.7494701+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:24:56.7494701+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (106, CAST(N'2023-10-06T08:25:11.5666633+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:25:11.5666633+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (107, CAST(N'2023-10-06T08:25:38.8684428+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:25:38.8684428+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (108, CAST(N'2023-10-06T08:25:44.9898895+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:25:44.9898895+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'U', N'user:1 use function "EditPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (109, CAST(N'2023-10-06T08:25:49.9271721+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:25:49.9271721+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'U', N'user:1 use function "EditPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (110, CAST(N'2023-10-06T08:25:56.8650321+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:25:56.8650321+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'U', N'user:1 use function "EditPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (111, CAST(N'2023-10-06T08:26:14.3083546+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:26:14.3083546+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (112, CAST(N'2023-10-06T08:26:31.2316753+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:26:31.2316753+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (113, CAST(N'2023-10-06T08:26:40.1204548+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:26:40.1204548+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'U', N'user:1 use function "EditPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (114, CAST(N'2023-10-06T08:26:46.9591975+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:26:46.9591975+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'U', N'user:1 use function "EditPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (115, CAST(N'2023-10-06T08:27:18.1629385+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:27:18.1629385+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (116, CAST(N'2023-10-06T08:27:35.6935531+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:27:35.6935531+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (117, CAST(N'2023-10-06T08:27:49.1414300+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:27:49.1414300+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (118, CAST(N'2023-10-06T08:28:10.1111201+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:28:10.1111201+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (119, CAST(N'2023-10-06T08:28:45.1030861+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:28:45.1030861+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (120, CAST(N'2023-10-06T08:28:59.1653124+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:28:59.1653124+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (121, CAST(N'2023-10-06T08:29:16.2914606+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:29:16.2914606+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (122, CAST(N'2023-10-06T08:29:35.6216898+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:29:35.6216898+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (123, CAST(N'2023-10-06T08:29:47.1838272+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:29:47.1838272+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (124, CAST(N'2023-10-06T08:30:12.8563982+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:30:12.8563982+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (125, CAST(N'2023-10-06T08:30:23.9684374+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:30:23.9684374+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (126, CAST(N'2023-10-06T08:30:38.1200108+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:30:38.1200108+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (127, CAST(N'2023-10-06T08:53:11.1829907+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:53:11.1829907+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (128, CAST(N'2023-10-06T08:54:12.6165429+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:54:12.6165429+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (129, CAST(N'2023-10-06T08:54:37.1874563+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:54:37.1874563+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (130, CAST(N'2023-10-06T08:56:09.8596117+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:56:09.8596117+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (131, CAST(N'2023-10-06T08:56:26.9512528+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:56:26.9512528+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (132, CAST(N'2023-10-06T08:56:53.6211853+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:56:53.6211853+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (133, CAST(N'2023-10-06T08:57:18.6143830+08:00' AS DateTimeOffset), CAST(N'2023-10-06T08:57:18.6143830+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (134, CAST(N'2023-10-06T09:07:51.2236555+08:00' AS DateTimeOffset), CAST(N'2023-10-06T09:07:51.2236555+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (135, CAST(N'2023-10-06T09:08:02.5656782+08:00' AS DateTimeOffset), CAST(N'2023-10-06T09:08:02.5656782+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "SetRoleStatus()", set roles status to false id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (136, CAST(N'2023-10-06T09:08:03.0967327+08:00' AS DateTimeOffset), CAST(N'2023-10-06T09:08:03.0967327+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "SetRoleStatus()", set roles status to true id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (137, CAST(N'2023-10-06T09:08:15.3120401+08:00' AS DateTimeOffset), CAST(N'2023-10-06T09:08:15.3120401+08:00' AS DateTimeOffset), NULL, 1, N'factory', N'U', N'user:1 use function "SetFactoryStatus()", set factory status to false id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (138, CAST(N'2023-10-06T09:08:15.7115401+08:00' AS DateTimeOffset), CAST(N'2023-10-06T09:08:15.7115401+08:00' AS DateTimeOffset), NULL, 1, N'factory', N'U', N'user:1 use function "SetFactoryStatus()", set factory status to true id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (139, CAST(N'2023-10-06T09:10:16.4186608+08:00' AS DateTimeOffset), CAST(N'2023-10-06T09:10:16.4186608+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'U', N'user:1 use function "EditPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (140, CAST(N'2023-10-06T09:10:32.5146506+08:00' AS DateTimeOffset), CAST(N'2023-10-06T09:10:32.5146506+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'U', N'user:1 use function "EditPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (141, CAST(N'2023-10-06T09:10:46.4553856+08:00' AS DateTimeOffset), CAST(N'2023-10-06T09:10:46.4553856+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "DeleteFactoryType()", set factory type status to false id:7')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (142, CAST(N'2023-10-06T09:10:47.0723384+08:00' AS DateTimeOffset), CAST(N'2023-10-06T09:10:47.0723384+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "DeleteFactoryType()", set factory type status to true id:7')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (143, CAST(N'2023-10-06T09:11:06.8642706+08:00' AS DateTimeOffset), CAST(N'2023-10-06T09:11:06.8642706+08:00' AS DateTimeOffset), NULL, 1, N'rank', N'U', N'user:1 use function "SetRankStatus()", set rank status to true id 2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (144, CAST(N'2023-10-06T09:11:11.3360611+08:00' AS DateTimeOffset), CAST(N'2023-10-06T09:11:11.3360611+08:00' AS DateTimeOffset), NULL, 1, N'grade', N'U', N'user:1 use function "SetGradeStatus()", set grade status to false id 3')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (145, CAST(N'2023-10-06T09:11:11.8518980+08:00' AS DateTimeOffset), CAST(N'2023-10-06T09:11:11.8518980+08:00' AS DateTimeOffset), NULL, 1, N'grade', N'U', N'user:1 use function "SetGradeStatus()", set grade status to true id 3')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (146, CAST(N'2023-10-06T09:11:16.0847507+08:00' AS DateTimeOffset), CAST(N'2023-10-06T09:11:16.0847507+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'U', N'user:1 use function "SetPermissionStatus()", set permission status to false id 7')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (147, CAST(N'2023-10-06T09:11:16.5717760+08:00' AS DateTimeOffset), CAST(N'2023-10-06T09:11:16.5717760+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'U', N'user:1 use function "SetPermissionStatus()", set permission status to true id 7')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (148, CAST(N'2023-10-11T11:09:49.6902324+08:00' AS DateTimeOffset), CAST(N'2023-10-11T11:09:49.6902324+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:4')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (149, CAST(N'2023-10-11T11:28:51.1253962+08:00' AS DateTimeOffset), CAST(N'2023-10-11T11:28:51.1253962+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:15')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (150, CAST(N'2023-10-11T11:53:50.9833396+08:00' AS DateTimeOffset), CAST(N'2023-10-11T11:53:50.9833396+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:16')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (151, CAST(N'2023-10-11T11:53:59.7608846+08:00' AS DateTimeOffset), CAST(N'2023-10-11T11:53:59.7608846+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (152, CAST(N'2023-10-13T09:00:28.6825552+08:00' AS DateTimeOffset), CAST(N'2023-10-13T09:00:28.6825552+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:17')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (153, CAST(N'2023-10-13T09:00:37.7602040+08:00' AS DateTimeOffset), CAST(N'2023-10-13T09:00:37.7602040+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (154, CAST(N'2023-10-13T10:55:59.3866760+08:00' AS DateTimeOffset), CAST(N'2023-10-13T10:55:59.3866760+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:18')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (155, CAST(N'2023-10-13T10:56:14.0505232+08:00' AS DateTimeOffset), CAST(N'2023-10-13T10:56:14.0505232+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:18')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (156, CAST(N'2023-10-13T10:56:44.3794707+08:00' AS DateTimeOffset), CAST(N'2023-10-13T10:56:44.3794707+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (157, CAST(N'2023-10-13T10:58:25.0684732+08:00' AS DateTimeOffset), CAST(N'2023-10-13T10:58:25.0684732+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'C', N'user:1 use function "AddVacationType()", create factory type id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (158, CAST(N'2023-10-13T11:02:53.5927282+08:00' AS DateTimeOffset), CAST(N'2023-10-13T11:02:53.5927282+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'U', N'user:1 use function "EditVacationType()", edit factory type id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (159, CAST(N'2023-10-13T11:02:56.9048594+08:00' AS DateTimeOffset), CAST(N'2023-10-13T11:02:56.9048594+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'U', N'user:1 use function "EditVacationType()", edit factory type id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (160, CAST(N'2023-10-13T11:03:20.9994552+08:00' AS DateTimeOffset), CAST(N'2023-10-13T11:03:20.9994552+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'C', N'user:1 use function "AddVacationType()", create factory type id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (161, CAST(N'2023-10-13T11:03:30.4080377+08:00' AS DateTimeOffset), CAST(N'2023-10-13T11:03:30.4080377+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'C', N'user:1 use function "AddVacationType()", create factory type id:3')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (162, CAST(N'2023-10-16T09:48:12.3369870+08:00' AS DateTimeOffset), CAST(N'2023-10-16T09:48:12.3369870+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'U', N'user:1 use function "EditVacationType()", edit factory type id:3')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (163, CAST(N'2023-10-16T09:48:34.3410688+08:00' AS DateTimeOffset), CAST(N'2023-10-16T09:48:34.3410688+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'C', N'user:1 use function "AddVacationType()", create factory type id:4')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (164, CAST(N'2023-10-16T09:50:50.9810098+08:00' AS DateTimeOffset), CAST(N'2023-10-16T09:50:50.9810098+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'U', N'user:1 use function "EditVacationType()", edit factory type id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (165, CAST(N'2023-10-16T09:50:59.1469498+08:00' AS DateTimeOffset), CAST(N'2023-10-16T09:50:59.1469498+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'U', N'user:1 use function "EditVacationType()", edit factory type id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (166, CAST(N'2023-10-16T09:53:10.8820810+08:00' AS DateTimeOffset), CAST(N'2023-10-16T09:53:10.8820810+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'U', N'user:1 use function "EditVacationType()", edit factory type id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (167, CAST(N'2023-10-16T09:53:14.9795109+08:00' AS DateTimeOffset), CAST(N'2023-10-16T09:53:14.9795109+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'U', N'user:1 use function "EditVacationType()", edit factory type id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (168, CAST(N'2023-10-16T11:15:15.8583143+08:00' AS DateTimeOffset), CAST(N'2023-10-16T11:15:15.8583143+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'C', N'user:1 use function "AddVacationType()", create factory type id:5')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (169, CAST(N'2023-10-16T11:24:23.8640107+08:00' AS DateTimeOffset), CAST(N'2023-10-16T11:24:23.8640107+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (170, CAST(N'2023-10-16T11:25:09.9255621+08:00' AS DateTimeOffset), CAST(N'2023-10-16T11:25:09.9255621+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (171, CAST(N'2023-10-17T08:53:09.5219345+08:00' AS DateTimeOffset), CAST(N'2023-10-17T08:53:09.5219345+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'U', N'user:1 use function "EditVacationType()", edit factory type id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (172, CAST(N'2023-10-17T08:53:20.4174193+08:00' AS DateTimeOffset), CAST(N'2023-10-17T08:53:20.4174193+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'U', N'user:1 use function "EditVacationType()", edit factory type id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (173, CAST(N'2023-10-17T08:53:28.9695570+08:00' AS DateTimeOffset), CAST(N'2023-10-17T08:53:28.9695570+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'U', N'user:1 use function "EditVacationType()", edit factory type id:3')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (174, CAST(N'2023-10-17T08:53:39.9007126+08:00' AS DateTimeOffset), CAST(N'2023-10-17T08:53:39.9007126+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'U', N'user:1 use function "EditVacationType()", edit factory type id:4')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (175, CAST(N'2023-10-17T09:00:26.4928219+08:00' AS DateTimeOffset), CAST(N'2023-10-17T09:00:26.4928219+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'U', N'user:1 use function "SetVacationTypeStatus()", set vacation type status to false id:5')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (176, CAST(N'2023-10-17T09:28:28.1303731+08:00' AS DateTimeOffset), CAST(N'2023-10-17T09:28:28.1303731+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (177, CAST(N'2023-10-17T09:49:34.1952088+08:00' AS DateTimeOffset), CAST(N'2023-10-17T09:49:34.1952088+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (178, CAST(N'2023-10-17T09:50:57.0853232+08:00' AS DateTimeOffset), CAST(N'2023-10-17T09:50:57.0853232+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (179, CAST(N'2023-10-17T11:31:19.1080371+08:00' AS DateTimeOffset), CAST(N'2023-10-17T11:31:19.1080371+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (180, CAST(N'2023-10-17T14:01:50.7340227+08:00' AS DateTimeOffset), CAST(N'2023-10-17T14:01:50.7340227+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (181, CAST(N'2023-10-17T14:04:50.1878065+08:00' AS DateTimeOffset), CAST(N'2023-10-17T14:04:50.1878065+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (182, CAST(N'2023-10-17T16:42:06.7936947+08:00' AS DateTimeOffset), CAST(N'2023-10-17T16:42:06.7936947+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'D', N'user:1 use function "DeleteVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (183, CAST(N'2023-10-17T16:42:31.2111335+08:00' AS DateTimeOffset), CAST(N'2023-10-17T16:42:31.2111335+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'D', N'user:1 use function "DeleteVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (184, CAST(N'2023-10-17T16:49:34.4553438+08:00' AS DateTimeOffset), CAST(N'2023-10-17T16:49:34.4553438+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'D', N'user:1 use function "DeleteVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (185, CAST(N'2023-10-17T16:50:44.3697652+08:00' AS DateTimeOffset), CAST(N'2023-10-17T16:50:44.3697652+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'D', N'user:1 use function "DeleteVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (186, CAST(N'2023-10-20T09:51:28.0285939+08:00' AS DateTimeOffset), CAST(N'2023-10-20T09:51:28.0285939+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (187, CAST(N'2023-10-20T09:57:48.2025169+08:00' AS DateTimeOffset), CAST(N'2023-10-20T09:57:48.2025169+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (188, CAST(N'2023-10-20T09:58:17.8844787+08:00' AS DateTimeOffset), CAST(N'2023-10-20T09:58:17.8844787+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (189, CAST(N'2023-10-20T09:58:33.8648406+08:00' AS DateTimeOffset), CAST(N'2023-10-20T09:58:33.8648406+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'D', N'user:1 use function "DeleteVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (190, CAST(N'2023-10-20T10:39:41.5333192+08:00' AS DateTimeOffset), CAST(N'2023-10-20T10:39:41.5333192+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (191, CAST(N'2023-10-20T10:42:36.5298426+08:00' AS DateTimeOffset), CAST(N'2023-10-20T10:42:36.5298426+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (192, CAST(N'2023-10-20T10:42:50.2599239+08:00' AS DateTimeOffset), CAST(N'2023-10-20T10:42:50.2599239+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (193, CAST(N'2023-10-20T10:47:11.6580472+08:00' AS DateTimeOffset), CAST(N'2023-10-20T10:47:11.6580472+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (194, CAST(N'2023-10-20T10:47:37.2151237+08:00' AS DateTimeOffset), CAST(N'2023-10-20T10:47:37.2151237+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (195, CAST(N'2023-10-20T10:48:07.2922422+08:00' AS DateTimeOffset), CAST(N'2023-10-20T10:48:07.2922422+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (196, CAST(N'2023-10-20T10:49:55.7741520+08:00' AS DateTimeOffset), CAST(N'2023-10-20T10:49:55.7741520+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (197, CAST(N'2023-10-20T10:50:27.7631691+08:00' AS DateTimeOffset), CAST(N'2023-10-20T10:50:27.7631691+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (198, CAST(N'2023-10-20T10:59:20.1631536+08:00' AS DateTimeOffset), CAST(N'2023-10-20T10:59:20.1631536+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (199, CAST(N'2023-10-20T10:59:51.0970309+08:00' AS DateTimeOffset), CAST(N'2023-10-20T10:59:51.0970309+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
GO
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (200, CAST(N'2023-10-20T11:00:06.5830270+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:00:06.5830270+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (201, CAST(N'2023-10-20T11:00:29.1786951+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:00:29.1786951+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (202, CAST(N'2023-10-20T11:01:05.0510416+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:05.0510416+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (203, CAST(N'2023-10-20T11:01:47.0724890+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:47.0724890+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (204, CAST(N'2023-10-20T11:02:37.0896830+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:02:37.0896830+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (205, CAST(N'2023-10-20T11:20:59.7060487+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:20:59.7060487+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (206, CAST(N'2023-10-20T11:22:16.9729397+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:22:16.9729397+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (207, CAST(N'2023-10-20T11:26:34.4221577+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:26:34.4221577+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (208, CAST(N'2023-10-20T11:26:44.7067686+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:26:44.7067686+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (209, CAST(N'2023-10-20T11:46:01.2758139+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:46:01.2758139+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:19')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (210, CAST(N'2023-10-20T11:51:21.2595802+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:51:21.2595802+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:19')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (211, CAST(N'2023-10-20T11:52:01.4514406+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:52:01.4514406+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (212, CAST(N'2023-10-20T11:52:35.4126967+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:52:35.4126967+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:19')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (213, CAST(N'2023-10-20T13:43:12.9124508+08:00' AS DateTimeOffset), CAST(N'2023-10-20T13:43:12.9124508+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:19')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (214, CAST(N'2023-10-20T15:05:42.3911397+08:00' AS DateTimeOffset), CAST(N'2023-10-20T15:05:42.3911397+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (215, CAST(N'2023-10-23T08:33:31.7066771+08:00' AS DateTimeOffset), CAST(N'2023-10-23T08:33:31.7066771+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:17')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (216, CAST(N'2023-10-23T09:12:36.1483768+08:00' AS DateTimeOffset), CAST(N'2023-10-23T09:12:36.1483768+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:20')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (217, CAST(N'2023-10-23T09:12:52.6944495+08:00' AS DateTimeOffset), CAST(N'2023-10-23T09:12:52.6944495+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (218, CAST(N'2023-10-23T09:17:10.5125856+08:00' AS DateTimeOffset), CAST(N'2023-10-23T09:17:10.5125856+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (219, CAST(N'2023-10-23T09:17:23.0727259+08:00' AS DateTimeOffset), CAST(N'2023-10-23T09:17:23.0727259+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (220, CAST(N'2023-10-23T09:17:37.4941847+08:00' AS DateTimeOffset), CAST(N'2023-10-23T09:17:37.4941847+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (221, CAST(N'2023-10-23T09:20:51.9227008+08:00' AS DateTimeOffset), CAST(N'2023-10-23T09:20:51.9227008+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (222, CAST(N'2023-10-23T10:01:56.0923705+08:00' AS DateTimeOffset), CAST(N'2023-10-23T10:01:56.0923705+08:00' AS DateTimeOffset), NULL, 1, N'work_classes', N'D', N'user:1 use function "DeleteWorkClass()", delete work class id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (223, CAST(N'2023-10-23T10:23:26.2033010+08:00' AS DateTimeOffset), CAST(N'2023-10-23T10:23:26.2033010+08:00' AS DateTimeOffset), NULL, 1, N'work_classes', N'U', N'user:1 use function "SetWorkClassStatus()", set work_classes status to false id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (224, CAST(N'2023-10-23T10:23:27.5162278+08:00' AS DateTimeOffset), CAST(N'2023-10-23T10:23:27.5162278+08:00' AS DateTimeOffset), NULL, 1, N'work_classes', N'U', N'user:1 use function "SetWorkClassStatus()", set work_classes status to true id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (225, CAST(N'2023-10-23T10:23:39.4594402+08:00' AS DateTimeOffset), CAST(N'2023-10-23T10:23:39.4594402+08:00' AS DateTimeOffset), NULL, 1, N'work_classes', N'U', N'user:1 use function "SetWorkClassStatus()", set work_classes status to false id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (226, CAST(N'2023-10-23T10:23:43.5372351+08:00' AS DateTimeOffset), CAST(N'2023-10-23T10:23:43.5372351+08:00' AS DateTimeOffset), NULL, 1, N'work_classes', N'U', N'user:1 use function "SetWorkClassStatus()", set work_classes status to true id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (227, CAST(N'2023-10-23T10:29:04.0840354+08:00' AS DateTimeOffset), CAST(N'2023-10-23T10:29:04.0840354+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (228, CAST(N'2023-10-23T10:29:17.7587022+08:00' AS DateTimeOffset), CAST(N'2023-10-23T10:29:17.7587022+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (229, CAST(N'2023-10-23T10:29:49.2960145+08:00' AS DateTimeOffset), CAST(N'2023-10-23T10:29:49.2960145+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (230, CAST(N'2023-10-23T10:29:59.5397049+08:00' AS DateTimeOffset), CAST(N'2023-10-23T10:29:59.5397049+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (231, CAST(N'2023-10-23T13:59:22.2501423+08:00' AS DateTimeOffset), CAST(N'2023-10-23T13:59:22.2501423+08:00' AS DateTimeOffset), NULL, 1, N'work_classes', N'U', N'user:1 use function "SetWorkClassStatus()", set work_classes status to false id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (232, CAST(N'2023-10-23T13:59:23.2773194+08:00' AS DateTimeOffset), CAST(N'2023-10-23T13:59:23.2773194+08:00' AS DateTimeOffset), NULL, 1, N'work_classes', N'U', N'user:1 use function "SetWorkClassStatus()", set work_classes status to true id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (233, CAST(N'2023-10-23T16:37:05.4912875+08:00' AS DateTimeOffset), CAST(N'2023-10-23T16:37:05.4912875+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:21')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (234, CAST(N'2023-10-23T16:39:13.5196215+08:00' AS DateTimeOffset), CAST(N'2023-10-23T16:39:13.5196215+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:22')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (235, CAST(N'2023-10-23T16:39:21.5252194+08:00' AS DateTimeOffset), CAST(N'2023-10-23T16:39:21.5252194+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (236, CAST(N'2023-10-24T13:49:11.5381862+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:49:11.5381862+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:20')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (237, CAST(N'2023-10-24T13:50:03.2208337+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:50:03.2208337+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'U', N'user:1 use function "EditPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (238, CAST(N'2023-10-24T13:50:08.8761752+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:50:08.8761752+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'U', N'user:1 use function "EditPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (239, CAST(N'2023-10-24T13:50:16.1630016+08:00' AS DateTimeOffset), CAST(N'2023-10-24T13:50:16.1630016+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'U', N'user:1 use function "EditPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (240, CAST(N'2023-10-24T14:38:04.3778542+08:00' AS DateTimeOffset), CAST(N'2023-10-24T14:38:04.3778542+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (241, CAST(N'2023-10-24T14:38:30.9012755+08:00' AS DateTimeOffset), CAST(N'2023-10-24T14:38:30.9012755+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (242, CAST(N'2023-10-24T14:39:49.7635438+08:00' AS DateTimeOffset), CAST(N'2023-10-24T14:39:49.7635438+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'U', N'user:1 use function "SetPermissionStatus()", set permission status to false id 47')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (243, CAST(N'2023-10-24T14:40:17.4246903+08:00' AS DateTimeOffset), CAST(N'2023-10-24T14:40:17.4246903+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'U', N'user:1 use function "SetPermissionStatus()", set permission status to true id 47')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (244, CAST(N'2023-10-24T15:38:49.5760743+08:00' AS DateTimeOffset), CAST(N'2023-10-24T15:38:49.5760743+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'U', N'user:1 use function "SetPermissionStatus()", set permission status to false id 47')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (245, CAST(N'2023-10-27T11:55:19.3568045+08:00' AS DateTimeOffset), CAST(N'2023-10-27T11:55:19.3568045+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:15')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (246, CAST(N'2023-10-27T11:55:39.0178749+08:00' AS DateTimeOffset), CAST(N'2023-10-27T11:55:39.0178749+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:15')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (247, CAST(N'2023-10-27T13:13:03.5162854+08:00' AS DateTimeOffset), CAST(N'2023-10-27T13:13:03.5162854+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:23')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (248, CAST(N'2023-10-27T13:13:26.8167991+08:00' AS DateTimeOffset), CAST(N'2023-10-27T13:13:26.8167991+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (249, CAST(N'2023-10-27T13:13:37.9681515+08:00' AS DateTimeOffset), CAST(N'2023-10-27T13:13:37.9681515+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (250, CAST(N'2023-10-27T13:16:57.1479871+08:00' AS DateTimeOffset), CAST(N'2023-10-27T13:16:57.1479871+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (251, CAST(N'2023-10-27T13:18:15.0311943+08:00' AS DateTimeOffset), CAST(N'2023-10-27T13:18:15.0311943+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (252, CAST(N'2023-10-27T13:18:22.7061870+08:00' AS DateTimeOffset), CAST(N'2023-10-27T13:18:22.7061870+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (253, CAST(N'2023-10-27T13:18:50.4123061+08:00' AS DateTimeOffset), CAST(N'2023-10-27T13:18:50.4123061+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:23')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (254, CAST(N'2023-10-27T13:19:44.4730933+08:00' AS DateTimeOffset), CAST(N'2023-10-27T13:19:44.4730933+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'D', N'user:1 use function "AddMenu()", delete menu id:3')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (255, CAST(N'2023-10-27T13:19:51.6545088+08:00' AS DateTimeOffset), CAST(N'2023-10-27T13:19:51.6545088+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'D', N'user:1 use function "AddMenu()", delete menu id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (256, CAST(N'2023-10-27T13:19:54.7782556+08:00' AS DateTimeOffset), CAST(N'2023-10-27T13:19:54.7782556+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'D', N'user:1 use function "AddMenu()", delete menu id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (257, CAST(N'2023-10-27T13:20:59.5665322+08:00' AS DateTimeOffset), CAST(N'2023-10-27T13:20:59.5665322+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:23')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (258, CAST(N'2023-10-27T13:21:40.3473217+08:00' AS DateTimeOffset), CAST(N'2023-10-27T13:21:40.3473217+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:17')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (259, CAST(N'2023-10-27T13:22:20.8107173+08:00' AS DateTimeOffset), CAST(N'2023-10-27T13:22:20.8107173+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:23')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (260, CAST(N'2023-10-27T13:22:24.6268188+08:00' AS DateTimeOffset), CAST(N'2023-10-27T13:22:24.6268188+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:17')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (261, CAST(N'2023-10-27T14:17:53.7321318+08:00' AS DateTimeOffset), CAST(N'2023-10-27T14:17:53.7321318+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:24')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (262, CAST(N'2023-10-27T14:18:05.3916397+08:00' AS DateTimeOffset), CAST(N'2023-10-27T14:18:05.3916397+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:18')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (263, CAST(N'2023-10-27T14:18:15.0207683+08:00' AS DateTimeOffset), CAST(N'2023-10-27T14:18:15.0207683+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (264, CAST(N'2023-10-27T14:23:11.2871763+08:00' AS DateTimeOffset), CAST(N'2023-10-27T14:23:11.2871763+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (265, CAST(N'2023-10-27T14:23:23.2251134+08:00' AS DateTimeOffset), CAST(N'2023-10-27T14:23:23.2251134+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (266, CAST(N'2023-10-27T14:23:33.8412243+08:00' AS DateTimeOffset), CAST(N'2023-10-27T14:23:33.8412243+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'C', N'user:1 use function "AddPermission()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (267, CAST(N'2023-10-27T14:23:43.9005863+08:00' AS DateTimeOffset), CAST(N'2023-10-27T14:23:43.9005863+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (268, CAST(N'2023-10-27T14:26:54.6056622+08:00' AS DateTimeOffset), CAST(N'2023-10-27T14:26:54.6056622+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'C', N'user:1 use function "AddLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (269, CAST(N'2023-10-27T14:30:56.9954930+08:00' AS DateTimeOffset), CAST(N'2023-10-27T14:30:56.9954930+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'D', N'user:1 use function "SetLeaveCoreStatus()", set status to false id = 1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (270, CAST(N'2023-10-27T14:30:58.0739603+08:00' AS DateTimeOffset), CAST(N'2023-10-27T14:30:58.0739603+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'D', N'user:1 use function "SetLeaveCoreStatus()", set status to true id = 1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (271, CAST(N'2023-10-27T14:31:03.5738565+08:00' AS DateTimeOffset), CAST(N'2023-10-27T14:31:03.5738565+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'D', N'user:1 use function "SetLeaveCoreStatus()", set status to false id = 1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (272, CAST(N'2023-10-27T14:31:08.3007741+08:00' AS DateTimeOffset), CAST(N'2023-10-27T14:31:08.3007741+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'D', N'user:1 use function "SetLeaveCoreStatus()", set status to true id = 1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (273, CAST(N'2023-10-27T14:31:20.5546656+08:00' AS DateTimeOffset), CAST(N'2023-10-27T14:31:20.5546656+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (274, CAST(N'2023-10-27T14:31:26.3389509+08:00' AS DateTimeOffset), CAST(N'2023-10-27T14:31:26.3389509+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (275, CAST(N'2023-10-30T10:59:37.3750954+08:00' AS DateTimeOffset), CAST(N'2023-10-30T10:59:37.3750954+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (276, CAST(N'2023-10-30T11:00:56.4018597+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:00:56.4018597+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (277, CAST(N'2023-10-30T11:00:59.7536710+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:00:59.7536710+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (278, CAST(N'2023-10-30T11:02:34.8828645+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:02:34.8828645+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (279, CAST(N'2023-10-30T11:03:01.1696190+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:03:01.1696190+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (280, CAST(N'2023-10-30T11:04:31.7772513+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:04:31.7772513+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (281, CAST(N'2023-10-30T11:10:01.4498937+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:10:01.4498937+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (282, CAST(N'2023-10-30T11:10:45.1455428+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:10:45.1455428+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (283, CAST(N'2023-10-30T11:13:24.5512784+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:13:24.5512784+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (284, CAST(N'2023-10-30T11:13:32.9135191+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:13:32.9135191+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'U', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (285, CAST(N'2023-10-30T11:14:02.6733992+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:14:02.6733992+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'U', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (286, CAST(N'2023-10-30T11:14:10.0219863+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:14:10.0219863+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'U', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (287, CAST(N'2023-10-30T11:19:48.1285272+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:19:48.1285272+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'U', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (288, CAST(N'2023-10-30T11:20:00.6399640+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:20:00.6399640+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'U', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (289, CAST(N'2023-10-30T11:22:34.0671166+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:22:34.0671166+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'U', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (290, CAST(N'2023-10-30T11:22:42.7661843+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:22:42.7661843+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'U', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (291, CAST(N'2023-10-30T11:39:36.5816258+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:39:36.5816258+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'D', N'user:1 use function "DeleteOvertime()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (292, CAST(N'2023-10-30T11:42:37.7481226+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:42:37.7481226+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (293, CAST(N'2023-10-30T11:43:08.6995182+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:43:08.6995182+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'U', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (294, CAST(N'2023-10-30T11:47:45.8205454+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:47:45.8205454+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'D', N'user:1 use function "DeleteOvertime()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (295, CAST(N'2023-10-30T11:48:38.7784459+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:48:38.7784459+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (296, CAST(N'2023-10-30T11:48:43.8907594+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:48:43.8907594+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'U', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (297, CAST(N'2023-10-30T11:48:50.1523871+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:48:50.1523871+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'D', N'user:1 use function "DeleteOvertime()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (298, CAST(N'2023-10-30T11:52:36.3570667+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:52:36.3570667+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (299, CAST(N'2023-10-30T11:52:47.4119689+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:52:47.4119689+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'D', N'user:1 use function "DeleteOvertime()"')
GO
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (300, CAST(N'2023-10-30T11:54:44.6078856+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:54:44.6078856+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (301, CAST(N'2023-10-30T11:54:51.4526879+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:54:51.4526879+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'D', N'user:1 use function "DeleteOvertime()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (302, CAST(N'2023-10-30T11:55:56.4936505+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:55:56.4936505+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (303, CAST(N'2023-10-30T11:55:59.2887009+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:55:59.2887009+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'D', N'user:1 use function "DeleteOvertime()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (304, CAST(N'2023-10-30T11:56:16.8377314+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:56:16.8377314+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (305, CAST(N'2023-10-30T11:56:22.5888005+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:56:22.5888005+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'U', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (306, CAST(N'2023-10-30T11:57:58.5922507+08:00' AS DateTimeOffset), CAST(N'2023-10-30T11:57:58.5922507+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'U', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (307, CAST(N'2023-10-30T13:14:55.5331863+08:00' AS DateTimeOffset), CAST(N'2023-10-30T13:14:55.5331863+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (308, CAST(N'2023-10-30T13:15:03.9616507+08:00' AS DateTimeOffset), CAST(N'2023-10-30T13:15:03.9616507+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (309, CAST(N'2023-10-30T13:15:07.2789435+08:00' AS DateTimeOffset), CAST(N'2023-10-30T13:15:07.2789435+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'U', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (310, CAST(N'2023-10-30T13:18:37.2227384+08:00' AS DateTimeOffset), CAST(N'2023-10-30T13:18:37.2227384+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (311, CAST(N'2023-10-30T13:18:49.8082348+08:00' AS DateTimeOffset), CAST(N'2023-10-30T13:18:49.8082348+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (312, CAST(N'2023-10-30T13:20:14.3766025+08:00' AS DateTimeOffset), CAST(N'2023-10-30T13:20:14.3766025+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (313, CAST(N'2023-10-30T14:10:49.5407516+08:00' AS DateTimeOffset), CAST(N'2023-10-30T14:10:49.5407516+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:21')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (314, CAST(N'2023-10-30T14:10:54.3567083+08:00' AS DateTimeOffset), CAST(N'2023-10-30T14:10:54.3567083+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:15')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (315, CAST(N'2023-10-30T14:10:59.4055187+08:00' AS DateTimeOffset), CAST(N'2023-10-30T14:10:59.4055187+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:4')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (316, CAST(N'2023-10-30T14:13:26.8383168+08:00' AS DateTimeOffset), CAST(N'2023-10-30T14:13:26.8383168+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:25')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (317, CAST(N'2023-10-30T14:14:36.1947935+08:00' AS DateTimeOffset), CAST(N'2023-10-30T14:14:36.1947935+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:26')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (318, CAST(N'2023-10-30T14:14:57.2116357+08:00' AS DateTimeOffset), CAST(N'2023-10-30T14:14:57.2116357+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (319, CAST(N'2023-10-30T14:15:32.9037628+08:00' AS DateTimeOffset), CAST(N'2023-10-30T14:15:32.9037628+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:25')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (320, CAST(N'2023-10-31T08:41:41.9745182+08:00' AS DateTimeOffset), CAST(N'2023-10-31T08:41:41.9745182+08:00' AS DateTimeOffset), NULL, 1, N'users', N'U', N'user:1 use function "SetStatus()", set user status to false id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (321, CAST(N'2023-10-31T08:41:42.6811584+08:00' AS DateTimeOffset), CAST(N'2023-10-31T08:41:42.6811584+08:00' AS DateTimeOffset), NULL, 1, N'users', N'U', N'user:1 use function "SetStatus()", set user status to true id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (322, CAST(N'2023-10-31T11:08:46.3152113+08:00' AS DateTimeOffset), CAST(N'2023-10-31T11:08:46.3152113+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (323, CAST(N'2023-10-31T11:08:49.1408932+08:00' AS DateTimeOffset), CAST(N'2023-10-31T11:08:49.1408932+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'D', N'user:1 use function "DeleteOvertime()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (324, CAST(N'2023-10-31T14:00:49.5365333+08:00' AS DateTimeOffset), CAST(N'2023-10-31T14:00:49.5365333+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'U', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (325, CAST(N'2023-10-31T14:00:58.1132110+08:00' AS DateTimeOffset), CAST(N'2023-10-31T14:00:58.1132110+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (326, CAST(N'2023-10-31T14:01:18.3565324+08:00' AS DateTimeOffset), CAST(N'2023-10-31T14:01:18.3565324+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'D', N'user:1 use function "DeleteOvertime()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (327, CAST(N'2023-10-31T14:07:48.4364257+08:00' AS DateTimeOffset), CAST(N'2023-10-31T14:07:48.4364257+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (328, CAST(N'2023-10-31T14:08:20.1663399+08:00' AS DateTimeOffset), CAST(N'2023-10-31T14:08:20.1663399+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'D', N'user:1 use function "DeleteOvertime()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (329, CAST(N'2023-10-31T14:08:50.6399166+08:00' AS DateTimeOffset), CAST(N'2023-10-31T14:08:50.6399166+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'D', N'user:1 use function "DeleteOvertime()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (330, CAST(N'2023-10-31T14:09:00.3616364+08:00' AS DateTimeOffset), CAST(N'2023-10-31T14:09:00.3616364+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (331, CAST(N'2023-11-01T09:19:40.7869654+08:00' AS DateTimeOffset), CAST(N'2023-11-01T09:19:40.7869654+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'D', N'user:1 use function "DeleteLeaveCore()", id = 1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (332, CAST(N'2023-11-01T09:20:16.5250384+08:00' AS DateTimeOffset), CAST(N'2023-11-01T09:20:16.5250384+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'C', N'user:1 use function "AddLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (333, CAST(N'2023-11-01T09:21:20.7719617+08:00' AS DateTimeOffset), CAST(N'2023-11-01T09:21:20.7719617+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'C', N'user:1 use function "AddLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (334, CAST(N'2023-11-01T09:22:46.0227707+08:00' AS DateTimeOffset), CAST(N'2023-11-01T09:22:46.0227707+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'C', N'user:1 use function "AddLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (335, CAST(N'2023-11-01T09:23:40.5798136+08:00' AS DateTimeOffset), CAST(N'2023-11-01T09:23:40.5798136+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'C', N'user:1 use function "AddLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (336, CAST(N'2023-11-01T09:24:31.5497260+08:00' AS DateTimeOffset), CAST(N'2023-11-01T09:24:31.5497260+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'C', N'user:1 use function "AddLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (337, CAST(N'2023-11-01T09:25:07.3703406+08:00' AS DateTimeOffset), CAST(N'2023-11-01T09:25:07.3703406+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'C', N'user:1 use function "AddLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (338, CAST(N'2023-11-01T09:26:01.6801296+08:00' AS DateTimeOffset), CAST(N'2023-11-01T09:26:01.6801296+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'C', N'user:1 use function "AddLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (339, CAST(N'2023-11-01T09:29:27.8532406+08:00' AS DateTimeOffset), CAST(N'2023-11-01T09:29:27.8532406+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'C', N'user:1 use function "AddLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (340, CAST(N'2023-11-01T09:29:47.5407639+08:00' AS DateTimeOffset), CAST(N'2023-11-01T09:29:47.5407639+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'C', N'user:1 use function "AddLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (341, CAST(N'2023-11-01T09:30:28.7126943+08:00' AS DateTimeOffset), CAST(N'2023-11-01T09:30:28.7126943+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'C', N'user:1 use function "AddLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (342, CAST(N'2023-11-01T09:33:30.6984055+08:00' AS DateTimeOffset), CAST(N'2023-11-01T09:33:30.6984055+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'C', N'user:1 use function "AddLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (343, CAST(N'2023-11-01T09:49:08.7176735+08:00' AS DateTimeOffset), CAST(N'2023-11-01T09:49:08.7176735+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:27')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (344, CAST(N'2023-11-01T09:51:09.4504207+08:00' AS DateTimeOffset), CAST(N'2023-11-01T09:51:09.4504207+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:19')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (345, CAST(N'2023-11-01T09:51:56.4807425+08:00' AS DateTimeOffset), CAST(N'2023-11-01T09:51:56.4807425+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (346, CAST(N'2023-11-01T10:59:34.2014671+08:00' AS DateTimeOffset), CAST(N'2023-11-01T10:59:34.2014671+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (347, CAST(N'2023-11-01T11:14:22.2863846+08:00' AS DateTimeOffset), CAST(N'2023-11-01T11:14:22.2863846+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (348, CAST(N'2023-11-01T11:14:29.3177914+08:00' AS DateTimeOffset), CAST(N'2023-11-01T11:14:29.3177914+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (349, CAST(N'2023-11-01T11:15:35.3818210+08:00' AS DateTimeOffset), CAST(N'2023-11-01T11:15:35.3818210+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (350, CAST(N'2023-11-01T11:24:43.2620062+08:00' AS DateTimeOffset), CAST(N'2023-11-01T11:24:43.2620062+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (351, CAST(N'2023-11-01T11:24:49.7407408+08:00' AS DateTimeOffset), CAST(N'2023-11-01T11:24:49.7407408+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (352, CAST(N'2023-11-01T11:24:56.7750341+08:00' AS DateTimeOffset), CAST(N'2023-11-01T11:24:56.7750341+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (353, CAST(N'2023-11-01T11:25:02.1560656+08:00' AS DateTimeOffset), CAST(N'2023-11-01T11:25:02.1560656+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (354, CAST(N'2023-11-06T13:28:21.3970976+08:00' AS DateTimeOffset), CAST(N'2023-11-06T13:28:21.3970976+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:23')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (355, CAST(N'2023-11-06T13:28:35.8812412+08:00' AS DateTimeOffset), CAST(N'2023-11-06T13:28:35.8812412+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:25')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (356, CAST(N'2023-11-06T13:28:42.3005216+08:00' AS DateTimeOffset), CAST(N'2023-11-06T13:28:42.3005216+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:26')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (357, CAST(N'2023-11-06T14:48:27.2252078+08:00' AS DateTimeOffset), CAST(N'2023-11-06T14:48:27.2252078+08:00' AS DateTimeOffset), NULL, 1, N'users', N'U', N'user:1 use function "EditUser()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (358, CAST(N'2023-11-06T14:48:53.9662923+08:00' AS DateTimeOffset), CAST(N'2023-11-06T14:48:53.9662923+08:00' AS DateTimeOffset), NULL, 1, N'users', N'U', N'user:1 use function "EditUser()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (359, CAST(N'2023-11-06T14:49:37.2429910+08:00' AS DateTimeOffset), CAST(N'2023-11-06T14:49:37.2429910+08:00' AS DateTimeOffset), NULL, 1, N'users', N'U', N'user:1 use function "EditUser()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (360, CAST(N'2023-11-06T14:51:45.1519231+08:00' AS DateTimeOffset), CAST(N'2023-11-06T14:51:45.1519231+08:00' AS DateTimeOffset), NULL, 1, N'users', N'U', N'user:1 use function "EditUser()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (361, CAST(N'2023-11-06T15:02:00.8816720+08:00' AS DateTimeOffset), CAST(N'2023-11-06T15:02:00.8816720+08:00' AS DateTimeOffset), NULL, 1, N'users', N'U', N'user:1 use function "EditUser()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (362, CAST(N'2023-11-06T15:02:24.0296667+08:00' AS DateTimeOffset), CAST(N'2023-11-06T15:02:24.0296667+08:00' AS DateTimeOffset), NULL, 1, N'users', N'U', N'user:1 use function "EditUser()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (363, CAST(N'2023-11-06T15:02:29.1683769+08:00' AS DateTimeOffset), CAST(N'2023-11-06T15:02:29.1683769+08:00' AS DateTimeOffset), NULL, 1, N'users', N'U', N'user:1 use function "EditUser()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (364, CAST(N'2023-11-07T10:28:16.3672327+08:00' AS DateTimeOffset), CAST(N'2023-11-07T10:28:16.3672327+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:23')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (365, CAST(N'2023-11-07T10:35:01.2458985+08:00' AS DateTimeOffset), CAST(N'2023-11-07T10:35:01.2458985+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:4')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (366, CAST(N'2023-11-07T10:36:26.9370197+08:00' AS DateTimeOffset), CAST(N'2023-11-07T10:36:26.9370197+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:11')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (367, CAST(N'2023-11-07T10:45:34.6550024+08:00' AS DateTimeOffset), CAST(N'2023-11-07T10:45:34.6550024+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:4')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (368, CAST(N'2023-11-07T10:46:36.4996748+08:00' AS DateTimeOffset), CAST(N'2023-11-07T10:46:36.4996748+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:11')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (369, CAST(N'2023-11-07T10:48:06.4208223+08:00' AS DateTimeOffset), CAST(N'2023-11-07T10:48:06.4208223+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:28')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (370, CAST(N'2023-11-07T10:48:59.6857903+08:00' AS DateTimeOffset), CAST(N'2023-11-07T10:48:59.6857903+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'C', N'user:1 use function "AddRole()", add role id:8')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (371, CAST(N'2023-11-07T10:49:10.4139060+08:00' AS DateTimeOffset), CAST(N'2023-11-07T10:49:10.4139060+08:00' AS DateTimeOffset), NULL, 1, N'users', N'U', N'user:1 use function "EditUser()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (372, CAST(N'2023-11-07T10:50:55.7987369+08:00' AS DateTimeOffset), CAST(N'2023-11-07T10:50:55.7987369+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:23')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (373, CAST(N'2023-11-07T11:30:04.0467247+08:00' AS DateTimeOffset), CAST(N'2023-11-07T11:30:04.0467247+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (374, CAST(N'2023-11-07T15:03:45.1136631+08:00' AS DateTimeOffset), CAST(N'2023-11-07T15:03:45.1136631+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:8')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (375, CAST(N'2023-11-07T15:07:22.8949677+08:00' AS DateTimeOffset), CAST(N'2023-11-07T15:07:22.8949677+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:28')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (376, CAST(N'2023-11-07T15:13:05.3622271+08:00' AS DateTimeOffset), CAST(N'2023-11-07T15:13:05.3622271+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:30')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (377, CAST(N'2023-11-07T15:14:13.4915671+08:00' AS DateTimeOffset), CAST(N'2023-11-07T15:14:13.4915671+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:28')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (378, CAST(N'2023-11-07T15:14:43.2769623+08:00' AS DateTimeOffset), CAST(N'2023-11-07T15:14:43.2769623+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:8')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (379, CAST(N'2023-11-07T15:15:40.4168112+08:00' AS DateTimeOffset), CAST(N'2023-11-07T15:15:40.4168112+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:30')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (380, CAST(N'2023-11-07T15:16:05.3328164+08:00' AS DateTimeOffset), CAST(N'2023-11-07T15:16:05.3328164+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:28')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (381, CAST(N'2023-11-07T15:16:36.5189110+08:00' AS DateTimeOffset), CAST(N'2023-11-07T15:16:36.5189110+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:28')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (382, CAST(N'2023-11-08T09:17:32.8603081+08:00' AS DateTimeOffset), CAST(N'2023-11-08T09:17:32.8603081+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:31')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (383, CAST(N'2023-11-08T09:23:59.2598557+08:00' AS DateTimeOffset), CAST(N'2023-11-08T09:23:59.2598557+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:32')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (384, CAST(N'2023-11-08T09:25:12.3272157+08:00' AS DateTimeOffset), CAST(N'2023-11-08T09:25:12.3272157+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:8')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (385, CAST(N'2023-11-10T13:13:47.6114651+08:00' AS DateTimeOffset), CAST(N'2023-11-10T13:13:47.6114651+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:9')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (386, CAST(N'2023-11-10T13:14:03.2085101+08:00' AS DateTimeOffset), CAST(N'2023-11-10T13:14:03.2085101+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:9')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (387, CAST(N'2023-11-10T13:15:53.7382048+08:00' AS DateTimeOffset), CAST(N'2023-11-10T13:15:53.7382048+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:9')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (388, CAST(N'2023-11-10T13:16:07.4319176+08:00' AS DateTimeOffset), CAST(N'2023-11-10T13:16:07.4319176+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:5')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (389, CAST(N'2023-11-10T13:32:11.8679639+08:00' AS DateTimeOffset), CAST(N'2023-11-10T13:32:11.8679639+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:33')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (390, CAST(N'2023-11-10T13:32:29.5455467+08:00' AS DateTimeOffset), CAST(N'2023-11-10T13:32:29.5455467+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (391, CAST(N'2023-11-13T15:13:17.1726943+08:00' AS DateTimeOffset), CAST(N'2023-11-13T15:13:17.1726943+08:00' AS DateTimeOffset), NULL, 1, N'work_classes', N'U', N'user:1 use function "SetWorkShiftStatus()", set work_classes status to false id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (392, CAST(N'2023-11-13T15:13:18.1440181+08:00' AS DateTimeOffset), CAST(N'2023-11-13T15:13:18.1440181+08:00' AS DateTimeOffset), NULL, 1, N'work_classes', N'U', N'user:1 use function "SetWorkShiftStatus()", set work_classes status to true id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (393, CAST(N'2023-11-22T10:05:10.0818412+08:00' AS DateTimeOffset), CAST(N'2023-11-22T10:05:10.0818412+08:00' AS DateTimeOffset), NULL, 1, N'departments', N'C', N'user:1 use function "AddDepartment()", add department id:0')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (394, CAST(N'2023-11-22T10:07:05.0717646+08:00' AS DateTimeOffset), CAST(N'2023-11-22T10:07:05.0717646+08:00' AS DateTimeOffset), NULL, 1, N'departments', N'U', N'user:1 use function "EditDepartment()", edit department id:10')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (395, CAST(N'2023-11-22T10:14:09.6460011+08:00' AS DateTimeOffset), CAST(N'2023-11-22T10:14:09.6460011+08:00' AS DateTimeOffset), NULL, 1, N'departments', N'U', N'user:1 use function "EditDepartment()", edit department id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (396, CAST(N'2023-11-22T10:55:18.3567422+08:00' AS DateTimeOffset), CAST(N'2023-11-22T10:55:18.3567422+08:00' AS DateTimeOffset), NULL, 1, N'departments', N'U', N'user:1 use function "EditDepartment()", edit department id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (397, CAST(N'2023-11-24T09:51:43.5138600+08:00' AS DateTimeOffset), CAST(N'2023-11-24T09:51:43.5138600+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:34')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (398, CAST(N'2023-11-24T09:52:05.6460640+08:00' AS DateTimeOffset), CAST(N'2023-11-24T09:52:05.6460640+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (399, CAST(N'2023-11-29T15:04:49.5227026+08:00' AS DateTimeOffset), CAST(N'2023-11-29T15:04:49.5227026+08:00' AS DateTimeOffset), NULL, 1, N'departments', N'U', N'user:1 use function "EditDepartment()", edit department id:1')
GO
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (400, CAST(N'2023-12-11T15:28:06.3701869+08:00' AS DateTimeOffset), CAST(N'2023-12-11T15:28:06.3701869+08:00' AS DateTimeOffset), NULL, 1, N'users', N'U', N'user:1 use function "EditUser()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (401, CAST(N'2023-12-11T15:28:12.2355728+08:00' AS DateTimeOffset), CAST(N'2023-12-11T15:28:12.2355728+08:00' AS DateTimeOffset), NULL, 1, N'users', N'U', N'user:1 use function "EditUser()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (402, CAST(N'2023-12-11T15:28:24.0986515+08:00' AS DateTimeOffset), CAST(N'2023-12-11T15:28:24.0986515+08:00' AS DateTimeOffset), NULL, 1, N'departments', N'U', N'user:1 use function "EditDepartment()", edit department id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (403, CAST(N'2023-12-18T15:03:40.2889984+08:00' AS DateTimeOffset), CAST(N'2023-12-18T15:03:40.2889984+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (404, CAST(N'2023-12-19T10:21:36.6989477+08:00' AS DateTimeOffset), CAST(N'2023-12-19T10:21:36.6989477+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:35')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (405, CAST(N'2023-12-19T10:23:48.7066626+08:00' AS DateTimeOffset), CAST(N'2023-12-19T10:23:48.7066626+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:8')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (406, CAST(N'2023-12-20T08:54:28.1027106+08:00' AS DateTimeOffset), CAST(N'2023-12-20T08:54:28.1027106+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:32')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (407, CAST(N'2023-12-20T08:54:43.3042197+08:00' AS DateTimeOffset), CAST(N'2023-12-20T08:54:43.3042197+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:32')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (408, CAST(N'2023-12-20T08:54:57.9488647+08:00' AS DateTimeOffset), CAST(N'2023-12-20T08:54:57.9488647+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:35')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (409, CAST(N'2023-12-20T09:13:04.5338564+08:00' AS DateTimeOffset), CAST(N'2023-12-20T09:13:04.5338564+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:32')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (410, CAST(N'2023-12-20T09:24:54.5028364+08:00' AS DateTimeOffset), CAST(N'2023-12-20T09:24:54.5028364+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:32')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (411, CAST(N'2023-12-20T09:25:14.7620161+08:00' AS DateTimeOffset), CAST(N'2023-12-20T09:25:14.7620161+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:32')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (412, CAST(N'2023-12-20T09:25:28.1709212+08:00' AS DateTimeOffset), CAST(N'2023-12-20T09:25:28.1709212+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:32')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (413, CAST(N'2023-12-20T09:39:28.4438081+08:00' AS DateTimeOffset), CAST(N'2023-12-20T09:39:28.4438081+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (414, CAST(N'2023-12-20T09:47:57.4584351+08:00' AS DateTimeOffset), CAST(N'2023-12-20T09:47:57.4584351+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (415, CAST(N'2023-12-20T09:48:10.4624017+08:00' AS DateTimeOffset), CAST(N'2023-12-20T09:48:10.4624017+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (416, CAST(N'2023-12-20T09:48:53.6069955+08:00' AS DateTimeOffset), CAST(N'2023-12-20T09:48:53.6069955+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (417, CAST(N'2023-12-20T09:49:10.2405035+08:00' AS DateTimeOffset), CAST(N'2023-12-20T09:49:10.2405035+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'D', N'user:1 use function "DeleteVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (418, CAST(N'2023-12-20T10:39:22.8783856+08:00' AS DateTimeOffset), CAST(N'2023-12-20T10:39:22.8783856+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'D', N'user:1 use function "DeletePermission()", delete permission id 44')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (419, CAST(N'2023-12-20T10:39:24.1658613+08:00' AS DateTimeOffset), CAST(N'2023-12-20T10:39:24.1658613+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'D', N'user:1 use function "DeletePermission()", delete permission id 45')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (420, CAST(N'2023-12-20T10:39:25.5366353+08:00' AS DateTimeOffset), CAST(N'2023-12-20T10:39:25.5366353+08:00' AS DateTimeOffset), NULL, 1, N'permission', N'D', N'user:1 use function "DeletePermission()", delete permission id 46')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (421, CAST(N'2023-12-20T10:47:30.3991216+08:00' AS DateTimeOffset), CAST(N'2023-12-20T10:47:30.3991216+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:36')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (422, CAST(N'2023-12-20T10:47:43.7836450+08:00' AS DateTimeOffset), CAST(N'2023-12-20T10:47:43.7836450+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:34')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (423, CAST(N'2023-12-20T10:48:18.1814293+08:00' AS DateTimeOffset), CAST(N'2023-12-20T10:48:18.1814293+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:2')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (424, CAST(N'2023-12-20T10:51:44.2271582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T10:51:44.2271582+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:34')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (425, CAST(N'2023-12-20T10:52:23.9108761+08:00' AS DateTimeOffset), CAST(N'2023-12-20T10:52:23.9108761+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:34')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (426, CAST(N'2023-12-20T14:37:59.9204375+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.9204375+08:00' AS DateTimeOffset), NULL, 1, N'vacations', N'C', N'user:1 use function "AddVacation()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (427, CAST(N'2023-12-22T15:18:24.8409393+08:00' AS DateTimeOffset), CAST(N'2023-12-22T15:18:24.8409393+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'U', N'user:1 use function "SetVacationTypeStatus()", set vacation type status to true id:5')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (428, CAST(N'2023-12-22T15:18:39.8395856+08:00' AS DateTimeOffset), CAST(N'2023-12-22T15:18:39.8395856+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'U', N'user:1 use function "SetVacationTypeStatus()", set vacation type status to false id:5')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (429, CAST(N'2023-12-22T15:18:41.6050020+08:00' AS DateTimeOffset), CAST(N'2023-12-22T15:18:41.6050020+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'D', N'user:1 use function "DeleteVacationType()", delete vacation type id:5')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (430, CAST(N'2023-12-22T15:20:41.8672515+08:00' AS DateTimeOffset), CAST(N'2023-12-22T15:20:41.8672515+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'C', N'user:1 use function "AddVacationType()", create factory type id:6')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (431, CAST(N'2023-12-22T15:20:43.4121405+08:00' AS DateTimeOffset), CAST(N'2023-12-22T15:20:43.4121405+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'U', N'user:1 use function "SetVacationTypeStatus()", set vacation type status to false id:6')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (432, CAST(N'2023-12-22T15:21:55.4585699+08:00' AS DateTimeOffset), CAST(N'2023-12-22T15:21:55.4585699+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (433, CAST(N'2023-12-22T15:23:06.4952596+08:00' AS DateTimeOffset), CAST(N'2023-12-22T15:23:06.4952596+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (434, CAST(N'2023-12-22T15:25:09.8162354+08:00' AS DateTimeOffset), CAST(N'2023-12-22T15:25:09.8162354+08:00' AS DateTimeOffset), NULL, 1, N'overtime_rate', N'C', N'user:1 use function "SaveOvertimeRate()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (435, CAST(N'2023-12-22T16:00:16.8281962+08:00' AS DateTimeOffset), CAST(N'2023-12-22T16:00:16.8281962+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'U', N'user:1 use function "EditVacationType()", edit factory type id:6')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (436, CAST(N'2023-12-25T09:40:42.1988921+08:00' AS DateTimeOffset), CAST(N'2023-12-25T09:40:42.1988921+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'U', N'user:1 use function "EditVacationType()", edit factory type id:6')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (437, CAST(N'2023-12-25T09:40:58.1382329+08:00' AS DateTimeOffset), CAST(N'2023-12-25T09:40:58.1382329+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'U', N'user:1 use function "EditVacationType()", edit factory type id:6')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (438, CAST(N'2023-12-25T09:41:53.3313993+08:00' AS DateTimeOffset), CAST(N'2023-12-25T09:41:53.3313993+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'U', N'user:1 use function "EditVacationType()", edit factory type id:6')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (439, CAST(N'2023-12-25T09:42:04.5534115+08:00' AS DateTimeOffset), CAST(N'2023-12-25T09:42:04.5534115+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'U', N'user:1 use function "EditVacationType()", edit factory type id:6')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (440, CAST(N'2023-12-25T09:42:48.7194156+08:00' AS DateTimeOffset), CAST(N'2023-12-25T09:42:48.7194156+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'U', N'user:1 use function "EditVacationType()", edit factory type id:6')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (441, CAST(N'2023-12-25T09:42:58.4584469+08:00' AS DateTimeOffset), CAST(N'2023-12-25T09:42:58.4584469+08:00' AS DateTimeOffset), NULL, 1, N'vacation_type', N'U', N'user:1 use function "EditVacationType()", edit factory type id:6')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (442, CAST(N'2023-12-25T10:00:00.5017673+08:00' AS DateTimeOffset), CAST(N'2023-12-25T10:00:00.5017673+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "DeleteFactoryType()", set factory type status to false id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (443, CAST(N'2023-12-25T10:00:01.2109222+08:00' AS DateTimeOffset), CAST(N'2023-12-25T10:00:01.2109222+08:00' AS DateTimeOffset), NULL, 1, N'factory_type', N'U', N'user:1 use function "DeleteFactoryType()", set factory type status to true id:1')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (444, CAST(N'2023-12-25T10:04:28.7862223+08:00' AS DateTimeOffset), CAST(N'2023-12-25T10:04:28.7862223+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'C', N'user:1 use function "AddRole()", add role id:9')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (445, CAST(N'2023-12-25T10:04:41.2911388+08:00' AS DateTimeOffset), CAST(N'2023-12-25T10:04:41.2911388+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:9')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (446, CAST(N'2023-12-25T10:04:51.3573771+08:00' AS DateTimeOffset), CAST(N'2023-12-25T10:04:51.3573771+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:9')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (447, CAST(N'2023-12-25T10:05:10.9894748+08:00' AS DateTimeOffset), CAST(N'2023-12-25T10:05:10.9894748+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'D', N'user:1 use function "DeleteRole()", delete roles id:9')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (448, CAST(N'2023-12-25T11:15:48.9663927+08:00' AS DateTimeOffset), CAST(N'2023-12-25T11:15:48.9663927+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (449, CAST(N'2023-12-25T11:16:12.6171605+08:00' AS DateTimeOffset), CAST(N'2023-12-25T11:16:12.6171605+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (450, CAST(N'2024-01-02T09:40:10.4862146+08:00' AS DateTimeOffset), CAST(N'2024-01-02T09:40:10.4862146+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:38')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (451, CAST(N'2024-01-02T09:40:46.1606926+08:00' AS DateTimeOffset), CAST(N'2024-01-02T09:40:46.1606926+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:8')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (452, CAST(N'2024-01-02T11:11:14.5968830+08:00' AS DateTimeOffset), CAST(N'2024-01-02T11:11:14.5968830+08:00' AS DateTimeOffset), NULL, 1, N'users', N'U', N'user:1 use function "ResetPassword()", reset user password id:3')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (453, CAST(N'2024-01-02T11:12:10.2950328+08:00' AS DateTimeOffset), CAST(N'2024-01-02T11:12:10.2950328+08:00' AS DateTimeOffset), NULL, 1, N'users', N'U', N'user:1 use function "ResetPassword()", reset user password id:3')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (454, CAST(N'2024-01-02T13:04:14.5950601+08:00' AS DateTimeOffset), CAST(N'2024-01-02T13:04:14.5950601+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'C', N'user:1 use function "AddMenu()", add menu id:39')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (455, CAST(N'2024-01-02T13:04:52.0687202+08:00' AS DateTimeOffset), CAST(N'2024-01-02T13:04:52.0687202+08:00' AS DateTimeOffset), NULL, 1, N'roles', N'U', N'user:1 use function "EditRole()", edit role id:8')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (456, CAST(N'2024-01-02T13:05:01.7663337+08:00' AS DateTimeOffset), CAST(N'2024-01-02T13:05:01.7663337+08:00' AS DateTimeOffset), NULL, 1, N'users', N'U', N'user:1 use function "EditUser()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (457, CAST(N'2024-01-02T16:45:41.4236459+08:00' AS DateTimeOffset), CAST(N'2024-01-02T16:45:41.4236459+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:38')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (458, CAST(N'2024-01-02T16:45:48.7240286+08:00' AS DateTimeOffset), CAST(N'2024-01-02T16:45:48.7240286+08:00' AS DateTimeOffset), NULL, 1, N'menus', N'U', N'user:1 use function "EditMenu()", edit menu id:38')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (459, CAST(N'2024-01-17T15:13:46.2043257+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:13:46.2043257+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (460, CAST(N'2024-01-17T15:13:58.5646461+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:13:58.5646461+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (461, CAST(N'2024-01-17T15:14:06.7090656+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:14:06.7090656+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (462, CAST(N'2024-01-17T15:47:18.2285540+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:47:18.2285540+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (463, CAST(N'2024-01-17T15:47:37.6964555+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:47:37.6964555+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (464, CAST(N'2024-01-17T15:49:13.6580793+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:49:13.6580793+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (465, CAST(N'2024-01-17T15:49:20.0587513+08:00' AS DateTimeOffset), CAST(N'2024-01-17T15:49:20.0587513+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (466, CAST(N'2024-01-22T09:06:09.9326550+08:00' AS DateTimeOffset), CAST(N'2024-01-22T09:06:09.9326550+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (467, CAST(N'2024-01-22T09:07:31.4549072+08:00' AS DateTimeOffset), CAST(N'2024-01-22T09:07:31.4549072+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (468, CAST(N'2024-01-22T09:07:46.2883862+08:00' AS DateTimeOffset), CAST(N'2024-01-22T09:07:46.2883862+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (469, CAST(N'2024-01-22T09:12:33.4047767+08:00' AS DateTimeOffset), CAST(N'2024-01-22T09:12:33.4047767+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (470, CAST(N'2024-01-22T09:13:04.1522952+08:00' AS DateTimeOffset), CAST(N'2024-01-22T09:13:04.1522952+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (471, CAST(N'2024-01-22T09:31:37.8323358+08:00' AS DateTimeOffset), CAST(N'2024-01-22T09:31:37.8323358+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (472, CAST(N'2024-01-22T10:07:03.0035329+08:00' AS DateTimeOffset), CAST(N'2024-01-22T10:07:03.0035329+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (473, CAST(N'2024-01-22T10:07:11.4238357+08:00' AS DateTimeOffset), CAST(N'2024-01-22T10:07:11.4238357+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (474, CAST(N'2024-01-22T10:07:21.2996317+08:00' AS DateTimeOffset), CAST(N'2024-01-22T10:07:21.2996317+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (475, CAST(N'2024-01-22T10:07:29.0819759+08:00' AS DateTimeOffset), CAST(N'2024-01-22T10:07:29.0819759+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (476, CAST(N'2024-01-22T10:07:52.3831674+08:00' AS DateTimeOffset), CAST(N'2024-01-22T10:07:52.3831674+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (477, CAST(N'2024-01-22T10:07:58.8492367+08:00' AS DateTimeOffset), CAST(N'2024-01-22T10:07:58.8492367+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (478, CAST(N'2024-01-22T10:08:49.1346862+08:00' AS DateTimeOffset), CAST(N'2024-01-22T10:08:49.1346862+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
INSERT [dbo].[system_logs] ([id], [created_at], [updated_at], [deleted_at], [user_id], [table], [crud], [message]) VALUES (479, CAST(N'2024-01-25T11:03:57.5845438+08:00' AS DateTimeOffset), CAST(N'2024-01-25T11:03:57.5845438+08:00' AS DateTimeOffset), NULL, 1, N'leave_core', N'U', N'user:1 use function "EditLeaveCore()"')
SET IDENTITY_INSERT [dbo].[system_logs] OFF
GO
INSERT [dbo].[user_role] ([role_id], [user_id]) VALUES (1, 1)
INSERT [dbo].[user_role] ([role_id], [user_id]) VALUES (1, 3)
INSERT [dbo].[user_role] ([role_id], [user_id]) VALUES (2, 1)
INSERT [dbo].[user_role] ([role_id], [user_id]) VALUES (8, 1)
INSERT [dbo].[user_role] ([role_id], [user_id]) VALUES (8, 3)
GO
SET IDENTITY_INSERT [dbo].[users] ON 

INSERT [dbo].[users] ([id], [created_at], [updated_at], [deleted_at], [username], [password], [real_name], [card], [birth], [arrived], [leaved], [status], [remark], [email], [mobile], [department_id], [factory_id], [factory_type_id], [rank_id], [grade_id], [desc], [avatar], [last_mover], [check_in_code], [backend]) VALUES (1, CAST(N'2023-09-25T16:02:16.8700000+00:00' AS DateTimeOffset), CAST(N'2023-11-07T16:05:48.6875801+08:00' AS DateTimeOffset), NULL, N'admin', N'$2a$10$ZEsBnQxmnh6u1CJ2dlXBlOQAsn5.lcrg7lo7A./HjDJgThEHlxwIS', N'管理員', N'G122123456', CAST(N'2023-09-24T00:00:00.0000000+00:00' AS DateTimeOffset), CAST(N'2023-09-25T00:00:00.0000000+00:00' AS DateTimeOffset), CAST(N'0001-01-01T00:00:00.0000000+00:00' AS DateTimeOffset), 1, N'管理員帳號', N'mucglliadia@mail.gj.com.tw', N'0900000000', 1, 0, 0, 0, 0, N'我是管理員', N'', 1, N'0000010163', 1)
INSERT [dbo].[users] ([id], [created_at], [updated_at], [deleted_at], [username], [password], [real_name], [card], [birth], [arrived], [leaved], [status], [remark], [email], [mobile], [department_id], [factory_id], [factory_type_id], [rank_id], [grade_id], [desc], [avatar], [last_mover], [check_in_code], [backend]) VALUES (3, CAST(N'2023-10-02T11:45:26.7529223+08:00' AS DateTimeOffset), CAST(N'2024-01-02T13:38:46.3965610+08:00' AS DateTimeOffset), NULL, N'mucgll0328', N'$2a$10$EBxybmOXmxw44uiBr6LIT.gnKrWJwH4jNRlBF4rqq6d4NeoQJZmVO', N'張育維', N'G122478901', CAST(N'0001-01-01T00:00:00.0000000+00:00' AS DateTimeOffset), CAST(N'2022-06-26T00:00:00.0000000+00:00' AS DateTimeOffset), CAST(N'0001-01-01T00:00:00.0000000+00:00' AS DateTimeOffset), 1, N'13256446', N'mucglliadia@mail.gj.com.tw', N'0900123456', 1, 0, 0, 0, 0, N'', N'', 1, N'', 0)
SET IDENTITY_INSERT [dbo].[users] OFF
GO
SET IDENTITY_INSERT [dbo].[vacation_types] ON 

INSERT [dbo].[vacation_types] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover], [color]) VALUES (1, CAST(N'2023-10-13T10:58:25.0635263+08:00' AS DateTimeOffset), CAST(N'2023-10-17T08:53:09.5122652+08:00' AS DateTimeOffset), NULL, N'rest', N'休息日', 1, N'', 1, N'#0BFA49')
INSERT [dbo].[vacation_types] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover], [color]) VALUES (2, CAST(N'2023-10-13T11:03:20.9914260+08:00' AS DateTimeOffset), CAST(N'2023-10-17T08:53:20.4092981+08:00' AS DateTimeOffset), NULL, N'holiday', N'例假日', 1, N'', 1, N'#65B2F3')
INSERT [dbo].[vacation_types] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover], [color]) VALUES (3, CAST(N'2023-10-13T11:03:30.4043683+08:00' AS DateTimeOffset), CAST(N'2023-10-17T08:53:28.9657399+08:00' AS DateTimeOffset), NULL, N'national', N'國定假日', 1, N'', 1, N'#FC7F01')
INSERT [dbo].[vacation_types] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover], [color]) VALUES (4, CAST(N'2023-10-16T09:48:34.3358190+08:00' AS DateTimeOffset), CAST(N'2023-10-17T08:53:39.8853811+08:00' AS DateTimeOffset), NULL, N'bridge', N'彈性放假日', 1, N'', 1, N'#9E255E')
INSERT [dbo].[vacation_types] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover], [color]) VALUES (5, CAST(N'2023-10-16T11:15:15.8531388+08:00' AS DateTimeOffset), CAST(N'2023-12-22T15:18:39.8340677+08:00' AS DateTimeOffset), CAST(N'2023-12-22T15:18:41.5902104+08:00' AS DateTimeOffset), N'customize', N'自訂', 0, N'', 1, N'')
INSERT [dbo].[vacation_types] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [last_mover], [color]) VALUES (6, CAST(N'2023-12-22T15:20:41.8628511+08:00' AS DateTimeOffset), CAST(N'2023-12-25T09:42:58.4447131+08:00' AS DateTimeOffset), NULL, N'normal', N'平日', 0, N'此欄為平日加班倍率設定用，請勿刪除、編輯', 1, N'#131313')
SET IDENTITY_INSERT [dbo].[vacation_types] OFF
GO
SET IDENTITY_INSERT [dbo].[vacations] ON 

INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (166, CAST(N'2023-10-20T11:00:06.5729292+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:00:06.5729292+08:00' AS DateTimeOffset), NULL, N'4b6326dd-44f6-4b5d-bff3-328b4424a480', N'9c9c9c01-d9b3-4496-8c90-f789acc57a3e', 3, 2023, 10, 10, N'雙十節放假', 1, N'2023-10-10', N'2023-10-10', N'year', N'2030-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (167, CAST(N'2023-10-20T11:00:06.5729292+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:00:06.5729292+08:00' AS DateTimeOffset), NULL, N'4b6326dd-44f6-4b5d-bff3-328b4424a480', N'9c9c9c01-d9b3-4496-8c90-f789acc57a3e', 3, 2024, 10, 10, N'雙十節放假', 1, N'2023-10-10', N'2023-10-10', N'year', N'2030-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (168, CAST(N'2023-10-20T11:00:06.5729292+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:00:06.5729292+08:00' AS DateTimeOffset), NULL, N'4b6326dd-44f6-4b5d-bff3-328b4424a480', N'9c9c9c01-d9b3-4496-8c90-f789acc57a3e', 3, 2025, 10, 10, N'雙十節放假', 1, N'2023-10-10', N'2023-10-10', N'year', N'2030-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (169, CAST(N'2023-10-20T11:00:06.5729292+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:00:06.5729292+08:00' AS DateTimeOffset), NULL, N'4b6326dd-44f6-4b5d-bff3-328b4424a480', N'9c9c9c01-d9b3-4496-8c90-f789acc57a3e', 3, 2026, 10, 10, N'雙十節放假', 1, N'2023-10-10', N'2023-10-10', N'year', N'2030-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (170, CAST(N'2023-10-20T11:00:06.5729292+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:00:06.5729292+08:00' AS DateTimeOffset), NULL, N'4b6326dd-44f6-4b5d-bff3-328b4424a480', N'9c9c9c01-d9b3-4496-8c90-f789acc57a3e', 3, 2027, 10, 10, N'雙十節放假', 1, N'2023-10-10', N'2023-10-10', N'year', N'2030-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (171, CAST(N'2023-10-20T11:00:06.5729292+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:00:06.5729292+08:00' AS DateTimeOffset), NULL, N'4b6326dd-44f6-4b5d-bff3-328b4424a480', N'9c9c9c01-d9b3-4496-8c90-f789acc57a3e', 3, 2028, 10, 10, N'雙十節放假', 1, N'2023-10-10', N'2023-10-10', N'year', N'2030-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (172, CAST(N'2023-10-20T11:00:06.5729292+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:00:06.5729292+08:00' AS DateTimeOffset), NULL, N'4b6326dd-44f6-4b5d-bff3-328b4424a480', N'9c9c9c01-d9b3-4496-8c90-f789acc57a3e', 3, 2029, 10, 10, N'雙十節放假', 1, N'2023-10-10', N'2023-10-10', N'year', N'2030-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (236, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 1, 7, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (237, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 1, 14, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (238, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 1, 21, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (239, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 1, 28, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (240, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 2, 4, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (241, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 2, 11, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (242, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 2, 18, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (243, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 2, 25, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (244, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 3, 4, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (245, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 3, 11, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (246, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 3, 18, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (247, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 3, 25, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (248, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 4, 1, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (249, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 4, 8, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (250, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 4, 15, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (251, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 4, 22, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (252, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 4, 29, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (253, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 5, 6, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (254, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 5, 13, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (255, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 5, 20, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (256, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 5, 27, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (257, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 6, 3, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (258, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 6, 10, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (259, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 6, 17, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (260, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 6, 24, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (261, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 7, 1, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (262, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 7, 8, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (263, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 7, 15, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (264, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 7, 22, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (265, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 7, 29, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (266, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 8, 5, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (267, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 8, 12, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (268, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 8, 19, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (269, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 8, 26, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (270, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 9, 2, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (271, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 9, 9, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (272, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 9, 16, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (273, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 9, 23, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (274, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 9, 30, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (275, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 10, 7, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (276, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 10, 14, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (277, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 10, 21, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (278, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 10, 28, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (279, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 11, 4, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (280, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 11, 11, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (281, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 11, 18, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (282, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 11, 25, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (283, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 12, 2, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (284, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 12, 9, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (285, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 12, 16, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (286, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 12, 23, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (287, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2023, 12, 30, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (288, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 1, 6, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (289, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 1, 13, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (290, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 1, 20, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (291, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 1, 27, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (292, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 2, 3, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (293, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 2, 10, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (294, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 2, 17, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (295, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 2, 24, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (296, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 3, 2, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (297, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 3, 9, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (298, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 3, 16, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (299, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 3, 23, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (300, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 3, 30, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (301, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 4, 6, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (302, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 4, 13, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (303, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 4, 20, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (304, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 4, 27, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (305, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 5, 4, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (306, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 5, 11, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (307, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 5, 18, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (308, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 5, 25, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (309, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 6, 1, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (310, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 6, 8, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (311, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 6, 15, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (312, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 6, 22, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (313, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 6, 29, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (314, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 7, 6, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (315, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 7, 13, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (316, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 7, 20, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (317, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 7, 27, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (318, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 8, 3, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (319, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 8, 10, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (320, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 8, 17, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (321, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 8, 24, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (322, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 8, 31, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (323, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 9, 7, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (324, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 9, 14, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (325, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 9, 21, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (326, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 9, 28, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (327, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 10, 5, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
GO
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (328, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 10, 12, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (329, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 10, 19, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (330, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 10, 26, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (331, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 11, 2, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (332, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 11, 9, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (333, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 11, 16, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (334, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 11, 23, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (335, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 11, 30, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (336, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 12, 7, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (337, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 12, 14, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (338, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 12, 21, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (339, CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:01:04.7442794+08:00' AS DateTimeOffset), NULL, N'17b08fee-c8ed-4704-b155-e5908aa75b61', N'aca04368-bbf2-4d0b-9ad7-b97309770875', 1, 2024, 12, 28, N'休息日通常為周六', 1, N'2023-01-07', N'2023-01-07', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (445, CAST(N'2023-10-20T11:02:37.0771486+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:02:37.0771486+08:00' AS DateTimeOffset), NULL, N'b4a74996-2e9c-4544-b08a-a750d637739b', N'5afa49ac-1bd3-4158-84a1-49b5698ed633', 3, 2023, 1, 1, N'中華民國開國紀念日', 1, N'2023-01-01', N'2023-01-01', N'year', N'2030-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (446, CAST(N'2023-10-20T11:02:37.0771486+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:02:37.0771486+08:00' AS DateTimeOffset), NULL, N'b4a74996-2e9c-4544-b08a-a750d637739b', N'5afa49ac-1bd3-4158-84a1-49b5698ed633', 3, 2024, 1, 1, N'中華民國開國紀念日', 1, N'2023-01-01', N'2023-01-01', N'year', N'2030-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (447, CAST(N'2023-10-20T11:02:37.0771486+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:02:37.0771486+08:00' AS DateTimeOffset), NULL, N'b4a74996-2e9c-4544-b08a-a750d637739b', N'5afa49ac-1bd3-4158-84a1-49b5698ed633', 3, 2025, 1, 1, N'中華民國開國紀念日', 1, N'2023-01-01', N'2023-01-01', N'year', N'2030-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (448, CAST(N'2023-10-20T11:02:37.0771486+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:02:37.0771486+08:00' AS DateTimeOffset), NULL, N'b4a74996-2e9c-4544-b08a-a750d637739b', N'5afa49ac-1bd3-4158-84a1-49b5698ed633', 3, 2026, 1, 1, N'中華民國開國紀念日', 1, N'2023-01-01', N'2023-01-01', N'year', N'2030-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (449, CAST(N'2023-10-20T11:02:37.0771486+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:02:37.0771486+08:00' AS DateTimeOffset), NULL, N'b4a74996-2e9c-4544-b08a-a750d637739b', N'5afa49ac-1bd3-4158-84a1-49b5698ed633', 3, 2027, 1, 1, N'中華民國開國紀念日', 1, N'2023-01-01', N'2023-01-01', N'year', N'2030-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (450, CAST(N'2023-10-20T11:02:37.0771486+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:02:37.0771486+08:00' AS DateTimeOffset), NULL, N'b4a74996-2e9c-4544-b08a-a750d637739b', N'5afa49ac-1bd3-4158-84a1-49b5698ed633', 3, 2028, 1, 1, N'中華民國開國紀念日', 1, N'2023-01-01', N'2023-01-01', N'year', N'2030-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (451, CAST(N'2023-10-20T11:02:37.0771486+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:02:37.0771486+08:00' AS DateTimeOffset), NULL, N'b4a74996-2e9c-4544-b08a-a750d637739b', N'5afa49ac-1bd3-4158-84a1-49b5698ed633', 3, 2029, 1, 1, N'中華民國開國紀念日', 1, N'2023-01-01', N'2023-01-01', N'year', N'2030-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (452, CAST(N'2023-10-20T11:02:37.0771486+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:02:37.0771486+08:00' AS DateTimeOffset), NULL, N'b4a74996-2e9c-4544-b08a-a750d637739b', N'5afa49ac-1bd3-4158-84a1-49b5698ed633', 3, 2030, 1, 1, N'中華民國開國紀念日', 1, N'2023-01-01', N'2023-01-01', N'year', N'2030-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (454, CAST(N'2023-10-20T11:22:16.9557700+08:00' AS DateTimeOffset), CAST(N'2023-10-20T11:22:16.9557700+08:00' AS DateTimeOffset), NULL, N'10e9c310-18ba-4c5e-afb3-7b864d614cea', N'00000000-0000-0000-0000-000000000000', 4, 2023, 10, 9, N'國慶彈性放假日', 1, N'2023-10-09', N'2023-10-09', N'none', N'2023-10-09')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (665, CAST(N'2023-10-20T15:05:42.3799229+08:00' AS DateTimeOffset), CAST(N'2023-10-20T15:05:42.3799229+08:00' AS DateTimeOffset), NULL, N'39eca04e-1e97-4000-988d-fc8e6e6d88a0', N'27353b2e-e8d2-450d-916d-2efbd380c0c2', 3, 2024, 2, 28, N'228', 1, N'2024-02-28', N'2024-02-28', N'year', N'2030-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (666, CAST(N'2023-10-20T15:05:42.3799229+08:00' AS DateTimeOffset), CAST(N'2023-10-20T15:05:42.3799229+08:00' AS DateTimeOffset), NULL, N'39eca04e-1e97-4000-988d-fc8e6e6d88a0', N'27353b2e-e8d2-450d-916d-2efbd380c0c2', 3, 2025, 2, 28, N'228', 1, N'2024-02-28', N'2024-02-28', N'year', N'2030-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (667, CAST(N'2023-10-20T15:05:42.3799229+08:00' AS DateTimeOffset), CAST(N'2023-10-20T15:05:42.3799229+08:00' AS DateTimeOffset), NULL, N'39eca04e-1e97-4000-988d-fc8e6e6d88a0', N'27353b2e-e8d2-450d-916d-2efbd380c0c2', 3, 2026, 2, 28, N'228', 1, N'2024-02-28', N'2024-02-28', N'year', N'2030-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (668, CAST(N'2023-10-20T15:05:42.3799229+08:00' AS DateTimeOffset), CAST(N'2023-10-20T15:05:42.3799229+08:00' AS DateTimeOffset), NULL, N'39eca04e-1e97-4000-988d-fc8e6e6d88a0', N'27353b2e-e8d2-450d-916d-2efbd380c0c2', 3, 2027, 2, 28, N'228', 1, N'2024-02-28', N'2024-02-28', N'year', N'2030-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (669, CAST(N'2023-10-20T15:05:42.3799229+08:00' AS DateTimeOffset), CAST(N'2023-10-20T15:05:42.3799229+08:00' AS DateTimeOffset), NULL, N'39eca04e-1e97-4000-988d-fc8e6e6d88a0', N'27353b2e-e8d2-450d-916d-2efbd380c0c2', 3, 2028, 2, 28, N'228', 1, N'2024-02-28', N'2024-02-28', N'year', N'2030-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (670, CAST(N'2023-10-20T15:05:42.3799229+08:00' AS DateTimeOffset), CAST(N'2023-10-20T15:05:42.3799229+08:00' AS DateTimeOffset), NULL, N'39eca04e-1e97-4000-988d-fc8e6e6d88a0', N'27353b2e-e8d2-450d-916d-2efbd380c0c2', 3, 2029, 2, 28, N'228', 1, N'2024-02-28', N'2024-02-28', N'year', N'2030-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (741, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 1, 1, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (742, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 1, 8, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (743, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 1, 15, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (744, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 1, 22, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (745, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 1, 29, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (746, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 2, 5, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (747, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 2, 12, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (748, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 2, 19, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (749, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 2, 26, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (750, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 3, 5, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (751, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 3, 12, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (752, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 3, 19, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (753, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 3, 26, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (754, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 4, 2, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (755, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 4, 9, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (756, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 4, 16, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (757, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 4, 23, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (758, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 4, 30, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (759, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 5, 7, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (760, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 5, 14, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (761, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 5, 21, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (762, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 5, 28, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (763, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 6, 4, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (764, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 6, 11, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (765, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 6, 18, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (766, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 6, 25, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (767, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 7, 2, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (768, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 7, 9, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (769, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 7, 16, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (770, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 7, 23, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (771, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 7, 30, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (772, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 8, 6, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (773, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 8, 13, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (774, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 8, 20, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (775, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 8, 27, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (776, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 9, 3, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (777, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 9, 10, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (778, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 9, 17, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (779, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 9, 24, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (780, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 10, 1, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (781, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 10, 8, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (782, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 10, 15, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (783, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 10, 22, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (784, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 10, 29, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (785, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 11, 5, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (786, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 11, 12, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (787, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 11, 19, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (788, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 11, 26, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (789, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 12, 3, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (790, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 12, 10, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (791, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 12, 17, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (792, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 12, 24, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (793, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2023, 12, 31, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (794, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 1, 7, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (795, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 1, 14, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (796, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 1, 21, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (797, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 1, 28, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (798, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 2, 4, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (799, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 2, 11, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (800, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 2, 18, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (801, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 2, 25, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (802, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 3, 3, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (803, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 3, 10, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (804, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 3, 17, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (805, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 3, 24, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (806, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 3, 31, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (807, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 4, 7, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (808, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 4, 14, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (809, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 4, 21, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (810, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 4, 28, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (811, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 5, 5, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (812, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 5, 12, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (813, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 5, 19, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
GO
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (814, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 5, 26, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (815, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 6, 2, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (816, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 6, 9, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (817, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 6, 16, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (818, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 6, 23, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (819, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 6, 30, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (820, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 7, 7, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (821, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 7, 14, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (822, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 7, 21, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (823, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 7, 28, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (824, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 8, 4, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (825, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 8, 11, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (826, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 8, 18, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (827, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 8, 25, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (828, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 9, 1, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (829, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 9, 8, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (830, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 9, 15, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (831, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 9, 22, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (832, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 9, 29, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (833, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 10, 6, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (834, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 10, 13, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (835, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 10, 20, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (836, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 10, 27, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (837, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 11, 3, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (838, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 11, 10, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (839, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 11, 17, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (840, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 11, 24, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (841, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 12, 1, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (842, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 12, 8, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (843, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 12, 15, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (844, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 12, 22, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
INSERT [dbo].[vacations] ([id], [created_at], [updated_at], [deleted_at], [general_key], [internal_key], [type_id], [year], [month], [day], [remark], [last_mover], [start_date], [end_date], [repeat], [end_repeat]) VALUES (845, CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:37:59.5904582+08:00' AS DateTimeOffset), NULL, N'6f079bf2-2b6e-4341-b82e-b596dad4ed99', N'0f14ce2b-c645-44b2-bd10-8e21e444962e', 3, 2024, 12, 29, N'例假日通常為周日', 1, N'2023-01-01', N'2023-01-01', N'week', N'2025-01-01')
SET IDENTITY_INSERT [dbo].[vacations] OFF
GO
SET IDENTITY_INSERT [dbo].[work_schedule] ON 

INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (892, CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), NULL, 2024, 1, 17, 3, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (893, CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), NULL, 2024, 1, 1, 3, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (894, CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), NULL, 2024, 1, 12, 3, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (895, CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), NULL, 2024, 1, 9, 3, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (896, CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), NULL, 2024, 1, 16, 3, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (897, CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), NULL, 2024, 1, 3, 3, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (898, CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), NULL, 2024, 1, 5, 3, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (899, CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), NULL, 2024, 1, 2, 3, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (900, CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), NULL, 2024, 1, 4, 3, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (901, CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), NULL, 2024, 1, 8, 3, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (902, CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), NULL, 2024, 1, 15, 3, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (903, CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), NULL, 2024, 1, 10, 3, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (904, CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), CAST(N'2023-12-15T15:07:58.2803292+08:00' AS DateTimeOffset), NULL, 2024, 1, 11, 3, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (959, CAST(N'2023-12-18T15:22:49.1173348+08:00' AS DateTimeOffset), CAST(N'2023-12-18T15:22:49.1173348+08:00' AS DateTimeOffset), NULL, 2023, 11, 30, 1, 3)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (960, CAST(N'2023-12-18T15:22:49.1173348+08:00' AS DateTimeOffset), CAST(N'2023-12-18T15:22:49.1173348+08:00' AS DateTimeOffset), NULL, 2023, 11, 30, 1, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (961, CAST(N'2023-12-18T15:22:49.1173348+08:00' AS DateTimeOffset), CAST(N'2023-12-18T15:22:49.1173348+08:00' AS DateTimeOffset), NULL, 2023, 11, 30, 1, 5)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (979, CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), NULL, 2023, 12, 12, 3, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (981, CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), NULL, 2023, 12, 15, 3, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (983, CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), NULL, 2023, 12, 20, 3, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (985, CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), NULL, 2023, 12, 14, 3, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (987, CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), NULL, 2023, 12, 16, 3, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (989, CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), NULL, 2023, 12, 21, 3, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (991, CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), NULL, 2023, 12, 5, 3, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (993, CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), NULL, 2023, 12, 6, 3, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (995, CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), NULL, 2023, 12, 8, 3, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (997, CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), NULL, 2023, 12, 1, 3, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (999, CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), NULL, 2023, 12, 22, 3, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1001, CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), NULL, 2023, 12, 7, 3, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1003, CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), NULL, 2023, 12, 9, 3, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1005, CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), NULL, 2023, 12, 13, 3, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1007, CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), NULL, 2023, 12, 2, 3, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1009, CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), CAST(N'2023-12-20T14:41:26.2083861+08:00' AS DateTimeOffset), NULL, 2023, 12, 27, 3, 3)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1010, CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), NULL, 2023, 12, 14, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1011, CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), NULL, 2023, 12, 22, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1012, CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), NULL, 2023, 12, 27, 1, 3)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1013, CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), NULL, 2023, 12, 9, 1, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1014, CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), NULL, 2023, 12, 28, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1015, CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), NULL, 2023, 12, 1, 1, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1016, CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), NULL, 2023, 12, 20, 1, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1017, CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), NULL, 2023, 12, 5, 1, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1018, CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), NULL, 2023, 12, 7, 1, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1019, CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), NULL, 2023, 12, 8, 1, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1020, CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), NULL, 2023, 12, 12, 1, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1021, CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), NULL, 2023, 12, 13, 1, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1022, CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), NULL, 2023, 12, 15, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1023, CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), NULL, 2023, 12, 16, 1, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1024, CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), NULL, 2023, 12, 2, 1, 2)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1025, CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), NULL, 2023, 12, 21, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1026, CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), NULL, 2023, 12, 6, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1027, CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), CAST(N'2023-12-25T08:40:15.1828560+08:00' AS DateTimeOffset), NULL, 2023, 12, 29, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1088, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 17, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1089, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 19, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1090, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 21, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1091, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 31, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1092, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 4, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1093, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 14, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1094, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 2, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1095, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 23, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1096, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 24, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1097, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 5, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1098, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 8, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1099, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 9, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1100, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 10, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1101, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 13, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1102, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 27, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1103, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 11, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1104, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 20, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1105, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 22, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1106, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 1, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1107, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 12, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1108, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 28, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1109, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 3, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1110, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 18, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1111, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 30, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1112, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 26, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1113, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 16, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1114, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 25, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1115, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 29, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1116, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 7, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1117, CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), CAST(N'2024-01-12T15:29:57.6839354+08:00' AS DateTimeOffset), NULL, 2024, 1, 15, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1118, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 3, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1119, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 9, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1120, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 29, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1121, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 12, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1122, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 5, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1123, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 14, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1124, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 18, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1125, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 13, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1126, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 2, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1127, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 23, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1128, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 20, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1129, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 7, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1130, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 6, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1131, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 1, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1132, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 17, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1133, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 22, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1134, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 28, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1135, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 19, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1136, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 8, 1, 1)
GO
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1137, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 15, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1138, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 24, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1139, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 26, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1140, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 11, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1141, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 4, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1142, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 10, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1143, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 16, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1144, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 21, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1145, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 27, 1, 1)
INSERT [dbo].[work_schedule] ([id], [created_at], [updated_at], [deleted_at], [year], [month], [day], [user_id], [work_shift_id]) VALUES (1146, CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), CAST(N'2024-02-16T09:01:15.5954126+08:00' AS DateTimeOffset), NULL, 2024, 2, 25, 1, 1)
SET IDENTITY_INSERT [dbo].[work_schedule] OFF
GO
SET IDENTITY_INSERT [dbo].[work_shift] ON 

INSERT [dbo].[work_shift] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [start_time], [end_time], [rest_start_time], [rest_end_time], [last_mover], [color]) VALUES (1, CAST(N'2023-10-24T13:50:48.0164764+08:00' AS DateTimeOffset), CAST(N'2023-11-22T11:24:43.8546090+08:00' AS DateTimeOffset), NULL, N'morning', N'早班', 1, N'', N'08:00', N'17:00', N'12:00', N'13:00', 1, N'#0BFA49')
INSERT [dbo].[work_shift] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [start_time], [end_time], [rest_start_time], [rest_end_time], [last_mover], [color]) VALUES (2, CAST(N'2023-11-13T16:48:07.1924612+08:00' AS DateTimeOffset), CAST(N'2023-12-15T15:41:21.0864378+08:00' AS DateTimeOffset), NULL, N'night', N'晚班', 1, N'', N'17:00', N'21:00', N'18:00', N'19:00', 1, N'#65B2F3')
INSERT [dbo].[work_shift] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [start_time], [end_time], [rest_start_time], [rest_end_time], [last_mover], [color]) VALUES (3, CAST(N'2023-11-14T10:48:24.5994036+08:00' AS DateTimeOffset), CAST(N'2023-11-14T10:48:24.6333724+08:00' AS DateTimeOffset), NULL, N'test1', N'大夜班', 1, N'', N'21:00', N'23:59', N'22:00', N'23:00', 1, N'#FC7F01')
INSERT [dbo].[work_shift] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [start_time], [end_time], [rest_start_time], [rest_end_time], [last_mover], [color]) VALUES (4, CAST(N'2023-11-14T11:13:32.4450875+08:00' AS DateTimeOffset), CAST(N'2023-11-15T08:24:32.0532894+08:00' AS DateTimeOffset), NULL, N'test2', N'全天班', 1, N'', N'00:00', N'23:59', N'03:00', N'04:00', 1, N'#847CC2')
INSERT [dbo].[work_shift] ([id], [created_at], [updated_at], [deleted_at], [code], [name], [status], [remark], [start_time], [end_time], [rest_start_time], [rest_end_time], [last_mover], [color]) VALUES (5, CAST(N'2023-12-15T16:28:15.6542350+08:00' AS DateTimeOffset), CAST(N'2023-12-15T16:28:16.0225938+08:00' AS DateTimeOffset), NULL, N'test3', N'下午班', 1, N'', N'15:00', N'17:00', N'15:00', N'16:00', 1, N'#8305A1')
SET IDENTITY_INSERT [dbo].[work_shift] OFF
GO
SET ANSI_PADDING ON
GO
/****** Object:  Index [UQ__departme__357D4CF9B7B5F09A]    Script Date: 2024/2/26 上午 10:00:27 ******/
ALTER TABLE [dbo].[department] ADD UNIQUE NONCLUSTERED 
(
	[code] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
GO
SET ANSI_PADDING ON
GO
/****** Object:  Index [UQ__factory__357D4CF997DBC82F]    Script Date: 2024/2/26 上午 10:00:27 ******/
ALTER TABLE [dbo].[factory] ADD UNIQUE NONCLUSTERED 
(
	[code] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
GO
SET ANSI_PADDING ON
GO
/****** Object:  Index [UQ__factory___357D4CF9F4573E10]    Script Date: 2024/2/26 上午 10:00:27 ******/
ALTER TABLE [dbo].[factory_type] ADD UNIQUE NONCLUSTERED 
(
	[code] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
GO
SET ANSI_PADDING ON
GO
/****** Object:  Index [UQ__grades__357D4CF908CAA2F0]    Script Date: 2024/2/26 上午 10:00:27 ******/
ALTER TABLE [dbo].[grades] ADD UNIQUE NONCLUSTERED 
(
	[code] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
GO
SET ANSI_PADDING ON
GO
/****** Object:  Index [UQ__leave_co__357D4CF96B61C62E]    Script Date: 2024/2/26 上午 10:00:27 ******/
ALTER TABLE [dbo].[leave_core] ADD UNIQUE NONCLUSTERED 
(
	[code] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
GO
SET ANSI_PADDING ON
GO
/****** Object:  Index [UQ__permissi__357D4CF9BBC75B51]    Script Date: 2024/2/26 上午 10:00:27 ******/
ALTER TABLE [dbo].[permissions] ADD UNIQUE NONCLUSTERED 
(
	[code] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
GO
SET ANSI_PADDING ON
GO
/****** Object:  Index [UQ__ranks__357D4CF9F05C8630]    Script Date: 2024/2/26 上午 10:00:27 ******/
ALTER TABLE [dbo].[ranks] ADD UNIQUE NONCLUSTERED 
(
	[code] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
GO
SET ANSI_PADDING ON
GO
/****** Object:  Index [UQ__roles__357D4CF9E8DA137E]    Script Date: 2024/2/26 上午 10:00:27 ******/
ALTER TABLE [dbo].[roles] ADD UNIQUE NONCLUSTERED 
(
	[code] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
GO
SET ANSI_PADDING ON
GO
/****** Object:  Index [UQ__vacation__357D4CF9544A43A4]    Script Date: 2024/2/26 上午 10:00:27 ******/
ALTER TABLE [dbo].[vacation_types] ADD UNIQUE NONCLUSTERED 
(
	[code] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
GO
SET ANSI_PADDING ON
GO
/****** Object:  Index [UQ__work_shi__357D4CF91A3ADDC1]    Script Date: 2024/2/26 上午 10:00:27 ******/
ALTER TABLE [dbo].[work_shift] ADD UNIQUE NONCLUSTERED 
(
	[code] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
GO
ALTER TABLE [dbo].[check-in_data] ADD  DEFAULT ((0)) FOR [work_shift_id]
GO
ALTER TABLE [dbo].[department] ADD  DEFAULT ((1)) FOR [status]
GO
ALTER TABLE [dbo].[department] ADD  DEFAULT ((0)) FOR [parent_id]
GO
ALTER TABLE [dbo].[department] ADD  DEFAULT ((0)) FOR [last_mover]
GO
ALTER TABLE [dbo].[department] ADD  DEFAULT ((0)) FOR [manager_id]
GO
ALTER TABLE [dbo].[factory] ADD  DEFAULT ((1)) FOR [status]
GO
ALTER TABLE [dbo].[factory] ADD  DEFAULT ((0)) FOR [last_mover]
GO
ALTER TABLE [dbo].[factory_type] ADD  DEFAULT ((1)) FOR [status]
GO
ALTER TABLE [dbo].[factory_type] ADD  DEFAULT ((0)) FOR [last_mover]
GO
ALTER TABLE [dbo].[grades] ADD  DEFAULT ((1)) FOR [status]
GO
ALTER TABLE [dbo].[grades] ADD  DEFAULT ((0)) FOR [last_mover]
GO
ALTER TABLE [dbo].[leave_core] ADD  DEFAULT ((1)) FOR [status]
GO
ALTER TABLE [dbo].[leave_core] ADD  DEFAULT ((0)) FOR [last_mover]
GO
ALTER TABLE [dbo].[leave_core] ADD  DEFAULT ((0)) FOR [deferred]
GO
ALTER TABLE [dbo].[leave_core] ADD  DEFAULT ((0)) FOR [exchange]
GO
ALTER TABLE [dbo].[leave_core] ADD  DEFAULT ((1)) FOR [cycle]
GO
ALTER TABLE [dbo].[leave_core] ADD  DEFAULT ((0)) FOR [maturity]
GO
ALTER TABLE [dbo].[leave_core] ADD  DEFAULT ((0)) FOR [is_maturity]
GO
ALTER TABLE [dbo].[leave_correct] ADD  DEFAULT ((0)) FOR [user_id]
GO
ALTER TABLE [dbo].[leave_correct] ADD  DEFAULT ((0)) FOR [leave_core_id]
GO
ALTER TABLE [dbo].[leave_defer] ADD  DEFAULT ((0)) FOR [user_id]
GO
ALTER TABLE [dbo].[leave_defer] ADD  DEFAULT ((0)) FOR [leave_core_id]
GO
ALTER TABLE [dbo].[leave_defer] ADD  DEFAULT ((0)) FOR [previous_id]
GO
ALTER TABLE [dbo].[leave_defer] ADD  DEFAULT ((0)) FOR [next_id]
GO
ALTER TABLE [dbo].[leave_request_form] ADD  DEFAULT ((0)) FOR [sign_status]
GO
ALTER TABLE [dbo].[leave_request_form] ADD  DEFAULT ((0)) FOR [leave_core_id]
GO
ALTER TABLE [dbo].[leave_request_form] ADD  DEFAULT ((0)) FOR [user_id]
GO
ALTER TABLE [dbo].[leave_request_form] ADD  DEFAULT ((0)) FOR [department_id]
GO
ALTER TABLE [dbo].[leave_request_form] ADD  DEFAULT ((0)) FOR [proxy_user_id]
GO
ALTER TABLE [dbo].[leave_request_form] ADD  DEFAULT ((0)) FOR [proxy_department_id]
GO
ALTER TABLE [dbo].[menus] ADD  DEFAULT ((0)) FOR [parent_id]
GO
ALTER TABLE [dbo].[menus] ADD  DEFAULT ((0)) FOR [sort]
GO
ALTER TABLE [dbo].[menus] ADD  DEFAULT ((1)) FOR [status]
GO
ALTER TABLE [dbo].[menus] ADD  DEFAULT ((0)) FOR [last_mover]
GO
ALTER TABLE [dbo].[menus] ADD  DEFAULT ((0)) FOR [show]
GO
ALTER TABLE [dbo].[overtime_rate] ADD  DEFAULT ((0)) FOR [last_mover]
GO
ALTER TABLE [dbo].[overtime_request_form] ADD  DEFAULT ((0)) FOR [sign_status]
GO
ALTER TABLE [dbo].[overtime_request_form] ADD  DEFAULT ((0)) FOR [user_id]
GO
ALTER TABLE [dbo].[overtime_request_form] ADD  DEFAULT ((0)) FOR [department_id]
GO
ALTER TABLE [dbo].[permissions] ADD  DEFAULT ((1)) FOR [status]
GO
ALTER TABLE [dbo].[permissions] ADD  DEFAULT ((0)) FOR [last_mover]
GO
ALTER TABLE [dbo].[ranks] ADD  DEFAULT ((1)) FOR [status]
GO
ALTER TABLE [dbo].[ranks] ADD  DEFAULT ((0)) FOR [last_mover]
GO
ALTER TABLE [dbo].[roles] ADD  DEFAULT ((1)) FOR [status]
GO
ALTER TABLE [dbo].[roles] ADD  DEFAULT ((0)) FOR [last_mover]
GO
ALTER TABLE [dbo].[users] ADD  DEFAULT ((1)) FOR [status]
GO
ALTER TABLE [dbo].[users] ADD  DEFAULT ((0)) FOR [rank_id]
GO
ALTER TABLE [dbo].[users] ADD  DEFAULT ((0)) FOR [grade_id]
GO
ALTER TABLE [dbo].[users] ADD  DEFAULT ((0)) FOR [last_mover]
GO
ALTER TABLE [dbo].[users] ADD  DEFAULT ((0)) FOR [backend]
GO
ALTER TABLE [dbo].[vacation_types] ADD  DEFAULT ((1)) FOR [status]
GO
ALTER TABLE [dbo].[vacation_types] ADD  DEFAULT ((0)) FOR [last_mover]
GO
ALTER TABLE [dbo].[vacations] ADD  DEFAULT ((0)) FOR [type_id]
GO
ALTER TABLE [dbo].[vacations] ADD  DEFAULT ((0)) FOR [year]
GO
ALTER TABLE [dbo].[vacations] ADD  DEFAULT ((0)) FOR [month]
GO
ALTER TABLE [dbo].[vacations] ADD  DEFAULT ((0)) FOR [day]
GO
ALTER TABLE [dbo].[vacations] ADD  DEFAULT ((0)) FOR [last_mover]
GO
ALTER TABLE [dbo].[work_schedule] ADD  DEFAULT ((0)) FOR [year]
GO
ALTER TABLE [dbo].[work_schedule] ADD  DEFAULT ((0)) FOR [month]
GO
ALTER TABLE [dbo].[work_schedule] ADD  DEFAULT ((0)) FOR [day]
GO
ALTER TABLE [dbo].[work_schedule] ADD  DEFAULT ((0)) FOR [user_id]
GO
ALTER TABLE [dbo].[work_schedule] ADD  DEFAULT ((0)) FOR [work_shift_id]
GO
ALTER TABLE [dbo].[work_shift] ADD  DEFAULT ((1)) FOR [status]
GO
ALTER TABLE [dbo].[work_shift] ADD  DEFAULT ((0)) FOR [last_mover]
GO
