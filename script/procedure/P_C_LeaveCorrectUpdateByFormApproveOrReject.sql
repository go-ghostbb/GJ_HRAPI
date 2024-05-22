drop procedure if exists [dbo].[P_C_LeaveCorrectUpdateByFormApproveOrReject]
go

create procedure [dbo].[P_C_LeaveCorrectUpdateByFormApproveOrReject] (@leave_request_form_id bigint)
as
begin
    -- 查詢表單基本訊息
    declare @temp_sign_status tinyint,
            @temp_deleted_at datetimeoffset;

    select @temp_sign_status = sign_status,
           @temp_deleted_at = deleted_at
    from leave_request_form
    where id = @leave_request_form_id;

    -- 遍歷對照表
    declare myCursor cursor for
    select leave_correct_id, day from leave_request_form_correct_map where leave_request_form_id = @leave_request_form_id;

    -- cursor暫存
    declare @cur_leave_correct_id bigint,
            @cur_day float;

    open myCursor;
    fetch next from myCursor into @cur_leave_correct_id, @cur_day;
    while (@@fetch_status <> -1)
    begin
        -- approve
        if (@temp_sign_status = 1 and @temp_deleted_at is null)
        begin
            update leave_correct
            set signing = signing - @cur_day,
                used = used + @cur_day
            where id = @cur_leave_correct_id;
        end;

        -- reject or deleted but not approve
        if (@temp_sign_status = 2 or (@temp_deleted_at is not null and @temp_sign_status != 1))
        begin
            update leave_correct
            set signing = signing - @cur_day
            where id = @cur_leave_correct_id;
        end;

        -- deleted and approve
        if (@temp_deleted_at is not null and @temp_sign_status = 1)
        begin
            update leave_correct
            set used = used - @cur_day
            where id = @cur_leave_correct_id;
        end;

        -- step
        fetch next from myCursor into @cur_leave_correct_id, @cur_day;
    end
    close myCursor;
    deallocate myCursor;
end