-- employee
set identity_insert employee on
insert into employee (id, created_at, updated_at, deleted_at, hire_date, termination_date, employment_status, backend, real_name, national_id, birth, email, mobile, department_id, rank_id, grade_id, code, salary, salary_cycle, card_number)
values (1, sysdatetimeoffset(), sysdatetimeoffset(), null, getdate(), null, 'active', 1, 'admin', 'admin', getdate(), N'project@mail.gj.com.tw', '0900026201', 1, 0, 0, 'T10000', 45000, 'month', '0000000000')
set identity_insert employee off
go

-- role
set identity_insert role on
insert into role (id, created_at, updated_at, deleted_at, code, name, remark)
values (1, sysdatetimeoffset(), sysdatetimeoffset(), null, 'root', N'超級管理員', '')
set identity_insert role off
go

-- employee_role
insert into employee_role (role_id, employee_id)
values (1, 1)
go

-- menu
set identity_insert menu on
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (1, N'2024-02-28 20:38:28.3520000 +00:00', N'2024-03-02 13:46:02.0380040 +08:00', null, N'menu', N'software', 0, N'/home', N'HomePage', N'/home/index', N'', 1, 1, N'首頁', 0, N'', 0, 1, N'ant-design:home-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (2, N'2024-02-28 20:38:28.3520000 +00:00', N'2024-02-28 20:38:29.1390000 +00:00', null, N'menu', N'web', 0, N'/home', N'HomePage', N'/home/index', N'', 1, 1, N'首頁', 0, N'', 0, 1, N'ant-design:home-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (10002, N'2024-02-29 10:01:41.6850000 +00:00', N'2024-05-08 14:00:28.4034140 +08:00', null, N'directory', N'software', 0, N'/manager', N'Manager', N'LAYOUT', N'', 3, 1, N'系統管理', 0, N'', 0, 0, N'ant-design:setting-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (10004, N'2024-02-29 10:04:34.5890000 +00:00', N'2024-03-04 10:43:02.9124390 +08:00', null, N'menu', N'software', 10002, N'menu', N'MenuManager', N'/manager/menu/index', N'', 4, 1, N'菜單管理', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (20004, N'2024-03-02 15:27:13.4557760 +08:00', N'2024-03-03 18:07:20.0579810 +08:00', null, N'menu', N'software', 10002, N'department', N'DepartmentManager', N'/manager/department/index', N'', 2, 1, N'部門管理', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (20005, N'2024-03-03 18:07:02.4667520 +08:00', N'2024-03-04 10:42:53.9287780 +08:00', null, N'menu', N'software', 10002, N'role', N'RoleManager', N'/manager/role/index', N'', 3, 1, N'角色管理', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (30005, N'2024-03-04 10:42:32.8527320 +08:00', N'2024-03-04 10:42:41.4903070 +08:00', null, N'menu', N'software', 10002, N'perm', N'PermManager', N'/manager/permission/index', N'', 5, 1, N'權限標示管理', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (30006, N'2024-03-05 14:42:43.2246400 +08:00', N'2024-03-05 14:42:43.2246400 +08:00', null, N'menu', N'software', 10002, N'employee', N'EmployeeManager', N'/manager/employee/index', N'', 1, 1, N'員工管理', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (40006, N'2024-03-06 10:42:22.0415420 +08:00', N'2024-03-06 10:42:22.0415420 +08:00', null, N'menu', N'software', 10002, N'employee/:id', N'EmployeeDetail', N'/manager/employee/detail/index', N'', 0, 1, N'詳情', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'/manager/employee', 0, 1, 0, 0);
set identity_insert menu off
go

set identity_insert menu on
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (40007, N'2024-03-06 15:16:52.4211420 +08:00', N'2024-05-08 14:00:33.8275880 +08:00', null, N'directory', N'software', 0, N'/setting', N'Setting', N'LAYOUT', N'', 4, 1, N'設定', 0, N'', 0, 0, N'ant-design:apartment-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (40008, N'2024-03-06 15:17:22.7663920 +08:00', N'2024-03-06 15:17:22.7663920 +08:00', null, N'menu', N'software', 40007, N'Leave', N'LeaveSetting', N'/setting/leave/index', N'', 1, 1, N'假別設定', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (40009, N'2024-03-08 15:33:47.9727410 +08:00', N'2024-03-14 13:06:54.7538940 +08:00', null, N'directory', N'software', 40007, N'vacation', N'Vacation', N'', N'', 3, 1, N'休假日設定', 0, N'', 0, 0, N'ant-design:book-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (40010, N'2024-03-08 15:34:31.0037630 +08:00', N'2024-05-02 13:21:16.8888270 +08:00', null, N'menu', N'software', 40009, N'type', N'VacationSetting', N'/setting/vacation/index', N'', 1, 1, N'休假類別設定', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (40011, N'2024-03-11 13:15:17.3552510 +08:00', N'2024-03-11 13:15:17.3552510 +08:00', null, N'menu', N'software', 40009, N'vacationSchedule', N'VacationScheduleSetting', N'/setting/vacation-schedule/index', N'', 2, 1, N'休假設定', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (40012, N'2024-03-11 22:39:29.9018760 +08:00', N'2024-03-14 13:06:59.1323750 +08:00', null, N'directory', N'software', 40007, N'work', N'Work', N'', N'', 4, 1, N'排班設定', 0, N'', 0, 0, N'ant-design:calendar-twotone', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (40013, N'2024-03-11 22:40:19.9886440 +08:00', N'2024-05-02 13:21:46.5811170 +08:00', null, N'menu', N'software', 40012, N'workShift', N'WorkShiftSetting', N'/setting/workShift/index', N'', 1, 1, N'班別類別設定', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (40014, N'2024-03-12 11:55:54.0839190 +08:00', N'2024-03-12 11:55:54.0839190 +08:00', null, N'menu', N'software', 40012, N'workSchedule', N'WorkScheduleSetting', N'/setting/work-schedule/index', N'', 2, 1, N'班表設定', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (40015, N'2024-03-14 11:47:27.4511380 +08:00', N'2024-03-14 13:42:59.8335340 +08:00', null, N'menu', N'software', 40007, N'rank', N'PositionRankSetting', N'/setting/rank/index', N'', 2, 1, N'職級/職等設定', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
set identity_insert menu off
go

set identity_insert menu on
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (50015, N'2024-03-19 11:52:34.6797510 +08:00', N'2024-03-19 11:52:34.6797510 +08:00', null, N'directory', N'software', 40007, N'sign-off', N'SignOff', N'', N'', 5, 1, N'簽核設定', 0, N'', 0, 0, N'ant-design:apartment-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (50016, N'2024-03-19 11:53:54.4610800 +08:00', N'2024-03-19 11:54:32.7536660 +08:00', null, N'menu', N'software', 50015, N'leave', N'SignOffLeaveSetting', N'/setting/sign-off-leave/index', N'', 1, 1, N'請假簽核設定', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (50017, N'2024-03-20 10:18:25.6881220 +08:00', N'2024-03-20 10:18:25.6881220 +08:00', null, N'menu', N'software', 50015, N'overtime', N'SignOffOvertimeSetting', N'/setting/sign-off-overtime/index', N'', 2, 1, N'加班簽核設定', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (50018, N'2024-03-20 13:10:59.8849110 +08:00', N'2024-03-20 13:13:33.1231180 +08:00', null, N'directory', N'software', 40007, N'salary', N'SalarySetting', N'', N'', 6, 1, N'薪資設定', 0, N'', 0, 0, N'ant-design:credit-card-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (50019, N'2024-03-20 13:16:07.1531340 +08:00', N'2024-03-20 13:16:59.3260890 +08:00', null, N'menu', N'software', 50018, N'add', N'SalaryAddSetting', N'/setting/salary/add/index', N'', 1, 1, N'加項設定', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (50020, N'2024-03-20 13:16:51.8270170 +08:00', N'2024-03-20 13:16:51.8270170 +08:00', null, N'menu', N'software', 50018, N'reduce', N'SalaryReduceSetting', N'/setting/salary/reduce/index', N'', 2, 1, N'減項設定', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (50021, N'2024-03-21 13:42:05.5749000 +08:00', N'2024-05-08 14:01:00.5135020 +08:00', null, N'menu', N'software', 0, N'/salaryWork', N'SalaryWork', N'/salary-work/index', N'', 5, 1, N'薪資作業', 0, N'', 0, 0, N'ant-design:fund-projection-screen-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (50022, N'2024-03-25 15:16:14.3536450 +08:00', N'2024-03-28 14:06:48.7162820 +08:00', null, N'directory', N'software', 0, N'/upload', N'Upload', N'LAYOUT', N'', 6, 1, N'上傳', 0, N'', 0, 0, N'ant-design:cloud-upload-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (50023, N'2024-03-25 15:19:23.1277640 +08:00', N'2024-03-25 15:19:38.6446870 +08:00', null, N'menu', N'software', 50022, N'checkInData', N'CheckInDataUpload', N'/upload/check-in-data/index', N'', 1, 1, N'打卡資料上傳', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
set identity_insert menu off
go


set identity_insert menu on
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (60022, N'2024-03-28 14:06:41.0368170 +08:00', N'2024-03-28 14:08:29.5413550 +08:00', null, N'menu', N'software', 0, N'/salaryWork/setting/:id', N'SalaryWorkSetting', N'/salary-work/stage-view/index', N'', 5, 1, N'薪資作業設定', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'/salaryWork', 0, 1, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (60023, N'2024-04-08 09:25:15.5056940 +08:00', N'2024-05-06 09:52:59.7890600 +08:00', null, N'directory', N'web', 0, N'/daily', N'Daily', N'LAYOUT', N'', 3, 1, N'日常作業', 0, N'', 0, 0, N'ant-design:account-book-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (60024, N'2024-04-08 09:26:23.8627760 +08:00', N'2024-04-08 09:26:23.8627760 +08:00', null, N'menu', N'web', 60023, N'leave', N'LeaveDaily', N'/daily/leave/index', N'', 1, 1, N'請假單', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (60025, N'2024-04-12 10:13:40.0307050 +08:00', N'2024-04-12 10:13:40.0307050 +08:00', null, N'menu', N'software', 50015, N'checkIn', N'SignOffCheckInSetting', N'/setting/sign-off-check-in/index', N'', 3, 1, N'補打卡簽核設定', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (60026, N'2024-04-17 09:01:46.6807070 +08:00', N'2024-04-30 15:32:58.4503070 +08:00', null, N'menu', N'web', 60023, N'checkIn', N'CheckInDaily', N'/daily/check-in/index', N'', 3, 1, N'補打卡單', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (60027, N'2024-04-30 15:32:53.2732950 +08:00', N'2024-04-30 15:32:53.2732950 +08:00', null, N'menu', N'web', 60023, N'overtime', N'OvertimeDaily', N'/daily/overtime/index', N'', 2, 1, N'加班單', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (60028, N'2024-05-06 09:52:54.9617710 +08:00', N'2024-05-06 09:52:54.9617710 +08:00', null, N'directory', N'web', 0, N'/query', N'Query', N'LAYOUT', N'', 2, 1, N'查詢', 0, N'', 0, 0, N'ant-design:search-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (60029, N'2024-05-06 09:53:52.3033280 +08:00', N'2024-05-06 09:53:52.3033280 +08:00', null, N'menu', N'web', 60028, N'checkIn', N'CheckInQuery', N'/query/check-in/index', N'', 1, 1, N'出缺勤查詢', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (60030, N'2024-05-08 14:00:17.8192810 +08:00', N'2024-05-31 10:04:39.7042680 +08:00', null, N'menu', N'software', 60032, N'checkIn', N'CheckInStatusQuery', N'/query/check_in/index', N'', 2, 1, N'出缺勤查詢', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
set identity_insert menu off
go

set identity_insert menu on
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (60031, N'2024-05-16 14:19:48.5411050 +08:00', N'2024-05-16 14:19:48.5411050 +08:00', null, N'menu', N'web', 60028, N'queryCorrect', N'LeaveCorrectQuery', N'/query/leave/index', N'', 2, 1, N'假別查詢', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (60032, N'2024-05-31 10:02:39.8931410 +08:00', N'2024-05-31 10:02:39.8931410 +08:00', null, N'directory', N'software', 0, N'/query', N'Query', N'LAYOUT', N'', 2, 1, N'查詢', 0, N'', 0, 0, N'ant-design:search-outlined', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
insert into menu (id, created_at, updated_at, deleted_at, type, show, parent_id, path, name, component, redirect, sort, status, title, dynamic_level, real_path, ignore_keep_alive, affix, icon, frame_src, transition_name, hide_breadcrumb, carry_param, hide_children_in_menu, current_active_menu, hide_tab, hide_menu, ignore_route, hide_path_for_children) VALUES (60033, N'2024-05-31 10:04:31.7857960 +08:00', N'2024-05-31 10:04:31.7857960 +08:00', null, N'menu', N'software', 60032, N'leaveCorrect', N'LeaveCorrectQuery', N'/query/leave/index', N'', 2, 1, N'假別查詢', 0, N'', 0, 0, N'', N'', N'', 0, 0, 0, N'', 0, 0, 0, 0);
set identity_insert menu off
go

-- role_menu
insert into role_menu (menu_id, role_id) VALUES (1, 1);
insert into role_menu (menu_id, role_id) VALUES (2, 1);
insert into role_menu (menu_id, role_id) VALUES (10002, 1);
insert into role_menu (menu_id, role_id) VALUES (10004, 1);
insert into role_menu (menu_id, role_id) VALUES (20004, 1);
insert into role_menu (menu_id, role_id) VALUES (20005, 1);
insert into role_menu (menu_id, role_id) VALUES (30005, 1);
insert into role_menu (menu_id, role_id) VALUES (30006, 1);
insert into role_menu (menu_id, role_id) VALUES (40006, 1);
insert into role_menu (menu_id, role_id) VALUES (40007, 1);
insert into role_menu (menu_id, role_id) VALUES (40008, 1);
insert into role_menu (menu_id, role_id) VALUES (40009, 1);
insert into role_menu (menu_id, role_id) VALUES (40010, 1);
insert into role_menu (menu_id, role_id) VALUES (40011, 1);
insert into role_menu (menu_id, role_id) VALUES (40012, 1);
insert into role_menu (menu_id, role_id) VALUES (40013, 1);
insert into role_menu (menu_id, role_id) VALUES (40014, 1);
insert into role_menu (menu_id, role_id) VALUES (40015, 1);
insert into role_menu (menu_id, role_id) VALUES (50015, 1);
insert into role_menu (menu_id, role_id) VALUES (50016, 1);
insert into role_menu (menu_id, role_id) VALUES (50017, 1);
insert into role_menu (menu_id, role_id) VALUES (50018, 1);
insert into role_menu (menu_id, role_id) VALUES (50019, 1);
insert into role_menu (menu_id, role_id) VALUES (50020, 1);
insert into role_menu (menu_id, role_id) VALUES (50021, 1);
insert into role_menu (menu_id, role_id) VALUES (50022, 1);
insert into role_menu (menu_id, role_id) VALUES (50023, 1);
insert into role_menu (menu_id, role_id) VALUES (60022, 1);
insert into role_menu (menu_id, role_id) VALUES (60023, 1);
insert into role_menu (menu_id, role_id) VALUES (60024, 1);
insert into role_menu (menu_id, role_id) VALUES (60025, 1);
insert into role_menu (menu_id, role_id) VALUES (60026, 1);
insert into role_menu (menu_id, role_id) VALUES (60027, 1);
insert into role_menu (menu_id, role_id) VALUES (60028, 1);
insert into role_menu (menu_id, role_id) VALUES (60029, 1);
insert into role_menu (menu_id, role_id) VALUES (60030, 1);
insert into role_menu (menu_id, role_id) VALUES (60031, 1);
insert into role_menu (menu_id, role_id) VALUES (60032, 1);
insert into role_menu (menu_id, role_id) VALUES (60033, 1);
go

-- department
set identity_insert department on
insert into department (id, created_at, updated_at, deleted_at, code, name, remark)
values (1, sysdatetimeoffset(), sysdatetimeoffset(), null, 'Test', N'測試部門', 'test')
set identity_insert department off
go

-- permission
set identity_insert permission on
insert into permission (id, created_at, updated_at, deleted_at, perm, name, status, remark) VALUES (5, N'2024-03-04 15:57:16.2171740 +08:00', N'2024-03-04 15:57:16.2171740 +08:00', null, N'permission:create', N'權限標示新增', 1, N'');
insert into permission (id, created_at, updated_at, deleted_at, perm, name, status, remark) VALUES (6, N'2024-03-04 16:01:06.8393040 +08:00', N'2024-03-04 16:01:06.8393040 +08:00', null, N'permission:edit', N'權限標示編輯', 1, N'');
insert into permission (id, created_at, updated_at, deleted_at, perm, name, status, remark) VALUES (7, N'2024-03-04 16:01:30.7015620 +08:00', N'2024-03-04 16:01:30.7015620 +08:00', null, N'permission:delete', N'權限標示刪除', 1, N'');
insert into permission (id, created_at, updated_at, deleted_at, perm, name, status, remark) VALUES (8, N'2024-03-04 16:07:07.8565630 +08:00', N'2024-03-04 16:07:07.8565630 +08:00', null, N'role:create', N'角色新增', 1, N'');
insert into permission (id, created_at, updated_at, deleted_at, perm, name, status, remark) VALUES (9, N'2024-03-04 16:07:25.9724410 +08:00', N'2024-03-04 16:07:25.9724410 +08:00', null, N'role:edit', N'角色編輯', 1, N'');
insert into permission (id, created_at, updated_at, deleted_at, perm, name, status, remark) VALUES (10, N'2024-03-04 16:07:39.8983290 +08:00', N'2024-03-04 16:07:39.8983290 +08:00', null, N'role:delete', N'角色刪除', 1, N'');
insert into permission (id, created_at, updated_at, deleted_at, perm, name, status, remark) VALUES (11, N'2024-03-04 16:08:34.9813580 +08:00', N'2024-03-04 16:08:34.9813580 +08:00', null, N'role:set', N'角色設定權限', 1, N'');
insert into permission (id, created_at, updated_at, deleted_at, perm, name, status, remark) VALUES (12, N'2024-03-04 16:18:32.1008250 +08:00', N'2024-03-04 16:18:32.1008250 +08:00', null, N'menu:create', N'菜單新增', 1, N'');
insert into permission (id, created_at, updated_at, deleted_at, perm, name, status, remark) VALUES (13, N'2024-03-04 16:18:45.9068210 +08:00', N'2024-03-04 16:18:45.9068210 +08:00', null, N'menu:edit', N'菜單編輯', 1, N'');
set identity_insert permission off
go

set identity_insert permission on
insert into permission (id, created_at, updated_at, deleted_at, perm, name, status, remark) VALUES (14, N'2024-03-04 16:19:02.3259600 +08:00', N'2024-03-04 16:19:02.3259600 +08:00', null, N'menu:delete', N'菜單刪除', 1, N'');
insert into permission (id, created_at, updated_at, deleted_at, perm, name, status, remark) VALUES (15, N'2024-03-04 16:23:29.5202530 +08:00', N'2024-03-04 16:23:29.5202530 +08:00', null, N'department:create', N'部門新增', 1, N'');
insert into permission (id, created_at, updated_at, deleted_at, perm, name, status, remark) VALUES (16, N'2024-03-04 16:23:44.8699720 +08:00', N'2024-03-04 16:23:44.8699720 +08:00', null, N'department:edit', N'部門編輯', 1, N'');
insert into permission (id, created_at, updated_at, deleted_at, perm, name, status, remark) VALUES (17, N'2024-03-04 16:23:57.5918910 +08:00', N'2024-03-04 16:23:57.5918910 +08:00', null, N'department:delete', N'部門刪除', 1, N'');
insert into permission (id, created_at, updated_at, deleted_at, perm, name, status, remark) VALUES (18, N'2024-03-04 16:39:51.7749010 +08:00', N'2024-03-04 16:41:35.3169650 +08:00', null, N'role:update:status', N'角色設定狀態', 1, N'');
insert into permission (id, created_at, updated_at, deleted_at, perm, name, status, remark) VALUES (19, N'2024-03-04 16:43:14.7264510 +08:00', N'2024-03-04 16:43:14.7264510 +08:00', null, N'permission:update:status', N'權限標示設定狀態', 1, N'');
insert into permission (id, created_at, updated_at, deleted_at, perm, name, status, remark) VALUES (20, N'2024-03-04 16:43:45.9217090 +08:00', N'2024-03-04 16:43:45.9217090 +08:00', null, N'menu:update:status', N'菜單設定狀態', 1, N'');
insert into permission (id, created_at, updated_at, deleted_at, perm, name, status, remark) VALUES (21, N'2024-03-04 16:44:09.2712170 +08:00', N'2024-03-04 16:44:09.2712170 +08:00', null, N'department:update:status', N'部門設定狀態', 1, N'');
insert into permission (id, created_at, updated_at, deleted_at, perm, name, status, remark) VALUES (22, N'2024-03-05 14:45:17.8466130 +08:00', N'2024-03-05 14:45:17.8466130 +08:00', null, N'employee:create', N'員工新增', 1, N'');
set identity_insert permission off
go

set identity_insert permission on
insert into permission (id, created_at, updated_at, deleted_at, perm, name, status, remark) VALUES (23, N'2024-03-05 14:45:30.1598880 +08:00', N'2024-03-05 14:45:30.1598880 +08:00', null, N'employee:edit', N'員工編輯', 1, N'');
insert into permission (id, created_at, updated_at, deleted_at, perm, name, status, remark) VALUES (24, N'2024-03-05 14:45:45.2502730 +08:00', N'2024-03-05 14:45:45.2502730 +08:00', null, N'employee:delete', N'員工刪除', 1, N'');
set identity_insert permission off
go

-- role_permission
insert into role_permission (permission_id, role_id) VALUES (5, 1);
insert into role_permission (permission_id, role_id) VALUES (6, 1);
insert into role_permission (permission_id, role_id) VALUES (7, 1);
insert into role_permission (permission_id, role_id) VALUES (8, 1);
insert into role_permission (permission_id, role_id) VALUES (9, 1);
insert into role_permission (permission_id, role_id) VALUES (10, 1);
insert into role_permission (permission_id, role_id) VALUES (11, 1);
insert into role_permission (permission_id, role_id) VALUES (12, 1);
insert into role_permission (permission_id, role_id) VALUES (13, 1);
insert into role_permission (permission_id, role_id) VALUES (14, 1);
insert into role_permission (permission_id, role_id) VALUES (15, 1);
insert into role_permission (permission_id, role_id) VALUES (16, 1);
insert into role_permission (permission_id, role_id) VALUES (17, 1);
insert into role_permission (permission_id, role_id) VALUES (18, 1);
insert into role_permission (permission_id, role_id) VALUES (19, 1);
insert into role_permission (permission_id, role_id) VALUES (20, 1);
insert into role_permission (permission_id, role_id) VALUES (21, 1);
insert into role_permission (permission_id, role_id) VALUES (22, 1);
insert into role_permission (permission_id, role_id) VALUES (23, 1);
insert into role_permission (permission_id, role_id) VALUES (24, 1);
go

