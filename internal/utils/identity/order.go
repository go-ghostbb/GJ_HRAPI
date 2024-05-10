package identity

import (
	"context"
	"fmt"
	gberror "ghostbb.io/gb/errors/gb_error"
	"ghostbb.io/gb/frame/g"
	"hrapi/internal/query"
	"time"
)

var (
	leaveOrderPrefix    string
	overtimeOrderPrefix string
	checkInOrderPrefix  string
)

func formOrderInit() error {
	var (
		qConfig = query.ConfigMap
		err     error
	)
	// 請假
	err = qConfig.WithContext(context.Background()).Select(qConfig.Value).Where(qConfig.Key.Eq("LeaveOrderPrefix")).Scan(&leaveOrderPrefix)
	if err != nil {
		return err
	}
	if leaveOrderPrefix == "" {
		return gberror.New("config: LeaveOrderPrefix not found")
	}

	// 加班
	err = qConfig.WithContext(context.Background()).Select(qConfig.Value).Where(qConfig.Key.Eq("OvertimeOrderPrefix")).Scan(&overtimeOrderPrefix)
	if err != nil {
		return err
	}
	if overtimeOrderPrefix == "" {
		return gberror.New("config: OvertimeOrderPrefix not found")
	}

	// 補打卡
	err = qConfig.WithContext(context.Background()).Select(qConfig.Value).Where(qConfig.Key.Eq("CheckInOrderPrefix")).Scan(&checkInOrderPrefix)
	if err != nil {
		return err
	}
	if checkInOrderPrefix == "" {
		return gberror.New("config: CheckInOrderPrefix not found")
	}

	return nil
}

// LeaveOrder 請假單單號產生
func LeaveOrder(ctx context.Context) (string, error) {
	var (
		todayStr = time.Now().Format("20060102")
		key      = fmt.Sprintf("hrms:leaveOrder:%s", todayStr)
	)

	count, err := incr(ctx, key)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%s%06d", leaveOrderPrefix, todayStr, count), nil
}

// OvertimeOrder 加班單單號產生
func OvertimeOrder(ctx context.Context) (string, error) {
	var (
		todayStr = time.Now().Format("20060102")
		key      = fmt.Sprintf("hrms:overtimeOrder:%s", todayStr)
	)

	count, err := incr(ctx, key)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%s%06d", overtimeOrderPrefix, todayStr, count), nil
}

// CheckInOrder 補打卡單單號產生
func CheckInOrder(ctx context.Context) (string, error) {
	var (
		todayStr = time.Now().Format("20060102")
		key      = fmt.Sprintf("hrms:checkInOrder:%s", todayStr)
	)

	count, err := incr(ctx, key)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%s%06d", checkInOrderPrefix, todayStr, count), nil
}

// incr
func incr(ctx context.Context, key string) (int64, error) {
	count, err := g.Redis().Incr(ctx, key)
	if err != nil {
		return 0, err
	}
	if count <= 1 {
		tomorrow := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, 0, 0, 0, 0, time.Local)
		// 設定過期時間
		_, err = g.Redis().Expire(ctx, key, int64(tomorrow.Sub(time.Now()).Seconds()))
		if err != nil {
			return 0, err
		}
	}

	return count, nil
}
