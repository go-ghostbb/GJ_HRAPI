-- 寫進gorm gen裡, 需把註釋刪掉
update check_in_status
set
    -- 上班
    work_attend_status = (
        select
            case
                -- 沒刷卡
                when check_in_status.work_attend_time = '00:00:00' then 'not swiped'
                -- 正常(沒遲到)
                when work_start >= convert(varchar(5), check_in_status.work_attend_time, 108) then 'normal'
                -- 遲到
                when work_start < convert(varchar(5), check_in_status.work_attend_time, 108) then 'late'
                end
        from work_shift
        where id = work_shift_id
    ),
    work_attend_proc_status = (
        select
            case
                -- 沒刷卡
                when check_in_status.work_attend_time = '00:00:00' then 'not processed'
                -- 正常(不需處理)
                when work_start >= convert(varchar(5), check_in_status.work_attend_time, 108) then 'normal'
                -- 未處理
                when work_start < convert(varchar(5), check_in_status.work_attend_time, 108) then 'not processed'
                end
        from work_shift
        where id = work_shift_id
    ),
    -- 下班
    off_work_attend_status = (
        select
            case
                -- 沒刷卡
                when check_in_status.off_work_attend_time = '00:00:00' then 'not swiped'
                -- 正常(沒早退)
                when work_end <= check_in_status.off_work_attend_time then 'normal'
                -- 早退
                when work_end > check_in_status.off_work_attend_time then 'early'
                end
        from work_shift
        where id = work_shift_id
    ),
    off_work_attend_proc_status = (
        select
            case
                -- 沒刷卡
                when check_in_status.off_work_attend_time = '00:00:00' then 'not processed'
                -- 正常(沒早退)
                when work_end <= check_in_status.off_work_attend_time then 'normal'
                -- 早退
                when work_end > check_in_status.off_work_attend_time then 'not processed'
                end
        from work_shift
        where id = work_shift_id
    ),
    -- 缺勤時數
    absence_hours = (
        select
            iif(check_in_status.work_attend_time != '00:00:00' and check_in_status.off_work_attend_time != '00:00:00',
                -- 上下班如果都有刷卡
                iif(work_start < convert(varchar(5), check_in_status.work_attend_time, 108),
                    -- 如果遲到, 計算遲到多久(單位：小時)
                    convert(float, datediff(second, work_start, work_attend_time)) / 60 / 60,
                    -- 否則0, 代表缺勤時數0
                    0)
                    + -- 上班遲到 + 下班早退
                iif(work_end > check_in_status.off_work_attend_time,
                    -- 如果遲到, 計算遲到多久(單位：小時)
                    convert(float, datediff(second , off_work_attend_time, '17:00:00')) / 60 / 60,
                    -- 否則0, 代表缺勤時數0
                    0)
                , absence_hours)
        from work_shift
        where id = work_shift_id
    )
where work_check_in_date in ('2024-01-02');