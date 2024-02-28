DROP FUNCTION IF EXISTS [dbo].[FN_C_Seniority]
GO

CREATE FUNCTION [dbo].[FN_C_Seniority] (@startDate SMALLDATETIME, @endDate SMALLDATETIME)
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

