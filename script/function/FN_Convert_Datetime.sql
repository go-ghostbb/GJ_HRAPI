drop function if exists [dbo].[FN_Convert_Datetime]
go

create function [dbo].[FN_Convert_Datetime] (@date date, @time time(0))
    returns datetime
as
begin
    declare @result datetime = convert(datetime, convert(varchar, @date) + ' ' + convert(varchar, @time))
    return @result
end