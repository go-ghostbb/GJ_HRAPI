DROP FUNCTION IF EXISTS [dbo].[FN_C_SeniorityMonth]
GO

CREATE FUNCTION [dbo].[FN_C_SeniorityMonth] (@startDate date, @endDate date)
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
go

