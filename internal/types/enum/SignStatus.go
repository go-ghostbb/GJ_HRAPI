package enum

import "database/sql/driver"

type SignStatus int64

const (
	// 送簽中
	SignStatusSending SignStatus = iota
	// 核准
	SignStatusApprove
	// 駁回
	SignStatusReject
	// 簽核中
	SignStatusSigning
	// 僅通知，通知成功
	SignStatusOnlyNotifySuc
	// 僅通知，通知失敗
	SignStatusOnlyNotifyFail
)

func (s *SignStatus) Scan(value interface{}) error {
	*s = SignStatus(value.(int64))
	return nil
}

func (s SignStatus) Value() (driver.Value, error) {
	return int64(s), nil
}
