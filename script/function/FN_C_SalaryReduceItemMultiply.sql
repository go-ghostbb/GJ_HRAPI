DROP FUNCTION IF EXISTS [dbo].[FN_C_SalaryReduceItemMultiply]
GO

create function [dbo].[FN_C_SalaryReduceItemMultiply] (@employee_ids varchar(max), @startDate date, @endDate date)
    returns @result table
                    (
                        employee_id           bigint,
                        salary_reduce_item_id bigint,
                        income_tax            bit,
                        amount                float(53),
                        multiply              float(53)
                    )
as
begin
    insert into @result (employee_id, salary_reduce_item_id, income_tax, amount, multiply)
    select
        e.id as 'employee_id', s.id, s.income_tax, s.amount,
        case
            when s.calc_type = 'default' then
                -- 在計算條件為預設金額時
                iif((select f.month
                     from employee i_e
                              join FN_C_EmployeeSeniorityWithEndDate(@endDate) f on (i_e.id = f.employee_id)
                     where id = e.id) > 0 or s.month_calc = 0 or s.amount = 0,
                    -- 如果年資滿一個月或是沒開啟按比例計算或是金額為0(金額為0代表使用者自行輸入金額), 直接回傳倍率1
                    1,
                    -- 否則按比例計算
                    (convert(float, datediff(day, e.hire_date, @endDate)) + 1) / (convert(float, datediff(day, @startDate, @endDate)) + 1)
                )
            when s.unit = 'hour' then
                -- 在計算條件不為預設金額且計算單位為小時
                (
                    case s.calc_type
                        when 'condition' then
                            (
                                -- 如果是條件帶入且條件為小時
                                case s.operator
                                    when '>=' then
                                        (
                                            select count(*)
                                            from check_in_status c
                                                     join work_shift w on (c.work_shift_id = w.id)
                                            where
                                                c.work_shift_id != 0 and
                                                work_check_in_date between @startDate and @endDate and
                                                c.employee_id = e.id and
                                                w.total_hours - c.absence_hours - c.leave_hours >= s.argument
                                        )
                                    when '<=' then
                                        (
                                            select count(*)
                                            from check_in_status c
                                                     join work_shift w on (c.work_shift_id = w.id)
                                            where
                                                c.work_shift_id != 0 and
                                                work_check_in_date between @startDate and @endDate and
                                                c.employee_id = e.id and
                                                w.total_hours - c.absence_hours - c.leave_hours <= s.argument
                                        )
                                    end
                                )
                        when 'unit' then
                            (
                                -- 如果是單位計算且是小時
                                select floor(sum(w.total_hours) - sum(c.absence_hours) - sum(c.leave_hours))
                                from check_in_status c
                                         join work_shift w on (c.work_shift_id = w.id)
                                         join employee e on (c.employee_id = e.id)
                                where
                                    c.work_shift_id != 0 and
                                    work_check_in_date between @startDate and @endDate and
                                    c.employee_id = e.id and
                                    e.salary_cycle = 'hour'
                            )
                        end
                    )
            when s.unit = 'day' then
                (
                    case s.calc_type
                        when 'condition' then
                            (
                                -- 如果是條件帶入且條件為天
                                case s.operator
                                    when '>=' then
                                        (
                                            select iif(count(*) >= s.argument, 1, 0)
                                            from check_in_status c
                                            where
                                                c.work_shift_id != 0 and
                                                work_check_in_date between @startDate and @endDate and
                                                c.employee_id = e.id and
                                                c.absence_hours = 0
                                        )
                                    when '<=' then
                                        (
                                            select iif(count(*) <= s.argument, 1, 0)
                                            from check_in_status c
                                            where
                                                c.work_shift_id != 0 and
                                                work_check_in_date between @startDate and @endDate and
                                                c.employee_id = e.id and
                                                c.absence_hours = 0
                                        )
                                    end
                                )
                        when 'unit' then
                            (
                                select count(*)
                                from check_in_status c
                                where
                                    c.work_shift_id != 0 and
                                    work_check_in_date between @startDate and @endDate and
                                    c.employee_id = e.id and
                                    c.absence_hours = 0
                            )
                        end
                    )
            end as 'multiply'
    from salary_reduce_item s
             join salary_reduce_item_employee se on (se.salary_reduce_item_id = s.id)
             join employee e on (se.employee_id = e.id)
    where e.id in (select Value from string_split(@employee_ids, ',')) and s.deleted_at is null

    -- 回傳
    return
end
go