drop function if exists [dbo].[OwnForm]
go

create function [dbo].[OwnForm] (@emp_id bigint, @keyword varchar(max), @offset int, @limit int)
    returns @result table (
        id bigint,
        created_at datetimeoffset,
        [order] nvarchar(50),
        sign_status tinyint,
        remark nvarchar(max)
    )
begin
    ;with form as (
        select id, created_at, [order], sign_status, remark from leave_request_form
            where employee_id = @emp_id and [order] like '%' + @keyword + '%' and deleted_at is null
        union all
        select id, created_at, [order], sign_status, remark from overtime_request_form
            where employee_id = @emp_id and [order] like '%' + @keyword + '%' and deleted_at is null
        union all
        select id, created_at, [order], sign_status, remark from check_in_request_form
            where employee_id = @emp_id and [order] like '%' + @keyword + '%' and deleted_at is null
     )
     insert into @result (id, created_at, [order], sign_status, remark)
     select *
     from form
     order by created_at desc
     offset @offset rows
     fetch next @limit rows only;

    return;
end