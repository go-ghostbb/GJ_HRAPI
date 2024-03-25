-- 更新上班
update check_in_status
set work_attend_time = ?, updated_at = getdate()
where
    work_shift_id = (select id from work_shift where code = ?) and
    employee_id = (select id from employee where card_number = ?) and
    work_check_in_date = ?;

-- 更新下班
update check_in_status
set off_work_attend_time = ?, updated_at = getdate()
where
    work_shift_id = (select id from work_shift where code = ?) and
    employee_id = (select id from employee where card_number = ?) and
    off_work_check_in_date = ?;