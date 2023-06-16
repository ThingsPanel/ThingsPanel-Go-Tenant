package models

const (
	NotificationType_Members = 1
	NotificationType_Email   = 2
	NotificationType_Webhook = 3
)

const (
	NotificationSwitch_Open  = 1
	NotificationSwitch_Close = 2
)

type TpNotificationGroups struct {
	Id                 string `json:"id" gorm:"primaryKey"`
	GroupName          string `json:"group_name" `
	Desc               string `json:"desc" `
	NotificationType   int    `json:"notification_type"`
	Status             int    `json:"status"`
	NotificationConfig string `json:"notification_config"`
	TenantId           string `json:"tenant_id"`
}

func (TpNotificationGroups) TableName() string {
	return "tp_notification_groups"
}
